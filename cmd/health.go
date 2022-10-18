package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/health"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func HealthCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Synchronize health",
		Long:  "Health synchronization related commands.",
	}
	cmd.AddCommand(health.HealthListCommand(cli))
	cmd.AddCommand(health.HealthDeleteCommand(cli))
	cmd.AddCommand(health.HealthClearErrorCommand(cli))
	cmd.AddCommand(health.HealthStatusCommand(cli))

	return cmd
}
