package otelrelationmapping_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestOtelRelationMappingDescribe_YAMLOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:describe-test"
	mockResult := stackstate_api.OtelRelationMapping{
		Identifier: mockIdentifier,
		Name:       "test-relation",
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = mockResult

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.PrintLnCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], mockIdentifier)
}

func TestOtelRelationMappingDescribe_JsonOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:describe-test"
	mockResult := stackstate_api.OtelRelationMapping{
		Identifier: mockIdentifier,
		Name:       "test-relation",
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = mockResult

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier, "-o", "json")

	require.NoError(t, err)
	require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	// jsonOutput["data"] is a struct. Marshal to JSON and unmarshal to map for assertions
	jsonBytes, err := json.Marshal(jsonOutput["data"])
	require.NoError(t, err)
	var data map[string]interface{}
	require.NoError(t, json.Unmarshal(jsonBytes, &data))
	assert.Equal(t, mockResult.Name, data["name"])
	assert.Equal(t, mockResult.Identifier, data["identifier"])
}

func TestOtelRelationMappingDescribe_ToFileYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:file-test"
	mockResult := stackstate_api.OtelRelationMapping{
		Identifier: mockIdentifier,
		Name:       "file-yaml-relation",
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = mockResult

	file, err := os.CreateTemp(os.TempDir(), "otelrelationmapping_yaml_")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", mockIdentifier, "--file", file.Name())

	// Success message output
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], file.Name())

	body, err := os.ReadFile(file.Name())
	assert.NoError(t, err)
	var actualMapping stackstate_api.OtelRelationMapping
	assert.NoError(t, yaml.Unmarshal(body, &actualMapping))
	// Compare on relevant fields for robustness
	assert.Equal(t, mockResult.Name, actualMapping.Name)
	assert.Equal(t, mockResult.Identifier, actualMapping.Identifier)
}

func TestOtelRelationMappingDescribe_ToFileJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:file-test-json"
	mockResult := stackstate_api.OtelRelationMapping{
		Identifier: mockIdentifier,
		Name:       "file-json-relation",
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = mockResult

	file, err := os.CreateTemp(os.TempDir(), "otelrelationmapping_json_")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", mockIdentifier, "--file", file.Name(), "-o", "json")
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, file.Name(), jsonOutput["filepath"])

	body, err := os.ReadFile(file.Name())
	assert.NoError(t, err)
	var actualMapping stackstate_api.OtelRelationMapping
	assert.NoError(t, yaml.Unmarshal(body, &actualMapping))
	// Compare on relevant fields for robustness
	assert.Equal(t, mockResult.Name, actualMapping.Name)
	assert.Equal(t, mockResult.Identifier, actualMapping.Identifier)
}

func TestOtelRelationMappingDescribe_MissingIdentifierArg(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelRelationMappingDescribe_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDescribeCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:error"
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Error = assert.AnError
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier, "-o", "json")
	assert.NotNil(t, err)
}
