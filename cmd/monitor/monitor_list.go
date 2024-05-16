package monitor

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func MonitorListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all monitors",
		Long:  "List all monitors.",
		RunE:  cli.CmdRunEWithApi(RunMonitorListCommand),
	}
	return cmd
}

func RunMonitorListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	monitors, resp, err := api.MonitorAPI.GetAllMonitors(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(monitors.Monitors, func(i, j int) bool {
		return *monitors.Monitors[i].Identifier < *monitors.Monitors[j].Identifier
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"monitors": monitors.Monitors,
		})
	} else {
		tableData := [][]interface{}{}
		for _, monitor := range monitors.Monitors {
			tableData = append(tableData, []interface{}{monitor.Id, monitor.Status, *monitor.Identifier, monitor.Name, monitor.FunctionId, monitor.Tags})
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Id", "Status", "Identifier", "Name", "Function Id", "Tags"},
			Data:                tableData,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "monitors"},
		})
	}

	return nil
}
