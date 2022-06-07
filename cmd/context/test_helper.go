package context

import (
	"path/filepath"
	"testing"

	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupConfig(t *testing.T, cli *di.MockDeps) {
	cli.StsConfig.CurrentContext = "foo"
	cli.StsConfig.Contexts = []*config.NamedContext{
		{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "/api")},
		{Name: "bar", Context: newContext("http://bar.com", "", "svctok-xxxx", "/api/v1")},
	}
	cli.ConfigPath = filepath.Join(t.TempDir(), "config.yaml")
}

func newContext(url, apiToken, serviceToken, apiPath string) *config.StsContext {
	return &config.StsContext{
		URL:          url,
		APIToken:     apiToken,
		ServiceToken: serviceToken,
		APIPath:      apiPath,
	}
}
