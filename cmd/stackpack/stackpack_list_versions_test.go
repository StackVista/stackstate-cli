package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var listVersionsMockResponse = []stackstate_api.StackPackVersionInfo{
	{Version: "1.0.0", IsDev: false, IsInUse: true},
	{Version: "1.1.0", IsDev: false, IsInUse: false},
	{Version: "2.0.0-SNAPSHOT.1", IsDev: true, IsInUse: false},
}

func setupStackPackListVersionsCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackListVersionsCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListVersionsResponse.Result = listVersionsMockResponse
	return &cli, cmd
}

func TestStackpackListVersionsPrintToTable(t *testing.T) {
	cli, cmd := setupStackPackListVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-versions", "--name", "kubernetes")

	assert.Equal(t, []printer.TableData{
		{
			Header: []string{"version", "dev", "in use"},
			Data: [][]interface{}{
				{"1.0.0", false, true},
				{"1.1.0", false, false},
				{"2.0.0-SNAPSHOT.1", true, false},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "versions"},
		},
	}, *cli.MockPrinter.TableCalls)

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackListVersionsCalls))
	assert.Equal(t, "kubernetes", (*cli.MockClient.ApiMocks.StackpackApi.StackPackListVersionsCalls)[0].PstackPackName)
}

func TestStackpackListVersionsPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackListVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-versions", "--name", "kubernetes", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"versions": listVersionsMockResponse,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestStackpackListVersionsEmpty(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackListVersionsCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListVersionsResponse.Result = []stackstate_api.StackPackVersionInfo{}
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-versions", "--name", "kubernetes")

	assert.Equal(t, []printer.TableData{
		{
			Header:              []string{"version", "dev", "in use"},
			Data:                [][]interface{}{},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "versions"},
		},
	}, *cli.MockPrinter.TableCalls)
}
