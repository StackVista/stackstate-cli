package conf

import (
	"fmt"
	"os"
	"strings"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// readable by all the user groups, but writable by the user only.
	ConfFilePermission = 0644
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

func ReadConf(cmd *cobra.Command) (Conf, error) {
	homeFolder, err := home.Expand(HomePath)
	if err != nil {
		return Conf{}, err
	}
	configPaths := []string{
		homeFolder,
		".",
	}

	return readConfWithPaths(cmd, configPaths)
}

func readConfWithPaths(cmd *cobra.Command, paths []string) (Conf, error) {
	viper.Reset()
	// try read config file
	viper.SetConfigName(ViperConfigName)
	viper.SetConfigType(ViperConfigType)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
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

	conf := bind(cmd)

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

func WriteConf(conf Conf) error {
	path, err := home.Expand(HomePath)
	if err != nil {
		return err
	}

	return WriteConfTo(conf, path)
}

func WriteConfTo(conf Conf, path string) error {
	filename := path + "/" + ViperConfigName + "." + ViperConfigType
	err := ValidateConf(conf) // only want to write validated conf
	if err != nil {
		return nil
	}

	confYaml := convertConfToYaml(conf)
	err = os.WriteFile(filename, []byte(confYaml), ConfFilePermission)
	if err != nil {
		return nil
	}
	return nil
}