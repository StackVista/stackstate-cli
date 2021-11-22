package kops

import (
	"gitlab.com/stackvista/stackstate-cli2/internal/utils"
)

const KopsConfFile = "/static/config/kops.yaml"

type Config struct {
	Kops *Kops `yaml:"kops"`
	AWS  *AWS  `yaml:"aws"`
}

type AWS struct {
	ProfilePrefix string        `yaml:"profilePrefix"`
	Profiles      []*AWSProfile `yaml:"profiles"`
}

type AWSProfile struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

type Kops struct {
	BucketSuffix string `yaml:"bucketSuffix"`
}

// LoadKopsConfig loads the embedded config file
func LoadKopsConfig() (*Config, error) {
	cfg := &Config{}
	err := utils.LoadEmbeddedConfig(KopsConfFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
