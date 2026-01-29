package otelrelationmapping_test

import (
	"errors"
	"io"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

type mockEditorStub struct {
	content []byte
	err     error
}

func (m *mockEditorStub) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.content, nil
}

func TestOtelRelationMappingEdit_SuccessYAML(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:edit-success"
	orig := stackstate_api.OtelRelationMapping{
		Identifier: mockIdentifier,
		Name:       "orig",
	}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = orig
	// YAML of edited mapping
	editedYAML := "identifier: " + mockIdentifier + "\nname: changed\n"
	cli.Editor = &mockEditorStub{content: []byte(editedYAML)}
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Result = stackstate_api.OtelMappingItem{Name: "changed", Identifier: &mockIdentifier}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.Nil(t, err)
	assert.Greater(t, len(*cli.MockPrinter.SuccessCalls), 0)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], "changed")
}

func TestOtelRelationMappingEdit_SuccessJSON(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:edit-success-json"
	orig := stackstate_api.OtelRelationMapping{Identifier: mockIdentifier, Name: "orig-json"}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = orig
	editedYAML := "identifier: " + mockIdentifier + "\nname: changed-json\n"
	cli.Editor = &mockEditorStub{content: []byte(editedYAML)}
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Result = stackstate_api.OtelMappingItem{Name: "changed-json", Identifier: &mockIdentifier}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier, "-o", "json")

	assert.Nil(t, err)
	assert.Greater(t, len(*cli.MockPrinter.PrintJsonCalls), 0)
}

func TestOtelRelationMappingEdit_NoChange(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:edit-unchanged"
	orig := stackstate_api.OtelRelationMapping{Identifier: mockIdentifier, Name: "test-unchanged"}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = orig
	// Echo original as-is, using canonical YAML
	origYAML, _ := yaml.Marshal(orig)
	cli.Editor = &mockEditorStub{content: origYAML}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.Nil(t, err)
	assert.Contains(t, *cli.MockPrinter.PrintWarnCalls, "No changes made")
}

func TestOtelRelationMappingEdit_MissingIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelRelationMappingEdit_ApiGetError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:error-on-get"
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Error = errors.New("not found")

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.NotNil(t, err)
}

func TestOtelRelationMappingEdit_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:error-on-upsert"

	yamlContent := "identifier: " + mockIdentifier + "\nname: apierr\n"
	cli.Editor = &mockEditorStub{content: []byte(yamlContent)}
	cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsResponse.Error = errors.New("upsert failed")

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "upsert failed")
}

func TestOtelRelationMappingEdit_EditorError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:editorfail"
	orig := stackstate_api.OtelRelationMapping{Identifier: mockIdentifier, Name: "orig"}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = orig
	cli.Editor = &mockEditorStub{err: errors.New("editor failure")}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to open editor")

	// Verify no upsert call was made
	assert.Len(t, *cli.MockClient.ApiMocks.OtelMappingApi.UpsertOtelRelationMappingsCalls, 0)
}

func TestOtelRelationMappingEdit_InvalidYAML(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingEditCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:bad-yaml"
	orig := stackstate_api.OtelRelationMapping{Identifier: mockIdentifier, Name: "orig"}
	cli.MockClient.ApiMocks.OtelMappingApi.GetOtelRelationMappingResponse.Result = orig
	cli.Editor = &mockEditorStub{content: []byte("not: valid:: yaml::")}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to parse edited YAML")
}
