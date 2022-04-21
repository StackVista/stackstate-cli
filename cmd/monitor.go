package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/monitor"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func MonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "manage and develop monitors",
	}
	cmd.AddCommand(monitor.ListMonitorsCommand(cli))
	cmd.AddCommand(monitor.MonitorApplyCommand(cli))
	cmd.AddCommand(monitor.DeleteMonitorCommand(cli))
	cmd.AddCommand(monitor.RunMonitorCommand(cli))
	cmd.AddCommand(monitor.MonitorStatusCommand(cli))

	return cmd
}
