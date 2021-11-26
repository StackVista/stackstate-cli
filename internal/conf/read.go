package conf

import (
	"fmt"
	"strings"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	ViperConfigName        = "cli-config"
	ViperConfigType        = "yaml"
	ViperEnvPrefix         = "STS_CLI"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
)

type ReadConfError struct {
	prefixMsg string
	RootCause error
}

func (s ReadConfError) Error() string {
	return fmt.Sprintf("%s %s", s.prefixMsg, s.RootCause)
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
		"Potential config file locations: %v\n"+
		"Or set environment variables: %v\n", strings.Join(configFilenames, ", "), MinimumRequiredEnvVars)
	return msg
}

/*
TODO:
- Flags
*/

func ReadConf() (Conf, error) {
	homeFolder, err := home.Expand("~/.stackstate")
	if err != nil {
		return Conf{}, err
	}
	configPaths := []string{
		homeFolder,
		".",
	}

	return ReadConfWithPaths(configPaths)
}

func ReadConfWithPaths(paths []string) (Conf, error) {
	// try read config file
	viper.SetConfigName(ViperConfigName)
	viper.SetConfigType(ViperConfigType)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	var configFileNotFoundError *viper.ConfigFileNotFoundError
	if err := viper.ReadInConfig(); err != nil {
		if fileNotFoundErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			configFileNotFoundError = &fileNotFoundErr
		} else {
			// Config file was found but another error was produced
			return Conf{}, ReadConfError{
				prefixMsg: "Error while loading config file.",
				RootCause: err,
			}
		}
	}

	// set environment variable config
	viper.SetEnvPrefix("STS_CLI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// read config
	conf := Conf{
		ApiUrl:   viper.GetString("api.url"),
		ApiToken: viper.GetString("api.token"),
	}

	// if at this point there is still no config and there was no config file
	// blame it on the config file
	if (conf == Conf{} && configFileNotFoundError != nil) {
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
