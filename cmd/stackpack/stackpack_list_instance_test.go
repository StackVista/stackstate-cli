package stackpack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

var (
	testName    = "ucmdb"
	testVersion = "0.1.1"
)

func setupStackpackListInstanceFn() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := StackpackListInstanceCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestStackpackListInstancePrintToTable(t *testing.T) {
	statusInstalled := "INSTALLED"
	id := int64(12345)

	unknownName := "unknown"
	expectedUpdateTime := util.TimeToString(time.UnixMilli(1438167001716))
	mockResponse := []stackstate_client.Sstackpack{
		{
			Name: &testName,
			Configurations: &[]stackstate_client.SstackpackConfigurations{
				{
					Id:                  &id,
					Status:              &statusInstalled,
					StackPackVersion:    &testVersion,
					LastUpdateTimestamp: 1438167001716,
				},
			},
		},
		{
			Name: &unknownName,
			Configurations: &[]stackstate_client.SstackpackConfigurations{
				{
					Id:                  &id,
					Status:              &statusInstalled,
					StackPackVersion:    &testVersion,
					LastUpdateTimestamp: 1438167001716,
				},
			},
		},
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(mockResponse); err != nil {
		t.Fatal(err)
	}
	cli, cmd := setupStackpackListInstanceFn()
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "list-instances", "--name", testName)
	expectedTableCall := []printer.TableCall{
		{
			Header: []string{"id", "status", "version", "last updated"},
			Data:   [][]string{{fmt.Sprintf("%d", id), statusInstalled, testVersion, expectedUpdateTime}},
			StructData: []stackstate_client.SstackpackConfigurations{
				{
					Id:                  &id,
					Status:              &statusInstalled,
					StackPackVersion:    &testVersion,
					LastUpdateTimestamp: 1438167001716,
				},
			},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
