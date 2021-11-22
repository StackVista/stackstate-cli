package values

import (
	"crypto/md5" //nolint:gosec
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/internal/yaml"
)

const (
	TokenLength = 32
)

const (
	None string = "None"
	KOPS string = "KOPS"
	K3D  string = "K3D"
)

func IngressTypes() []string {
	return []string{None, KOPS, K3D}
}

type HelmValues struct {
	PullSecretUsername   string `yaml:"pull-user"`
	PullSecretPassword   string `yaml:"pull-password"`
	PullPolicy           string `yaml:"pull-policy" default:"default"`
	LicenseKey           string `yaml:"license"`
	BaseURL              string `yaml:"base-url"`
	DefaultAdminPassword string `yaml:"default-password"`
	APIPassword          string `yaml:"api-password"`
	ExtraAPIKey          string `yaml:"extra-api-key"`
	K8sClusterName       string `yaml:"cluster-name"`
	Ingress              string `yaml:"ingress" default:"None"`
	InstallAgent         bool   `yaml:"install-agent"`
	OverrideStsImageTag  string `yaml:"override-sts-image-tag"`
}

func BuildValuesFile(helmValues HelmValues) (yaml.YamlDoc, error) {
	yamlDoc := yaml.YamlDoc{}

	apiKey, err := RandomToken(TokenLength)
	if err != nil {
		return nil, err
	}

	adminPasswordHash, err := Md5Sum([]byte(helmValues.DefaultAdminPassword))
	if err != nil {
		return nil, err
	}

	apiPasswordHash, err := Md5Sum([]byte(helmValues.APIPassword))
	if err != nil {
		return nil, err
	}

	yamlDoc.SetValue("global.receiverApiKey", apiKey)
	yamlDoc.SetValue("hbase.all.image.pullSecretUsername", helmValues.PullSecretUsername)
	yamlDoc.SetValue("hbase.all.image.pullSecretPassword", helmValues.PullSecretPassword)
	yamlDoc.SetValue("stackstate.components.all.image.pullSecretUsername", helmValues.PullSecretUsername)
	yamlDoc.SetValue("stackstate.components.all.image.pullSecretPassword", helmValues.PullSecretPassword)
	if helmValues.PullPolicy != "default" {
		yamlDoc.SetValue("stackstate.components.all.image.pullPolicy", helmValues.PullPolicy)
	}
	yamlDoc.SetValue("stackstate.baseUrl", helmValues.BaseURL)
	yamlDoc.SetValue("stackstate.license.key", helmValues.LicenseKey)
	yamlDoc.SetValue("stackstate.authentication.adminPassword", adminPasswordHash)
	yamlDoc.SetValue("stackstate.admin.authentication.password", apiPasswordHash)

	if helmValues.ExtraAPIKey != "" {
		yamlDoc.SetValue("stackstate.components.receiver.extraEnv.secret.EXTRA_API_KEY", helmValues.ExtraAPIKey)
	}

	if helmValues.OverrideStsImageTag != "" {
		yamlDoc.SetValue("stackstate.components.all.image.tag", helmValues.OverrideStsImageTag)
	}

	if helmValues.InstallAgent {
		if err := addClusterAgent(yamlDoc, helmValues); err != nil {
			return nil, err
		}
	}

	switch helmValues.Ingress {
	case None:
	case KOPS:
		if err := addKopsIngress(yamlDoc, helmValues); err != nil {
			return nil, err
		}
	case K3D:
		addK3dIngress(yamlDoc)
	default:
		return nil, errors.New("supported ingress types: None, KOPS, K3D")
	}

	return yamlDoc, nil
}

func addKopsIngress(yamlDoc yaml.YamlDoc, helmValues HelmValues) error {
	yamlDoc.SetValue("ingress.annotations", map[string]string{
		"kubernetes.io/ingress.class":                 "ingress-nginx-external",
		"cert-manager.io/cluster-issuer":              "letsencrypt-prod",
		"nginx.ingress.kubernetes.io/ingress.class":   "ingress-nginx-external",
		"nginx.ingress.kubernetes.io/proxy-body-size": "100m",
	})
	yamlDoc.SetValue("ingress.enabled", true)

	err := yamlDoc.AppendArray("ingress.hosts", map[string]interface{}{
		"host": fmt.Sprintf("{{ .Release.Namespace }}.%s", helmValues.K8sClusterName),
	})
	if err != nil {
		return err
	}

	return yamlDoc.AppendArray("ingress.tls", map[string]interface{}{
		"hosts": []string{
			fmt.Sprintf("{{ .Release.Namespace }}.%s", helmValues.K8sClusterName),
		},
		"secretName": "{{ .Release.Namespace }}-ecc-tls",
	})
}

func addK3dIngress(yamlDoc yaml.YamlDoc) {
	yamlDoc.SetValue("ingress.annotations", map[string]string{
		"ingress.kubernetes.io/ssl-redirect": "false",
	})
	yamlDoc.SetValue("ingress.enabled", true)
}

func addClusterAgent(yamlDoc yaml.YamlDoc, helmValues HelmValues) error {
	err := yamlDoc.AppendArray("stackstate.stackpacks.installed", map[string]interface{}{
		"name": "kubernetes",
		"configuration": map[string]interface{}{
			"kubernetes_cluster_name": helmValues.K8sClusterName,
		},
	})
	if err != nil {
		return err
	}

	yamlDoc.SetValue("cluster-agent.enabled", true)
	yamlDoc.SetValue("cluster-agent.stackstate.cluster.name", helmValues.K8sClusterName)

	authToken, err := RandomToken(TokenLength)
	if err != nil {
		return err
	}

	yamlDoc.SetValue("cluster-agent.stackstate.cluster.authToken", authToken)

	return nil
}

func RandomToken(size int) (string, error) {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return Md5Sum(key)
}

func Md5Sum(s []byte) (string, error) {
	hasher := md5.New() //nolint:gosec
	if _, err := hasher.Write(s); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
