package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type LoadedConfiguration struct {
	StsConfig      *Config
	CurrentContext *StsContext
}

func LoadCurrentContext(cmd *cobra.Command, viper *viper.Viper, configPath string) (*StsContext, common.CLIError) {
	cfg, loadError := ReadConfig(configPath)

	vpConfig := Bind(cmd, viper)
	currCtx := vpConfig.CurrentContext
	currentContext := vpConfig.Context

	if cfg != nil { // Potentially no config file was found
		currCtx = util.DefaultIfEmpty(currCtx, cfg.CurrentContext)

		ctx, errx := cfg.GetContext(currCtx)
		if errx != nil {
			return nil, common.NewNotFoundError(errx)
		}

		// The Viper config is leading, i.e. it overrides the config file
		currentContext = currentContext.Merge(ctx.Context)
	} else {
		// Ensure default api-path is set
		currentContext.APIPath = util.DefaultIfEmpty(currentContext.APIPath, "/api")
	}

	if err := currentContext.Validate(currCtx); err != nil {
		if loadError != nil {
			return nil, loadError
		} else {
			return nil, err
		}
	}

	return currentContext, nil
}
