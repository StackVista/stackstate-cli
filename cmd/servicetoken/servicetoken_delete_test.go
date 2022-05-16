package servicetoken

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestDeleteShouldFailOnNonIntID(t *testing.T) {
	cli := di.NewMockDeps()
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "foo")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid id: foo", err.Error())
}

func TestDelete(t *testing.T) {
	cli := di.NewMockDeps()
	cmd := DeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1")

	assert.Len(t, *cli.MockClient.ApiMocks.ServiceTokenApi.DeleteServiceTokenCalls, 1)
	assert.Equal(t, int64(1), (*cli.MockClient.ApiMocks.ServiceTokenApi.DeleteServiceTokenCalls)[0].PserviceTokenId)

	assert.Equal(t, []string{"Service token deleted: 1"}, *cli.MockPrinter.SuccessCalls)
}
