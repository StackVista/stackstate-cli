package conf

import (
	"fmt"
	"os"
	"strings"

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
	NoColor  bool
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
	viper.BindEnv("no_color", "STS_CLI_NO_COLOR")
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		viper.Set("no_color", true)
	}

	// bind cmd flags
	viper.BindPFlag("api.url", cmd.Flags().Lookup("api-url"))
	viper.BindPFlag("api.token", cmd.Flags().Lookup("api-token"))
	viper.BindPFlag("no_color", cmd.Flags().Lookup("no-color"))

	// read config from Viper
	return Conf{
		ApiUrl:   viper.GetString("api.url"),
		ApiToken: viper.GetString("api.token"),
		NoColor:  viper.GetBool("no_color"),
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
  token: %s
no_color: false
`, conf.ApiUrl, conf.ApiToken)
}
