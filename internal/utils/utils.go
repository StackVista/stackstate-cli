package utils

import (
	"errors"
	"os"
	"reflect"

	"github.com/markbates/pkger"
	"gitlab.com/stackvista/stackstate-cli2/internal/yaml"
)

func IsTty() bool {
	stat, _ := os.Stdin.Stat()
	return stat.Mode()&os.ModeCharDevice != 0
}

func StdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return stat.Mode()&os.ModeCharDevice == 0
}

func LoadEmbeddedConfig(configFile string, config interface{}) error {
	if reflect.TypeOf(config).Kind() != reflect.Ptr {
		return errors.New("expected a Pointer kind in LoadEmbeddedConfig")
	}

	f, err := pkger.Open(configFile)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := yaml.DecodeObjectFromReader(f, config); err != nil {
		return err
	}

	return nil
}
