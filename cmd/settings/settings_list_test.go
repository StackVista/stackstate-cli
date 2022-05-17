package settings

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
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

func setupSettingsListCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := SettingsListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult
	return &cli, cmd
}

func TestSettingsListPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type \"ComponentType\""},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListWithNamespaeAndOwnerPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType",
		"--namespace", "component",
		"--owned-by", "urn:stackpack:stackstate-self-health:shared")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type \"ComponentType\""},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingsListCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType", "--json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"settings": nodeApiResult,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
