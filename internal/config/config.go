package config

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-defaults"
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
	APIPath      string `yaml:"api-path" default:"/api" json:"api-path"`
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
func (c *Config) GetContext(name string) (*NamedContext, error) {
	for _, context := range c.Contexts {
		if context.Name == name {
			return context, nil
		}
	}

	return nil, fmt.Errorf("Context with name '%s' not found", name)
}

// UnmarshalYAML unmarshals the StsContext YAML part into a struct, ensuring that any defaults are set.
func (s *StsContext) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(s)

	type cfg StsContext

	if err := unmarshal((*cfg)(s)); err != nil {
		return err
	}

	return nil
}

// Merge merges the StsContext with a fallback object.
func (c *StsContext) Merge(fallback *StsContext) *StsContext {
	return &StsContext{
		URL:          util.DefaultIfEmpty(c.URL, fallback.URL),
		APIToken:     util.DefaultIfEmpty(c.APIToken, fallback.APIToken),
		ServiceToken: util.DefaultIfEmpty(c.ServiceToken, fallback.ServiceToken),
		APIPath:      util.DefaultIfEmpty(c.APIPath, fallback.APIPath),
	}
}

func (c *StsContext) Validate() error {
	errors := []error{}

	if c.URL == "" {
		errors = append(errors, MissingFieldError{FieldName: "url"})
	} else if !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://") {
		errors = append(errors, fmt.Errorf("URL %s must start with \"https://\" or \"http://\"", c.URL))
	}

	if c.APIToken == "" && c.ServiceToken == "" {
		errors = append(errors, MissingFieldError{FieldName: "{api-token | service-token}"})
	}

	if c.APIToken != "" && c.ServiceToken != "" {
		errors = append(errors, MissingFieldError{FieldName: "can only specify one of api-token an service-token"})
	}

	if len(errors) > 0 {
		return ValidateConfError{ValidationErrors: errors}
	}

	return nil
}
