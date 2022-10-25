package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/monitor"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func MonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Manage monitors",
		Long:  "Manage, test and develop monitors.",
	}
	cmd.AddCommand(monitor.MonitorListCommand(cli))
	cmd.AddCommand(monitor.MonitorApplyCommand(cli))
	cmd.AddCommand(monitor.MonitorDeleteCommand(cli))
	cmd.AddCommand(monitor.MonitorRunCommand(cli))
	cmd.AddCommand(monitor.MonitorStatusCommand(cli))
	cmd.AddCommand(monitor.MonitorEditCommand(cli))
	cmd.AddCommand(monitor.MonitorDisableCommand(cli))
	cmd.AddCommand(monitor.MonitorEnableCommand(cli))
	cmd.AddCommand(monitor.MonitorDescribeCommand(cli))

	return cmd
}
