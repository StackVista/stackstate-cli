package kubernetes

import "os"

const DefaultServiceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token" //nolint:gosec

func GetServiceAccountToken(path string) (string, error) {
	if path == "" {
		path = DefaultServiceAccountTokenPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
