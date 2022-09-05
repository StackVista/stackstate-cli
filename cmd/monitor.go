package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/monitor"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
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

	return cmd
}
