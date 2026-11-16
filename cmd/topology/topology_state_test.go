package topology

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	testType        = "type"
	testState       = "state"
	testHealthState = "healthState"
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
	assert.NotNil(t, jsonData[componentsKey])

	// Check the component has the health state
	components := jsonData[componentsKey].([]ComponentState)
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
	assert.Equal(t, []string{testName, typeColumn, "Health State", identifiersColumn}, table.Header)
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
		Type: querySnapshotResultType,
		ViewSnapshotResponse: map[string]interface{}{
			testTypeKey:   viewSnapshotType,
			componentsKey: []interface{}{},
			testMetadata: map[string]interface{}{
				testComponentTypes: []interface{}{},
				"layers":           []interface{}{},
				"domains":          []interface{}{},
			},
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "nonexistent")

	// Verify empty table is displayed
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Equal(t, []string{testName, typeColumn, "Health State", identifiersColumn}, table.Header)
	assert.Len(t, table.Data, 0)
	assert.Equal(t, printer.NotFoundMsg{Types: componentsKey}, table.MissingTableDataMsg)
}

func TestTopologyStateDefaultsToUnknown(t *testing.T) {
	cli, cmd := setStateCmd(t)
	// Response without state field
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: querySnapshotResultType,
		ViewSnapshotResponse: map[string]interface{}{
			testTypeKey: viewSnapshotType,
			componentsKey: []interface{}{
				map[string]interface{}{
					"id":            float64(229404307680647),
					nameKey:         "test-component",
					testType:        float64(239975151751041),
					testIdentifiers: []interface{}{"urn:test:component:1"},
					// No state field
				},
			},
			testMetadata: map[string]interface{}{
				testComponentTypes: []interface{}{
					map[string]interface{}{
						"id":    float64(239975151751041),
						nameKey: "test type",
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
		Type: querySnapshotResultType,
		ViewSnapshotResponse: map[string]interface{}{
			testTypeKey: viewSnapshotType,
			componentsKey: []interface{}{
				map[string]interface{}{
					"id":            float64(1),
					nameKey:         "component-1",
					testType:        float64(239975151751041),
					testIdentifiers: []interface{}{"urn:test:1"},
					testState: map[string]interface{}{
						testHealthState: "CLEAR",
					},
				},
				map[string]interface{}{
					"id":            float64(2),
					nameKey:         "component-2",
					testType:        float64(239975151751041),
					testIdentifiers: []interface{}{"urn:test:2"},
					testState: map[string]interface{}{
						testHealthState: "CRITICAL",
					},
				},
				map[string]interface{}{
					"id":            float64(3),
					nameKey:         "component-3",
					testType:        float64(239975151751041),
					testIdentifiers: []interface{}{"urn:test:3"},
					testState: map[string]interface{}{
						testHealthState: "DEVIATING",
					},
				},
			},
			testMetadata: map[string]interface{}{
				testComponentTypes: []interface{}{
					map[string]interface{}{
						"id":    float64(239975151751041),
						nameKey: "test type",
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
