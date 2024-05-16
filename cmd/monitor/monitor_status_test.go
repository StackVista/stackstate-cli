package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
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
		RuntimeStatus:   sts.MONITORRUNTIMESTATUSVALUE_ERROR,
	}
	healthStateCount         int32 = 10
	unmappedHealthStateCount int32 = 1
	monitorError                   = sts.MonitorError{
		Error: "Error",
		Count: 11,
	}
	metrics = sts.MonitorMetrics{
		RuntimeMetrics: sts.MonitorRuntimeMetrics{
			HealthStatesCount: &healthStateCount,
		},
		HealthSyncServiceMetrics: &sts.HealthStreamMetrics{},
	}

	monitorRuntimeMetrics = sts.MonitorRuntimeMetrics{
		HealthStatesCount:          &healthStateCount,
		UnmappedHealthStatesCount:  &unmappedHealthStateCount,
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

	unmapped = sts.UnmatchedCheckState{
		CheckStateId:              "136106883530514-db-shard-4",
		TopologyElementIdentifier: "urn:service:/database/db-shard-4wrong",
	}
	unmappedTopologyMatchResult = sts.TopologyMatchResult{
		MatchedCheckStates:         0,
		UnmatchedCheckStates:       []sts.UnmatchedCheckState{unmapped},
		MultipleMatchesCheckStates: nil,
	}

	monitorStatusResult = &sts.MonitorStatus{
		Monitor:             monitor,
		Errors:              []sts.MonitorError{monitorError},
		Metrics:             metrics,
		TopologyMatchResult: &topologyMatchResult,
	}

	monitorStatusResultWithHealthCounts = &sts.MonitorStatus{
		Monitor:             monitor,
		Errors:              []sts.MonitorError{monitorError},
		Metrics:             metricsWithHealthStateCounts,
		TopologyMatchResult: &unmappedTopologyMatchResult,
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
	cli.MockClient.ApiMocks.MonitorAPI.GetMonitorWithStatusResponse.Result = *monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: 10", "Monitor Status: ERROR", "", "Monitor Stream errors:", "", "Monitor Stream metrics:", "",
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
	cli.MockClient.ApiMocks.MonitorAPI.GetMonitorWithStatusResponse.Result = *monitorStatusResultWithHealthCounts

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: 10", "Monitor Status: ERROR", "Monitor last run: 2022-09-13 14:34:35.007 +0000 UTC",
		"", "Monitor Stream errors:", "", "Monitor health states mapped to topology:", "",
		"Monitor Stream metrics:", "", "Monitor health states with identifier matching exactly 1 topology element: 0", "",
		"Monitor health states with identifier which has no matching topology element (1):"}

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
		{
			Header: []string{"monitor health state id", "topology element identifier"},
			Data:   [][]interface{}{{"136106883530514-db-shard-4", "urn:service:/database/db-shard-4wrong"}},
		},
	}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestMonitorStatusForDisabled(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorAPI.GetMonitorWithStatusResponse.Result = *monitorStatusEmpty

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: -", "Monitor Status: ERROR"}
	expectedTableCall := []printer.TableData{}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsStatusPrintsToJson(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorAPI.GetMonitorWithStatusResponse.Result = *monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306", "-o", "json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"monitor-status": monitorStatusResult,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
