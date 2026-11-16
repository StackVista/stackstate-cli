package otelcomponentmapping_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelcomponentmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func createTempYamlFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp(os.TempDir(), "test_componentmapping_*.yaml")
	if err != nil {
		panic(err)
	}

	_, err = tmpfile.Write([]byte(content))
	assert.NoError(t, err)
	tmpfile.Close()
	return tmpfile.Name()
}

func toOtelMappingItem(m stackstate_api.OtelComponentMapping) stackstate_api.OtelMappingItem {
	return stackstate_api.OtelMappingItem{
		Name:       m.Name,
		Identifier: &m.Identifier,
	}
}

func TestOtelComponentMappingApply_SuccessYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.OtelComponentMapping{
		Name:       "my-mapping",
		Identifier: "urn:otel:component:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)

	want := toOtelMappingItem(mapping)
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsResponse.Result = want

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 1)
	upsertCall := (*cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls)[0]
	assert.Equal(t, mapping.Identifier, upsertCall.PupsertOtelComponentMappingsRequest.Identifier)
	assert.Equal(t, mapping.Name, upsertCall.PupsertOtelComponentMappingsRequest.Name)
}

func TestOtelComponentMappingApply_SuccessJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.OtelComponentMapping{
		Name:       "my-mapping",
		Identifier: "urn:otel:component:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)

	want := toOtelMappingItem(mapping)
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsResponse.Result = want

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 1)

	output := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, output, 1)
	val, ok := output[0]["otel_component_mapping"]
	assert.True(t, ok)
	actualPtr, ok := val.(*stackstate_api.OtelMappingItem)
	if ok {
		assert.Equal(t, &want, actualPtr)
	} else {
		actualVal, ok := val.(stackstate_api.OtelMappingItem)
		require.True(t, ok, "PrintJsonCalls type is not OtelMappingItem or *OtelMappingItem")
		assert.Equal(t, want, actualVal)
	}
}

func TestOtelComponentMappingApply_FileNotFound(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	badPath := "/not/a/real/path.yaml"
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", badPath)
	}, "expected panic for missing file during apply command execution")
}

func TestOtelComponentMappingApply_WrongExtension(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	tmpFile := createTempYamlFile(t, "irrelevant-content")
	newPath := tmpFile + ".txt"
	err := os.Rename(tmpFile, newPath)
	assert.NoError(t, err)
	defer os.Remove(newPath)
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", newPath)
	}, "expected panic for wrong extension during apply command execution")
}

func TestOtelComponentMappingApply_InvalidYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	tmpFile := createTempYamlFile(t, "not: [valid, yaml,,,")
	defer os.Remove(tmpFile)
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)
	}, "expected panic for invalid YAML during apply command execution")
}

func TestOtelComponentMappingApply_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.OtelComponentMapping{
		Name:       "my-mapping",
		Identifier: "urn:otel:component:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)

	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsResponse.Error = errors.New("api failure")

	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)
	}, "expected panic for API error during apply command execution")
}
