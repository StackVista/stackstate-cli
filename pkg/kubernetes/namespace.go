package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NamespaceExists(ctx context.Context, c client.Client, name string) (bool, error) {
	nl := &corev1.NamespaceList{}
	if err := c.List(ctx, nl); err != nil {
		return false, err
	}

	for _, ns := range nl.Items {
		if ns.Name == name {
			return true, nil
		}
	}

	return false, nil
}
