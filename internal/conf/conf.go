package conf

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
This config is used throughout the CLI.

Note: when updating this struct, please also update the:
 1. Conf struct
 2. Constants (e.g. MinimumRequiredEnvVars)
 2. Bindings (i.e. flags, environment variables and YAML)
 3. Validations
 4. and last but not least, the tests!
*/
type Conf struct {
	ApiURL   string
	ApiToken string
}

const (
	XDGConfigSubPath       = "stackstate-cli"
	ViperConfigName        = "config"
	ViperConfigType        = "yaml"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
	MinimumRequiredFlags   = "api-url, api-token"
)

//nolint:golint,errcheck
func bind(cmd *cobra.Command, vp *viper.Viper) Conf {
	// bind environment variables
	vp.BindEnv("api-url", "STS_CLI_API_URL")
	vp.BindEnv("api-token", "STS_CLI_API_TOKEN")

	// bind flags
	vp.BindPFlag("api-url", cmd.Flags().Lookup("api-url"))
	vp.BindPFlag("api-token", cmd.Flags().Lookup("api-token"))

	// bind YAML
	return Conf{
		ApiURL:   vp.GetString("api-url"),
		ApiToken: vp.GetString("api-token"),
	}
}

func validate(conf Conf, errors *[]error) {
	if conf.ApiURL == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-url"})
	}

	if conf.ApiToken == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-token"})
	}
}

func convertConfToYaml(conf Conf) string {
	return fmt.Sprintf(`api-url: %s
api-token: %s
`, conf.ApiURL, conf.ApiToken)
}
