package monitor

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type StatusArgs struct {
	ID         int64
	Identifier string
}

func MonitorStatusCommand(cli *di.Deps) *cobra.Command {
	args := &StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "get the status of a monitor",
		Long:  "Get the satus of a single monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorStatusCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	mutex_flags.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

const (
	THOUSAND = 1000
	TWO      = 2
	THREE    = 3
)

func RunMonitorStatusCommand(args *StatusArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		var resp *http.Response
		var monitorStatus *stackstate_api.MonitorStatus
		var err error

		if args.ID != 0 {
			monitorStatus, resp, err = api.MonitorApi.GetMonitorWithStatus(cli.Context, args.ID).Execute()
		} else {
			monitorStatus, resp, err = api.MonitorUrnApi.GetMonitorWithStatusByURN(cli.Context, args.Identifier).Execute()
		}

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"monitor-status": monitorStatus,
			})
		} else {
			cli.Printer.PrintLn("")
			cli.Printer.PrintLn("Monitor Health State count: " + util.ToString(monitorStatus.MonitorHealthStateStateCount))
			if monitorStatus.HasErrors() {
				PrintErrors(cli, monitorStatus.GetErrors())
			}

			metricsData := make([][]interface{}, 0)
			bucketSizeS := monitorStatus.GetMetrics().HealthSyncServiceMetrics.BucketSizeSeconds
			bucketSizeSDouble := bucketSizeS * TWO
			bucketSizeSTriple := bucketSizeS * THREE
			metrics := monitorStatus.GetMetrics().HealthSyncServiceMetrics
			metricsData = append(metricsData, CreateMetricRows("latency (Seconds)", metrics.GetLatencySeconds()))
			metricsData = append(metricsData, CreateMetricRows("messages processed (per second)", metrics.GetMessagePerSecond()))
			metricsData = append(metricsData, CreateMetricRows("monitor health states created (per second)", metrics.GetCreatesPerSecond()))
			metricsData = append(metricsData, CreateMetricRows("monitor health states updated (per second)", metrics.GetUpdatesPerSecond()))
			metricsData = append(metricsData, CreateMetricRows("monitor health states deleted (per second)", metrics.GetDeletesPerSecond()))

			cli.Printer.PrintLn("")
			cli.Printer.PrintLn("Monitor Stream metrics:")
			cli.Printer.Table(printer.TableData{
				Header: []string{"metric",
					"value between now and " + util.ToString(bucketSizeS) + " seconds ago",
					"value between " + util.ToString(bucketSizeS) + " and " + util.ToString(bucketSizeSDouble) + " seconds ago",
					"value between " + util.ToString(bucketSizeSDouble) + " and " + util.ToString(bucketSizeSTriple) + " seconds ago"},
				Data: metricsData,
			})

			topologyMatchResult := monitorStatus.TopologyMatchResult
			cli.Printer.PrintLn("")
			cli.Printer.PrintLn("Monitor health states with identifier matching exactly 1 topology element: " + util.ToString(topologyMatchResult.MatchedCheckStates))

			if len(topologyMatchResult.GetUnmatchedCheckStates()) > 0 {
				PrintTopologyMatchResultUnmatched(cli, topologyMatchResult.GetUnmatchedCheckStates(), monitorStatus)
			}

			if len(topologyMatchResult.GetMultipleMatchesCheckStates()) > 0 {
				PrintTopologyMatchResultMultipleMatched(cli, topologyMatchResult.GetMultipleMatchesCheckStates())
			}
		}

		return nil
	}
}
func PrintErrors(cli *di.Deps,
	errors []stackstate_api.MonitorError,
) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Monitor Stream errors:")
	data := make([][]interface{}, 0)
	for _, e := range errors {
		data = append(data, []interface{}{
			e.Error,
			e.Count,
		})
	}

	cli.Printer.Table(printer.TableData{
		Header: []string{"message", "occurrence count"},
		Data:   data,
	})
}

func PrintTopologyMatchResultUnmatched(cli *di.Deps,
	unmatchedCheckStates []stackstate_api.UnmatchedCheckState,
	monitorStatus *stackstate_api.MonitorStatus) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Monitor health states with identifier which has no matching topology element:")

	unmatchedData := make([][]interface{}, 0)
	for _, u := range unmatchedCheckStates {
		unmatchedData = append(unmatchedData, []interface{}{
			u.GetCheckStateId(),
			u.GetTopologyElementIdentifier(),
		})
	}

	cli.Printer.Table(printer.TableData{
		Header: []string{"monitor health state id", "topology element identifier"},
		Data:   unmatchedData,
	})
}

func PrintTopologyMatchResultMultipleMatched(cli *di.Deps,
	multipleMatchesCheckState []stackstate_api.MultipleMatchesCheckState,
) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Monitor health  states with identifier which has multiple matching topology elements:")

	multipleMatched := make([][]interface{}, 0)
	for _, m := range multipleMatchesCheckState {
		multipleMatched = append(multipleMatched, []interface{}{
			m.GetCheckStateId(),
			m.GetTopologyElementIdentifier(),
			m.GetMatchCount(),
		})
	}

	cli.Printer.Table(printer.TableData{
		Header: []string{"monitor health state id", "topology element identifier", "number of matched topology elements"},
		Data:   multipleMatched,
	})
}

func CreateMetricRows(metric string, metrics []stackstate_api.MetricBucketValue) []interface{} {
	metricField := make([]interface{}, 0)
	metricField = append(metricField, metric)

	metricRow := make([]interface{}, 0)
	for _, m := range metrics {
		metricRow = append(metricRow, m.Value)
	}
	metricRow = append(metricField, metricRow...)

	return metricRow
}
