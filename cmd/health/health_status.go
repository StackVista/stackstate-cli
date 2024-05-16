package health

import (
	"fmt"
	"net/http"

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
		topologies, resp, err = api.HealthSynchronizationAPI.GetHealthSynchronizationStreamTopologyMatches(cli.Context, args.Urn).Execute()
	} else {
		topologies, resp, err = api.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamTopologyMatches(cli.Context, args.Urn, args.SubStream).Execute()
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
	status, resp, err := api.HealthSynchronizationAPI.GetHealthSynchronizationStreamStatus(cli.Context, args.Urn).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		json := map[string]interface{}{
			"consistency-model": status.ConsistencyModel,
			"metrics":           streamMetricsToJson(status.AggregateMetrics),
			"errors":            streamErrorsToJson(status.GlobalErrors),
		}
		if status.HasMainStreamStatus() {
			json["main-stream"] = subStreamStatusToJson(*status.MainStreamStatus)
		}
		cli.Printer.PrintJson(json)
	} else {
		cli.Printer.PrintLn(fmt.Sprintf("Stream consistency model: %s", status.ConsistencyModel))
		cli.Printer.PrintLn("\nAggregate metrics for the stream and all substreams:")
		cli.Printer.Table(streamMetricsToTable(status.AggregateMetrics))
		cli.Printer.PrintLn("\nErrors for non-existing substreams:")
		cli.Printer.Table(streamErrorsToTable(status.GlobalErrors))
		if status.HasMainStreamStatus() {
			cli.Printer.PrintLn("\nMain stream status:")
			printSubStreamStatus(cli, *status.MainStreamStatus)
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
		"name":                               name,
		fmt.Sprintf("now-%d", size):          metricValueOrDash(bucket, 0),
		fmt.Sprintf("%d-%d", size, 2*size):   metricValueOrDash(bucket, 1), //nolint:gomnd
		fmt.Sprintf("%d-%d", 2*size, 3*size): metricValueOrDash(bucket, 2), //nolint:gomnd
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
		metricValueOrDash(bucket, 2), //nolint:gomnd
	}
}

func streamMetricsToTable(metrics stackstate_api.HealthStreamMetrics) printer.TableData {
	size := metrics.BucketSizeSeconds
	return printer.TableData{
		Header: []string{"Metric", fmt.Sprintf("%ds ago", size), fmt.Sprintf("%d-%ds ago", size, 2*size), fmt.Sprintf("%d-%ds ago", 2*size, 3*size)}, //nolint:gomnd
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

func printSubStreamStatus(cli *di.Deps, status stackstate_api.HealthSubStreamStatus) {
	cli.Printer.PrintLn(fmt.Sprintf("Synchronized check state count: %d", status.CheckStateCount))
	cli.Printer.PrintLn("\nConsistency state:")
	cli.Printer.Table(streamConsistencyStateToTable(status.SubStreamState))
	cli.Printer.PrintLn("\nMetrics:")
	cli.Printer.Table(streamMetricsToTable(status.Metrics))
	cli.Printer.PrintLn("\nErrors:")
	cli.Printer.Table(streamErrorsToTable(status.Errors))
}

func subStreamStatusToJson(status stackstate_api.HealthSubStreamStatus) map[string]interface{} {
	return map[string]interface{}{
		"check-state-count": status.CheckStateCount,
		"consistency-state": streamConsistencyStateToJson(status.SubStreamState),
		"metrics":           streamMetricsToJson(status.Metrics),
		"errors":            streamErrorsToJson(status.Errors),
	}
}

func RunSubStreamStatus(cli *di.Deps, api *stackstate_api.APIClient, args *StatusArgs) common.CLIError {
	status, resp, err := api.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamStatus(cli.Context, args.Urn, args.SubStream).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(subStreamStatusToJson(*status))
	} else {
		printSubStreamStatus(cli, *status)
	}

	return nil
}
