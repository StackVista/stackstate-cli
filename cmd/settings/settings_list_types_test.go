package settings

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	testWorld = "world"
	testHello = "hello"
)

func setupSettingListTypesCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SettingsListTypesCommand(&cli.Deps)
	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypesInner{
			{TypeName: testWorld, Description: testHello},
			{TypeName: testHello, Description: testWorld},
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
		Data:                [][]interface{}{{testHello, testWorld}, {testWorld, testHello}},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "setting types"},
	}}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestListTypesPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingListTypesCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	ordered := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypesInner{
			{TypeName: testHello, Description: testWorld},
			{TypeName: testWorld, Description: testHello},
		},
	}

	expectedJsonCalls := []map[string]interface{}{{
		"setting-types": ordered,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
