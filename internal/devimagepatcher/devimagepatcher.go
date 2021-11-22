package devimagepatcher

import (
	"context"
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/pkg/docker"
	stsKubernetes "gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
	k8sApiV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

const stackstateImageRegistry = "quay.io"
const stackstateImageRepository = "stackstate"

type DevImagePatcher struct {
	context      context.Context
	config       *config.Dev
	docker       *docker.Docker
	k8sClientset *kubernetes.Clientset
}

func Services() ([]string, error) {
	services, err := LoadServicesConfig()
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(*services))
	for k := range *services {
		keys = append(keys, k)
	}
	return keys, nil
}

func Groups() ([]string, error) {
	services, err := LoadServicesConfig()
	if err != nil {
		return nil, err
	}
	groups := map[string]bool{}
	for _, description := range *services {
		for _, g := range description.Groups {
			groups[g] = true
		}
	}

	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	return keys, nil
}

func New(ctx context.Context, config *config.Dev) (*DevImagePatcher, error) {
	docker, err := docker.New()
	if err != nil {
		return nil, err
	}

	k8sClientset, err := stsKubernetes.NewClientSet()
	if err != nil {
		return nil, err
	}

	return &DevImagePatcher{
		context:      ctx,
		config:       config,
		docker:       docker,
		k8sClientset: k8sClientset,
	}, nil
}

func (p *DevImagePatcher) PushNewDockerImages(serviceName *string, group *string, tag string) error {
	services, err := LoadServicesConfig()
	if err != nil {
		return err
	}

	imagesToPush := map[string]bool{}

	for service, description := range *services {
		if (serviceName != nil && *serviceName == service) || description.IsInGroup(group) {
			imagesToPush[description.ImageName] = true
		}
	}

	for image := range imagesToPush {
		err := p.pushNewImage(image, tag)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *DevImagePatcher) UpdateDeployments(serviceName *string, group *string, tag string, namespace string) error {
	services, err := LoadServicesConfig()
	if err != nil {
		return err
	}

	for service, description := range *services {
		if (serviceName != nil && *serviceName == service) || description.IsInGroup(group) {
			err := p.updateDeployment(service, description, tag, p.fallbackToNamespaceFromContext(namespace))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *DevImagePatcher) pushNewImage(image string, tag string) error {
	source := fmt.Sprintf("%s/%s/%s:%s", stackstateImageRegistry, stackstateImageRepository, image, tag)
	target := fmt.Sprintf("%s/%s/%s:%s", p.config.Registry, stackstateImageRepository, image, tag)

	fmt.Printf("Retagging %s to %s\n", source, target)
	err := p.docker.Tag(p.context, source, target)
	if err != nil {
		return err
	}

	fmt.Printf("Pushing image to registry %s\n", target)
	errPush := p.docker.PushImageAnonymous(p.context, target)

	if errPush != nil {
		return errPush
	}

	return nil
}

func (p *DevImagePatcher) updateDeployment(serviceName string, service ServiceDescription, tag string, namespace string) error {
	image := fmt.Sprintf("%s/%s/%s:%s", p.config.Registry, stackstateImageRepository, service.ImageName, tag)

	fmt.Printf("Patching deployment '%s' in namespace '%s' to use image '%s'\n", service.Deployment, namespace, image)

	patch := []byte(fmt.Sprintf(`{"spec":{ "template": {"spec":{"containers":[{"name":"%s","image":"%s"}]}}}}`, service.Container, image))
	_, err := p.k8sClientset.AppsV1().Deployments(namespace).Patch(p.context, service.Deployment, k8sTypes.StrategicMergePatchType, patch, k8sApiV1.PatchOptions{})

	if err != nil {
		return err
	}

	labelSelector := fmt.Sprintf("app.kubernetes.io/component=%s", serviceName)

	fmt.Printf("Removing all existing replicasets for deployment '%s' in namespace '%s' to ensure quick and smooth startup of the new pod\n", service.Deployment, namespace)
	err = p.k8sClientset.AppsV1().ReplicaSets(namespace).DeleteCollection(p.context, k8sApiV1.DeleteOptions{}, k8sApiV1.ListOptions{LabelSelector: labelSelector})

	return err
}

func (p *DevImagePatcher) fallbackToNamespaceFromContext(argumentNamespace string) string {
	if argumentNamespace != "" {
		return argumentNamespace
	}
	namespace, err := stsKubernetes.Namespace()
	if err != nil {
		return "default"
	}
	return namespace
}
