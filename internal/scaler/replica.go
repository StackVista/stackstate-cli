package scaler

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	DaemonSetDummyLabel      = "com.stackstate.saas.scaler.daemonset"
	DaemonSetDummyLabelValue = "off"
)

type ReplicaCount map[string]int

type FetchReplicas func(ctx context.Context, c client.Client, nsName string) (ReplicaCount, error)

type SetReplicas func(ctx context.Context, c client.Client, nsName string, repls ReplicaCount) error

func RetrieveReplicas(ctx context.Context, c client.Client, nsName string) (map[string]ReplicaCount, error) {
	m := map[string]ReplicaCount{}

	funcs := map[string]FetchReplicas{
		"deployment":  fetchDeploymentReplicas,
		"replicaset":  fetchReplicaSetReplicas,
		"statefulset": fetchStatefulSetReplicas,
		"daemonset":   fetchDaemonSetReplicas,
	}

	for t, f := range funcs {
		r, err := f(ctx, c, nsName)
		if err != nil {
			return nil, err
		}

		m[t] = r
	}
	return m, nil
}

func ScaleReplicas(ctx context.Context, c client.Client, nsName string, m map[string]ReplicaCount) error {
	funcs := map[string]SetReplicas{
		"deployment":  setDeploymentReplicas,
		"replicaset":  setReplicaSetReplicas,
		"statefulset": setStatefulSetReplicas,
		"daemonset":   setDaemonSetReplicas,
	}

	for t, f := range funcs {
		err := f(ctx, c, nsName, m[t])
		if err != nil {
			return err
		}
	}

	return nil
}

func fetchDeploymentReplicas(ctx context.Context, c client.Client, nsName string) (ReplicaCount, error) {
	r := ReplicaCount{}
	err := WithDeployments(ctx, c, nsName, func(d appsv1.Deployment) error {
		r[d.Name] = intOrDefault(d.Spec.Replicas, 1)
		return nil
	})

	return r, err
}

func fetchStatefulSetReplicas(ctx context.Context, c client.Client, nsName string) (ReplicaCount, error) {
	r := ReplicaCount{}
	err := WithStatefulSets(ctx, c, nsName, func(d appsv1.StatefulSet) error {
		r[d.Name] = intOrDefault(d.Spec.Replicas, 1)
		return nil
	})

	return r, err
}

func fetchReplicaSetReplicas(ctx context.Context, c client.Client, nsName string) (ReplicaCount, error) {
	r := ReplicaCount{}
	err := WithReplicaSets(ctx, c, nsName, func(d appsv1.ReplicaSet) error {
		r[d.Name] = intOrDefault(d.Spec.Replicas, 1)
		return nil
	})

	return r, err
}

func fetchDaemonSetReplicas(ctx context.Context, c client.Client, nsName string) (ReplicaCount, error) {
	r := ReplicaCount{}
	err := WithDaemonSets(ctx, c, nsName, func(d appsv1.DaemonSet) error {
		r[d.Name] = int(d.Status.DesiredNumberScheduled) // The Desired count is taken as the replicas.
		return nil
	})

	return r, err
}

func intOrDefault(i *int32, def int) int {
	if i == nil {
		return def
	}

	return int(*i)
}

func setDeploymentReplicas(ctx context.Context, c client.Client, nsName string, r ReplicaCount) error {
	return WithDeployments(ctx, c, nsName, func(d appsv1.Deployment) error {
		return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			res := &appsv1.Deployment{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: d.Namespace, Name: d.Name}, res); err != nil {
				return err
			}

			i := int32(r[d.Name])
			res.Spec.Replicas = &i

			return c.Update(ctx, res)
		})
	})
}

func setReplicaSetReplicas(ctx context.Context, c client.Client, nsName string, r ReplicaCount) error {
	return WithReplicaSets(ctx, c, nsName, func(d appsv1.ReplicaSet) error {
		return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			res := &appsv1.ReplicaSet{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: d.Namespace, Name: d.Name}, res); err != nil {
				return err
			}

			i := int32(r[d.Name])
			res.Spec.Replicas = &i

			return c.Update(ctx, res)
		})
	})
}

func setStatefulSetReplicas(ctx context.Context, c client.Client, nsName string, r ReplicaCount) error {
	return WithStatefulSets(ctx, c, nsName, func(d appsv1.StatefulSet) error {
		return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			res := &appsv1.StatefulSet{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: d.Namespace, Name: d.Name}, res); err != nil {
				return err
			}

			i := int32(r[d.Name])
			res.Spec.Replicas = &i

			return c.Update(ctx, res)
		})
	})
}

func setDaemonSetReplicas(ctx context.Context, c client.Client, nsName string, r ReplicaCount) error {
	return WithDaemonSets(ctx, c, nsName, func(d appsv1.DaemonSet) error {
		return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			res := &appsv1.DaemonSet{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: d.Namespace, Name: d.Name}, res); err != nil {
				return err
			}

			if r[d.Name] == 0 {
				if res.Spec.Template.Spec.NodeSelector == nil {
					res.Spec.Template.Spec.NodeSelector = make(map[string]string)
				}
				// Scale DaemonSet down by adding dummy label
				res.Spec.Template.Spec.NodeSelector[DaemonSetDummyLabel] = DaemonSetDummyLabelValue
			} else {
				// Scale DaemonSet back up by removing dummy label
				delete(res.Spec.Template.Spec.NodeSelector, DaemonSetDummyLabel)
			}

			return c.Update(ctx, res)
		})
	})
}
