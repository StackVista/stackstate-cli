package otelrelationmapping_test

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping_test"
	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestOtelRelationMappingStatusJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusResponse.Result = *otelmapping_test.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier", "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel-relation-mapping": otelmapping_test.TestSomeOtelMappingStatus,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelRelationMappingStatusTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusResponse.Result = *otelmapping_test.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusCalls
	assert.Len(t, calls, 1)

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Name", "Identifier", "Components", "Relations"},
			Data: [][]interface{}{
				{otelmapping_test.TestSomeOtelMappingStatusItem.Name, "identifier", otelmapping_test.TestComponentCount, otelmapping_test.TestRelationCount},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
		},
		{
			Header: []string{"Metric", "10s ago", "10-20s ago", "20-30s ago"},
			Data: [][]interface{}{
				{"latency seconds", otelmapping_test.TestMetricValue, otelmapping_test.TestMetricValue, otelmapping_test.TestMetricValue},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "metrics"},
		},
		{
			Header: []string{"Issue Id", "Level", "Message"},
			Data: [][]interface{}{
				{"-", otelmapping_test.TestError1.Level, otelmapping_test.TestError1.Message},
				{"-", otelmapping_test.TestError2.Level, otelmapping_test.TestError2.Message},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel relation mapping errors"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
