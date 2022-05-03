package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	ucmdbName       = "ucmdb"
	ucmdDisplayName = "HP UCMDB"
	ucmdbVersion    = "0.1.1"
	mockResponse    = []stackstate_api.Sstackpack{
		{
			Name:        &ucmdbName,
			DisplayName: &ucmdDisplayName,
			Version:     &ucmdbVersion,
			Configurations: &[]stackstate_api.SstackpackConfigurations{
				{
					StackPackVersion: &ucmdbVersion,
				},
			},
			NextVersion: &stackstate_api.SstackpackLatestVersion{
				Version: &ucmdbVersion,
			},
			LatestVersion: &stackstate_api.SstackpackLatestVersion{
				Version: &ucmdbVersion,
			},
		},
	}
)

func setupStackPackListCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackListCommand(&cli.Deps)
	return &cli, cmd
}

func TestStackpackListPrintToTable(t *testing.T) {
	cli, cmd := setupStackPackListCmd()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                [][]interface{}{{&ucmdbName, &ucmdDisplayName, &ucmdbVersion, ucmdbVersion, ucmdbVersion, 1}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackListCmd()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--json")
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
	installedStackPack := stackstate_api.Sstackpack{
		Name:        &name,
		DisplayName: &displayName,
		Version:     &version,
		Configurations: &[]stackstate_api.SstackpackConfigurations{
			{
				StackPackVersion: &version,
			},
		},
		NextVersion: &stackstate_api.SstackpackLatestVersion{
			Version: &nextVersion,
		},
		LatestVersion: &stackstate_api.SstackpackLatestVersion{
			Version: &latestVersion,
		},
	}
	mockResponse := []stackstate_api.Sstackpack{
		installedStackPack,
		{
			Name:           &name,
			DisplayName:    &displayName,
			Version:        &version,
			Configurations: nil,
		},
	}
	cli, cmd := setupStackPackListCmd()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--installed")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                [][]interface{}{{&name, &displayName, &version, nextVersion, latestVersion, 1}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
