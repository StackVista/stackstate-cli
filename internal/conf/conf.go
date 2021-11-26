package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Conf struct {
	ApiUrl   string
	ApiToken string
}

const (
	ViperConfigName        = "cli-config"
	ViperConfigType        = "yaml"
	ViperEnvPrefix         = "STS_CLI"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
	MinimumRequiredFlags   = "api-url, api-token"
)

func BindConfig(cmd *cobra.Command) Conf {
	// bind environment variable config
	viper.BindEnv("api.url", "STS_CLI_API_URL")
	viper.BindEnv("api.token", "STS_CLI_API_TOKEN")

	// bind cmd flags
	viper.BindPFlag("api.url", cmd.Flags().Lookup("api-url"))
	viper.BindPFlag("api.token", cmd.Flags().Lookup("api-token"))

	// read config
	return Conf{
		ApiUrl:   viper.GetString("api.url"),
		ApiToken: viper.GetString("api.token"),
	}
}
