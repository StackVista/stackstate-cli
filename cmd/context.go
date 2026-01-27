package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/context"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ContextCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "Manage CLI connection contexts for SUSE Observability servers",
		Long:  `Manage CLI connection contexts. A context stores the URL and authentication credentials for a SUSE Observability server, allowing you to switch between different servers or environments without re-entering connection details.`,
	}

	cmd.AddCommand(context.SaveCommand(cli))
	cmd.AddCommand(context.ListCommand(cli))
	cmd.AddCommand(context.ShowCommand(cli))
	cmd.AddCommand(context.SetCommand(cli))
	cmd.AddCommand(context.DeleteCommand(cli))
	cmd.AddCommand(context.ValidateCommand(cli))

	return cmd
}
