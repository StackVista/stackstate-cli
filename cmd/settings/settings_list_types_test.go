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

func setupSettingListTypesCmd() (di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := SettingsListTypesCommand(&cli.Deps)
	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypes{
			{TypeName: "hello", Description: "world"},
		},
	}
	cli.MockClient.ApiMocks.NodeApi.NodeListTypesResponse.Result = nodeApiResult
	return cli, cmd
}

func TestListTypesPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd()

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd)

	expectedTableCall := []printer.TableData{{
		Header:              []string{"name", "description"},
		Data:                [][]interface{}{{"hello", "world"}},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "setting types"},
	}}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestListTypesPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd()

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--json")

	expectedJsonCalls := []interface{}{cli.MockClient.ApiMocks.NodeApi.NodeListTypesResponse.Result}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
