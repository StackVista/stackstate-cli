package config

import (
	"github.com/mcuadros/go-defaults"
)

type Config struct {
}

func (s *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(s)

	type cfg Config

	if err := unmarshal((*cfg)(s)); err != nil {
		return err
	}

	return nil
}
