package monitor

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/internal/util"
)

type StatusArgs struct {
	ID         int64
	Identifier string
}

func MonitorStatusCommand(cli *di.Deps) *cobra.Command {
	args := &StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of a monitor",
		Long:  "Get the status of a single monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorStatusCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

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

		identifier := IdOrIdentifier(args.ID, args.Identifier)
		monitorStatus, resp, err = api.MonitorApi.GetMonitorWithStatus(cli.Context, identifier).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"monitor-status": monitorStatus,
			})
		} else {
			monitorMetrics := monitorStatus.GetMetrics()
			monitorRuntimeMetrics := monitorMetrics.GetRuntimeMetrics()
			cli.Printer.PrintLn("")
			cli.Printer.PrintLn("Monitor Health State count: " + util.ToString(monitorRuntimeMetrics.HealthStatesCount))
			cli.Printer.PrintLn("Monitor Status: " + util.ToString(monitorStatus.Monitor.RuntimeStatus))
			lastRunTimestamp, lastRunTimestampOk := monitorRuntimeMetrics.GetLastRunTimestampOk()
			if lastRunTimestampOk {
				cli.Printer.PrintLn("Monitor last run: " + time.UnixMilli(*lastRunTimestamp).UTC().String())
			}

			if monitorStatus.HasErrors() {
				PrintErrors(cli, monitorStatus.GetErrors())
			}

			if lastRunTimestampOk {
				cli.Printer.PrintLn("")
				cli.Printer.PrintLn("Monitor health states mapped to topology:")
				topologyMappedMetricsData := make([][]interface{}, 0)
				topologyMappedMetricsData = append(topologyMappedMetricsData, []interface{}{"CLEAR", monitorRuntimeMetrics.GetClearCount()})
				topologyMappedMetricsData = append(topologyMappedMetricsData, []interface{}{"DEVIATING", monitorRuntimeMetrics.GetDeviatingCount()})
				topologyMappedMetricsData = append(topologyMappedMetricsData, []interface{}{"CRITICAL", monitorRuntimeMetrics.GetCriticalCount()})
				topologyMappedMetricsData = append(topologyMappedMetricsData, []interface{}{"UNKNOWN", monitorRuntimeMetrics.GetUnknownCount()})

				cli.Printer.Table(printer.TableData{
					Header: []string{"HealthState", "count"},
					Data:   topologyMappedMetricsData,
				})
			}

			healthSyncMetricsData := make([][]interface{}, 0)
			healthSyncMetrics, healthSyncMetricsOk := monitorMetrics.GetHealthSyncServiceMetricsOk()

			if healthSyncMetricsOk {
				bucketSizeS := healthSyncMetrics.BucketSizeSeconds
				bucketSizeSDouble := bucketSizeS * TWO
				bucketSizeSTriple := bucketSizeS * THREE
				healthSyncMetricsData = append(healthSyncMetricsData, CreateMetricRows("latency (Seconds)", healthSyncMetrics.GetLatencySeconds()))
				healthSyncMetricsData = append(healthSyncMetricsData, CreateMetricRows("messages processed (per second)", healthSyncMetrics.GetMessagePerSecond()))
				healthSyncMetricsData = append(healthSyncMetricsData, CreateMetricRows("monitor health states created (per second)", healthSyncMetrics.GetCreatesPerSecond()))
				healthSyncMetricsData = append(healthSyncMetricsData, CreateMetricRows("monitor health states updated (per second)", healthSyncMetrics.GetUpdatesPerSecond()))
				healthSyncMetricsData = append(healthSyncMetricsData, CreateMetricRows("monitor health states deleted (per second)", healthSyncMetrics.GetDeletesPerSecond()))

				cli.Printer.PrintLn("")
				cli.Printer.PrintLn("Monitor Stream metrics:")
				cli.Printer.Table(printer.TableData{
					Header: []string{"metric",
						"value between now and " + util.ToString(bucketSizeS) + " seconds ago",
						"value between " + util.ToString(bucketSizeS) + " and " + util.ToString(bucketSizeSDouble) + " seconds ago",
						"value between " + util.ToString(bucketSizeSDouble) + " and " + util.ToString(bucketSizeSTriple) + " seconds ago"},
					Data: healthSyncMetricsData,
				})
			}

			topologyMatchResult, topologyMatchResultOk := monitorStatus.GetTopologyMatchResultOk()
			if topologyMatchResultOk {
				cli.Printer.PrintLn("")
				cli.Printer.PrintLn("Monitor health states with identifier matching exactly 1 topology element: " + util.ToString(topologyMatchResult.MatchedCheckStates))

				if monitorRuntimeMetrics.UnmappedHealthStatesCount != nil && *monitorRuntimeMetrics.UnmappedHealthStatesCount > 0 {
					PrintTopologyMatchResultUnmatched(cli, *monitorRuntimeMetrics.UnmappedHealthStatesCount, topologyMatchResult.GetUnmatchedCheckStates(), monitorStatus)
				}

				if len(topologyMatchResult.GetMultipleMatchesCheckStates()) > 0 {
					PrintTopologyMatchResultMultipleMatched(cli, topologyMatchResult.GetMultipleMatchesCheckStates())
				}
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
	unmappedElementsCount int32,
	unmatchedCheckStates []stackstate_api.UnmatchedCheckState,
	monitorStatus *stackstate_api.MonitorStatus) {
	cli.Printer.PrintLn("")
	cli.Printer.PrintLn("Monitor health states with identifier which has no matching topology element (" + util.ToString(unmappedElementsCount) + "):")

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
