package stackpack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommandFn() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := StackpackListCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestStackpackListPrintToTable(t *testing.T) {
	name := "ucmdb"
	displayName := "HP UCMDB"
	version := "0.1.1"
	mockResponse := []stackstate_client.Stackpack{
		{
			Name:        &name,
			DisplayName: &displayName,
			Version:     &version,
			Configurations: []stackstate_client.StackpackConfigurationsInner{
				{
					StackPackVersion: &version,
				},
			},
			NextVersion: &stackstate_client.StackpackLatestVersion{
				Version: &version,
			},
			LatestVersion: &stackstate_client.StackpackLatestVersion{
				Version: &version,
			},
		},
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(mockResponse); err != nil {
		t.Fatal(err)
	}
	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                [][]interface{}{{&name, &displayName, &version, version, version, 1}},
			StructData:          mockResponse,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListWithInstalledPrintToTable(t *testing.T) {
	name := "ucmdb"
	displayName := "HP UCMDB"
	version := "0.1.1"
	nextVersion := "0.1.2"
	latestVersion := "0.2.0"
	installedStackPack := stackstate_client.Stackpack{
		Name:        &name,
		DisplayName: &displayName,
		Version:     &version,
		Configurations: []stackstate_client.StackpackConfigurationsInner{
			{
				StackPackVersion: &version,
			},
		},
		NextVersion: &stackstate_client.StackpackLatestVersion{
			Version: &nextVersion,
		},
		LatestVersion: &stackstate_client.StackpackLatestVersion{
			Version: &latestVersion,
		},
	}
	mockResponse := []stackstate_client.Stackpack{
		installedStackPack,
		{
			Name:           &name,
			DisplayName:    &displayName,
			Version:        &version,
			Configurations: nil,
		},
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(mockResponse); err != nil {
		t.Fatal(err)
	}
	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "list", "--installed")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                [][]interface{}{{&name, &displayName, &version, nextVersion, latestVersion, 1}},
			StructData:          []stackstate_client.Stackpack{installedStackPack},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
