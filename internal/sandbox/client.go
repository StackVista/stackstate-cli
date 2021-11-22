package sandbox

import (
	"context"
	"fmt"

	devopsv1 "github.com/stackvista/sandbox-operator/apis/devops/v1"
	"github.com/stackvista/sandbox-operator/pkg/client/versioned"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Sandboxer struct {
	client *versioned.Clientset
	config *rest.Config
}

func New() (*Sandboxer, error) {
	cfg, client, err := connectToKubernetesSandboxAPI()
	if err != nil {
		return nil, err
	}

	return &Sandboxer{
		client: client,
		config: cfg,
	}, nil
}

func (sb *Sandboxer) Get(ctx context.Context, name string) (*devopsv1.Sandbox, error) {
	return sb.client.DevopsV1().Sandboxes().Get(ctx, name, v1.GetOptions{})
}

func (sb *Sandboxer) Create(ctx context.Context, name string, spec *devopsv1.SandboxSpec) (*devopsv1.Sandbox, error) {
	box := &devopsv1.Sandbox{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: *spec,
	}

	return sb.client.DevopsV1().Sandboxes().Create(ctx, box, v1.CreateOptions{})
}

func (sb *Sandboxer) Delete(ctx context.Context, name, username string) error {
	boxes := sb.client.DevopsV1().Sandboxes()

	b, err := boxes.Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return err
	}

	if b.Spec.User != username {
		return fmt.Errorf("you do not own the sandbox named '%s'", name)
	}

	return boxes.Delete(ctx, name, v1.DeleteOptions{})
}

func (sb *Sandboxer) ListMySandboxes(ctx context.Context, username string) ([]string, error) {
	list, err := sb.client.DevopsV1().Sandboxes().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sandboxes := []string{}
	for _, s := range list.Items {
		if s.Spec.User == username {
			sandboxes = append(sandboxes, s.Name)
		}
	}

	return sandboxes, nil
}

func connectToKubernetesSandboxAPI() (*rest.Config, *versioned.Clientset, error) {
	f, err := kubernetes.KubeConfigFile()
	if err != nil {
		return nil, nil, err
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", f)
	if err != nil {
		return nil, nil, err
	}

	api, err := versioned.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	if err := verifySupportsSandboxes(cfg, api); err != nil {
		return nil, nil, err
	}

	return cfg, api, nil
}

func verifySupportsSandboxes(cfg *rest.Config, client *versioned.Clientset) error {
	serverGroups, err := client.Discovery().ServerGroups()
	if err != nil {
		return err
	}

	for _, g := range serverGroups.Groups {
		if g.Name == devopsv1.GroupVersion.Group {
			return nil
		}
	}

	return fmt.Errorf("the cluster at %s does not support '%s'", cfg.Host, devopsv1.GroupVersion)
}

func (sb *Sandboxer) GetKubernetesHost() string {
	return sb.config.Host
}
