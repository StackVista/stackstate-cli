package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/agent"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func AgentCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Manage the StackState agents",
		Long:  "Manage the StackState agents.",
	}

	cmd.AddCommand(agent.ListCommand(deps))
	return cmd
}
