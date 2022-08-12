package config

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-defaults"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type Config struct {
	Contexts       []*NamedContext `yaml:"contexts" json:"contexts"`
	CurrentContext string          `yaml:"current-context" json:"current-context"`
}

type NamedContext struct {
	Name    string      `yaml:"name" json:"name"`
	Context *StsContext `yaml:"context" json:"context"`
}

type StsContext struct {
	URL          string `yaml:"url" json:"url"`
	APIToken     string `yaml:"api-token,omitempty" json:"api-token,omitempty"`
	ServiceToken string `yaml:"service-token,omitempty" json:"service-token,omitempty"`
	K8sSAToken   string `yaml:"k8s-sa-token,omitempty" json:"k8s-sa-token,omitempty"`
	APIPath      string `yaml:"api-path" default:"/api" json:"api-path"`
}

func EmptyConfig() *Config {
	return &Config{
		Contexts: []*NamedContext{},
	}
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	type cfg Config

	if err := unmarshal(&m); err != nil {
		return err
	}

	// Check whether this is in the new format
	if _, ok := m["contexts"]; ok {
		if err := unmarshal((*cfg)(c)); err != nil {
			return err
		}
	} else {
		// Fallback to the old format
		ctx := &StsContext{}
		if err := unmarshal(ctx); err != nil {
			return err
		}
		c.Contexts = []*NamedContext{
			{
				Name:    "default",
				Context: ctx,
			},
		}
		c.CurrentContext = "default"
	}

	return nil
}

// GetContext gets the context with the given name
func (c *Config) GetContext(name string) (*NamedContext, common.CLIError) {
	for _, context := range c.Contexts {
		if context.Name == name {
			return context, nil
		}
	}

	return nil, common.NewNotFoundError(fmt.Errorf("Context with name '%s' not found", name))
}

// UnmarshalYAML unmarshals the StsContext YAML part into a struct, ensuring that any defaults are set.
func (c *StsContext) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(c)

	type cfg StsContext

	if err := unmarshal((*cfg)(c)); err != nil {
		return err
	}

	return nil
}

// Merges the StsContext with a fallback object.
func (c *StsContext) Merge(fallback *StsContext) *StsContext {
	newCtx := &StsContext{
		URL:     util.DefaultIfEmpty(c.URL, fallback.URL),
		APIPath: util.DefaultIfEmpty(util.DefaultIfEmpty(c.APIPath, fallback.APIPath), "/api"),
	}

	tokenFromOverride := len(util.RemoveEmpty([]string{c.APIToken, c.ServiceToken, c.K8sSAToken})) > 0
	if !tokenFromOverride {
		newCtx.APIToken = fallback.APIToken
		newCtx.ServiceToken = fallback.ServiceToken
		newCtx.K8sSAToken = fallback.K8sSAToken
	} else {
		newCtx.APIToken = c.APIToken
		newCtx.ServiceToken = c.ServiceToken
		newCtx.K8sSAToken = c.K8sSAToken
	}

	return newCtx
}

func (c *StsContext) Validate(contextName string) common.CLIError {
	errors := []error{}

	switch {
	case c.URL == "":
		errors = append(errors, MissingFieldError{FieldName: "url"})
	case !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://"):
		errors = append(errors, fmt.Errorf("URL %s must start with \"https://\" or \"http://\"", c.URL))
	}

	if c.APIToken == "" && c.ServiceToken == "" && c.K8sSAToken == "" {
		errors = append(errors, MissingFieldError{FieldName: "{api-token | service-token | k8s-sa-token}"})
	}

	authenticationTokens := util.RemoveEmpty([]string{c.APIToken, c.ServiceToken, c.K8sSAToken})
	if len(authenticationTokens) > 1 {
		errors = append(errors, fmt.Errorf("Can only specify one of {api-token | service-token | k8s-sa-token}"))
	}

	if len(errors) > 0 {
		return ValidateContextError{
			ContextName:      contextName,
			ValidationErrors: errors,
		}
	}

	return nil
}
