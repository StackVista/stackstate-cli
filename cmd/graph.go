package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/graph"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func GraphCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "Manage the SUSE Observability topology graph data",
		Long:  "Manage the SUSE Observability topology graph, including data retention settings and cleanup operations.",
	}

	cmd.AddCommand(graph.RetentionCommand(deps))
	cmd.AddCommand(graph.DeleteExpiredDataCommand(deps))
	return cmd
}
