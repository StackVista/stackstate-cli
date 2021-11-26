package conf

type Conf struct {
	ApiUrl   string
	ApiToken string
}

const (
	ViperConfigName        = "cli-config"
	ViperConfigType        = "yaml"
	ViperEnvPrefix         = "STS_CLI"
	MinimumRequiredEnvVars = "STS_CLI_API_URL, STS_CLI_API_TOKEN"
)
