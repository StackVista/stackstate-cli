package config

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-defaults"
)

type Config struct {
	Contexts       []NamedContext `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context"`
}

type NamedContext struct {
	Name    string     `yaml:"name"`
	Context StsContext `yaml:"context"`
}

type StsContext struct {
	URL          string `yaml:"url" json:"url"`
	ApiToken     string `yaml:"api-token,omitempty" json:"api-token,omitempty"`
	ServiceToken string `yaml:"service-token,omitempty" json:"service-token,omitempty"`
	ApiPath      string `yaml:"api-path" default:"/api" json:"api-path"`
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
		c.Contexts = []NamedContext{
			{
				Name:    "default",
				Context: *ctx,
			},
		}
		c.CurrentContext = "default"
	}

	return nil
}

// GetCurrentContext returns the current context
func (c *Config) GetCurrentContext() *StsContext {
	for _, context := range c.Contexts {
		if context.Name == c.CurrentContext {
			return &context.Context
		}
	}

	return nil
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
func (c *StsContext) Merge(fallback StsContext) StsContext {
	merged := StsContext{
		URL:          c.URL,
		ApiToken:     c.ApiToken,
		ServiceToken: c.ServiceToken,
		ApiPath:      c.ApiPath,
	}

	if merged.URL == "" {
		merged.URL = fallback.URL
	}

	if merged.ApiToken == "" {
		merged.ApiToken = fallback.ApiToken
	}

	if merged.ServiceToken == "" {
		merged.ServiceToken = fallback.ServiceToken
	}

	if merged.ApiPath == "" {
		merged.ApiPath = fallback.ApiPath
	}

	return merged
}

// Validate validates the Configuration
func (c *Config) Validate() error {
	errors := []error{}
	if c.CurrentContext == "" {
		errors = append(errors, MissingFieldError{FieldName: "current-context"})
	}

	for _, context := range c.Contexts {
		context.Context.validate(&errors)
	}

	if len(errors) > 0 {
		return ValidateConfError{ValidationErrors: errors}
	}

	return nil
}

func (c StsContext) validate(errors *[]error) {
	if c.URL == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "url"})
	} else if !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://") {
		*errors = append(*errors, fmt.Errorf("URL %s must start with \"https://\" or \"http://\"", c.URL))
	}

	if c.ApiToken == "" && c.ServiceToken == "" {
		*errors = append(*errors, MissingFieldError{FieldName: "{api-token | service-token}"})
	}

	if c.ApiToken != "" && c.ServiceToken != "" {
		*errors = append(*errors, MissingFieldError{FieldName: "can only specify one of api-token an service-token"})
	}
}
