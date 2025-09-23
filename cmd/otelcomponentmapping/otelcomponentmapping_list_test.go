package otelcomponentmapping

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestListOtelComponentMappingsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelComponentMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingsResponse.Result = otelmapping.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingsCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel component mappings": otelmapping.TestAllMappingItems,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelComponentMappingListTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := OtelComponentMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingsResponse.Result = otelmapping.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingsCalls
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
