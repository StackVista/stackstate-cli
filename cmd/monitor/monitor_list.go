package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func ListMonitorsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all monitors.",
		RunE:  di.CmdRunEWithDeps(cli, RunListMonitorsCommand),
	}
	return cmd
}

func RunListMonitorsCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	monitors, resp, err := cli.Client.MonitorApi.GetAllMonitors(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	tableData := [][]string{}
	for _, monitor := range monitors.Monitors {
		tableData = append(tableData, []string{util.ToString(monitor.Id), util.ToString(*monitor.Identifier), util.ToString(monitor.Name)})
	}
	cli.Printer.Table([]string{"Id", "Identifier", "Name"}, tableData, monitors.Monitors)

	return nil
}
