package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/pkg/xdg"
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
			return nil, ReadConfError{err, false}
		}

		file = path
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ReadConfError{err, true}
		}
		return nil, ReadConfError{err, false}
	}

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

	return ioutil.WriteFile(file, b, ConfigFilePermission)
}
