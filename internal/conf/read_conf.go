package conf

import (
	"fmt"
	"strings"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func ReadConf(cmd *cobra.Command) (Conf, error) {
	homeFolder, err := home.Expand("~/.stackstate")
	if err != nil {
		return Conf{}, err
	}
	configPaths := []string{
		homeFolder,
		".",
	}

	return ReadConfWithPaths(cmd, configPaths)
}

func ReadConfWithPaths(cmd *cobra.Command, paths []string) (Conf, error) {
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

	// set environment variable config
	viper.SetEnvPrefix(ViperEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// cmd flags
	viper.BindPFlag("api.url", cmd.Flags().Lookup("api-url"))
	viper.BindPFlag("api.token", cmd.Flags().Lookup("api-token"))

	// read config
	conf := Conf{
		ApiUrl:   viper.GetString("api.url"),
		ApiToken: viper.GetString("api.token"),
	}

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
