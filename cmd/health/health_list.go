package health

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type ListArgs struct {
	Urn string
}

func HealthListCommand(cli *di.Deps) *cobra.Command {
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all streams",
		Long:  "List all health streams.",
		RunE:  cli.CmdRunEWithApi(RunHealthListCommand(args)),
	}
	common.AddUrnFlagVar(cmd, &args.Urn, "Urn of the health synchronization stream")

	return cmd
}

func RunHealthListCommand(args *ListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if args.Urn != "" {
			subList, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverview(cli.Context, args.Urn).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}
			printSubStreamList(cli, subList)
		} else {
			streamList, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationStreamsOverview(cli.Context).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}
			printStreamList(cli, streamList)
		}
		return nil
	}
}

func printSubStreamList(cli *di.Deps, subList *stackstate_api.SubStreamList) {
	data := make([][]interface{}, 0)
	jsonMap := make([]map[string]interface{}, 0)
	for _, v := range subList.GetSubStreams() {
		data = append(data, []interface{}{v.GetSubStreamId(), v.GetCheckStateCount()})
		jsonMap = append(jsonMap, map[string]interface{}{
			"sub_stream_id":     v.GetSubStreamId(),
			"check_state_count": v.GetCheckStateCount(),
		})
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"sub-stream": jsonMap,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Sub stream id", "Check state count"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health sub-stream"},
		})
	}
}

func printStreamList(cli *di.Deps, streamList *stackstate_api.StreamList) {
	data := make([][]interface{}, 0)
	jsonMap := make([]map[string]interface{}, 0)
	for _, v := range streamList.GetItems() {
		data = append(data, []interface{}{v.GetUrn(), v.GetConsistencyModel(), v.GetSubStreams()})
		jsonMap = append(jsonMap, map[string]interface{}{
			"stream_urn":               v.GetUrn(),
			"stream_consistency_model": v.GetConsistencyModel(),
			"sub_stream_count":         v.GetSubStreams(),
		})
	}
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"streams": jsonMap,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Stream urn", "Stream consistency model", "Sub stream count"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health stream"},
		})
	}
}
