package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldUnmarshalConfig(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
- name: nightly
  context:
    url: http://nightly:8080
    k8s-sa-token: foobar
    api-path: /hidden/api
- name: prod
  context:
    url: http://prod:8080
    service-token: foo
    api-path: /hidden/api
current-context: prod
`

	cfg, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.Len(t, cfg.Contexts, 3)
	assert.Equal(t, "prod", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
	assert.Equal(t, "http://nightly:8080", cfg.Contexts[1].Context.URL)
	assert.Equal(t, "foobar", cfg.Contexts[1].Context.K8sSAToken)
	assert.Equal(t, "/hidden/api", cfg.Contexts[2].Context.APIPath)
	assert.Equal(t, "http://prod:8080", cfg.Contexts[2].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[2].Context.ServiceToken)
	assert.Equal(t, "/hidden/api", cfg.Contexts[1].Context.APIPath)
	assert.Empty(t, cfg.Contexts[1].Context.APIToken)
}

func TestShouldUnmarshalOldConfigFormat(t *testing.T) {
	config := `
url: http://localhost:8080
api-token: foo
`
	cfg, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.Len(t, cfg.Contexts, 1)
	assert.Equal(t, "default", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
}

func TestValidateValidStsContext(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.NoError(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name))
}

func TestValidateStsContextWithMissingTokens(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* Missing field '{api-token | service-token | k8s-sa-token}'")
}

func TestValidateStsContextWithMissingURL(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    api-token: foo
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* Missing field 'url'")
}

func TestValidateStsContextWithMalformedURL(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: localhost:8080
    api-token: foo
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* URL localhost:8080 must start with \"https://\" or \"http://\"")
}

func TestValidateWithMultipleErrors(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: localhost:8080
    api-token: foo
    service-token: bar
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), `Failed to validate the 'default' context:
* URL localhost:8080 must start with "https://" or "http://"
* Can only specify one of {api-token | service-token | k8s-sa-token}`)
}

func TestMergeWithNoTokenOverride(t *testing.T) {
	c := &StsContext{
		URL: "http://localhost:8080",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "foo", n.APIToken)
}

func TestMergeWithSameTokenOverride(t *testing.T) {
	c := &StsContext{
		URL:      "http://localhost:8080",
		APIToken: "bar",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "bar", n.APIToken)
}

func TestMergeWithOtherTokenOverride(t *testing.T) {
	c := &StsContext{
		URL:          "http://localhost:8080",
		ServiceToken: "bar",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "bar", n.ServiceToken)
	assert.Equal(t, "", n.APIToken)
}
