package scaler

import (
	"context"
	"sort"
	"strings"

	appsv1 "k8s.io/api/apps/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// WithDeployments loops over all Deployment objects in the Namespace
func WithDeployments(ctx context.Context, c client.Client, nsName string, f func(d appsv1.Deployment) error) error {
	l := &appsv1.DeploymentList{}
	if err := c.List(ctx, l, &client.ListOptions{Namespace: nsName}); err != nil {
		return err
	}

	items := l.Items
	sort.Slice(items, func(i, j int) bool { return strings.Compare(items[i].Name, items[j].Name) < 0 })

	for _, i := range items {
		if err := f(i); err != nil {
			return err
		}
	}

	return nil
}

// WithStatefulSets loops over all StatefulSet objects in the Namespace
func WithStatefulSets(ctx context.Context, c client.Client, nsName string, f func(d appsv1.StatefulSet) error) error {
	l := &appsv1.StatefulSetList{}
	if err := c.List(ctx, l, &client.ListOptions{Namespace: nsName}); err != nil {
		return err
	}

	items := l.Items
	sort.Slice(items, func(i, j int) bool { return strings.Compare(items[i].Name, items[j].Name) < 0 })

	for _, i := range items {
		if err := f(i); err != nil {
			return err
		}
	}

	return nil
}

// WithDaemonSets loops over all DaemonSet objects in the Namespace
func WithDaemonSets(ctx context.Context, c client.Client, nsName string, f func(d appsv1.DaemonSet) error) error {
	l := &appsv1.DaemonSetList{}
	if err := c.List(ctx, l, &client.ListOptions{Namespace: nsName}); err != nil {
		return err
	}

	items := l.Items
	sort.Slice(items, func(i, j int) bool { return strings.Compare(items[i].Name, items[j].Name) < 0 })

	for _, i := range items {
		if err := f(i); err != nil {
			return err
		}
	}

	return nil
}

// WithReplicaSets loops over all ReplicaSet objects in the Namespace
func WithReplicaSets(ctx context.Context, c client.Client, nsName string, f func(d appsv1.ReplicaSet) error) error {
	l := &appsv1.ReplicaSetList{}
	if err := c.List(ctx, l, &client.ListOptions{Namespace: nsName}); err != nil {
		return err
	}

	items := l.Items
	sort.Slice(items, func(i, j int) bool { return strings.Compare(items[i].Name, items[j].Name) < 0 })

	for _, i := range items {
		if err := f(i); err != nil {
			return err
		}
	}

	return nil
}
