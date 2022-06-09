package stackpack

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	testName        = "ucmdb"
	testVersion     = "0.1.1"
	statusInstalled = "INSTALLED"
	id              = int64(12345)

	unknownName          = "unknown"
	expectedUpdateTimeLi = time.UnixMilli(1438167001716)
)

func setupStackpackListInstanceFn(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackListInstanceCommand(&cli.Deps)

	mockResponse := []stackstate_api.Sstackpack{
		{
			Name: &testName,
			Configurations: []stackstate_api.SstackpackConfigurations{
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
			Configurations: []stackstate_api.SstackpackConfigurations{
				{
					Id:                  &id,
					Status:              &statusInstalled,
					StackPackVersion:    &testVersion,
					LastUpdateTimestamp: 1438167001716,
				},
			},
		},
	}
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse

	return &cli, cmd
}

func TestStackpackListInstancePrintToTable(t *testing.T) {
	cli, cmd := setupStackpackListInstanceFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-instances", "-n", testName)
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"id", "status", "version", "last updated"},
			Data:                [][]interface{}{{&id, &statusInstalled, &testVersion, expectedUpdateTimeLi}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListInstancePrintToJson(t *testing.T) {
	cli, cmd := setupStackpackListInstanceFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-instances", "--name", testName, "-o", "json")
	expectedJsonCalls := []map[string]interface{}{{
		"instances": []stackstate_api.SstackpackConfigurations{
			{
				Id:                  &id,
				Status:              &statusInstalled,
				StackPackVersion:    &testVersion,
				LastUpdateTimestamp: 1438167001716,
			}}},
	}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
