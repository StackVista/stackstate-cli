package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
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
		TopologyMapping: "urn:service1:/${tags.name}",
		IntervalSeconds: 2,
	}

	expiry int32 = 7200
	repeat int32 = 2
	error        = sts.HealthStreamError{
		ErrorCode: "SubStreamStopWithoutStart",
		Level:     "ERROR",
		Error:     "Error",
		Count:     11,
	}
	healthStreamStatus = sts.HealthStreamStatus{
		Partition:        1,
		ConsistencyModel: "REPEAT_SNAPSHOTS",
		RecoverMessage:   nil,
		GlobalErrors:     nil,
		AggregateMetrics: sts.HealthStreamMetrics{},
		MainStreamStatus: &sts.HealthSubStreamStatus{
			Errors:  &[]sts.HealthStreamError{error},
			Metrics: sts.HealthStreamMetrics{},
			SubStreamState: sts.HealthSubStreamConsistencyState{
				HealthSubStreamSnapshot: &sts.HealthSubStreamSnapshot{
					Type: "HealthSubStreamSnapshot", ExpiryIntervalMs: &expiry, RepeatIntervalMs: repeat,
				},
			},
		},
	}

	topologyMatchResult = sts.TopologyMatchResult{
		MatchedCheckStates:         0,
		UnmatchedCheckStates:       nil,
		MultipleMatchesCheckStates: nil,
	}

	monitorStatusResult = sts.MonitorStatus{
		Monitor:             monitor,
		Status:              healthStreamStatus,
		TopologyMatchResult: topologyMatchResult,
	}
)

func setMonitorStatusCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := MonitorStatusCommand(&cli.Deps)
	return &cli, cmd
}

func TestSettingsStatusPrintsToTable(t *testing.T) {
	cli, cmd := setMonitorStatusCmd()
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306")

	expectedPrintlnCalls := []string{"", "Synchronized check state count: 0", "Repeat interval (Seconds): 0",
		"Expiry (Seconds): 7", "", "Synchronization errors:", "", "Synchronization metrics:", "",
		"Check states with identifier matching exactly 1 topology element: 0"}
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"code", "level", "message", "occurrence count"},
			Data:   [][]interface{}{{"SubStreamStopWithoutStart", "ERROR", "Error", int32(11)}},
		},
		{
			Header: []string{"metric", "value between now and 0 seconds ago", "value between 0 and 0 seconds ago",
				"value between 0 and 0 seconds ago"},
			Data: [][]interface{}{{"latency (Seconds)"}, {"messages processed (per second)"},
				{"check states created (per second)"}, {"check states updated (per second)"},
				{"check states deleted (per second)"}},
		},
	}

	assert.Equal(t, expectedPrintlnCalls, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsStatusPrintsToJson(t *testing.T) {
	cli, cmd := setMonitorStatusCmd()
	cli.MockClient.ApiMocks.MonitorApi.GetMonitorWithStatusResponse.Result = monitorStatusResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "211684343791306", "--json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"monitor-status": monitorStatusResult,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
