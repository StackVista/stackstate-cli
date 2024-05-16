package settings

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func setupSettingListTypesCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SettingsListTypesCommand(&cli.Deps)
	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypesInner{
			{TypeName: "world", Description: "hello"},
			{TypeName: "hello", Description: "world"},
		},
	}
	cli.MockClient.ApiMocks.NodeAPI.NodeListTypesResponse.Result = nodeApiResult
	return &cli, cmd
}

func TestListTypesPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	expectedTableCall := []printer.TableData{{
		Header:              []string{"name", "description"},
		Data:                [][]interface{}{{"hello", "world"}, {"world", "hello"}},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "setting types"},
	}}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestListTypesPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	ordered := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypesInner{
			{TypeName: "hello", Description: "world"},
			{TypeName: "world", Description: "hello"},
		},
	}

	expectedJsonCalls := []map[string]interface{}{{
		"setting-types": ordered,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
