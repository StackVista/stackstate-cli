package settings

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func setupSettingListTypesCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SettingsListTypesCommand(&cli.Deps)
	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypes{
			{TypeName: "hello", Description: "world"},
		},
	}
	cli.MockClient.ApiMocks.NodeApi.NodeListTypesResponse.Result = nodeApiResult
	return &cli, cmd
}

func TestListTypesPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	expectedTableCall := []printer.TableData{{
		Header:              []string{"name", "description"},
		Data:                [][]interface{}{{"hello", "world"}},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "setting types"},
	}}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestListTypesPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"setting-types": cli.MockClient.ApiMocks.NodeApi.NodeListTypesResponse.Result,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
