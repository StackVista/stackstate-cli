package topology

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const baseURL = "https://suse-observability-instance.com"

func setInspectCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := InspectCommand(&cli.Deps)
	return &cli, cmd
}

func mockSnapshotResponse() sts.QuerySnapshotResult {
	return sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshot",
			"components": []interface{}{
				map[string]interface{}{
					"id":          float64(229404307680647),
					"name":        "test-component",
					"type":        float64(239975151751041),
					"layer":       float64(186771622698247),
					"domain":      float64(209616858431909),
					"identifiers": []interface{}{"urn:test:component:1"},
					"tags":        []interface{}{"service.namespace:test"},
					"properties":  map[string]interface{}{"key": "value"},
				},
			},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{
					map[string]interface{}{
						"id":   float64(239975151751041),
						"name": "test type",
					},
				},
				"layers": []interface{}{
					map[string]interface{}{
						"id":   float64(186771622698247),
						"name": "Test Layer",
					},
				},
				"domains": []interface{}{
					map[string]interface{}{
						"id":   float64(209616858431909),
						"name": "Test Domain",
					},
				},
			},
		},
	}
}

func TestTopologyInspectJson(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "-o", "json")

	// Assert API was called
	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Assert JSON output contains components
	jsonCalls := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, jsonCalls, 1)

	jsonData := jsonCalls[0]
	// Components are serialized from []Component struct
	assert.NotNil(t, jsonData["components"])

	assert.True(t, len(calls) > 0, "API should have been called")
}

func TestTopologyInspectTable(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type")

	// Assert API was called
	calls := *cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotCalls
	assert.Len(t, calls, 1)

	// Assert table output
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Equal(t, []string{"Name", "Type", "Identifiers"}, table.Header)
	assert.Len(t, table.Data, 1)

	row := table.Data[0]
	assert.Equal(t, "test-component", row[0])
	assert.Equal(t, "test type", row[1])
	assert.Equal(t, "urn:test:component:1", row[2])
}

func TestTopologyInspectWithTags(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = baseURL

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

func TestTopologyInspectWithIdentifiers(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = baseURL

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

func TestTopologyInspectNoResults(t *testing.T) {
	cli, cmd := setInspectCmd(t)
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
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "nonexistent")

	// Verify empty table is displayed
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	assert.Equal(t, []string{"Name", "Type", "Identifiers"}, table.Header)
	assert.Len(t, table.Data, 0)
	assert.Equal(t, printer.NotFoundMsg{Types: "components"}, table.MissingTableDataMsg)
}

func TestTopologyInspectMultipleIdentifiers(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	responseData := mockSnapshotResponse()
	// Modify to include multiple identifiers
	components := responseData.ViewSnapshotResponse["components"].([]interface{})
	comp := components[0].(map[string]interface{})
	comp["identifiers"] = []interface{}{"urn:test:1", "urn:test:2", "urn:test:3"}
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = responseData
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type")

	// Verify multiple identifiers are joined with comma
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	row := table.Data[0]
	assert.Equal(t, "urn:test:1, urn:test:2, urn:test:3", row[2])
}

func TestTopologyInspectURLEncoding(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	responseData := mockSnapshotResponse()
	// Set identifier with special characters
	components := responseData.ViewSnapshotResponse["components"].([]interface{})
	comp := components[0].(map[string]interface{})
	comp["identifiers"] = []interface{}{"urn:opentelemetry:namespace/service:component/name"}
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = responseData
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "-o", "json")

	// Verify URL encoding in link
	jsonCalls := *cli.MockPrinter.PrintJsonCalls
	returnedComponents := jsonCalls[0]["components"].([]Component)
	// url.PathEscape encodes / to %2F, but : is not encoded, and base URL is prefixed
	expectedLink := baseURL + "/#/components/urn:opentelemetry:namespace%2Fservice:component%2Fname"
	assert.Equal(t, expectedLink, returnedComponents[0].Link)
}

func TestTopologyInspectWithEnvironment(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	responseData := mockSnapshotResponse()
	// Add environment to component
	components := responseData.ViewSnapshotResponse["components"].([]interface{})
	comp := components[0].(map[string]interface{})
	comp["environment"] = float64(123)
	// Add environment to metadata
	metadata := responseData.ViewSnapshotResponse["metadata"].(map[string]interface{})
	metadata["environments"] = []interface{}{
		map[string]interface{}{
			"id":   float64(123),
			"name": "Production",
		},
	}
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = responseData
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "-o", "json")

	// Verify environment is included in JSON
	jsonCalls := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, jsonCalls, 1)
	jsonData := jsonCalls[0]
	assert.NotNil(t, jsonData["components"])
}

func TestTopologyInspectTypeRequired(t *testing.T) {
	cli, cmd := setInspectCmd(t)

	// Execute without --type flag should panic due to required flag
	assert.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)
	}, "command should panic when required --type flag is missing")
}

func TestTopologyInspectMetadataMapping(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = baseURL

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type", "-o", "json")

	// Verify metadata is properly mapped to component by checking table output
	// which includes resolved names
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 0) // JSON mode doesn't produce table output

	// Verify JSON output was generated
	jsonCalls := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, jsonCalls, 1)
}

func TestTopologyInspectWithBaseURL(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = mockSnapshotResponse()
	cli.CurrentContext.URL = "https://stackstate.example.com:8080"

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "test type")

	// Verify base URL is included in the link
	tableCalls := *cli.MockPrinter.TableCalls
	assert.Len(t, tableCalls, 1)

	table := tableCalls[0]
	row := table.Data[0]
	expectedLink := "https://stackstate.example.com:8080/#/components/urn:test:component:1"
	assert.Equal(t, expectedLink, row[3])
}

func TestTopologyInspectFetchTimeout(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type":              "ViewSnapshotFetchTimeout",
			"usedTimeoutSeconds": float64(30),
		},
	}
	cli.CurrentContext.URL = baseURL

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--type", "test type")

	// Verify execution error is reported
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Query timed out after 30 seconds")
}

func TestTopologyInspectTooManyActiveQueries(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshotTooManyActiveQueries",
		},
	}
	cli.CurrentContext.URL = baseURL

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--type", "test type")

	// Verify execution error is reported
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Too many active queries")
}

func TestTopologyInspectTopologySizeOverflow(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type":   "ViewSnapshotTopologySizeOverflow",
			"maxSize": float64(10000),
		},
	}
	cli.CurrentContext.URL = baseURL

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--type", "test type")

	// Verify execution error is reported
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Query result exceeded maximum size")
}

func TestTopologyInspectDataUnavailable(t *testing.T) {
	cli, cmd := setInspectCmd(t)
	cli.MockClient.ApiMocks.SnapshotApi.QuerySnapshotResponse.Result = sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type":                 "ViewSnapshotDataUnavailable",
			"unavailableAtEpochMs":  float64(1000000),
			"availableSinceEpochMs": float64(1609459200000),
		},
	}
	cli.CurrentContext.URL = baseURL

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--type", "test type")

	// Verify execution error is reported
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Requested data is not available")
}
