package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	testUnknownHealthState = ":shrug:"
	testHealthStreamError  = "health stream error"
	testName               = "name"
	testLastTenSeconds     = "now-10"
	testFirstMetricRange   = "10-20"
	testSecondMetricRange  = "20-30"
	testMetricLatency      = "2000 ms"
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
			Error: testUnknownHealthState,
		},
		{
			Error: testHealthStreamError,
		},
	}

	ExpectedErrorsJson = []map[string]interface{}{
		{
			"count":   int32(0),
			"message": testUnknownHealthState,
		},
		{
			"count":   int32(0),
			"message": testHealthStreamError,
		},
	}
	ExpectedErrorsTable = printer.TableData{
		Header: []string{"Message", "Count"},
		Data: [][]interface{}{
			{testUnknownHealthState, int32(0)},
			{testHealthStreamError, int32(0)},
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: errorsKey},
	}

	ExpectedMetricsJson = []map[string]interface{}{
		{
			testName:              "latency seconds",
			testLastTenSeconds:    float64(23.5),
			testFirstMetricRange:  float64(23.5),
			testSecondMetricRange: float64(23.5),
		},
		{
			testName:              "messages per seconds",
			testLastTenSeconds:    float64(23.5),
			testFirstMetricRange:  float64(23.5),
			testSecondMetricRange: "-",
		},
		{
			testName:              "creates per seconds",
			testLastTenSeconds:    float64(23.5),
			testFirstMetricRange:  "-",
			testSecondMetricRange: "-",
		},
		{
			testName:              "updates per seconds",
			testLastTenSeconds:    "-",
			testFirstMetricRange:  "-",
			testSecondMetricRange: "-",
		},
		{
			testName:              "deletes per seconds",
			testLastTenSeconds:    "-",
			testFirstMetricRange:  "-",
			testSecondMetricRange: "-",
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
		MissingTableDataMsg: printer.NotFoundMsg{Types: metricsKey},
	}
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
		Errors:  SomeErrors,
		Metrics: SomeMetrics,
	}

	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamStatusResponse.Result = stackstate_api.HealthStreamStatus{
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
				{testMetricLatency, testMetricLatency, "-", "-", "-", "-"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: errorsKey},
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
			errorsKey:  ExpectedErrorsJson,
			metricsKey: ExpectedMetricsJson,
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
				{testMetricLatency, testMetricLatency, "-", "-", "-", "-"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: errorsKey},
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
				errorsKey:  ExpectedErrorsJson,
				metricsKey: ExpectedMetricsJson,
			},
			errorsKey:  ExpectedErrorsJson,
			metricsKey: ExpectedMetricsJson,
		},
	}, *cli.MockPrinter.PrintJsonCalls)
}
