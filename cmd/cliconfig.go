package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/cliconfig"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func CliConfigCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli-config",
		Short: "manage the configuration of the CLI",
		Long:  "Manage the configuration of the CLI.",
	}
	cmd.AddCommand(cliconfig.DescribeCommand(deps))
	cmd.AddCommand(cliconfig.SaveCommand(deps))
	cmd.AddCommand(cliconfig.ValidateCommand(deps))

	return cmd
}
