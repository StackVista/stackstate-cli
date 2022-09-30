package monitor

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupDescribeCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDescribeCommand(&cli.Deps)
	return &cli, cmd
}

func TestMonitorsDescribe(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common",
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd(t)
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "-i", "123")
	assert.Nil(t, err)
	assert.Equal(t, []string{expectedStr}, *cli.MockPrinter.PrintLnCalls)
}

func TestMonitorDescribeJson(t *testing.T) {
	expectedStr := `{"nodes": [{ "description": "description-1", "id": -214, "name": "name-1", "ownedBy": "urn:stackpack:common",
"parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd(t)
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr

	di.ExecuteCommandWithContext(&cli.Deps, cmd, "-i", "123", "-o", "json") //nolint:errcheck
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

func TestRunMonitorDescribeToFile(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_")
	if err != nil {
		panic(err)
	}
	filePath := file.Name()
	file.Close()

	expectedStr := `{"nodes": [{ "description": "description-1", "id": 123, "description": "description-1", "name": "name-1",
"ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}],
"script": { "scriptBody": "script-bdy-1"}}]}`
	cli, cmd := setupDescribeCmd(t)
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = expectedStr
	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-i", "123", "--file", filePath)
	assert.Nil(t, err)
	body, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(body))
}
