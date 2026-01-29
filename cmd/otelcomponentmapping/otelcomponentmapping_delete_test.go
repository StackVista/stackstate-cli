package otelcomponentmapping_test

import (
	"testing"

	"github.com/stackvista/stackstate-cli/cmd/otelcomponentmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestOtelComponentMappingDelete_Success(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDeleteCommand(&cli.Deps)
	mockUrn := "urn:otel:component:delete-test"

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn)

	assert.Nil(t, err)
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], mockUrn)
}

func TestOtelComponentMappingDelete_MissingIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"identifier\" not set")
}

func TestOtelComponentMappingDelete_ApiError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDeleteCommand(&cli.Deps)
	mockUrn := "urn:otel:component:delete-error"

	cli.MockClient.ApiMocks.OtelMappingApi.DeleteOtelComponentMappingResponse.Error = assert.AnError

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn)
	assert.NotNil(t, err)
}

func TestOtelComponentMappingDelete_SuccessJsonOutput(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := otelcomponentmapping.OtelComponentMappingDeleteCommand(&cli.Deps)
	mockUrn := "urn:otel:component:delete-test-json"

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", mockUrn, "--output", "json")

	assert.Nil(t, err)
	expected := []map[string]interface{}{
		{"deleted_identifier": mockUrn},
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
