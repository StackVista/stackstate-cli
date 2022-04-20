package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"net/http"
)

func MonitorStatusCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status -i ID",
		Short: "Status of a single monitor",
		RunE:  cli.CmdRunEWithApi(RunMonitorStatusCommand),
	}

	cmd.Flags().StringP(IDFlag, "i", "", IDFlag)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck

	return cmd
}

func RunMonitorStatusCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	identifier, err := cmd.Flags().GetString(IDFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	id, err := util.StringToInt64(identifier)
	var resp *http.Response
	var monitorStatus stackstate_client.MonitorStatus

	if err == nil {
		monitorStatus, resp, err = api.MonitorApi.GetMonitorWithStatus(cli.Context, id).Execute()
	} else {
		monitorStatus, resp, err = api.MonitorUrnApi.GetMonitorWithStatusByURN(cli.Context, identifier).Execute()
	}

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	mainStreamStatus := monitorStatus.Status.MainStreamStatus
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Synchronized check state count: " + util.ToString(mainStreamStatus.CheckStateCount))
	subStreamSnapshot := mainStreamStatus.SubStreamState.HealthSubStreamSnapshot
	cli.Printer.PrintLn("Repeat interval (Seconds): " + util.ToString(subStreamSnapshot.GetRepeatIntervalMs() / 1000))
	if subStreamSnapshot.HasExpiryIntervalMs() {
		cli.Printer.PrintLn("Expiry (Seconds): " + util.ToString(subStreamSnapshot.GetExpiryIntervalMs() / 1000))
	}

	if mainStreamStatus.HasErrors() {
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Synchronization errors:")
		data := make([][]interface{}, 0)
		for _, e := range mainStreamStatus.GetErrors() {
			data = append(data, []interface{}{
				e.ErrorCode,
				e.Level,
				e.Error,
				e.Count,
			})
		}

		cli.Printer.Table(
			[]string{"code", "level", "message", "occurrence count"},
			data,
			monitorStatus,
		)
	}

	metricsData := make([][]interface{}, 0)
	bucketSizeS := mainStreamStatus.GetMetrics().BucketSizeSeconds
	metrics := mainStreamStatus.GetMetrics()
	metricsData = append(metricsData, CreateMetricRows("latency (Seconds)", metrics.GetLatencySeconds()))
	metricsData = append(metricsData, CreateMetricRows("messages processed (per second)", metrics.GetMessagePerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states created (per second)", metrics.GetCreatesPerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states updated (per second)", metrics.GetUpdatesPerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states deleted (per second)", metrics.GetDeletesPerSecond()))

	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Synchronization metrics:")
	cli.Printer.Table(
		[]string{"metric", "value between now and " + util.ToString(bucketSizeS) + " seconds ago", "value between " + util.ToString(bucketSizeS) + " and " + util.ToString(bucketSizeS * 2) + " seconds ago", "value between " + util.ToString(bucketSizeS * 2)  + " and " + util.ToString(bucketSizeS * 3) + " seconds ago"},
		metricsData,
		monitorStatus,
	)

	topologyMatchResult := monitorStatus.TopologyMatchResult
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Check states with identifier matching exactly 1 topology element: " + util.ToString(topologyMatchResult.MatchedCheckStates))

	if len(topologyMatchResult.GetUnmatchedCheckStates()) > 0 {
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Check states with identifier which has no matching topology element:")

		unmatchedData := make([][]interface{}, 0)
		for _, u := range topologyMatchResult.GetUnmatchedCheckStates() {
			unmatchedData = append(unmatchedData, []interface{}{
				u.GetCheckStateId(),
				u.GetTopologyElementIdentifier(),
			})
		}

		cli.Printer.Table(
			[]string{"check state id", "topology element identifier"},
			unmatchedData,
			monitorStatus,
		)
	}


	if len(topologyMatchResult.GetMultipleMatchesCheckStates()) > 0 {
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Check states with identifier which has multiple matching topology elements:")

		multipleMatched := make([][]interface{}, 0)
		for _, m := range topologyMatchResult.GetMultipleMatchesCheckStates() {
			multipleMatched = append(multipleMatched, []interface{}{
				m.GetCheckStateId(),
				m.GetTopologyElementIdentifier(),
				m.GetMatchCount(),
			})
		}

		cli.Printer.Table(
			[]string{"check state id", "topology element identifier", "number of matched topology elements"},
			multipleMatched,
			monitorStatus,
		)
	}

	return nil
}

func CreateMetricRows(metric string, metrics []stackstate_client.MetricBucketValue) []interface{} {
	metricField:= make([]interface{}, 0)
	metricField = append(metricField, metric)

	metricRow := make([]interface{}, 0)
	for _, m := range metrics {
		metricRow = append(metricRow, m.Value)
	}
	metricRow = append(metricField, metricRow...)

	return metricRow
}