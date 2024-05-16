package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	nodeResult = stackstate_api.Node{
		TypeName:            "",
		Id:                  1234,
		LastUpdateTimestamp: 0,
		Identifier:          nil,
		Name:                nil,
		Description:         nil,
		OwnedBy:             nil,
		Manual:              nil,
		IsSettingsNode:      nil,
	}
)

func setMonitorCloneCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorCloneCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldCloneMonitor(t *testing.T) {
	cli, cmd := setMonitorCloneCmd(t)
	cli.MockClient.ApiMocks.NodeAPI.CloneResponse.Result = nodeResult
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--name", "newName")
	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.CloneCalls, 1)
	assert.Equal(t, "Monitor 1234 has been created", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldCloneMonitorWithJson(t *testing.T) {
	cli, cmd := setMonitorCloneCmd(t)
	cli.MockClient.ApiMocks.NodeAPI.CloneResponse.Result = nodeResult
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--name", "newName", "-o", "json")
	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.CloneCalls, 1)
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
}
