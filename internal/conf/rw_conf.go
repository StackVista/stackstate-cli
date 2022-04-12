package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ConfFilePath       = 0755 // rwxr-xr-x
	ConfFilePermission = 0644 // rw-r--r--
)

type ReadConfError struct {
	prefixMsg string
	RootCause error
}

func (s ReadConfError) Error() string {
	if s.prefixMsg != "" {
		return fmt.Sprintf("%s %s", s.prefixMsg, s.RootCause)
	} else {
		return s.RootCause.Error()
	}
}

type MissingConfError struct {
	Paths []string
}

func (s MissingConfError) Error() string {
	configFilenames := []string{}
	for _, path := range s.Paths {
		configFilenames = append(configFilenames, path+"/"+ViperConfigName+"."+ViperConfigType)
	}
	msg := fmt.Sprintf("Missing config. "+
		"Config can be provided via file, flags or environment variables.\n"+
		"Potential config file locations: %v.\n"+
		"Or use the command line flags: %v.\n"+
		"Or set environment variables: %v.",
		strings.Join(configFilenames, ", "),
		MinimumRequiredFlags,
		MinimumRequiredEnvVars,
	)
	return msg
}

// XDG spec https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func getXdgConfigHome() (xdgConfigHome string, err error) {
	xdgConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome, err = home.Expand("~/.config")
	}
	return
}

func getConfPath() (string, error) {
	xdgConfigHome, err := getXdgConfigHome()
	if err != nil {
		return "", err
	}

	return xdgConfigHome + "/" + XDGConfigSubPath, nil
}

func ReadConf(cmd *cobra.Command) (Conf, error) {
	confPath, err := getConfPath()
	if err != nil {
		return Conf{}, err
	}
	return readConfWithPaths(cmd, viper.GetViper(), []string{confPath})
}

func readConfWithPaths(cmd *cobra.Command, vp *viper.Viper, paths []string) (Conf, error) {
	// try read config file
	vp.SetConfigName(ViperConfigName)
	vp.SetConfigType(ViperConfigType)
	for _, path := range paths {
		vp.AddConfigPath(path)
	}

	if err := vp.ReadInConfig(); err != nil {
		// continue if config file not found
		// on missing config the paths are named in the error msg
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			return Conf{}, ReadConfError{
				prefixMsg: "Error while loading config file.",
				RootCause: err,
			}
		}
	}

	conf := bind(cmd, vp)

	// is config missing entirely?
	if (conf == Conf{}) {
		return Conf{}, ReadConfError{RootCause: MissingConfError{Paths: paths}}
	}

	// validate
	if err := ValidateConf(conf); err != nil {
		return Conf{}, ReadConfError{
			prefixMsg: "Could not load config, because of validation errors.",
			RootCause: err,
		}
	}

	return conf, nil
}

func WriteConf(conf Conf) (string, error) {
	confPath, err := getConfPath()
	if err != nil {
		return "", err
	}

	filename := confPath + "/" + ViperConfigName + "." + ViperConfigType
	err = WriteConfTo(conf, filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func WriteConfTo(conf Conf, filename string) error {
	err := ValidateConf(conf) // only want to write validated conf
	if err != nil {
		return nil
	}

	confYaml := convertConfToYaml(conf)

	if err := os.MkdirAll(filepath.Dir(filename), ConfFilePath); err != nil {
		return err
	}

	if err := os.WriteFile(filename, []byte(confYaml), ConfFilePermission); err != nil {
		return err
	}
	return nil
}
