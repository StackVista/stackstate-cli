package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/dev"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func DevCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Development tools for quickly patching development docker images into a running StackState deployment",
	}

	cmd.SetOut(os.Stdout)

	cmd.PersistentFlags().StringP(dev.RegistryFlag, "r", "", "Docker registry to push the locally build docker image to")

	cmd.PersistentFlags().StringP(dev.TagFlag, "t", "", "New tag for the service(s) (required)")
	cmd.PersistentFlags().StringP(dev.NamespaceFlag, "n", "", "Kubernetes namespace, if not specified uses the default namespace for the current context")
	cmd.PersistentFlags().BoolP(dev.NoPushFlag, "p", false, "Skip tagging and pushing of new docker image, requires image to be already available")

	err := cmd.MarkPersistentFlagRequired(dev.TagFlag)
	if err != nil {
		cmd.Printf("Flag '%s' not found.", dev.TagFlag)
	}

	err = cmd.MarkPersistentFlagRequired(dev.RegistryFlag)
	if err != nil {
		cmd.Printf("Flag '%s' not found.", dev.RegistryFlag)
	}

	cmd.AddCommand(dev.PatchDeploymentImageCommand(&cfg.Dev))
	cmd.AddCommand(dev.PatchGroupImageCommand(&cfg.Dev))

	return cmd
}
