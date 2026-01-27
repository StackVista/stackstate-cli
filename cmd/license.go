package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/license"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func LicenseCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license",
		Short: "View and update the SUSE Observability license",
		Long:  "View and update the SUSE Observability license. The license controls feature availability and usage limits.",
	}

	cmd.AddCommand(license.ShowCommand(deps))
	cmd.AddCommand(license.UpdateCommand(deps))
	return cmd
}
