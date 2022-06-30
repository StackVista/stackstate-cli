package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/health"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func HealthCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Synchronize health",
		Long:  "Health synchronization related commands.",
	}
	cmd.AddCommand(health.HealthListCommand(cli))

	return cmd
}
