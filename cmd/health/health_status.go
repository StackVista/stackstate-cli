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

const (
	// Magic numbers are frowned upon, so I'm introducing these two smelly ding-dongs because I don't give a fuck anymore.
	TWO   = 2
	THREE = 3
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
			return RunStreamStatus(cli, api, args)
		default:
			return RunSubStreamStatus(cli, api, args)
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

func RunStreamStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	streamStatus, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationStreamStatus(cli.Context, args.Urn).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(healthStreamStatusToJson(streamStatus))
	} else {
		for _, v := range healthStreamStatusToString(streamStatus) {
			cli.Printer.PrintLn(v)
		}
	}
	return nil
}

func metricValueOrDash(bucket []stackstate_api.MetricBucketValue, index int) interface{} {
	if index < len(bucket) && bucket[index].HasValue() {
		return bucket[index].GetValue()
	}
	return "-"
}

func metricBucketToJson(name string, bucket []stackstate_api.MetricBucketValue, size int32) map[string]interface{} {
	return map[string]interface{}{
		"name":                                     name,
		fmt.Sprintf("now-%d", size):                metricValueOrDash(bucket, 0),
		fmt.Sprintf("%d-%d", size, TWO*size):       metricValueOrDash(bucket, 1),
		fmt.Sprintf("%d-%d", TWO*size, THREE*size): metricValueOrDash(bucket, TWO),
	}
}

func streamMetricsToJson(metrics stackstate_api.HealthStreamMetrics) []map[string]interface{} {
	size := metrics.BucketSizeSeconds
	return []map[string]interface{}{
		metricBucketToJson("latency seconds", metrics.LatencySeconds, size),
		metricBucketToJson("messages per seconds", metrics.MessagePerSecond, size),
		metricBucketToJson("creates per seconds", metrics.CreatesPerSecond, size),
		metricBucketToJson("updates per seconds", metrics.UpdatesPerSecond, size),
		metricBucketToJson("deletes per seconds", metrics.DeletesPerSecond, size),
	}
}

func metricBucketToRow(name string, bucket []stackstate_api.MetricBucketValue) []interface{} {
	return []interface{}{
		name,
		metricValueOrDash(bucket, 0),
		metricValueOrDash(bucket, 1),
		metricValueOrDash(bucket, TWO),
	}
}

func streamMetricsToTable(metrics stackstate_api.HealthStreamMetrics) printer.TableData {
	size := metrics.BucketSizeSeconds
	return printer.TableData{
		Header: []string{"Metric", fmt.Sprintf("%ds ago", size), fmt.Sprintf("%d-%ds ago", size, TWO*size), fmt.Sprintf("%d-%ds ago", TWO*size, THREE*size)},
		Data: [][]interface{}{
			metricBucketToRow("latency seconds", metrics.LatencySeconds),
			metricBucketToRow("messages per seconds", metrics.MessagePerSecond),
			metricBucketToRow("creates per seconds", metrics.CreatesPerSecond),
			metricBucketToRow("updates per seconds", metrics.UpdatesPerSecond),
			metricBucketToRow("deletes per seconds", metrics.DeletesPerSecond),
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "metrics"},
	}
}

func streamErrorsToJson(errors []stackstate_api.HealthStreamError) []map[string]interface{} {
	data := make([]map[string]interface{}, len(errors))

	for i, error := range errors {
		data[i] = map[string]interface{}{
			"message": error.Error,
			"count":   error.Count,
		}
	}
	return data
}

func streamErrorsToTable(errors []stackstate_api.HealthStreamError) printer.TableData {
	data := make([][]interface{}, len(errors))

	for i, error := range errors {
		data[i] = []interface{}{error.Error, error.Count}
	}

	return printer.TableData{
		Header:              []string{"Message", "Count"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "errors"},
	}
}

func streamConsistencyStateToJson(state stackstate_api.HealthSubStreamConsistencyState) map[string]interface{} {
	data := make(map[string]interface{}, 0)

	if expiry := state.HealthSubStreamExpiry; expiry != nil {
		data["repeat-interval"] = expiry.RepeatIntervalMs
		data["expiry-interval"] = expiry.ExpiryIntervalMs
	}

	if snapshot := state.HealthSubStreamSnapshot; snapshot != nil {
		data["snapshot-repeat-interval"] = snapshot.RepeatIntervalMs
		if snapshot.ExpiryIntervalMs != nil {
			data["snapshot-expiry-interval"] = *snapshot.ExpiryIntervalMs
		}
	}

	if increments := state.HealthSubStreamTransactionalIncrements; increments != nil {
		data["checkpoint-offset"] = increments.GetOffset()
		if increments.BatchIndex != nil {
			data["checkpoint-batch-index"] = increments.BatchIndex
		}
	}
	return data
}

func streamConsistencyStateToTable(state stackstate_api.HealthSubStreamConsistencyState) printer.TableData {
	repeatInterval := "-"
	expiryInterval := "-"
	if expiry := state.HealthSubStreamExpiry; expiry != nil {
		repeatInterval = fmt.Sprintf("%d ms", expiry.RepeatIntervalMs)
		expiryInterval = fmt.Sprintf("%d ms", expiry.ExpiryIntervalMs)
	}

	snapshotRepeatInterval := "-"
	snapshotExpiryInterval := "-"
	if snapshot := state.HealthSubStreamSnapshot; snapshot != nil {
		snapshotRepeatInterval = fmt.Sprintf("%d ms", snapshot.RepeatIntervalMs)
		if snapshot.ExpiryIntervalMs != nil {
			snapshotExpiryInterval = fmt.Sprintf("%d ms", *snapshot.ExpiryIntervalMs)
		}
	}

	checkpointOffset := "-"
	checkpointBatchIndex := "-"
	if increments := state.HealthSubStreamTransactionalIncrements; increments != nil {
		checkpointOffset = fmt.Sprintf("%d", increments.GetOffset())
		if increments.BatchIndex != nil {
			checkpointBatchIndex = fmt.Sprintf("%d", increments.BatchIndex)
		}
	}

	return printer.TableData{
		Header: []string{"Repeat Interval", "Expiry Interval", "Snapshot Repeat Interval", "Snapshot Expiry Interval", "Checkpoint Offset", "Checkpoint Batch Index"},
		Data: [][]interface{}{
			{repeatInterval, expiryInterval, snapshotRepeatInterval, snapshotExpiryInterval, checkpointOffset, checkpointBatchIndex},
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "errors"},
	}
}

func RunSubStreamStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	status, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus(cli.Context, args.Urn, args.SubStream).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"check-state-count": status.CheckStateCount,
			"consistency-state": streamConsistencyStateToJson(status.SubStreamState),
			"metrics":           streamMetricsToJson(status.Metrics),
			"errors":            streamErrorsToJson(status.Errors),
		})
	} else {
		cli.Printer.PrintLn(fmt.Sprintf("Synchronized check state count: %d", status.CheckStateCount))
		cli.Printer.PrintLn("Consistency state:")
		cli.Printer.Table(streamConsistencyStateToTable(status.SubStreamState))
		cli.Printer.PrintLn("Aggregate metrics:")
		cli.Printer.Table(streamMetricsToTable(status.Metrics))
		cli.Printer.PrintLn("Errors:")
		cli.Printer.Table(streamErrorsToTable(status.Errors))
	}
	return nil
}

func healthStreamStatusToString(streamStatus *stackstate_api.HealthStreamStatus) []string {
	strValue := make([]string, 0)
	if streamStatus.GetRecoverMessage() != "" {
		str := fmt.Sprintf("This stream is in recovery mode.\n"+
			"This means stackstate is reconstructing the state of the health streams. "+
			"In this period no errors will be reported for the stream,\n"+
			"incoming data will be processed as usual.\n"+
			"The reason recovery mode was entered was because: %s", streamStatus.GetRecoverMessage())
		strValue = append(strValue, str)
	}
	str := fmt.Sprintf("Consistency model for the stream and all substreams: %s", streamStatus.GetConsistencyModel())
	strValue = append(strValue, str)
	if mainStream, ok := streamStatus.GetMainStreamStatusOk(); ok {
		strValue = append(strValue, fmt.Sprintf("Synchronized check state count: %d", mainStream.CheckStateCount))
		strOutput := consistencyStateToString(&mainStream.SubStreamState)
		strValue = append(strValue, strOutput...)
		strValue = append(strValue, fmt.Sprintf("\nSynchronization errors:\n %v", mainStream.GetErrors()),
			fmt.Sprintf("\nSynchronization metrics:\n %v", mainStream.GetMetrics()))
	} else {
		strValue = append(strValue, "Synchronized check state count: -")
	}
	strValue = append(strValue, fmt.Sprintf("\nAggregate metrics for the stream and all substreams:\n %v", streamStatus.GetAggregateMetrics()),
		fmt.Sprintf("\nErrors for non-existing sub streams:\n %v", streamStatus.GetGlobalErrors()))
	return strValue
}

func healthStreamStatusToJson(streamStatus *stackstate_api.HealthStreamStatus) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	if streamStatus.GetRecoverMessage() != "" {
		str := fmt.Sprintf("This stream is in recovery mode.\n"+
			"This means stackstate is reconstructing the state of the health streams. "+
			"In this period no errors will be reported for the stream,\n"+
			"incoming data will be processed as usual.\n"+
			"The reason recovery mode was entered was because: %s", streamStatus.GetRecoverMessage())
		jsonMap["recovery_message"] = str
	}
	str := fmt.Sprintf("Consistency model for the stream and all substreams: %s", streamStatus.GetConsistencyModel())
	jsonMap["stream_consistency_model"] = str
	if mainStream, ok := streamStatus.GetMainStreamStatusOk(); ok {
		jsonMap["state_check_count"] = fmt.Sprintf("%d", mainStream.CheckStateCount)
		jsonOutput := consistencyStateToJson(&mainStream.SubStreamState)
		util.ConcatMap(jsonMap, jsonOutput)
		jsonMap["errors"] = fmt.Sprintf("Synchronization errors: %v", mainStream.GetErrors())
		jsonMap["metrics"] = fmt.Sprintf("Synchronization metrics: %v", mainStream.GetMetrics())
	} else {
		jsonMap["synchronized_state_count"] = "-"
	}
	jsonMap["aggregate_metrics"] = fmt.Sprintf("Aggregate metrics for the stream and all substreams: %v", streamStatus.GetAggregateMetrics())
	jsonMap["non_existing_error"] = fmt.Sprintf("Errors for non-existing sub streams: %v", streamStatus.GetGlobalErrors())

	return jsonMap
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
