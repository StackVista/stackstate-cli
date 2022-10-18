package config

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/kubernetes"
	"github.com/stackvista/stackstate-cli/internal/util"
)

type LoadedConfiguration struct {
	StsConfig      *Config
	CurrentContext *StsContext
}

func LoadCurrentContext(ctx context.Context, cmd *cobra.Command, viper *viper.Viper, configPath string) (*StsContext, common.CLIError) {
	logger := log.Ctx(ctx)
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
		// Ensure default api-path is set, we don't want to specify the default in viper, as that would then always override the config file
		currentContext.APIPath = util.DefaultIfEmpty(currentContext.APIPath, "/api")
	}

	if !currentContext.HasAuthenticationTokenSet() {
		logger.Debug().Msg("Checking if running in Kubernetes")
		token, err := kubernetes.GetServiceAccountToken("")
		if err != nil {
			logger.Warn().Msg("Could not get Kubernetes ServiceAccount token")
		} else {
			currentContext.K8sSAToken = token

			logger.Info().Msg("Using Kubernetes ServiceAccount token for authentication")
		}
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
