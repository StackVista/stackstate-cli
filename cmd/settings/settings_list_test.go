package settings

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommandFn() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := SettingsListCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestSettingsListPrintsToTable(t *testing.T) {
	name := "One"
	description := "First component"
	owner := "owner-1"
	identifier := "identifier-1"
	nodeApiResult := []sts.Node{
		{
			Id:                  1,
			TypeName:            "ComponentType",
			Identifier:          &identifier,
			Name:                &name,
			Description:         &description,
			OwnedBy:             &owner,
			LastUpdateTimestamp: 1438167001716,
		},
	}
	expectedUpdateTime := time.UnixMilli(1438167001716)
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode([]map[string]interface{}{{"name": "ms_iis_ws"}}); err != nil {
		t.Fatal(err)
	}

	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--type", "ComponentType")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			StructData:          []map[string]interface{}{{"name": "ms_iis_ws"}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type ComponentType"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSettingsListWithNamespaeAndOwnerPrintsToTable(t *testing.T) {
	name := "One"
	description := "First component"
	owner := "owner-1"
	identifier := "urn:stackpack:stackstate-self-health:shared"
	nodeApiResult := []sts.Node{
		{
			Id:                  1,
			TypeName:            "ComponentType",
			Identifier:          &identifier,
			Name:                &name,
			Description:         &description,
			OwnedBy:             &owner,
			LastUpdateTimestamp: 1438167001716,
		},
	}
	expectedUpdateTime := time.UnixMilli(1438167001716)
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode([]map[string]interface{}{{"name": "ms_iis_ws"}}); err != nil {
		t.Fatal(err)
	}

	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--type", "ComponentType",
		"-n", "component",
		"-w", "urn:stackpack:stackstate-self-health:shared")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                [][]interface{}{{"ComponentType", int64(1), identifier, name, owner, expectedUpdateTime}},
			StructData:          []map[string]interface{}{{"name": "ms_iis_ws"}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "settings of type ComponentType"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
