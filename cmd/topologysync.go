package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/topologysync"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func TopologySyncCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topology-sync",
		Short: "Monitor topology synchronization streams",
		Long:  "Monitor and manage topology synchronization streams. View sync status, metrics, and clear errors from data sources.",
	}

	cmd.AddCommand(topologysync.ListCommand(deps))
	cmd.AddCommand(topologysync.ClearErrorsCommand(deps))
	cmd.AddCommand(topologysync.DescribeCommand(deps))

	return cmd
}
