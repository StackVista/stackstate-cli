package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	recoveryMessage = "status recovery message"

	noVal  = stackstate_api.NewMetricBucketValue()
	hasVal = (func() *stackstate_api.MetricBucketValue {
		v := stackstate_api.NewMetricBucketValue()
		v.SetValue(float64(23.5))
		return v
	})()

	SomeMetrics = stackstate_api.HealthStreamMetrics{
		BucketSizeSeconds: 10,
		LatencySeconds:    []stackstate_api.MetricBucketValue{*hasVal, *hasVal, *hasVal},
		MessagePerSecond:  []stackstate_api.MetricBucketValue{*hasVal, *hasVal, *noVal},
		CreatesPerSecond:  []stackstate_api.MetricBucketValue{*hasVal, *noVal, *noVal},
		UpdatesPerSecond:  []stackstate_api.MetricBucketValue{*noVal, *noVal, *noVal},
		// NOTE Last one is an empty array.
	}
	SomeErrors = []stackstate_api.HealthStreamError{
		{
			Error: ":shrug:",
		},
		{
			Error: "health stream error",
		},
	}

	ExpectedErrorsJson = []map[string]interface{}{
		{
			"count":   int32(0),
			"message": ":shrug:",
		},
		{
			"count":   int32(0),
			"message": "health stream error",
		},
	}
	ExpectedErrorsTable = printer.TableData{
		Header: []string{"Message", "Count"},
		Data: [][]interface{}{
			{":shrug:", int32(0)},
			{"health stream error", int32(0)},
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "errors"},
	}

	ExpectedMetricsJson = []map[string]interface{}{
		{
			"name":   "latency seconds",
			"now-10": float64(23.5),
			"10-20":  float64(23.5),
			"20-30":  float64(23.5),
		},
		{
			"name":   "messages per seconds",
			"now-10": float64(23.5),
			"10-20":  float64(23.5),
			"20-30":  "-",
		},
		{
			"name":   "creates per seconds",
			"now-10": float64(23.5),
			"10-20":  "-",
			"20-30":  "-",
		},
		{
			"name":   "updates per seconds",
			"now-10": "-",
			"10-20":  "-",
			"20-30":  "-",
		},
		{
			"name":   "deletes per seconds",
			"now-10": "-",
			"10-20":  "-",
			"20-30":  "-",
		},
	}
	ExpectedMetricsTable = printer.TableData{
		Header: []string{"Metric", "10s ago", "10-20s ago", "20-30s ago"},
		Data: [][]interface{}{
			{"latency seconds", float64(23.5), float64(23.5), float64(23.5)},
			{"messages per seconds", float64(23.5), float64(23.5), "-"},
			{"creates per seconds", float64(23.5), "-", "-"},
			{"updates per seconds", "-", "-", "-"},
			{"deletes per seconds", "-", "-", "-"},
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "metrics"},
	}
)

func setupHealthStatusCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthStatusCommand(&cli.Deps)

	cli.MockClient.ApiMocks.HealthSynchronizationAPI.GetHealthSynchronizationStreamTopologyMatchesResponse.Result = stackstate_api.TopologyMatchResult{
		MultipleMatchesCheckStates: []stackstate_api.MultipleMatchesCheckState{{
			CheckStateId:              "1",
			TopologyElementIdentifier: "id",
			MatchCount:                12,
		}},
	}

	cli.MockClient.ApiMocks.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamStatusResponse.Result = stackstate_api.HealthSubStreamStatus{
		CheckStateCount: 12,
		SubStreamState: stackstate_api.HealthSubStreamConsistencyState{
			HealthSubStreamExpiry: &stackstate_api.HealthSubStreamExpiry{
				RepeatIntervalMs: 2000,
				ExpiryIntervalMs: 2000,
			},
		},
		Errors:  SomeErrors,
		Metrics: SomeMetrics,
	}

	cli.MockClient.ApiMocks.HealthSynchronizationAPI.GetHealthSynchronizationStreamStatusResponse.Result = stackstate_api.HealthStreamStatus{
		RecoverMessage:   &recoveryMessage,
		ConsistencyModel: "stream consistency model",
		GlobalErrors:     SomeErrors,
		AggregateMetrics: SomeMetrics,
		MainStreamStatus: &stackstate_api.HealthSubStreamStatus{
			CheckStateCount: 2,
			SubStreamState: stackstate_api.HealthSubStreamConsistencyState{
				HealthSubStreamExpiry: &stackstate_api.HealthSubStreamExpiry{
					RepeatIntervalMs: 2000,
					ExpiryIntervalMs: 2000,
				},
			},
			Errors:  SomeErrors,
			Metrics: SomeMetrics,
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

func TestHealthStatusWithSubStreamPrintToText(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--sub-stream-urn", "dummy-sub-stream")
	expectedSuccessCalls := []string{
		"Synchronized check state count: 12",
		"\nConsistency state:",
		"\nMetrics:",
		"\nErrors:",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.PrintLnCalls)
	expectedTableCalls := []printer.TableData{
		{
			Header: []string{"Repeat Interval", "Expiry Interval", "Snapshot Repeat Interval", "Snapshot Expiry Interval", "Checkpoint Offset", "Checkpoint Batch Index"},
			Data: [][]interface{}{
				{"2000 ms", "2000 ms", "-", "-", "-", "-"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "errors"},
		},
		ExpectedMetricsTable,
		ExpectedErrorsTable,
	}
	assert.Equal(t, expectedTableCalls, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthStatusWithSubStreamPrintToJson(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--sub-stream-urn", "dummy-sub-stream", "-o", "json")
	assert.Equal(t, []map[string]interface{}{
		{
			"check-state-count": int32(12),
			"consistency-state": map[string]interface{}{
				"expiry-interval": int32(2000),
				"repeat-interval": int32(2000),
			},
			"errors":  ExpectedErrorsJson,
			"metrics": ExpectedMetricsJson,
		},
	}, *cli.MockPrinter.PrintJsonCalls)
}

func TestHealthStatusStreamPrintToText(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn")
	expectedSuccessCalls := []string{
		"Stream consistency model: stream consistency model",
		"\nAggregate metrics for the stream and all substreams:",
		"\nErrors for non-existing substreams:",
		"\nMain stream status:",
		"Synchronized check state count: 2",
		"\nConsistency state:",
		"\nMetrics:",
		"\nErrors:",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.PrintLnCalls)
	expectedTableCalls := []printer.TableData{
		ExpectedMetricsTable,
		ExpectedErrorsTable,
		{
			Header: []string{"Repeat Interval", "Expiry Interval", "Snapshot Repeat Interval", "Snapshot Expiry Interval", "Checkpoint Offset", "Checkpoint Batch Index"},
			Data: [][]interface{}{
				{"2000 ms", "2000 ms", "-", "-", "-", "-"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "errors"},
		},
		ExpectedMetricsTable,
		ExpectedErrorsTable,
	}
	assert.Equal(t, expectedTableCalls, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthStatusStreamPrintToJson(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "-o", "json")
	assert.Equal(t, []map[string]interface{}{
		{
			"consistency-model": "stream consistency model",
			"main-stream": map[string]interface{}{
				"check-state-count": int32(2),
				"consistency-state": map[string]interface{}{
					"expiry-interval": int32(2000),
					"repeat-interval": int32(2000),
				},
				"errors":  ExpectedErrorsJson,
				"metrics": ExpectedMetricsJson,
			},
			"errors":  ExpectedErrorsJson,
			"metrics": ExpectedMetricsJson,
		},
	}, *cli.MockPrinter.PrintJsonCalls)
}
