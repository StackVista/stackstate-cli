package scaler

import (
	"context"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
)

func (s *scaler) getStoredScaleStatus(ctx context.Context, nsName string) (*ScaleStatus, error) {
	cm := &corev1.ConfigMap{}
	if err := s.klient.Get(ctx, types.NamespacedName{Namespace: nsName, Name: ConfigMapKey}, cm); err != nil {
		if serr, ok := err.(*errors.StatusError); ok {
			if serr.ErrStatus.Code == http.StatusNotFound {
				return EmptyStatus(), nil
			}
		}

		return nil, err
	}

	return StatusFromConfigMapData(cm.Data)
}

func (s *scaler) storeNewScaleStatus(ctx context.Context, nsName string, status *ScaleStatus) error {
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		cm := &corev1.ConfigMap{}
		if err := s.klient.Get(ctx, types.NamespacedName{Namespace: nsName, Name: ConfigMapKey}, cm); err != nil {
			if serr, ok := err.(*errors.StatusError); ok {
				if serr.ErrStatus.Code == http.StatusNotFound {
					data, err := StatusToConfigMapData(status, map[string]string{})
					if err != nil {
						return err
					}

					return s.klient.Create(ctx, &corev1.ConfigMap{
						ObjectMeta: v1.ObjectMeta{
							Namespace: nsName,
							Name:      ConfigMapKey,
						},
						Data: data,
					})
				}
			}

			return err
		}

		_, err := StatusToConfigMapData(status, cm.Data)
		if err != nil {
			return err
		}

		return s.klient.Update(ctx, cm)
	})
}
