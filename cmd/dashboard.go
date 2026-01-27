package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/dashboard"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func DashboardCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Manage SUSE Observability dashboards",
		Long:  "Manage SUSE Observability dashboards. Dashboards provide customizable views for visualizing metrics, topology, and other observability data.",
	}
	cmd.AddCommand(dashboard.DashboardListCommand(cli))
	cmd.AddCommand(dashboard.DashboardDescribeCommand(cli))
	cmd.AddCommand(dashboard.DashboardCloneCommand(cli))
	cmd.AddCommand(dashboard.DashboardDeleteCommand(cli))
	cmd.AddCommand(dashboard.DashboardApplyCommand(cli))
	cmd.AddCommand(dashboard.DashboardEditCommand(cli))

	return cmd
}
