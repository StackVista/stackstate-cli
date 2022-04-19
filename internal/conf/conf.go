package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
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
	NoColor  bool
	Output   string
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
	vp.BindEnv("no-color", "STS_CLI_NO_COLOR")
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		vp.Set("no-color", true)
	}
	// Implementation of https://no-color.org/
	if _, no_color_env_exists := os.LookupEnv("NO_COLOR"); no_color_env_exists {
		vp.Set("no-color", true)
	}
	vp.BindEnv("output", "STS_CLI_OUTPUT")

	// bind flags
	vp.BindPFlag("api-url", cmd.Flags().Lookup("api-url"))
	vp.BindPFlag("api-token", cmd.Flags().Lookup("api-token"))
	vp.BindPFlag("no-color", cmd.Flags().Lookup("no-color"))
	vp.BindPFlag("output", cmd.Flags().Lookup("output"))

	// bind YAML
	return Conf{
		ApiURL:   vp.GetString("api-url"),
		ApiToken: vp.GetString("api-token"),
		NoColor:  vp.GetBool("no-color"),
		Output:   vp.GetString("output"),
	}
}

func validate(conf Conf, errors *[]error) {
	if conf.ApiURL == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-url"})
	}

	if conf.ApiToken == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "api-token"})
	}

	outputChoices := []string{"Auto", "JSON", "YAML", ""}
	if !util.StringInSliceIgnoreCase(conf.Output, outputChoices) {
		*errors = append(*errors, MustBeOneOfError{
			FieldName: "output",
			Value:     conf.Output,
			Choices:   outputChoices,
		})
	}
}

func convertConfToYaml(conf Conf) string {
	return fmt.Sprintf(`api-url: %s
api-token: %s
no-color: %v
`, conf.ApiURL, conf.ApiToken, conf.NoColor)
}
