package monitor

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func MonitorStatusCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status -i ID",
		Short: "get the status of a monitor",
		Long:  "Get the satus of a single monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorStatusCommand),
	}

	cmd.Flags().StringP(IDFlag, "i", "", IDFlag)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck

	return cmd
}

const (
	THOUSAND = 1000
	TWO      = 2
	THREE    = 3
)

func RunMonitorStatusCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	identifier, err := cmd.Flags().GetString(IDFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
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
	cli.Printer.PrintLn("Repeat interval (Seconds): " + util.ToString(subStreamSnapshot.GetRepeatIntervalMs()/THOUSAND))
	if subStreamSnapshot.HasExpiryIntervalMs() {
		cli.Printer.PrintLn("Expiry (Seconds): " + util.ToString(subStreamSnapshot.GetExpiryIntervalMs()/THOUSAND))
	}

	if mainStreamStatus.HasErrors() {
		PrintErrors(cli, mainStreamStatus.GetErrors(), monitorStatus)
	}

	metricsData := make([][]interface{}, 0)
	bucketSizeS := mainStreamStatus.GetMetrics().BucketSizeSeconds
	bucketSizeSDouble := bucketSizeS * TWO
	bucketSizeSTriple := bucketSizeS * THREE
	metrics := mainStreamStatus.GetMetrics()
	metricsData = append(metricsData, CreateMetricRows("latency (Seconds)", metrics.GetLatencySeconds()))
	metricsData = append(metricsData, CreateMetricRows("messages processed (per second)", metrics.GetMessagePerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states created (per second)", metrics.GetCreatesPerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states updated (per second)", metrics.GetUpdatesPerSecond()))
	metricsData = append(metricsData, CreateMetricRows("check states deleted (per second)", metrics.GetDeletesPerSecond()))

	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Synchronization metrics:")
	cli.Printer.Table(
		[]string{"metric",
			"value between now and " + util.ToString(bucketSizeS) + " seconds ago",
			"value between " + util.ToString(bucketSizeS) + " and " + util.ToString(bucketSizeSDouble) + " seconds ago",
			"value between " + util.ToString(bucketSizeSDouble) + " and " + util.ToString(bucketSizeSTriple) + " seconds ago"},
		metricsData,
		monitorStatus,
	)

	topologyMatchResult := monitorStatus.TopologyMatchResult
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Check states with identifier matching exactly 1 topology element: " + util.ToString(topologyMatchResult.MatchedCheckStates))

	if len(topologyMatchResult.GetUnmatchedCheckStates()) > 0 {
		PrintTopologyMatchResultUnmatched(cli, topologyMatchResult.GetUnmatchedCheckStates(), monitorStatus)
	}

	if len(topologyMatchResult.GetMultipleMatchesCheckStates()) > 0 {
		PrintTopologyMatchResultMultipleMatched(cli, topologyMatchResult.GetMultipleMatchesCheckStates(), monitorStatus)
	}

	return nil
}

func PrintErrors(cli *di.Deps,
	errors []stackstate_client.HealthStreamError,
	monitorStatus stackstate_client.MonitorStatus,
) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Synchronization errors:")
	data := make([][]interface{}, 0)
	for _, e := range errors {
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

func PrintTopologyMatchResultUnmatched(cli *di.Deps,
	unmatchedCheckStates []stackstate_client.UnmatchedCheckState,
	monitorStatus stackstate_client.MonitorStatus) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Check states with identifier which has no matching topology element:")

	unmatchedData := make([][]interface{}, 0)
	for _, u := range unmatchedCheckStates {
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

func PrintTopologyMatchResultMultipleMatched(cli *di.Deps,
	multipleMatchesCheckState []stackstate_client.MultipleMatchesCheckState,
	monitorStatus stackstate_client.MonitorStatus) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Check states with identifier which has multiple matching topology elements:")

	multipleMatched := make([][]interface{}, 0)
	for _, m := range multipleMatchesCheckState {
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

func CreateMetricRows(metric string, metrics []stackstate_client.MetricBucketValue) []interface{} {
	metricField := make([]interface{}, 0)
	metricField = append(metricField, metric)

	metricRow := make([]interface{}, 0)
	for _, m := range metrics {
		metricRow = append(metricRow, m.Value)
	}
	metricRow = append(metricField, metricRow...)

	return metricRow
}
