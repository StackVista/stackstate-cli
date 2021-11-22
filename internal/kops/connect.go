package kops

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/kubeconfig"

	"k8s.io/kops/cmd/kops/util"
)

func Connect(ctx context.Context, fqn string, admin bool, profile string) error {
	cfg, err := LoadKopsConfig()
	if err != nil {
		return err
	}

	if err := kopsBuildClusterConfig(ctx, cfg, fqn); err != nil {
		return err
	}

	if err := addRbacUser(ctx, cfg, fqn, admin, profile); err != nil {
		return err
	}

	return nil
}

func kopsBuildClusterConfig(ctx context.Context, cfg *Config, fqn string) error {
	s := strings.Split(fqn, ".")
	kopsRegistryPath := fmt.Sprintf("s3://%s-%s-%s", s[0], s[1], cfg.Kops.BucketSuffix)

	awsProfile := fmt.Sprintf("%s-%s", cfg.AWS.ProfilePrefix, s[1])

	os.Setenv("AWS_PROFILE", awsProfile)
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	f := util.NewFactory(&util.FactoryOptions{
		RegistryPath: kopsRegistryPath,
	})

	clientSet, err := f.Clientset()
	if err != nil {
		return err
	}

	cluster, err := clientSet.GetCluster(ctx, fqn)
	if err != nil {
		return err
	}

	keyStore, err := clientSet.KeyStore(cluster)
	if err != nil {
		return err
	}

	secretStore, err := clientSet.SecretStore(cluster)
	if err != nil {
		return err
	}

	b, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, &commands.CloudDiscoveryStatusStore{}, 0, "", false, f.KopsStateStore(), false)
	if err != nil {
		return err
	}

	if err := b.WriteKubecfg(clientcmd.NewDefaultPathOptions()); err != nil {
		return err
	}

	return nil
}

func addRbacUser(_ context.Context, cfg *Config, fqn string, admin bool, connectionProfile string) error {
	profileName := fmt.Sprintf("%s-%s", cfg.AWS.ProfilePrefix, strings.Split(fqn, ".")[1])

	var awsProfile *AWSProfile
	for _, p := range cfg.AWS.Profiles {
		if p.Name == profileName {
			awsProfile = p
		}
	}

	if awsProfile == nil {
		return fmt.Errorf("unknown AWS Profile '%s'", profileName)
	}

	role := "Developer"
	if admin {
		role = "Administrator"
	}

	user := api.NewAuthInfo()
	args := []string{"eks", "get-token", "--cluster-name", fqn, "--role-arn", fmt.Sprintf("arn:aws:iam::%s:role/%s", awsProfile.ID, role)}

	if strings.TrimSpace(connectionProfile) != "" {
		args = append(args, "--profile", connectionProfile)
	}

	user.Exec = &api.ExecConfig{
		APIVersion: "client.authentication.k8s.io/v1alpha1",
		Command:    "aws",
		Args:       args,
		Env:        nil,
	}

	return kubernetes.AddUser(fqn, user)
}
