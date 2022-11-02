package stackpack

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	testName        = "ucmdb"
	testVersion     = "0.1.1"
	statusInstalled = "INSTALLED"
	id              = int64(12345)
	lut             = int64(1438167001716)
	lut2            = int64(1418167001716)

	unknownName          = "unknown"
	expectedUpdateTimeLi = time.UnixMilli(1438167001716)
	beforeUpdateTimeLi   = time.UnixMilli(1418167001716)
)

func setupStackpackListInstanceFn(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackListInstanceCommand(&cli.Deps)

	mockResponse := []stackstate_api.FullStackPack{
		{
			Name: testName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &id,
					Status:              statusInstalled,
					StackPackVersion:    testVersion,
					LastUpdateTimestamp: &lut,
				},
				{
					Id:                  &id,
					Status:              statusInstalled,
					StackPackVersion:    testVersion,
					LastUpdateTimestamp: &lut2,
				},
			},
		},
		{
			Name: unknownName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &id,
					Status:              statusInstalled,
					StackPackVersion:    testVersion,
					LastUpdateTimestamp: &lut,
				},
			},
		},
	}
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse

	return &cli, cmd
}

func TestStackpackListInstancePrintToTable(t *testing.T) {
	cli, cmd := setupStackpackListInstanceFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-instances", "-n", testName)
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"id", "status", "version", "last updated"},
			Data: [][]interface{}{
				{&id, statusInstalled, testVersion, expectedUpdateTimeLi},
				{&id, statusInstalled, testVersion, beforeUpdateTimeLi},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListInstancePrintToJson(t *testing.T) {
	cli, cmd := setupStackpackListInstanceFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-instances", "--name", testName, "-o", "json")
	expectedJsonCalls := []map[string]interface{}{{
		"instances": []stackstate_api.StackPackConfiguration{
			{
				Id:                  &id,
				Status:              statusInstalled,
				StackPackVersion:    testVersion,
				LastUpdateTimestamp: &lut,
			}, {
				Id:                  &id,
				Status:              statusInstalled,
				StackPackVersion:    testVersion,
				LastUpdateTimestamp: &lut2,
			}}},
	}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
