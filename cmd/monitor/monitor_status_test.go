package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	count            = int32(3)
	lastRunTimestamp = int64(1663079675007)
	monitor          = sts.Monitor{
		Id:              211684343791306,
		Name:            "CPU Usage",
		Identifier:      nil,
		Description:     nil,
		FunctionId:      237346485409361,
		Arguments:       nil,
		RemediationHint: nil,
		IntervalSeconds: 2,
		Status:          sts.MONITORSTATUSVALUE_ENABLED,
	}
	monitorHealthStateCount int32 = 10
	monitorError                  = sts.MonitorError{
		Error: "Error",
		Count: 11,
	}
	metrics = sts.MonitorMetrics{
		HealthSyncServiceMetrics: &sts.HealthStreamMetrics{},
	}

	monitorRuntimeMetrics = sts.MonitorRuntimeMetrics{
		UnknownCount:               &count,
		ClearCount:                 &count,
		DeviatingCount:             &count,
		CriticalCount:              &count,
		LastRunTimestamp:           &lastRunTimestamp,
		LastFailedRunTimestamp:     &lastRunTimestamp,
		LastSuccessfulRunTimestamp: &lastRunTimestamp,
	}

	metricsWithHealthStateCounts = sts.MonitorMetrics{
		HealthSyncServiceMetrics: &sts.HealthStreamMetrics{},
		RuntimeMetrics:           monitorRuntimeMetrics,
	}

	topologyMatchResult = sts.TopologyMatchResult{
		MatchedCheckStates:         0,
		UnmatchedCheckStates:       nil,
		MultipleMatchesCheckStates: nil,
	}

	monitorStatusResult = &sts.MonitorStatus{
		Monitor:                      monitor,
		Errors:                       []sts.MonitorError{monitorError},
		Metrics:                      metrics,
		MonitorHealthStateStateCount: &monitorHealthStateCount,
		TopologyMatchResult:          &topologyMatchResult,
	}

	monitorStatusResultWithHealthCounts = &sts.MonitorStatus{
		Monitor:                      monitor,
		Errors:                       []sts.MonitorError{monitorError},
		Metrics:                      metricsWithHealthStateCounts,
		MonitorHealthStateStateCount: &monitorHealthStateCount,
		TopologyMatchResult:          &topologyMatchResult,
	}

	emptyMonitorRuntimeMetrics = sts.MonitorRuntimeMetrics{}

	emptyHealthStreamMetrics = sts.MonitorMetrics{
		RuntimeMetrics: emptyMonitorRuntimeMetrics,
	}
	monitorStatusEmpty = &sts.MonitorStatus{
		Monitor: monitor,
		Metrics: emptyHealthStreamMetrics,
	}
)

func setMonitorStatusCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorStatusCommand(&cli.Deps)
	return &cli, cmd
}

func TestMonitorStatusPrintsToTable(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = *monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: 10", "Monitor Status: ENABLED", "", "Monitor Stream errors:", "", "Monitor Stream metrics:", "",
		"Monitor health states with identifier matching exactly 1 topology element: 0"}
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"message", "occurrence count"},
			Data:   [][]interface{}{{"Error", int32(11)}},
		},
		{
			Header: []string{"metric", "value between now and 0 seconds ago", "value between 0 and 0 seconds ago",
				"value between 0 and 0 seconds ago"},
			Data: [][]interface{}{{"latency (Seconds)"}, {"messages processed (per second)"},
				{"monitor health states created (per second)"}, {"monitor health states updated (per second)"},
				{"monitor health states deleted (per second)"}},
		},
	}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestMonitorStatusWithHealthStatesCountsPrintsToTable(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = *monitorStatusResultWithHealthCounts

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: 10", "Monitor Status: ENABLED", "Monitor last run: 2022-09-13 14:34:35.007 +0000 UTC",
		"", "Monitor Stream errors:", "", "Monitor health states mapped to topology:", "",
		"Monitor Stream metrics:", "", "Monitor health states with identifier matching exactly 1 topology element: 0"}

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"message", "occurrence count"},
			Data:   [][]interface{}{{"Error", int32(11)}},
		},
		{
			Header: []string{"HealthState", "count"},
			Data:   [][]interface{}{{"CLEAR", count}, {"DEVIATING", count}, {"CRITICAL", count}, {"UNKNOWN", count}},
		},
		{
			Header: []string{"metric", "value between now and 0 seconds ago", "value between 0 and 0 seconds ago",
				"value between 0 and 0 seconds ago"},
			Data: [][]interface{}{{"latency (Seconds)"}, {"messages processed (per second)"},
				{"monitor health states created (per second)"}, {"monitor health states updated (per second)"},
				{"monitor health states deleted (per second)"}},
		},
	}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestMonitorStatusForDisabled(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = *monitorStatusEmpty

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: -", "Monitor Status: ENABLED"}
	expectedTableCall := []printer.TableData{}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsStatusPrintsToJson(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = *monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306", "-o", "json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"monitor-status": monitorStatusResult,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
