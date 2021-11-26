package conf

import (
	"os"

	home "github.com/mitchellh/go-homedir"
)

const (
	// readable by all the user groups, but writable by the user only.
	ConfFilePermission = 0644
)

func WriteConf(conf Conf) error {
	path, err := home.Expand(HomePath)
	if err != nil {
		return err
	}

	return WriteConfTo(conf, path)
}

func WriteConfTo(conf Conf, path string) error {
	filename := path + "/" + ViperConfigName + "." + ViperConfigType
	err := ValidateConf(conf) // only want to write validated conf
	if err != nil {
		return nil
	}

	confYaml := convertConfToYaml(conf)
	err = os.WriteFile(filename, []byte(confYaml), ConfFilePermission)
	if err != nil {
		return nil
	}
	return nil
}
