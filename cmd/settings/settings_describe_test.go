package settings

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

func setupDescribeCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := SettingsDescribeCommand(&cli.Deps)
	return &cli, cmd
}

func TestSettingsDescribe(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", 
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--ids", "-214")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsDescribeJson(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", 
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	di.ExecuteCommandWithContext(&cli.Deps, cmd, "--ids", "-214", "--json") //nolint:errcheck
	assert.Equal(
		t,
		[]map[string]interface{}{{
			"data":   expectedStr,
			"format": "stj",
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestSettingsDescribeIds(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common", 
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}},
{ "description": "description-1", "id": 314, "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": 
[{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--ids", "-214", "--ids", "314")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestSettingsDescribeMutuallyExclusiveFlags(t *testing.T) {
	cli, cmd := setupDescribeCmd()
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)

	assert.Equal(t, mutex_flags.NewMutuallyExclusiveFlagsRequiredError([]string{Ids, Namespace, TypeName}).Error(), err.Error())

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--ids", "-214", "--namespace", "default")
	assert.Equal(t, mutex_flags.NewMutuallyExclusiveFlagsMultipleError([]string{Ids, Namespace, TypeName}, []string{Ids, Namespace}).Error(), err.Error())
}

func TestRunSettingsDescribeWithReferencePrintToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, 
"identifier": "urn:stackpack:common:baseline-function:median-absolute-deviation", "name": "name-1", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}],
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--namespace", "default",
		"--allowed-namespace-refs", "urn:stackpack:common:baseline-function:median-absolute-deviation", "--allowed-namespace-refs", "urn:stackpack")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestRunSettingsDescribeToFile(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "test_")
	if err != nil {
		panic(err)
	}
	filePath := file.Name()
	file.Close()

	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "description": "description-1", "name": "name-1",
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], 
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr
	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--ids", "-214", "--file", filePath)
	assert.Nil(t, err)
	body, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(body))
}

func TestRunSettingsDescribeTypesPrintsToTable(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "Base line function", "id": -214, "name": "name-1", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "BaselineFunction"}], 
"script": { "scriptBody": "script-bdy-1"}},{ "description": "Check function", "id": 314, "name": "name 314", 
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "CheckFunction"}], 
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd()
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--type", "BaselineFunction", "--type", "CheckFunction")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}
