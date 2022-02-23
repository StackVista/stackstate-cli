package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/monitor"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func MonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Manage and develop monitors.",
	}
	cmd.AddCommand(monitor.ListMonitorsCommand(cli))
	cmd.AddCommand(monitor.CreateMonitorCommand(cli))

	return cmd
}
