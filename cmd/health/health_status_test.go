package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	recoveryMessage = "status recovery message"
)

func setupHealthStatusCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatchesResponse.Result = stackstate_api.TopologyMatchResult{
		MultipleMatchesCheckStates: []stackstate_api.MultipleMatchesCheckState{{
			CheckStateId:              "1",
			TopologyElementIdentifier: "id",
			MatchCount:                12,
		}},
	}

	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatusResponse.Result = stackstate_api.HealthSubStreamStatus{
		CheckStateCount: 12,
		SubStreamState: stackstate_api.HealthSubStreamConsistencyState{
			HealthSubStreamExpiry: &stackstate_api.HealthSubStreamExpiry{
				RepeatIntervalMs: 2000,
				ExpiryIntervalMs: 2000,
			},
		},
	}

	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamStatusResponse.Result = stackstate_api.HealthStreamStatus{
		RecoverMessage:   &recoveryMessage,
		ConsistencyModel: "stream consistency model",
		MainStreamStatus: &stackstate_api.HealthSubStreamStatus{
			CheckStateCount: 2,
			SubStreamState: stackstate_api.HealthSubStreamConsistencyState{
				HealthSubStreamExpiry: &stackstate_api.HealthSubStreamExpiry{
					RepeatIntervalMs: 2000,
					ExpiryIntervalMs: 2000,
				},
			},
			Errors: []stackstate_api.HealthStreamError{{
				Error: "health stream error",
			}},
			Metrics: stackstate_api.HealthStreamMetrics{
				BucketSizeSeconds: 10,
			},
		},
	}

	return &cli, cmd
}

func TestHealthStatusTopologyPrintToTable(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--topology-matches", "dummy")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Check state id", "Topology element identifier", "Number of matched topology elements"},
			Data:                [][]interface{}{{"1", "id", int32(12)}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health status"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthStatusTopologyPrintToJson(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--topology-matches", "dummy", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"status": [][]interface{}{{"1", "id", int32(12)}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}

func TestHealthStatusWithSubStream(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--sub-stream-urn", "dummy-sub-stream")
	expectedSuccessCalls := []string{"Synchronized check state count: 12", "Repeat interval (Seconds): 2", "Expiry (Seconds): 2"}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.PrintLnCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthStatusStream(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn")
	expectedSuccessCalls := []string{"This stream is in recovery mode.\n" +
		"This means stackstate is reconstructing the state of the health streams. " +
		"In this period no errors will be reported for the stream,\n" +
		"incoming data will be processed as usual.\n" +
		"The reason recovery mode was entered was because: status recovery message",
		"Consistency model for the stream and all substreams: stream consistency model",
		"Synchronized check state count: 2",
		"Repeat interval (Seconds): 2",
		"Expiry (Seconds): 2",
		"\nSynchronization errors:\n [{  health stream error 0}]",
		"\nSynchronization metrics:\n {10 [] [] [] [] []}",
		"\nAggregate metrics for the stream and all substreams:\n {0 [] [] [] [] []}",
		"\nErrors for non-existing sub streams:\n []",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.PrintLnCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}
