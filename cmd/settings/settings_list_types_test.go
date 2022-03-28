package settings

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommand() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := SettingsListTypesCommand(&mockCli.Deps)
	return mockCli, cmd
}

func TestListTypesPrintsToTable(t *testing.T) {
	cli, cmd := setupCommand()

	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypes{
			{TypeName: "hello", Description: "world"},
		},
	}
	cli.MockClient.ApiMocks.NodeApi.NodeListTypesResponse.Result = nodeApiResult

	util.ExecuteCommandWithContext(cli.Context, cmd)

	expectedTableCall := []printer.TableCall{{
		Header:     []string{"name", "description"},
		Data:       [][]string{{"hello", "world"}},
		StructData: nodeApiResult,
	}}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
