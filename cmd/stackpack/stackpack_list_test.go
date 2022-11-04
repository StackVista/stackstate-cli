package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	ucmdbName       = "ucmdb"
	ucmdDisplayName = "HP UCMDB"
	ucmdbVersion    = "0.1.1"
	awsName         = "aws"
	awsStackPack    = stackstate_api.FullStackPack{
		Name:        awsName,
		DisplayName: ucmdDisplayName,
		Version:     ucmdbVersion,
		Configurations: []stackstate_api.StackPackConfiguration{
			{
				StackPackVersion: ucmdbVersion,
			},
		},
		NextVersion: &stackstate_api.FullStackPack{
			Version: ucmdbVersion,
		},
		LatestVersion: &stackstate_api.FullStackPack{
			Version: ucmdbVersion,
		},
	}
	ucmdbStackPack = stackstate_api.FullStackPack{
		Name:        ucmdbName,
		DisplayName: ucmdDisplayName,
		Version:     ucmdbVersion,
		Configurations: []stackstate_api.StackPackConfiguration{
			{
				StackPackVersion: ucmdbVersion,
			},
		},
		NextVersion: &stackstate_api.FullStackPack{
			Version: ucmdbVersion,
		},
		LatestVersion: &stackstate_api.FullStackPack{
			Version: ucmdbVersion,
		},
	}
	mockResponse = []stackstate_api.FullStackPack{ucmdbStackPack, awsStackPack}
)

func setupStackPackListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackListCommand(&cli.Deps)
	return &cli, cmd
}

func TestStackpackListPrintToTable(t *testing.T) {
	cli, cmd := setupStackPackListCmd(t)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data: [][]interface{}{
				{awsName, ucmdDisplayName, ucmdbVersion, ucmdbVersion, ucmdbVersion, 1},
				{ucmdbName, ucmdDisplayName, ucmdbVersion, ucmdbVersion, ucmdbVersion, 1}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackListCmd(t)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "-o", "json")
	expectedJsonCalls := []map[string]interface{}{{
		"stackpacks": mockResponse,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestStackpackListWithInstalledPrintToTable(t *testing.T) {
	name := "ucmdb"
	displayName := "HP UCMDB"
	version := "0.1.1"
	nextVersion := "0.1.2"
	latestVersion := "0.2.0"
	installedStackPack := stackstate_api.FullStackPack{
		Name:        name,
		DisplayName: displayName,
		Version:     version,
		Configurations: []stackstate_api.StackPackConfiguration{
			{
				StackPackVersion: version,
			},
		},
		NextVersion: &stackstate_api.FullStackPack{
			Version: nextVersion,
		},
		LatestVersion: &stackstate_api.FullStackPack{
			Version: latestVersion,
		},
	}
	mockResponse := []stackstate_api.FullStackPack{
		installedStackPack,
		{
			Name:           name,
			DisplayName:    displayName,
			Version:        version,
			Configurations: nil,
		},
	}
	cli, cmd := setupStackPackListCmd(t)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--installed")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                [][]interface{}{{name, displayName, version, nextVersion, latestVersion, 1}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
