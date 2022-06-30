package health

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func HealthListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all streams",
		Long:  "List all health streams.",
		RunE:  cli.CmdRunEWithApi(RunHealthListCommand()),
	}

	return cmd
}

func RunHealthListCommand() di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		streamList, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationStreamsOverview(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		data := make([][]interface{}, 0)
		for _, v := range streamList.GetItems() {
			data = append(data, []interface{}{
				v.GetUrn(),
				v.GetConsistencyModel(),
				v.GetSubStreams(),
			})
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"streams": data,
			})
		} else {
			cli.Printer.Table(printer.TableData{
				Header:              []string{"Stream urn", "Stream consistency model", "Sub stream count"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "health stream"},
			})
		}
		return nil
	}
}
