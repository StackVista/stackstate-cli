package settings

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"io/ioutil"
	"testing"
)

func setupCommandExport() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := SettingsExportCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestSettingsExportPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsExportIdsPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}},a{ "description": "description-1", "id": 314, "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--ids", "314")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsExportMutuallyExclusiveFlags(t *testing.T) {
	cli, cmd := setupCommandExport()
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd)

	assert.Equal(t, common.NewMutuallyExclusiveFlagsRequiredError([]string{Ids, Namespace, TypeName}), err)

	_, err = util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--namespace", "default")
	assert.Equal(t, common.NewMutuallyExclusiveFlagsMultipleError([]string{Ids, Namespace, TypeName}, []string{Ids, Namespace}), err)
}

func TestRunSettingsExportToFile(t *testing.T) {
	filePath := "./result.sjt"
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "description": "description-1", "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupCommandExport()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--ids", "-214", "--file", filePath)
	assert.Nil(t, err)
	body, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(body))
}
