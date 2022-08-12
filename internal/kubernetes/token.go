package kubernetes

import "io/ioutil"

const DefaultServiceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"

func GetServiceAccountToken(path string) (string, error) {
	if path == "" {
		path = DefaultServiceAccountTokenPath
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
