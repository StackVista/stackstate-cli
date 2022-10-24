package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/topologysync"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func TopologySyncCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topology-sync",
		Short: "Manage the StackState Topology Synchronization",
		Long:  "Manage the StackState Topology Synchronization.",
	}

	cmd.AddCommand(topologysync.ListCommand(deps))
	return cmd
}
