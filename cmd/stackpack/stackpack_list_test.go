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
	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:       [][]string{{name, displayName, "0.1.1", "0.1.1", "0.1.1", "1"}},
			StructData: mockResponse,
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListWithInstalledPrintToTable(t *testing.T) {
	name := "ucmdb"
	displayName := "HP UCMDB"
	version := "0.1.1"
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
			Version: &version,
		},
		LatestVersion: &stackstate_client.StackpackLatestVersion{
			Version: &version,
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
	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:       [][]string{{name, displayName, "0.1.1", "0.1.1", "0.1.1", "1"}},
			StructData: []stackstate_client.Stackpack{installedStackPack},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
