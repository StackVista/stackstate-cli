package settings

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	name            = "One"
	otherName       = "Two"
	description     = "First component"
	owner           = "owner-1"
	identifier      = "identifier-1"
	otherIdentifier = "identifier-2"
	nodeApiResult   = []sts.Node{
		{
			Id:                  1,
			TypeName:            "ComponentType",
			Identifier:          &identifier,
			Name:                &name,
			Description:         &description,
			OwnedBy:             &owner,
			LastUpdateTimestamp: 1438167001716,
		},
		{
			Id:                  2,
			TypeName:            "ComponentType",
			Identifier:          &otherIdentifier,
			Name:                &otherName,
			Description:         &description,
			OwnedBy:             &owner,
			LastUpdateTimestamp: 1438167001716,
		},
	}
	expectedUpdateTime = time.UnixMilli(1438167001716)
)

func setupSettingsListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SettingsListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.NodeAPI.TypeListResponse.Result = nodeApiResult
	return &cli, cmd
}

func TestSettingsListPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}, {"ComponentType", int64(2), otherIdentifier, otherName, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type \"ComponentType\""},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListWithNamespaeAndOwnerPrintsToTable(t *testing.T) {
	cli, cmd := setupSettingsListCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType",
		"--namespace", "component",
		"--owned-by", "urn:stackpack:stackstate-self-health:shared")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}, {"ComponentType", int64(2), otherIdentifier, otherName, owner, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type \"ComponentType\""},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListPrintsToJson(t *testing.T) {
	cli, cmd := setupSettingsListCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--type", "ComponentType", "-o", "json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"settings": nodeApiResult,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
