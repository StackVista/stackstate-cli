package conf

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
This config is used throughout the CLI.

Note: when updating this struct, please also update the:
 - constants
 - bindings
 - and validations
 below.
*/
type Conf struct {
	ApiUrl   string
	ApiToken string
}

const (
	HomePath               = "~/.stackstate"
	ViperConfigName        = "cli-config"
	ViperConfigType        = "yaml"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
	MinimumRequiredFlags   = "api-url, api-token"
)

func bind(cmd *cobra.Command) Conf {
	// bind environment variable config
	viper.BindEnv("api.url", "STS_CLI_API_URL")
	viper.BindEnv("api.token", "STS_CLI_API_TOKEN")

	// bind cmd flags
	viper.BindPFlag("api.url", cmd.Flags().Lookup("api-url"))
	viper.BindPFlag("api.token", cmd.Flags().Lookup("api-token"))

	// read config from Viper
	return Conf{
		ApiUrl:   viper.GetString("api.url"),
		ApiToken: viper.GetString("api.token"),
	}
}

func validate(conf Conf, errors *[]error) {
	if conf.ApiUrl == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-url"})
	}
	if conf.ApiToken == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-token"})
	}
}

func convertConfToYaml(conf Conf) string {
	return fmt.Sprintf(`
api:
  url: %s
  token: %s`, conf.ApiUrl, conf.ApiToken)
}
