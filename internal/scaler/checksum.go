package scaler

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"hash"

	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ByteMax = 0xFF
)

// K8sChecksum is a function which calculates a K8s Checksum
type K8sChecksum func(ctx context.Context, c client.Client, nsName string, h hash.Hash) error

func CalculateChangeChecksum(ctx context.Context, c client.Client, nsName string) (string, error) {
	shaSum := sha512.New()

	for _, f := range []K8sChecksum{ChecksumDeployments, ChecksumDaemonSets, ChecksumReplicaSets, ChecksumStatefulSets} {
		if err := f(ctx, c, nsName, shaSum); err != nil {
			return "", err
		}
	}

	sum := shaSum.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum), nil
}

// ChecksumDeployments calculates the checksum of the Generation of all Deployment objects
func ChecksumDeployments(ctx context.Context, c client.Client, nsName string, h hash.Hash) error {
	return WithDeployments(ctx, c, nsName, func(d appsv1.Deployment) error {
		return hashResource(ctx, h, d.Name, d.Kind, d.Generation)
	})
}

// ChecksumReplicaSets calculates the checksum of the Generation of all ReplicaSet objects
func ChecksumReplicaSets(ctx context.Context, c client.Client, nsName string, h hash.Hash) error {
	return WithReplicaSets(ctx, c, nsName, func(d appsv1.ReplicaSet) error {
		return hashResource(ctx, h, d.Name, d.Kind, d.Generation)
	})
}

// ChecksumDaemonSets calculates the checksum of the Generation of all DaemonSet objects
func ChecksumDaemonSets(ctx context.Context, c client.Client, nsName string, h hash.Hash) error {
	return WithDaemonSets(ctx, c, nsName, func(d appsv1.DaemonSet) error {
		return hashResource(ctx, h, d.Name, d.Kind, d.Generation)
	})
}

// ChecksumStatefulSets calculates the checksum of the Generation of all StatefulSet objects
func ChecksumStatefulSets(ctx context.Context, c client.Client, nsName string, h hash.Hash) error {
	return WithStatefulSets(ctx, c, nsName, func(d appsv1.StatefulSet) error {
		return hashResource(ctx, h, d.Name, d.Kind, d.Generation)
	})
}

func hashResource(ctx context.Context, h hash.Hash, name, kind string, generation int64) error {
	log.Ctx(ctx).Trace().Int64("generation", generation).Str("name", name).Str("kind", kind).Msg("Adding generation to checksum")
	if _, err := h.Write([]byte(name)); err != nil {
		return err
	}

	if _, err := h.Write(int64AsBytes(generation)); err != nil {
		return err
	}

	return nil
}

func int64AsBytes(i int64) []byte {
	return []byte{
		byte(ByteMax & i),
		byte(ByteMax & (i >> 8)),
		byte(ByteMax & (i >> 16)),
		byte(ByteMax & (i >> 24)),
		byte(ByteMax & (i >> 32)),
		byte(ByteMax & (i >> 40)),
		byte(ByteMax & (i >> 48)),
		byte(ByteMax & (i >> 56)),
	}
}
