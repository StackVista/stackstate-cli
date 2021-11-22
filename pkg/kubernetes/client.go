package kubernetes

import (
	home "github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DefaultKubeConfig() (*rest.Config, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil && err != rest.ErrNotInCluster {
		return nil, err
	} else if err != nil {
		kubeconfig, err := home.Expand("~/.kube/config")
		if err != nil {
			return nil, err
		}

		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func Namespace() (string, error) {
	kubeconfig, err := home.Expand("~/.kube/config")
	if err != nil {
		return "", err
	}

	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{})

	namespace, _, err := config.Namespace()

	if namespace == "" && err != nil {
		namespace = "default"
	}

	return namespace, err
}

func NewClientForCluster(name string) (client.Client, error) {
	f, err := KubeConfigFile()
	if err != nil {
		return nil, err
	}

	cfg, err := clientcmd.LoadFromFile(f)
	if err != nil {
		return nil, err
	}

	for n, c := range cfg.Contexts {
		if c.Cluster == name {
			cfg.CurrentContext = n
		}
	}

	realCfg, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
		return cfg, nil
	})
	if err != nil {
		return nil, err
	}

	c, err := client.New(realCfg, client.Options{})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewClient() (client.Client, error) {
	cfg, err := DefaultKubeConfig()
	if err != nil {
		return nil, err
	}

	c, err := client.New(cfg, client.Options{})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewClientSet() (*kubernetes.Clientset, error) {
	config, err := DefaultKubeConfig()
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}
