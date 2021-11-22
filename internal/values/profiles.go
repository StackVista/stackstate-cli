package values

import (
	"fmt"
	"path"
	"strings"

	"github.com/markbates/pkger"
	"gitlab.com/stackvista/stackstate-cli2/internal/yaml"
)

const ProfileDir = "/static/profiles"

func ListProfiles() ([]string, error) {
	validProfiles, err := pkger.Open(ProfileDir)
	if err != nil {
		return nil, err
	}

	defer validProfiles.Close()

	profiles, err := validProfiles.Readdir(0)
	if err != nil {
		return nil, err
	}

	profileNames := []string{}
	for _, p := range profiles {
		name := p.Name()
		ext := path.Ext(name)

		profileNames = append(profileNames, strings.TrimSuffix(name, ext))
	}

	return profileNames, nil
}

func LoadProfile(name string) (yaml.YamlDoc, error) {
	profile, err := pkger.Open(fmt.Sprintf("%s/%s.yaml", ProfileDir, name))
	if err != nil {
		return nil, err
	}

	defer profile.Close()

	return yaml.DecodeFromReader(profile)
}

func ApplyProfile(values yaml.YamlDoc, profileName string) (yaml.YamlDoc, error) {
	profile, err := LoadProfile(profileName)
	if err != nil {
		return nil, err
	}

	return values.Merge(profile)
}
