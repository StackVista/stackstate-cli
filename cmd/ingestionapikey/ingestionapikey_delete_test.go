package ingestionapikey

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestDeleteShouldFailOnNonIntID(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "foo")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid argument \"foo\" for \"-i, --id\"")
}

func TestDelete(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1")

	assert.Len(t, *cli.MockClient.ApiMocks.IngestionApiKeyApi.DeleteIngestionApiKeyCalls, 1)
	assert.Equal(t, int64(1), (*cli.MockClient.ApiMocks.IngestionApiKeyApi.DeleteIngestionApiKeyCalls)[0].PingestionApiKeyId)

	assert.Equal(t, []string{"Ingestion Api Key deleted: 1"}, *cli.MockPrinter.SuccessCalls)
}
