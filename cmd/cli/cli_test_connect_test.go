package cli

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommand() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := CliTestCommandCommand(&mockCli.Deps)
	return mockCli, cmd
}

func TestConnectionFailure(t *testing.T) {
	cli, cmd := setupCommand()
	connectError := common.NewConnectError(fmt.Errorf("authentication error"), &http.Response{StatusCode: 401})
	cli.MockClient.ConnectError = connectError
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd)
	assert.Equal(t, connectError, err)
}
