package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/cli"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func CliCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "command to control this CLI",
	}
	cmd.AddCommand(cli.CliSaveConfigCommand(deps))
	cmd.AddCommand(cli.CliTestCommandCommand(deps))
	return cmd
}
