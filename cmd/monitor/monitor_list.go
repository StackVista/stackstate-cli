package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func MonitorListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all monitors",
		Long:  "List all monitors.",
		RunE:  cli.CmdRunEWithApi(RunMonitorListCommand),
	}
	return cmd
}

func RunMonitorListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo stackstate_api.ServerInfo,
) common.CLIError {
	monitors, resp, err := api.MonitorApi.GetAllMonitors(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"monitors": monitors.Monitors,
		})
	} else {
		tableData := [][]interface{}{}
		for _, monitor := range monitors.Monitors {
			tableData = append(tableData, []interface{}{monitor.Id, *monitor.Identifier, monitor.Name})
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Id", "Identifier", "Name"},
			Data:                tableData,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "monitors"},
		})
	}

	return nil
}
