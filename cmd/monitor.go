package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/monitor"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
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
	cmd.AddCommand(monitor.MonitorEnableDisableCommand(cli, "disable", "Disable a monitor", sts.MONITORSTATUSVALUE_DISABLED.Ptr(), "disabled"))
	cmd.AddCommand(monitor.MonitorEnableDisableCommand(cli, "enable", "Enable a monitor", sts.MONITORSTATUSVALUE_ENABLED.Ptr(), "enabled"))

	return cmd
}
