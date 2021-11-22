package scaler

import (
	"context"

	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ConfigMapKey = "com.stackstate.devops.scaler"
)

type Scaler interface {
	ScaleDownAll(ctx context.Context) error
	ScaleDownNamespace(ctx context.Context, name string, now bool) error
	ScaleUpNamespace(ctx context.Context, name string) error
}

type scaler struct {
	klient client.Client
	config config.ScalerConfig
}

func NewScaler(ctx context.Context, config config.ScalerConfig, c client.Client) (Scaler, error) {
	log.Info().Interface("config", config).Msg("Starting scaler")

	return &scaler{
		klient: c,
		config: config,
	}, nil
}
