package settings

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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
	expectedUpdateTime := time.UnixMilli(1438167001716).Format("Mon Jan _2 15:04:05 2006")
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode([]map[string]interface{}{{"name": "ms_iis_ws"}}); err != nil {
		t.Fatal(err)
	}

	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult

	util.ExecuteCommandWithContext(cli.Context, cmd, "--type", "ComponentType")

	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"Type", "Id", "Identifier", "Name", "description", "owned by", "last updated"},
			Data:       [][]string{{"ComponentType", "1", identifier, name, description, owner, expectedUpdateTime}},
			StructData: []map[string]interface{}{{"name": "ms_iis_ws"}},
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
	expectedUpdateTime := time.UnixMilli(1438167001716).Format("Mon Jan _2 15:04:05 2006")
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode([]map[string]interface{}{{"name": "ms_iis_ws"}}); err != nil {
		t.Fatal(err)
	}

	cli, cmd := setupCommandFn()
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.NodeApi.TypeListResponse.Result = nodeApiResult

	util.ExecuteCommandWithContext(cli.Context, cmd, "--type", "ComponentType",
		"-n", "component",
		"-w", "urn:stackpack:stackstate-self-health:shared")

	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"Type", "Id", "Identifier", "Name", "description", "owned by", "last updated"},
			Data:       [][]string{{"ComponentType", "1", identifier, name, description, owner, expectedUpdateTime}},
			StructData: []map[string]interface{}{{"name": "ms_iis_ws"}},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
