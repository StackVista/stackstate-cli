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
	monitor = sts.Monitor{
		Id:              211684343791306,
		Name:            "CPU Usage",
		Identifier:      nil,
		Description:     nil,
		FunctionId:      237346485409361,
		Arguments:       nil,
		RemediationHint: nil,
		IntervalSeconds: 2,
	}
	monitorHealthStateCount int32 = 10
	monitorError                  = sts.MonitorError{
		Error: "Error",
		Count: 11,
	}
	metrics = sts.MonitorMetrics{
		HealthSyncServiceMetrics: sts.HealthStreamMetrics{},
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
		MonitorHealthStateStateCount: monitorHealthStateCount,
		TopologyMatchResult:          topologyMatchResult,
	}
)

func setMonitorStatusCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorStatusCommand(&cli.Deps)
	return &cli, cmd
}

func TestSettingsStatusPrintsToTable(t *testing.T) {
	cli, cmd := setMonitorStatusCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = *monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Monitor Health State count: 10", "", "Monitor Stream errors:", "", "Monitor Stream metrics:", "",
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
