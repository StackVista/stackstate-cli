package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/context"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ContextCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "manage contexts",
		Long:  "Manage connections to different StackState servers.",
	}

	cmd.AddCommand(context.SaveCommand(cli))
	cmd.AddCommand(context.ListCommand(cli))
	cmd.AddCommand(context.ShowCommand(cli))
	// cmd.AddCommand(context.Command(cli))
	// cmd.AddCommand(context.Command(cli))
	// cmd.AddCommand(context.Command(cli))

	return cmd
}
