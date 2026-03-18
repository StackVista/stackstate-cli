package topology

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func setStateCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StateCommand(&cli.Deps)
	return &cli, cmd
}

func TestTopologyStateJson(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "-o", "json")

	// Assert API was called
	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Assert JSON output contains components
	jsonCalls := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, jsonCalls, 1)

	jsonData := jsonCalls[0]
	assert.NotNil(t, jsonData["components"])

	// Check the component has the health state
	components := jsonData["components"].([]ComponentState)
	assert.Len(t, components, 1)
	assert.Equal(t, "CRITICAL", components[0].HealthState)
}

func TestTopologyStateTable(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type")

	// Assert API was called
	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Assert table output
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Equal(t, []string{"Name", "Type", "Health State", "Identifiers"}, table.Header)
	assert.Len(t, table.Data, 1)

	row := table.Data[0]
	assert.Equal(t, "test-component", row[0])
	assert.Equal(t, "test type", row[1])
	assert.Equal(t, "CRITICAL", row[2])
	assert.Equal(t, "urn:test:component:1", row[3])
}

func TestTopologyStateWithTags(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd,
		"--type", "test type",
		"--tag", "service.namespace:test",
		"--tag", "service.name:myservice")

	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Verify query contains both tags ANDed
	call := calls[0]
	request := call.PviewSnapshotRequest
	expectedQuery := `type = "test type" AND label = "service.namespace:test" AND label = "service.name:myservice"`
	assert.Equal(t, expectedQuery, request.Query)
}

func TestTopologyStateWithIdentifiers(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd,
		"--type", "test type",
		"--identifier", "urn:test:1",
		"--identifier", "urn:test:2")

	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Verify query contains identifiers in IN clause
	call := calls[0]
	request := call.PviewSnapshotRequest
	expectedQuery := `type = "test type" AND identifier IN ("urn:test:1", "urn:test:2")`
	assert.Equal(t, expectedQuery, request.Query)
}

func TestTopologyStateNoResults(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type":      "ViewSnapshot",
			"components": []interface{}{},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{},
				"layers":         []interface{}{},
				"domains":        []interface{}{},
			},
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "nonexistent")

	// Verify empty table is displayed
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Equal(t, []string{"Name", "Type", "Health State", "Identifiers"}, table.Header)
	assert.Len(t, table.Data, 0)
	assert.Equal(t, printer.NotFoundMsg{Types: "components"}, table.MissingTableDataMsg)
}

func TestTopologyStateDefaultsToUnknown(t *testing.T) {
	cli, cmd := setStateCmd(t)
	// Response without state field
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshot",
			"components": []interface{}{
				map[string]interface{}{
					"id":          float64(229404307680647),
					"name":        "test-component",
					"type":        float64(239975151751041),
					"identifiers": []interface{}{"urn:test:component:1"},
					// No state field
				},
			},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{
					map[string]interface{}{
						"id":   float64(239975151751041),
						"name": "test type",
					},
				},
			},
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type")

	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	row := table.Data[0]
	assert.Equal(t, "UNKNOWN", row[2]) // Health state defaults to UNKNOWN
}

func TestTopologyStateLimit(t *testing.T) {
	cli, cmd := setStateCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshot",
			"components": []interface{}{
				map[string]interface{}{
					"id":          float64(1),
					"name":        "component-1",
					"type":        float64(239975151751041),
					"identifiers": []interface{}{"urn:test:1"},
					"state": map[string]interface{}{
						"healthState": "CLEAR",
					},
				},
				map[string]interface{}{
					"id":          float64(2),
					"name":        "component-2",
					"type":        float64(239975151751041),
					"identifiers": []interface{}{"urn:test:2"},
					"state": map[string]interface{}{
						"healthState": "CRITICAL",
					},
				},
				map[string]interface{}{
					"id":          float64(3),
					"name":        "component-3",
					"type":        float64(239975151751041),
					"identifiers": []interface{}{"urn:test:3"},
					"state": map[string]interface{}{
						"healthState": "DEVIATING",
					},
				},
			},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{
					map[string]interface{}{
						"id":   float64(239975151751041),
						"name": "test type",
					},
				},
			},
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "--limit", "2")

	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Len(t, table.Data, 2) // Only 2 components due to limit
}
