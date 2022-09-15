package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/license"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func LicenseCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license",
		Short: "Manage StackState license",
		Long:  "Manage StackState license.",
	}

	cmd.AddCommand(license.ShowCommand(deps))
	cmd.AddCommand(license.UpdateCommand(deps))
	return cmd
}
