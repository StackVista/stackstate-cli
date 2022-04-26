package settings

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

var (
	name          = "One"
	description   = "First component"
	owner         = "owner-1"
	identifier    = "identifier-1"
	nodeApiResult = []sts.Node{
		{
			Id:                  1,
			TypeName:            "ComponentType",
			Identifier:          &identifier,
			Name:                &name,
			Description:         &description,
			OwnedBy:             &owner,
			LastUpdateTimestamp: 1438167001716,
		},
	}
	expectedUpdateTime = time.UnixMilli(1438167001716)
)

func setupSettingsListCmd() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := SettingsListCommand(&mockCli.Deps)
	mockCli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult

	return mockCli, cmd
}

func TestSettingsListPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--type", "ComponentType")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type ComponentType"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListWithNamespaeAndOwnerPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--type", "ComponentType",
		"-n", "component",
		"-w", "urn:stackpack:stackstate-self-health:shared")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type ComponentType"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--type", "ComponentType", "--json")

	assert.Equal(t, []interface{}{nodeApiResult}, *cli.MockPrinter.PrintJsonCalls)
}
