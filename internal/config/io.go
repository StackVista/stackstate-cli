package config

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/pkg/xdg"
	"gopkg.in/yaml.v3"
)

const (
	ConfigFileDir            = "stackstate-cli"
	ConfigFileName           = "config.yaml"
	ConfigFilePathPermission = 0755 // rwxr-xr-x
	ConfigFilePermission     = 0644 // rw-r--r--

)

func DefaultConfigPath() (string, error) {
	path, err := xdg.GetXDGConfigHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, ConfigFileDir, ConfigFileName), nil
}

func ReadConfig(file string) (*Config, common.CLIError) {
	if file == "" {
		path, err := DefaultConfigPath()
		if err != nil {
			return nil, ReadConfError{err, true}
		}

		file = path
	}

	b, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ReadConfError{err, true}
		}
		return nil, ReadConfError{err, false}
	}

	return unmarshalYAMLConfig(b)
}

func unmarshalYAMLConfig(b []byte) (*Config, common.CLIError) {
	dec := yaml.NewDecoder(bytes.NewBuffer(b))
	cfg := &Config{}
	if err := dec.Decode(&cfg); err != nil {
		return nil, ReadConfError{err, false}
	}

	return cfg, nil
}

func WriteConfig(file string, cfg *Config) error {
	if file == "" {
		path, err := DefaultConfigPath()
		if err != nil {
			return err
		}

		file = path
	}

	if err := os.MkdirAll(filepath.Dir(file), ConfigFilePathPermission); err != nil {
		return err
	}

	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(file, b, ConfigFilePermission)
}
