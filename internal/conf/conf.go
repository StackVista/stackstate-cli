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
 1. Conf struct
 2. Constants (e.g. MinimumRequiredEnvVars)
 2. Bindings (i.e. flags, environment variables and YAML)
 3. Validations
 4. and last but not least, the tests!
*/
type Conf struct {
	ApiUrl   string
	ApiToken string
	NoColor  bool
	Output   bool
}

const (
	HomePath               = "~/.stackstate"
	ViperConfigName        = "cli-config"
	ViperConfigType        = "yaml"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
	MinimumRequiredFlags   = "api-url, api-token"
)

func bind(cmd *cobra.Command, vp *viper.Viper) Conf {
	// bind environment variables
	vp.BindEnv("api.url", "STS_CLI_API_URL")
	vp.BindEnv("api.token", "STS_CLI_API_TOKEN")
	vp.BindEnv("no_color", "STS_CLI_NO_COLOR")
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		vp.Set("no_color", true)
	}
	vp.BindEnv("output", "STS_CLI_OUTPUT")

	// bind flags
	vp.BindPFlag("api.url", cmd.Flags().Lookup("api-url"))
	vp.BindPFlag("api.token", cmd.Flags().Lookup("api-token"))
	vp.BindPFlag("no_color", cmd.Flags().Lookup("no-color"))
	vp.BindPFlag("output", cmd.Flags().Lookup("output"))

	// bind YAML
	return Conf{
		ApiUrl:   vp.GetString("api.url"),
		ApiToken: vp.GetString("api.token"),
		NoColor:  vp.GetBool("no_color"),
		Output:   vp.GetBool("output"),
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
