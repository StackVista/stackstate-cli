package scaler

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *scaler) ScaleUpNamespace(ctx context.Context, nsName string) error {
	logger := log.Ctx(ctx).With().Str("namespace", nsName).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("Restoring namespace")

	status, err := s.getStoredScaleStatus(ctx, nsName)
	if err != nil {
		return err
	}

	if !status.ScaledDown {
		return fmt.Errorf("namespace '%s' not scaled down, exiting", nsName)
	}

	if err := ScaleReplicas(ctx, s.klient, nsName, status.OriginalReplicas); err != nil {
		return err
	}

	logger.Info().Msg("All replicas restored to normal")

	// Retrieve new checksum as generations have changed
	newChecksum, err := CalculateChangeChecksum(ctx, s.klient, nsName)
	if err != nil {
		return err
	}

	// Update status
	status.Checksum = newChecksum
	status.Date = time.Now()
	status.ScaledDown = false

	return s.storeNewScaleStatus(ctx, nsName, status)
}
