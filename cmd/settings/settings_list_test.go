package settings

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func setupCommandFn(mockNodeApi sts.NodeApiMock) (*printer.MockPrinter, di.Deps, *cobra.Command) {
	client := sts.APIClient{}
	client.NodeApi = mockNodeApi
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &conf.Conf{},
		Printer: &mockPrinter,
		Context: context.Background(),
		Client:  &client,
	}
	cmd := SettingsListCommand(&cli)

	return &mockPrinter, cli, cmd
}

func TestSettingsListPrintsToTable(t *testing.T) {
	nodeApiResult := []sts.NodeListType{
		{
			Id:                  1,
			TypeName:            "ComponentType",
			Name:                "One",
			Description:         "First component",
			OwnedBy:             "owner-1",
			LastUpdateTimestamp: 1438167001716,
		},
	}
	expectedUpdateTime := time.UnixMilli(1438167001716).Format("Mon Jan _2 15:04:05 2006")
	nodeApiMock := sts.NewNodeApiMock()
	nodeApiMock.TypeListResponse.Result = nodeApiResult
	nodeApiMock.TypeListResponse.Response = &http.Response{Body: ioutil.NopCloser(strings.NewReader("no body"))}
	mockPrinter, cli, cmd := setupCommandFn(nodeApiMock)

	util.ExecuteCommandWithContext(cli.Context, cmd, "--type", "ComponentType")

	expectedTableCall := printer.TableCall{
		Header:     []string{"id", "type", "name", "description", "owned by", "last updated"},
		Data:       [][]string{{"1", "ComponentType", "One", "First component", "owner-1", expectedUpdateTime}},
		StructData: "no body",
	}

	assert.Equal(t, expectedTableCall, (*mockPrinter.TableCalls)[0])
}
