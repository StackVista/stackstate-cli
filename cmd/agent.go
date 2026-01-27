package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/agent"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func AgentCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Manage SUSE Observability agents and their registrations",
		Long:  `Manage SUSE Observability agents. Agents collect and send topology, telemetry, and traces data from monitored systems to SUSE Observability.`,
	}

	cmd.AddCommand(agent.ListCommand(deps))
	return cmd
}
