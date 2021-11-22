package config

import (
	"time"

	"github.com/mcuadros/go-defaults"
	"gitlab.com/stackvista/stackstate-cli2/internal/values"
)

type Config struct {
	Sandbox        SandboxConfig        `yaml:"sandbox"`
	GenerateValues GenerateValuesConfig `yaml:"generate-values"`
	Dev            Dev                  `yaml:"dev"`
	Scale          ScalerConfig         `yaml:"scale"`
	TrafficMirror  TrafficMirrorConfig  `yaml:"traffic-mirror"`
}

type SandboxConfig struct {
	Username     string `yaml:"username"`
	SlackID      string `yaml:"slack-id"`
	ManualExpiry bool   `yaml:"manual-expiry"`
}

type GenerateValuesConfig struct {
	values.HelmValues `yaml:",inline"`
	Profile           string `yaml:"profile"`
}

type Dev struct {
	Registry string `yaml:"registry"`
}

type ScalerConfig struct {
	ScaleInterval    time.Duration `yaml:"scale-interval" default:"4h"`
	SystemNamespaces []string      `yaml:"system-namespaces" default:"kube-system,kube-public,kube-node-lease"`
}

type TrafficMirrorConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
}

func (s *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(s)

	type cfg Config

	if err := unmarshal((*cfg)(s)); err != nil {
		return err
	}

	return nil
}
