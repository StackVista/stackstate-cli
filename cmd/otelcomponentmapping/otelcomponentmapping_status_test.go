package otelcomponentmapping

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestOtelComponentMappingStatusJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelComponentMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingStatusResponse.Result = *otelmapping.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier", "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingStatusCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel-component-mapping": otelmapping.TestSomeOtelMappingStatus,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelComponentMappingStatusTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelComponentMappingStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingStatusResponse.Result = *otelmapping.TestSomeOtelMappingStatus

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "identifier")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingStatusCalls
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
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel component mapping errors"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
