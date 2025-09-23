package otelrelationmapping

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestOtelRelationMappingStatusJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelRelationMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusResponse.Result = *otelmapping.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier", "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel-relation-mapping": otelmapping.TestSomeOtelMappingStatus,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelRelationMappingStatusTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelRelationMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusResponse.Result = *otelmapping.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingStatusCalls
	assert.Len(t, calls, 1)

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Name", "Identifier", "Components", "Relations"},
			Data: [][]interface{}{
				{otelmapping.TestSomeOtelMappingStatusItem.Name, "identifier", otelmapping.TestComponentCount, otelmapping.TestRelationCount},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
		},
		{
			Header: []string{"Metric", "10s ago", "10-20s ago", "20-30s ago"},
			Data: [][]interface{}{
				{"latency seconds", otelmapping.TestMetricValue, otelmapping.TestMetricValue, otelmapping.TestMetricValue},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "metrics"},
		},
		{
			Header: []string{"Issue Id", "Level", "Message"},
			Data: [][]interface{}{
				{"-", otelmapping.TestError1.Level, otelmapping.TestError1.Message},
				{"-", otelmapping.TestError2.Level, otelmapping.TestError2.Message},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel relation mapping errors"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
