package conf

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

const (
	ConfFilePath       = 0755 // rwxr-xr-x
	ConfFilePermission = 0644 // rw-r--r--
)

type ReadConfError struct {
	RootCause           error
	IsMissingConfigFile bool
}

func (s ReadConfError) Error() string {
	if s.IsMissingConfigFile {
		return fmt.Sprintf("could not load StackState CLI config\n%s\nYou do not have a config file. To create one try running `sts cli save-config --help`", s.RootCause)
	} else {
		return fmt.Sprintf("could not load StackState CLI config\n%s", s.RootCause)
	}
}

func (p ReadConfError) ExitCode() common.ExitCode {
	return common.ConfigErrorExitCode
}

func (p ReadConfError) ShowUsage() bool {
	return false
}

func (p ReadConfError) GetServerResponse() *http.Response {
	return nil
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

func ReadConf(cmd *cobra.Command) (Conf, common.CLIError) {
	confPath, err := getConfPath()
	if err != nil {
		return Conf{}, ReadConfError{
			RootCause:           err,
			IsMissingConfigFile: true,
		}
	}
	return readConfWithPaths(cmd, viper.GetViper(), []string{confPath})
}

func readConfWithPaths(cmd *cobra.Command, vp *viper.Viper, paths []string) (Conf, common.CLIError) {
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
				RootCause:           err,
				IsMissingConfigFile: false,
			}
		}
	}

	conf := bind(cmd, vp)

	// validate
	if err := ValidateConf(conf); err != nil {
		return Conf{}, ReadConfError{
			RootCause:           err,
			IsMissingConfigFile: vp.ConfigFileUsed() == "",
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
