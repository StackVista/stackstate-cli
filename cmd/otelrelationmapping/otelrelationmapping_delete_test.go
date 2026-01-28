package otelrelationmapping_test

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestOtelRelationMappingDelete_YAMLOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDeleteCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:delete-test"
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)
	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], mockIdentifier)
}

func TestOtelRelationMappingDelete_JsonOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDeleteCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:delete-json"

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier, "-o", "json")

	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	json := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, mockIdentifier, json["deleted_identifier"])
}

func TestOtelRelationMappingDelete_MissingIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDeleteCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelRelationMappingDelete_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelrelationmapping.OtelRelationMappingDeleteCommand(&cli.Deps)
	mockIdentifier := "urn:otel:relation:delete-error"
	cli.MockClient.ApiMocks.OtelMappingApi.DeleteOtelRelationMappingResponse.Error = assert.AnError
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockIdentifier)
	assert.NotNil(t, err)
}
