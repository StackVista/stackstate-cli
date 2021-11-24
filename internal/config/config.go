package config

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-defaults"
)

type Config struct {
	URL      string `yaml:"url"`
	ApiToken string `yaml:"api-token"`
}

type ConfigMissingFieldError struct {
	FieldName string
}

func (s *ConfigMissingFieldError) Error() string {
	return fmt.Sprintf("Config is missing field: %v", s.FieldName)
}

type MultiConfigErrors struct {
	Errors []error
}

func (s MultiConfigErrors) Error() string {
	strs := make([]string, len(s.Errors))
	for _, e := range s.Errors {
		strs = append(strs, e.Error())
	}
	return strings.Join(strs, "\n")
}

func (s *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(s)

	type cfg Config

	if err := unmarshal((*cfg)(s)); err != nil {
		return err
	}

	return s.Validate()
}

func (s *Config) Validate() error {
	var errors []error
	if s.URL == "" {
		errors = append(errors, &ConfigMissingFieldError{FieldName: "url"})
	}
	if s.ApiToken == "" {
		errors = append(errors, &ConfigMissingFieldError{FieldName: "api-token"})
	}

	if len(errors) != 0 {
		return MultiConfigErrors{Errors: errors}
	} else {
		return nil
	}
}
