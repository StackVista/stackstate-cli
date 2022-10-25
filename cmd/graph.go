package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/graph"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func GraphCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "Manage the StackState Graph",
		Long:  "Manage the StackState Graph.",
	}

	cmd.AddCommand(graph.RetentionCommand(deps))
	cmd.AddCommand(graph.DeleteExpiredDataCommand(deps))
	return cmd
}
