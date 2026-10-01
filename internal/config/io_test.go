package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigYAMLRoundTrip(t *testing.T) {
	const fixture = `contexts:
- name: production
  context: &connection
    url: https://suse-observability.example.com
    api-token: "0123"
    skip-ssl: yes
    ca-cert-base64-data: null
    ignored-field: ignored
- name: staging
  context: *connection
current-context: staging
`
	const expected = `contexts:
    - name: production
      context:
        url: https://suse-observability.example.com
        api-token: "0123"
        api-path: /api
        admin-api-path: /admin
        skip-ssl: true
    - name: staging
      context:
        url: https://suse-observability.example.com
        api-token: "0123"
        api-path: /api
        admin-api-path: /admin
        skip-ssl: true
current-context: staging
`
	path := filepath.Join(t.TempDir(), "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte(fixture), ConfigFilePermission))
	cfg, err := ReadConfig(path)
	require.NoError(t, err)
	require.Len(t, cfg.Contexts, 2)
	assert.Equal(t, cfg.Contexts[0].Context, cfg.Contexts[1].Context)

	cfg.Contexts[0].Context.K8sSAToken = "transient-token"
	cfg.Contexts[0].Context.K8sSATokenPath = "/transient/token"
	cfg.Contexts[0].Context.CaCertPath = "/transient/ca.crt"
	require.NoError(t, WriteConfig(path, cfg))
	saved, readErr := os.ReadFile(path)
	require.NoError(t, readErr)
	assert.Equal(t, expected, string(saved))

	reloaded, err := ReadConfig(path)
	require.NoError(t, err)
	assert.Equal(t, "staging", reloaded.CurrentContext)
	assert.Equal(t, "0123", reloaded.Contexts[0].Context.APIToken)
	assert.Empty(t, reloaded.Contexts[0].Context.K8sSAToken)
	assert.Empty(t, reloaded.Contexts[0].Context.K8sSATokenPath)
	assert.Empty(t, reloaded.Contexts[0].Context.CaCertPath)
}

func TestConfigYAMLRejectsInvalidInput(t *testing.T) {
	for _, fixture := range []string{
		"contexts: [",
		"contexts: []\ncontexts: []",
	} {
		t.Run(fixture, func(t *testing.T) {
			cfg, err := unmarshalYAMLConfig([]byte(fixture))
			require.Error(t, err)
			assert.Nil(t, cfg)
			assert.False(t, err.ShowUsage())
			assert.False(t, err.(ReadConfError).IsMissingConfigFile)
		})
	}
}
