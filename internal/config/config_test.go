package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestShouldUnmarshalConfig(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
- name: prod
  context:
    url: http://prod:8080
    service-token: foo
    api-path: /hidden/api
current-context: prod
`

	d := yaml.NewDecoder(bytes.NewBufferString(config))
	cfg := &Config{}
	assert.NoError(t, d.Decode(&cfg))
	assert.Len(t, cfg.Contexts, 2)
	assert.Equal(t, "prod", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
	assert.Equal(t, "http://prod:8080", cfg.Contexts[1].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[1].Context.ServiceToken)
	assert.Equal(t, "/hidden/api", cfg.Contexts[1].Context.APIPath)
	assert.Empty(t, cfg.Contexts[1].Context.APIToken)
}

func TestShouldUnmarshalOldConfigFormat(t *testing.T) {
	config := `
url: http://localhost:8080
api-token: foo
`
	d := yaml.NewDecoder(bytes.NewBufferString(config))
	cfg := &Config{}
	assert.NoError(t, d.Decode(&cfg))
	assert.Len(t, cfg.Contexts, 1)
	assert.Equal(t, "default", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
}
