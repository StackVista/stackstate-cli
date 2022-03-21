package settings

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommand(mockNodeApi sts.NodeApiMock) (*printer.MockPrinter, di.Deps, *cobra.Command) {
	client := sts.APIClient{}
	client.NodeApi = mockNodeApi
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &conf.Conf{},
		Printer: &mockPrinter,
		Context: context.Background(),
		Client:  &client,
	}
	cmd := SettingsListTypesCommand(&cli)

	return &mockPrinter, cli, cmd
}

func TestListTypesPrintsToTable(t *testing.T) {
	nodeApiResult := sts.NodeTypes{
		NodeTypes: []sts.NodeTypesNodeTypes{
			{TypeName: "hello", Description: "world"},
		},
	}
	nodeApiMock := sts.NewNodeApiMock()
	nodeApiMock.NodeSettingsResponse = sts.NodeSettingsMockResponse{Result: nodeApiResult}
	mockPrinter, cli, cmd := setupCommand(nodeApiMock)

	util.ExecuteCommandWithContext(cli.Context, cmd)

	expectedTableCall := printer.TableCall{
		Header:     []string{"Name", "Description"},
		Data:       [][]string{{"hello", "world"}},
		StructData: nodeApiResult,
	}

	assert.Equal(t, expectedTableCall, (*mockPrinter.TableCalls)[0])
}
