package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/stackpack"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func StackPackCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stackpack",
		Short: "Manage stackpacks",
		Long:  "Manage and upload StackPacks.",
	}
	cmd.AddCommand(stackpack.StackpackUploadCommand(cli))
	cmd.AddCommand(stackpack.StackpackListCommand(cli))
	cmd.AddCommand(stackpack.StackpackListInstanceCommand(cli))
	cmd.AddCommand(stackpack.StackpackInstallCommand(cli))
	cmd.AddCommand(stackpack.StackpackListParameterCommand(cli))
	cmd.AddCommand(stackpack.StackpackUninstallCommand(cli))
	cmd.AddCommand(stackpack.StackpackUpgradeCommand(cli))
	cmd.AddCommand(stackpack.StackpackConfirmManualStepsCommand(cli))

	return cmd
}
