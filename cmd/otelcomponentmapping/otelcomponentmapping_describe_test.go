package otelcomponentmapping_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelcomponentmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func deepNormalizeOtelComponentMapping(m *stackstate_api.OtelComponentMapping) {
	if m == nil {
		return
	}
	if m.Input.Signal == nil {
		m.Input.Signal = []stackstate_api.OtelInputSignal{}
	}
	if m.Vars == nil {
		m.Vars = []stackstate_api.OtelVariableMapping{}
	}
}

func TestOtelComponentMappingDescribe_YAMLOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	mockUrn := "urn:otel:component:cli-test"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "cli-yaml-test",
		Identifier: mockUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn)
	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.PrintLnCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], mockUrn)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], "cli-yaml-test")
}

func TestOtelComponentMappingDescribe_JSONOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	mockUrn := "urn:otel:component:cli-json-test"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "cli-json-test",
		Identifier: mockUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn, "-o", "json")
	require.NoError(t, err)

	require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]

	// jsonOutput["data"] is a struct, so marshal to JSON and unmarshal to map for assertions
	jsonBytes, err := json.Marshal(jsonOutput["data"])
	require.NoError(t, err)
	var data map[string]interface{}
	require.NoError(t, json.Unmarshal(jsonBytes, &data))
	assert.Equal(t, mockMapping.Name, data["name"])
	assert.Equal(t, mockMapping.Identifier, data["identifier"])
}

func TestOtelComponentMappingDescribe_ToFileYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	mockUrn := "urn:otel:component:file-test"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "file-yaml-test",
		Identifier: mockUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping

	file, err := os.CreateTemp(os.TempDir(), "otelmapping_yaml_")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", mockUrn, "--file", file.Name())

	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], file.Name())

	body, err := os.ReadFile(file.Name())
	assert.NoError(t, err)
	var actualMapping stackstate_api.OtelComponentMapping
	assert.NoError(t, yaml.Unmarshal(body, &actualMapping))
	deepNormalizeOtelComponentMapping(&mockMapping)
	deepNormalizeOtelComponentMapping(&actualMapping)
	assert.True(t, reflect.DeepEqual(mockMapping, actualMapping))
}

func TestOtelComponentMappingDescribe_ToFileJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	mockUrn := "urn:otel:component:file-test-json"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "file-json-test",
		Identifier: mockUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping

	file, err := os.CreateTemp(os.TempDir(), "otelcomponentmapping_json_")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", mockUrn, "--file", file.Name(), "-o", "json")

	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, file.Name(), jsonOutput["filepath"])

	body, err := os.ReadFile(file.Name())
	assert.NoError(t, err)
	var actualMapping stackstate_api.OtelComponentMapping
	assert.NoError(t, yaml.Unmarshal(body, &actualMapping))
	deepNormalizeOtelComponentMapping(&mockMapping)
	deepNormalizeOtelComponentMapping(&actualMapping)
	assert.True(t, reflect.DeepEqual(mockMapping, actualMapping))
}

func TestOtelComponentMappingDescribe_MissingIdentifierFlag(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd) // no --identifier flag
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelComponentMappingDescribe_FileError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDescribeCommand(&cli.Deps)
	mockUrn := "urn:otel:component:file-error"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "file-error-test",
		Identifier: mockUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping

	// Attempt to write to a directory (should error)
	dir := os.TempDir()
	badPath := filepath.Join(dir, "not_a_real_dir", "cannot_write_here.yaml")
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn, "--file", badPath)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), badPath)
}
