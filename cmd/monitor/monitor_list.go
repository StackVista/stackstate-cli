package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ListMonitorsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all monitors",
		Long:  "List all monitors.",
		RunE:  cli.CmdRunEWithApi(RunListMonitorsCommand),
	}
	return cmd
}

func RunListMonitorsCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo *stackstate_client.ServerInfo,
) common.CLIError {
	monitors, resp, err := api.MonitorApi.GetAllMonitors(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	tableData := [][]interface{}{}
	for _, monitor := range monitors.Monitors {
		tableData = append(tableData, []interface{}{monitor.Id, *monitor.Identifier, monitor.Name})
	}
	cli.Printer.Table(printer.TableData{
		Header:              []string{"Id", "Identifier", "Name"},
		Data:                tableData,
		StructData:          monitors.Monitors,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "monitors"},
	})

	return nil
}
