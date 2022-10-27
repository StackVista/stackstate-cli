package health

import (
	"fmt"
	"net/http"

	"github.com/stackvista/stackstate-cli/internal/util"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type StatusArgs struct {
	Urn       string
	SubStream string
	Topology  bool
}

func HealthStatusCommand(cli *di.Deps) *cobra.Command {
	args := &StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Status of the stream",
		Long:  "Status of an active health synchronization stream.",
		RunE:  cli.CmdRunEWithApi(RunHealthStatusCommand(args)),
	}
	common.AddRequiredUrnFlagVar(cmd, &args.Urn, "Urn of the stream")
	cmd.Flags().StringVar(&args.SubStream, SubStreamFlag, "", SubStreamFlagUsage)
	cmd.Flags().BoolVar(&args.Topology, TopologyFlag, false, TopologyFlagUsage)

	return cmd
}

func RunHealthStatusCommand(args *StatusArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		switch {
		case args.Topology:
			return RunTopologyMatchesStatus(cli, api, args)
		case args.SubStream == "":
			return RunSubStreamStatus(cli, api, args)
		default:
			return RunStreamStatus(cli, api, args)
		}
	}
}

func GetTopologyMatches(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) (*stackstate_api.TopologyMatchResult, common.CLIError) {
	var topologies *stackstate_api.TopologyMatchResult
	var resp *http.Response
	var err error
	if args.SubStream == "" {
		topologies, resp, err = api.HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatches(cli.Context, args.Urn).Execute()
	} else {
		topologies, resp, err = api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamTopologyMatches(cli.Context, args.Urn, args.SubStream).Execute()
	}

	if err != nil {
		return topologies, common.NewResponseError(err, resp)
	}
	return topologies, nil
}

func RunTopologyMatchesStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	topologies, err := GetTopologyMatches(cli, api, args)
	if err != nil {
		return err
	}

	data := make([][]interface{}, 0)
	for _, v := range topologies.GetMultipleMatchesCheckStates() {
		data = append(data, []interface{}{v.GetCheckStateId(), v.GetTopologyElementIdentifier(), v.GetMatchCount()})
	}
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{"status": data})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Check state id", "Topology element identifier", "Number of matched topology elements"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health status"},
		})
	}
	return nil
}

func RunSubStreamStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	streamStatus, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationStreamStatus(cli.Context, args.Urn).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	jsonMap, strValue := processSubStream(streamStatus)

	if cli.IsJson() {
		cli.Printer.PrintJson(jsonMap)
	} else {
		for _, v := range strValue {
			cli.Printer.PrintLn(v)
		}
	}
	return nil
}

func RunStreamStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	subSteamStatus, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus(cli.Context, args.Urn, args.SubStream).Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		jsonMap := make(map[string]interface{})
		if count, ok := subSteamStatus.GetCheckStateCountOk(); ok {
			jsonMap["state_check_count"] = fmt.Sprintf("%d", *count)
			jsonOutput := consistencyStateToJson(&subSteamStatus.SubStreamState)
			util.ConcatMap(jsonMap, jsonOutput)
		} else {
			jsonMap["synchronized_state_count"] = "-"
		}
		cli.Printer.PrintJson(jsonMap)
	} else {
		strValue := make([]string, 0)
		if count, ok := subSteamStatus.GetCheckStateCountOk(); ok {
			strValue = append(strValue, fmt.Sprintf("Synchronized check state count: %d", *count))
			strOutput := consistencyStateToString(&subSteamStatus.SubStreamState)
			strValue = append(strValue, strOutput...)
		} else {
			strValue = append(strValue, "Synchronized check state count: -")
		}
		for _, v := range strValue {
			cli.Printer.PrintLn(v)
		}
	}
	return nil
}

func processSubStream(streamStatus *stackstate_api.HealthStreamStatus) (map[string]interface{}, []string) {
	jsonMap := make(map[string]interface{})
	strValue := make([]string, 0)
	if streamStatus.GetRecoverMessage() != "" {
		str := fmt.Sprintf("This stream is in recovery mode.\n"+
			"This means stackstate is reconstructing the state of the health streams. "+
			"In this period no errors will be reported for the stream,\n"+
			"incoming data will be processed as usual.\n"+
			"The reason recovery mode was entered was because: %s", streamStatus.GetRecoverMessage())
		jsonMap["recovery_message"] = str
		strValue = append(strValue, str)
	}
	str := fmt.Sprintf("Consistency model for the stream and all substreams: %s", streamStatus.GetConsistencyModel())
	jsonMap["stream_consistency_model"] = str
	strValue = append(strValue, str)
	if mainStream, ok := streamStatus.GetMainStreamStatusOk(); ok {
		jsonMap["state_check_count"] = fmt.Sprintf("%d", mainStream.CheckStateCount)
		strValue = append(strValue, fmt.Sprintf("Synchronized check state count: %d", mainStream.CheckStateCount))

		strOutput := consistencyStateToString(&mainStream.SubStreamState)
		jsonOutput := consistencyStateToJson(&mainStream.SubStreamState)

		strValue = append(strValue, strOutput...)
		util.ConcatMap(jsonMap, jsonOutput)

		jsonMap["errors"] = fmt.Sprintf("Synchronization errors: %v", mainStream.GetErrors())
		jsonMap["metrics"] = fmt.Sprintf("Synchronization metrics: %v", mainStream.GetMetrics())
		strValue = append(strValue, fmt.Sprintf("\nSynchronization errors:\n %v", mainStream.GetErrors()),
			fmt.Sprintf("\nSynchronization metrics:\n %v", mainStream.GetMetrics()))
	} else {
		jsonMap["synchronized_state_count"] = "-"
		strValue = append(strValue, "Synchronized check state count: -")
	}
	jsonMap["aggregate_metrics"] = fmt.Sprintf("Aggregate metrics for the stream and all substreams: %v", streamStatus.GetAggregateMetrics())
	jsonMap["non_existing_error"] = fmt.Sprintf("Errors for non-existing sub streams: %v", streamStatus.GetGlobalErrors())
	strValue = append(strValue, fmt.Sprintf("\nAggregate metrics for the stream and all substreams:\n %v", streamStatus.GetAggregateMetrics()),
		fmt.Sprintf("\nErrors for non-existing sub streams:\n %v", streamStatus.GetGlobalErrors()))
	return jsonMap, strValue
}

func consistencyStateToJson(subStream *stackstate_api.HealthSubStreamConsistencyState) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	if subStream == nil {
		jsonMap["sub_stream_state"] = "Invalid Sub Stream State"
		return jsonMap
	}
	const divideBy = int32(1000)
	if subStreamStateType := subStream.HealthSubStreamExpiry; subStreamStateType != nil {
		jsonMap["repeat_interval"] = fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(subStreamStateType.RepeatIntervalMs/divideBy))
		jsonMap["expiry"] = fmt.Sprintf("Expiry (Seconds): %.0f", float64(subStreamStateType.ExpiryIntervalMs/divideBy))
	} else if statusType := subStream.HealthSubStreamSnapshot; statusType != nil {
		jsonMap["repeat_interval"] = fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(statusType.RepeatIntervalMs/divideBy))
		if expSecond := statusType.ExpiryIntervalMs; expSecond != nil {
			jsonMap["expiry"] = fmt.Sprintf("Expiry (Seconds): %.0f", float64(*expSecond/divideBy))
		}
	} else if statusType := subStream.HealthSubStreamTransactionalIncrements; statusType != nil {
		jsonMap["checkpoint_offset"] = fmt.Sprintf("Checkpoint offset: %d", statusType.GetOffset())
		if idx := statusType.BatchIndex; idx != nil {
			jsonMap["checkpoint_batch_index"] = fmt.Sprintf("Checkpoint batch index: %d", idx)
		}
	}
	return jsonMap
}

func consistencyStateToString(subStream *stackstate_api.HealthSubStreamConsistencyState) []string {
	strValue := make([]string, 0)
	if subStream == nil {
		strValue = append(strValue, "Invalid Sub Stream State")
		return strValue
	}
	const divideBy = int32(1000)
	if subStreamStateType := subStream.HealthSubStreamExpiry; subStreamStateType != nil {
		strValue = append(strValue,
			fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(subStreamStateType.RepeatIntervalMs/divideBy)),
			fmt.Sprintf("Expiry (Seconds): %.0f", float64(subStreamStateType.ExpiryIntervalMs/divideBy)),
		)
	} else if statusType := subStream.HealthSubStreamSnapshot; statusType != nil {
		strValue = append(strValue, fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(statusType.RepeatIntervalMs/divideBy)))
		if expSecond := statusType.ExpiryIntervalMs; expSecond != nil {
			strValue = append(strValue, fmt.Sprintf("Expiry (Seconds): %.0f", float64(*expSecond/divideBy)))
		}
	} else if statusType := subStream.HealthSubStreamTransactionalIncrements; statusType != nil {
		strValue = append(strValue, fmt.Sprintf("Checkpoint offset: %d", statusType.GetOffset()))
		if idx := statusType.BatchIndex; idx != nil {
			strValue = append(strValue, fmt.Sprintf("Checkpoint batch index: %d", idx))
		}
	}
	return strValue
}
