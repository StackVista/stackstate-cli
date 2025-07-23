package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	CurrentContext string
	Context        *StsContext
}

//nolint:golint,errcheck
func Bind(cmd *cobra.Command, vp *viper.Viper) *ViperConfig {
	// bind  variables
	vp.BindEnv("url", "STS_CLI_URL")
	vp.BindEnv("api-token", "STS_CLI_API_TOKEN")
	vp.BindEnv("service-token", "STS_CLI_SERVICE_TOKEN")
	vp.BindEnv("k8s-sa-token", "STS_CLI_K8S_SA_TOKEN")
	vp.BindEnv("k8s-sa-token-path", "STS_CLI_K8S_SA_TOKEN_PATH")
	vp.BindEnv("api-path", "STS_CLI_API_PATH")
	vp.BindEnv("context", "STS_CLI_CONTEXT")
	vp.BindEnv("skip-ssl", "STS_SKIP_SSL")
	vp.BindEnv("ca-cert-path", "STS_CA_CERT_PATH")
	vp.BindEnv("ca-cert-base64-data", "STS_CA_CERT_BASE64_PATH")

	// bind flags
	vp.BindPFlag("url", cmd.Flags().Lookup("url"))
	vp.BindPFlag("api-token", cmd.Flags().Lookup("api-token"))
	vp.BindPFlag("service-token", cmd.Flags().Lookup("service-token"))
	vp.BindPFlag("k8s-sa-token", cmd.Flags().Lookup("k8s-sa-token"))
	vp.BindPFlag("k8s-sa-token-path", cmd.Flags().Lookup("k8s-sa-token-path"))
	vp.BindPFlag("api-path", cmd.Flags().Lookup("api-path"))
	vp.BindPFlag("context", cmd.Flags().Lookup("context"))
	vp.BindPFlag("skip-ssl", cmd.Flags().Lookup("skip-ssl"))
	vp.BindPFlag("ca-cert-path", cmd.Flags().Lookup("ca-cert-path"))
	vp.BindPFlag("ca-cert-base64-data", cmd.Flags().Lookup("ca-cert-base64-data"))

	// bind YAML
	return &ViperConfig{
		CurrentContext: vp.GetString("context"),
		Context: &StsContext{
			URL:              vp.GetString("url"),
			APIToken:         vp.GetString("api-token"),
			ServiceToken:     vp.GetString("service-token"),
			K8sSAToken:       vp.GetString("k8s-sa-token"),
			K8sSATokenPath:   vp.GetString("k8s-sa-token-path"),
			APIPath:          vp.GetString("api-path"),
			SkipSSL:          vp.GetBool("skip-ssl"),
			CaCertBase64Data: vp.GetString("ca-cert-base64-data"),
			CaCertPath:       vp.GetString("ca-cert-path"),
		},
	}
}
