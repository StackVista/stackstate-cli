package kubernetes

import (
	home "github.com/mitchellh/go-homedir"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func KubeConfigFile() (string, error) {
	return home.Expand("~/.kube/config")
}

func SwitchNamespace(namespace string) error {
	return withConfig(func(c *api.Config) error {
		c.Contexts[c.CurrentContext].Namespace = namespace
		return nil
	})
}

func AddUser(name string, user *api.AuthInfo) error {
	return withConfig(func(c *api.Config) error {
		c.AuthInfos[name] = user
		return nil
	})
}

func ListClusters() ([]string, error) {
	clusters := []string{}

	if err := withConfig(func(c *api.Config) error {
		for k := range c.Clusters {
			clusters = append(clusters, k)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return clusters, nil
}

func CurrentCluster() (string, error) {
	cluster := ""
	if err := withConfig(func(c *api.Config) error {
		cluster = c.Contexts[c.CurrentContext].Cluster

		return nil
	}); err != nil {
		return "", err
	}

	return cluster, nil
}

func withConfig(apply func(*api.Config) error) error {
	f, err := KubeConfigFile()
	if err != nil {
		return err
	}

	cfg, err := clientcmd.LoadFromFile(f)
	if err != nil {
		return err
	}

	if err := apply(cfg); err != nil {
		return err
	}

	if err := clientcmd.WriteToFile(*cfg, f); err != nil {
		return err
	}

	return nil
}
