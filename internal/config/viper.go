package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//nolint:golint,errcheck
func Bind(cmd *cobra.Command, vp *viper.Viper) *StsContext {
	// bind  variables
	vp.BindEnv("url", "STS_CLI_URL")
	vp.BindEnv("api-token", "STS_CLI_API_TOKEN")
	vp.BindEnv("service-token", "STS_CLI_SERVICE_TOKEN")
	vp.BindEnv("api-path", "STS_CLI_API_PATH")

	// bind flags
	vp.BindPFlag("url", cmd.Flags().Lookup("url"))
	vp.BindPFlag("api-token", cmd.Flags().Lookup("api-token"))
	vp.BindPFlag("service-token", cmd.Flags().Lookup("service-token"))
	vp.BindPFlag("api-path", cmd.Flags().Lookup("api-path"))

	// bind YAML
	return &StsContext{
		URL:          vp.GetString("url"),
		APIToken:     vp.GetString("api-token"),
		ServiceToken: vp.GetString("service-token"),
		APIPath:      vp.GetString("api-path"),
	}
}
