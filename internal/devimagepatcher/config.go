package devimagepatcher

import (
	"gitlab.com/stackvista/stackstate-cli2/internal/utils"
)

const ServicesConfigFile = "/static/config/stackstate-services.yaml"

type ServicesConfig = map[string]ServiceDescription

type ServiceDescription struct {
	Deployment string
	Container  string
	ImageName  string
	Groups     []string
}

func (sd ServiceDescription) IsInGroup(group *string) bool {
	if group != nil {
		for _, serviceGroup := range sd.Groups {
			if serviceGroup == *group {
				return true
			}
		}
	}
	return false
}

// LoadServicesConfig loads the embedded config file
func LoadServicesConfig() (*ServicesConfig, error) {
	cfg := &ServicesConfig{}
	err := utils.LoadEmbeddedConfig(ServicesConfigFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
