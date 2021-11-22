package cmd

import (
	"gitlab.com/stackvista/stackstate-cli2/cmd/sandbox"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"

	"github.com/spf13/cobra"
)

func SandboxCommand(cfg *config.Config) *cobra.Command {
	namespaceCmd := &cobra.Command{
		Use:   "sandbox",
		Short: "CRUD you personal Kubernetes Sandbox",
	}

	namespaceCmd.PersistentFlags().StringP("username", "u", "", "Username that owns the sandbox")

	namespaceCmd.AddCommand(sandbox.CreateCommand(cfg))
	namespaceCmd.AddCommand(sandbox.DeleteCommand(cfg))
	namespaceCmd.AddCommand(sandbox.ListCommand(cfg))
	namespaceCmd.AddCommand(sandbox.SwitchCommand(cfg))
	return namespaceCmd
}
