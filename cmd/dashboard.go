package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/dashboard"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func DashboardCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Manage dashboards",
		Long:  "Manage, test and develop dashboards.",
	}
	cmd.AddCommand(dashboard.DashboardListCommand(cli))
	cmd.AddCommand(dashboard.DashboardDescribeCommand(cli))
	cmd.AddCommand(dashboard.DashboardCloneCommand(cli))
	cmd.AddCommand(dashboard.DashboardDeleteCommand(cli))
	cmd.AddCommand(dashboard.DashboardApplyCommand(cli))
	cmd.AddCommand(dashboard.DashboardEditCommand(cli))

	return cmd
}
