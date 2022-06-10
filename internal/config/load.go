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

func LoadCLIConfig(cmd *cobra.Command, viper *viper.Viper, configPath string) (*LoadedConfiguration, common.CLIError) {
	cfg, err := ReadConfig(configPath)
	if err != nil {
		return nil, err
	}

	// The Viper config is leading, i.e. it overrides the config file
	vpConfig := Bind(cmd, viper)
	currCtx := util.DefaultIfEmpty(vpConfig.CurrentContext, cfg.CurrentContext)

	ctx, errx := cfg.GetContext(currCtx)
	if errx != nil {
		return nil, common.NewNotFoundError(errx)
	}

	merged := vpConfig.Context.Merge(ctx.Context)
	currentContext := merged

	if err := currentContext.Validate(); err != nil {
		return nil, ReadConfError{
			RootCause:           err,
			IsMissingConfigFile: false,
		}
	}

	return &LoadedConfiguration{
		StsConfig:      cfg,
		CurrentContext: currentContext,
	}, nil
}
