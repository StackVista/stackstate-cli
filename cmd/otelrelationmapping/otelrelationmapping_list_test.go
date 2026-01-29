package otelrelationmapping_test

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelmapping_test"
	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestOtelRelationMappingListJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsResponse.Result = otelmapping_test.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsCalls
	assert.Len(t, calls, 1)

	expected := []map[string]interface{}{
		{
			"otel_relation_mappings": otelmapping_test.TestAllMappingItems,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestOtelRelationMappingListTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsResponse.Result = otelmapping_test.TestAllMappingItems

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	calls := *cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingsCalls
	assert.Len(t, calls, 1)

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Name", "Identifier"},
			Data: [][]interface{}{
				{otelmapping_test.TestSomeOtelMappingItem.Name, "identifier"},
				{otelmapping_test.TestSomeOtelMappingItem2.Name, "identifier2"},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
