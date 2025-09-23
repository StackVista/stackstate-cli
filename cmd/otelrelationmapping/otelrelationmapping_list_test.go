package otelrelationmapping

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestOtelRelationMappingListJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelRelationMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsResponse.Result = otelmapping.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel relation mappings": otelmapping.TestAllMappingItems,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelRelationMappingListTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelRelationMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsResponse.Result = otelmapping.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsCalls
	assert.Len(t, calls, 1)

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Name", "Identifier"},
			Data: [][]interface{}{
				{otelmapping.TestSomeOtelMappingItem.Name, "identifier"},
				{otelmapping.TestSomeOtelMappingItem2.Name, "identifier2"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
