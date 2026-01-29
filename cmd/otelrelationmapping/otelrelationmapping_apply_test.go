package otelrelationmapping_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func createTempYamlFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp(os.TempDir(), "test_relationmapping_*.yaml")
	if err != nil {
		panic(err)
	}

	_, err = tmpfile.Write([]byte(content))
	assert.NoError(t, err)
	tmpfile.Close()
	return tmpfile.Name()
}

func toOtelMappingItem(m stackstate_api.UpsertOtelRelationMappingsRequest) stackstate_api.OtelMappingItem {
	return stackstate_api.OtelMappingItem{
		Name:       m.Name,
		Identifier: &m.Identifier,
	}
}

func TestOtelRelationMappingApply_SuccessYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.UpsertOtelRelationMappingsRequest{
		Name:       "my-relation-mapping",
		Identifier: "urn:otel:relation:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)

	want := toOtelMappingItem(mapping)
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Result = want

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsCalls, 1)
	upsertCall := (*cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsCalls)[0]
	assert.Equal(t, mapping.Identifier, upsertCall.PupsertOtelRelationMappingsRequest.Identifier)
	assert.Equal(t, mapping.Name, upsertCall.PupsertOtelRelationMappingsRequest.Name)
}

func TestOtelRelationMappingApply_SuccessJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.UpsertOtelRelationMappingsRequest{
		Name:       "my-relation-mapping",
		Identifier: "urn:otel:relation:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)

	want := toOtelMappingItem(mapping)
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Result = want

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsCalls, 1)
	output := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, output, 1)
	val, ok := output[0]["otel_relation_mapping"]
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

func TestOtelRelationMappingApply_FileNotFound(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingApplyCommand(&cli.Deps)
	badPath := "/not/a/real/path.yaml"
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", badPath)
	}, "expected panic for missing file during apply command execution")
}

func TestOtelRelationMappingApply_InvalidYaml(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingApplyCommand(&cli.Deps)
	invalidYaml := "not: valid:: yaml::"
	tmpFile := createTempYamlFile(t, invalidYaml)
	defer os.Remove(tmpFile)
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)
	}, "expected panic for invalid YAML during apply command execution")
}

func TestOtelRelationMappingApply_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingApplyCommand(&cli.Deps)
	mapping := stackstate_api.UpsertOtelRelationMappingsRequest{
		Name:       "my-relation-mapping",
		Identifier: "urn:otel:relation:foo",
	}
	yamlContent, err := yaml.Marshal(mapping)
	assert.NoError(t, err)
	tmpFile := createTempYamlFile(t, string(yamlContent))
	defer os.Remove(tmpFile)
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Error = errors.New("api error")
	require.Panics(t, func() {
		di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-f", tmpFile)
	}, "expected panic for API error during apply command execution")
}
