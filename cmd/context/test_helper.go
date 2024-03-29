package context

import (
	"path/filepath"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func setupConfig(t *testing.T, cli *di.MockDeps) {
	cfg := &config.Config{
		CurrentContext: "foo",
		Contexts: []*config.NamedContext{
			{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "", "/api")},
			{Name: "bar", Context: newContext("http://bar.com", "", "svctok-xxxx", "", "/api/v1")},
			{Name: "foobar", Context: newContext("http://bar.com", "", "", "eyJhbGc", "/api/v1")},
		},
	}
	cli.ConfigPath = filepath.Join(t.TempDir(), "config.yaml")

	err := config.WriteConfig(cli.ConfigPath, cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func newContext(url, apiToken, serviceToken, k8sSAToken, apiPath string) *config.StsContext {
	return &config.StsContext{
		URL:          url,
		APIToken:     apiToken,
		ServiceToken: serviceToken,
		K8sSAToken:   k8sSAToken,
		APIPath:      apiPath,
	}
}
