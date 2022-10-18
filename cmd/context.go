package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/context"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ContextCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "Manage StackState contexts",
		Long:  "Manage connections to different StackState servers.",
	}

	cmd.AddCommand(context.SaveCommand(cli))
	cmd.AddCommand(context.ListCommand(cli))
	cmd.AddCommand(context.ShowCommand(cli))
	cmd.AddCommand(context.SetCommand(cli))
	cmd.AddCommand(context.DeleteCommand(cli))
	cmd.AddCommand(context.ValidateCommand(cli))

	return cmd
}
