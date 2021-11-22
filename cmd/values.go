package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/utils"
	"gitlab.com/stackvista/stackstate-cli2/internal/values"
	"gitlab.com/stackvista/stackstate-cli2/pkg/asp"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
)

func ValuesCommand(cfg *config.Config) *cobra.Command {
	var outputFile string

	var nonInteractive bool

	cmd := &cobra.Command{
		Use:   "generate-values",
		Short: "Generate a values.yaml for deploying StackState to Kubernetes with Helm 3",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if utils.IsTty() && !nonInteractive {
				if err := InteractiveValues(cmd, &cfg.GenerateValues); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := values.BuildValuesFile(cfg.GenerateValues.HelmValues)
			if err != nil {
				return err
			}

			if strings.TrimSpace(cfg.GenerateValues.Profile) != "" {
				merged, err := values.ApplyProfile(v, cfg.GenerateValues.Profile)
				if err != nil {
					return err
				}

				v = merged
			}

			return v.EncodeToFile(outputFile)
		},
	}

	cmd.SetOut(os.Stdout)

	cmd.Flags().StringP("pull-user", "u", "", "Username for Docker image pulling")
	cmd.Flags().StringP("pull-password", "p", "", "Password for Docker image pulling")
	cmd.Flags().StringP("pull-policy", "c", "", "Pull policy for StackState docker images")
	cmd.Flags().StringP("license", "l", "", "StackState license key")
	cmd.Flags().StringP("base-url", "b", "", "StackState base URL, externally (outside of the Kubernetes cluster) visible url of the StackState endpoints (i.e. https://my.stackstate.host)")
	cmd.Flags().StringP("default-password", "d", "", "Password that will be set for the default StackState 'admin' user")
	cmd.Flags().StringP("api-password", "a", "", "Password that will be set for StackState's admin api, access should be restricted (Dev)Ops")
	cmd.Flags().StringP("cluster-name", "k", "", "If provided will install the StackState Kubernetes agent. StackState and the agent will refer to the cluster by this name")
	cmd.Flags().StringP("ingress", "i", "", "Enable KOPS or K3D ingress for your installation. Be aware that KOPS Ingress exposes your installation on the public internet, please use strong passwords")
	cmd.Flags().StringP("extra-api-key", "e", "", "Provide an extra receiver API key for which traffic will be accepted. Use this for receiving data from traffic mirror.")
	cmd.Flags().StringP("override-sts-image-tag", "t", "", "Instead of the StackState image tag(s) in the helm chart deploy with this tag (i.e. tag for your development branch).")
	cmd.Flags().String("profile", "non-ha", "Apply a specific profile to the generated values file. The profiles set a range of default settings for the StackState deployment.")
	cmd.Flags().Bool("install-agent", false, "Install the StackState agent on the Kubernetes cluster.")

	cmd.Flags().StringVarP(&outputFile, "output", "o", "values.yaml", "Name of generated values file")
	cmd.Flags().BoolVarP(&nonInteractive, "non-interactive", "n", false, "Non-interactive mode")

	return cmd
}

func InteractiveValues(cmd *cobra.Command, cfg *config.GenerateValuesConfig) error { //nolint:funlen
	currentCluster, err := kubernetes.CurrentCluster()
	if err != nil {
		return err
	}

	err = asp.AskQuestions([]asp.InteractiveValue{
		asp.NewStringValue(cmd, "license", "Please provide the license key (-l)", &cfg.LicenseKey),
		asp.NewStringValue(cmd, "pull-user", "Please provide the username for pulling StackState Docker images (-u)", &cfg.PullSecretUsername),
		asp.NewPasswordValue(cmd, "pull-password", "Please provide the password for pulling StackState Docker images (-p)", &cfg.PullSecretPassword),
		asp.NewPasswordValue(cmd, "default-password", "Please provide the password for the 'admin' user that will be set for StackState (-d)", &cfg.DefaultAdminPassword),
		asp.NewPasswordValue(cmd, "api-password", "Please provide the password that will be set for StackState's admin api (-a)", &cfg.APIPassword),
		asp.NewStringChoiceP(cmd, "profile", "Please select a profile to apply", &cfg.Profile, values.ListProfiles),
		asp.NewOptionalStringValue(cmd, "override-sts-image-tag", "Override the StackState docker image tag that is deployed (press enter to skip)", &cfg.OverrideStsImageTag),
		asp.NewStringChoice(cmd, "pull-policy", "Please select a PullPolicy for pulling the StackState Docker images", &cfg.PullPolicy, []string{"default", "Always", "IfNotPresent", "Never"}),
		asp.NewStringChoice(cmd, "ingress", "Do you want to expose your installation using an Ingress (KOPS or K3D) (-i)?", &cfg.Ingress, values.IngressTypes()),
		asp.NewStringValue(cmd, "base-url", "Please provide the base URL for StackState, for example https://my-stackstate.sandbox-main.sandbox.stackstate.io (-b)", &cfg.BaseURL),
		asp.NewOptionalStringValue(cmd, "extra-api-key", "For usage with traffic mirror data provide its API key (press enter to skip) (-e)", &cfg.ExtraAPIKey),
		asp.NewBoolValue(cmd, "install-agent", "Do you want to install the StackState Kubernetes Agent on the cluster?", &cfg.InstallAgent),
	})
	if err != nil {
		return err
	}

	if cfg.InstallAgent || (cfg.Ingress == values.KOPS) {
		err = asp.AskQuestion(asp.NewStringChoiceDefP(cmd, "cluster-name", "Which cluster do you want to install StackState on?", &cfg.K8sClusterName, kubernetes.ListClusters, currentCluster))
		if err != nil {
			return err
		}
	}

	return nil
}
