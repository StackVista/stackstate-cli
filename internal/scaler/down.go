package scaler

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func (s *scaler) ScaleDownAll(ctx context.Context) error {
	nsList := &corev1.NamespaceList{}
	if err := s.klient.List(ctx, nsList); err != nil {
		return err
	}

	for _, ns := range nsList.Items {
		if util.ContainsString(s.config.SystemNamespaces, ns.Name) {
			log.Ctx(ctx).Debug().Str("namespace", ns.Name).Msg("Skipping system namespace")
			continue
		}

		if IsAnnotatedWithSkip(ns) {
			log.Ctx(ctx).Debug().Str("namespace", ns.Name).Msg("Skipping annotated namespace")
			continue
		}

		if err := s.ScaleDownNamespace(ctx, ns.Name, false); err != nil {
			return err
		}
	}

	return nil
}

func (s *scaler) ScaleDownNamespace(ctx context.Context, name string, now bool) error {
	logger := log.Ctx(ctx).With().Str("namespace", name).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("Inspecting namespace")

	status, err := s.getStoredScaleStatus(ctx, name)
	if err != nil {
		return err
	}

	if status.ScaledDown {
		logger.Info().Msg("Namespace already scaled down according to status")
		return nil
	}

	if now {
		logger.Warn().Msg("Immediately scaling down namespace as requested")
		return s.scaleDownNamespaceNow(ctx, name, status)
	}

	return s.scaleDownNamespaceIfNeeded(ctx, name, status)
}

func (s *scaler) scaleDownNamespaceIfNeeded(ctx context.Context, nsName string, status *ScaleStatus) error {
	logger := log.Ctx(ctx)

	newChecksum, err := CalculateChangeChecksum(ctx, s.klient, nsName)
	if err != nil {
		return err
	}

	if newChecksum != status.Checksum {
		logger.Info().Str("new-checksum", newChecksum).Str("old-checksum", status.Checksum).Msg("Detected checksum change")

		repls, err := RetrieveReplicas(ctx, s.klient, nsName)
		if err != nil {
			return err
		}

		status.Checksum = newChecksum
		status.Date = time.Now()
		status.OriginalReplicas = repls

		return s.storeNewScaleStatus(ctx, nsName, status)
	}

	logger.Debug().Time("last-modified", status.Date).Msg("No change detected")

	if status.Date.Add(s.config.ScaleInterval).After(time.Now()) {
		logger.Info().Time("last-modified", status.Date).Dur("scale-interval", s.config.ScaleInterval).Msg("Scale down time not reached")
		return nil
	}

	return s.scaleDownNamespaceNow(ctx, nsName, status)
}

func (s *scaler) scaleDownNamespaceNow(ctx context.Context, nsName string, status *ScaleStatus) error {
	downed := map[string]ReplicaCount{}
	for k, v := range status.OriginalReplicas {
		r := ReplicaCount{}
		downed[k] = r
		for n := range v {
			r[n] = 0
		}
	}

	if err := ScaleReplicas(ctx, s.klient, nsName, downed); err != nil {
		return err
	}

	// Retrieve new checksum as generations have changed
	newChecksum, err := CalculateChangeChecksum(ctx, s.klient, nsName)
	if err != nil {
		return err
	}

	// Update status
	status.Checksum = newChecksum
	status.Date = time.Now()
	status.ScaledDown = true

	return s.storeNewScaleStatus(ctx, nsName, status)
}
