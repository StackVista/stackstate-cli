package otelcomponentmapping_test

import (
	"errors"
	"io"
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelcomponentmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type mockEditorStub struct {
	content []byte
	error   error
}

func (m *mockEditorStub) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	if m.error != nil {
		return nil, m.error
	}
	return m.content, nil
}

func TestOtelComponentMappingEdit_SuccessNoChange(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:component:edit-test"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "edit-mapping",
		Identifier: mockIdentifier,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping
	// simulate the editor not changing the file
	marshal, _ := yaml.Marshal(mockMapping)
	// Simulate editing: set ReverseEditor content directly, as that's the default mock
	cli.Editor = &mockEditorStub{content: marshal}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 0)
	assert.Nil(t, err)
	assert.Greater(t, len(*cli.MockPrinter.PrintWarnCalls), 0)
	assert.Contains(t, (*cli.MockPrinter.PrintWarnCalls)[0], "No changes made")
}

func TestOtelComponentMappingEdit_MissingIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelComponentMappingEdit_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:component:edit-error"
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Error = errors.New("upsert failed")

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "upsert failed")
}

func TestOtelComponentMappingEdit_NoChangeJsonOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:component:edit-test"
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "edit-mapping",
		Identifier: mockIdentifier,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping
	marshal, _ := yaml.Marshal(mockMapping)
	cli.Editor = &mockEditorStub{content: marshal}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier, "--output", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 0)
	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, "No changes made", jsonOutput["message"])
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestOtelComponentMappingEdit_SuccessEdit(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:component:edit"
	// Original mapping
	origMapping := stackstate_api.OtelComponentMapping{
		Name:       "old-name",
		Identifier: mockIdentifier,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = origMapping
	// Simulate changing the name
	editedMapping := stackstate_api.OtelComponentMapping{
		Name:       "new-name",
		Identifier: mockIdentifier,
	}
	editedYaml, _ := yaml.Marshal(editedMapping)
	cli.Editor = &mockEditorStub{content: editedYaml}

	// Ensure the mock upsert returns the appropriate API type
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsResponse.Result = stackstate_api.OtelMappingItem{
		Name:       editedMapping.Name,
		Identifier: &editedMapping.Identifier,
	}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 1)
	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	successMsg := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successMsg, "OTel Component Mapping updated successfully!")
	assert.Contains(t, successMsg, mockIdentifier)
	assert.Contains(t, successMsg, "new-name")
}

func TestOtelComponentMappingEdit_SuccessJsonOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)
	oldUrn := "urn:otel:component:edit-test-success"
	// Original mapping
	mockMapping := stackstate_api.OtelComponentMapping{
		Name:       "orig-name",
		Identifier: oldUrn,
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelComponentMappingResponse.Result = mockMapping
	// Simulate edit with name change
	editedMapping := stackstate_api.OtelComponentMapping{
		Name:       "new-name",
		Identifier: oldUrn,
	}
	editedYaml, _ := yaml.Marshal(editedMapping)
	cli.Editor = &mockEditorStub{content: editedYaml}

	// Set the mock result with the right type
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsResponse.Result = stackstate_api.OtelMappingItem{
		Name:       editedMapping.Name,
		Identifier: &editedMapping.Identifier,
	}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", oldUrn, "--output", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 1)
	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	// The output should contain the otel_component_mapping key
	mappingVal, ok := jsonOutput["otel_component_mapping"]
	assert.True(t, ok)
	assert.NotNil(t, mappingVal)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestOtelComponentMappingEdit_EditorError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingEditCommand(&cli.Deps)

	mockEditor := &mockEditorStub{
		error: errors.New("editor failed to open"),
	}
	cli.Editor = mockEditor

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", "1234")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to open editor")

	// Verify no upsert call was made
	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelComponentMappingsCalls, 0)
}
