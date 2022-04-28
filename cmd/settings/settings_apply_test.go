package settings

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type Filename = string

func setupSettingsApplyCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := SettingsApplyCommand(&cli.Deps)

	cli.MockClient.ApiMocks.ImportApi.ImportSettingsResponse.Result = []map[string]interface{}{
		{
			"_type":      "Layer",
			"id":         12345,
			"identifier": "urn:stackpack:test:layer:test",
			"name":       "test",
		},
	}
	return &cli, cmd
}

func createTempFile() *os.File {
	file, err := ioutil.TempFile(os.TempDir(), "test_")
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	return file
}

func TestSetingsApplyFromFileNoOptionalArgs(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name())

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pbody, "hello world")
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pnamespace)
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].PtimeoutSeconds)
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Punlocked)
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Type", "Id", "Identifier", "Name"},
			Data:   [][]interface{}{{"Layer", 12345, "urn:stackpack:test:layer:test", "test"}},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSetingsApplyNamespace(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--namespace", "urn:test")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pnamespace, "urn:test")
}

func TestSetingsApplyWrongUnlockedStrategy(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--file", file.Name(), "--unlocked-strategy", "woopz")

	assert.Equal(t, common.NewCLIArgParseError(
		fmt.Errorf("invalid 'unlocked-strategy' flag value 'woopz' (must be { fail | skip | overwrite })")),
		err,
	)
}

func TestSetingsApplyUnlockedStrategyFail(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--unlocked-strategy", "fail")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Punlocked, "fail")
}

func TestSetingsApplyUnlockedTimeout(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--timeout", "16")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].PtimeoutSeconds, int64(16))
}

func TestSetingsApplyPrinJson(t *testing.T) {
	cli, cmd := setupSettingsApplyCmd()
	file := createTempFile()
	defer os.Remove(file.Name())

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--json")

	expectedJsonCalls := []interface{}{map[string]interface{}{
		"applied-settings": cli.MockClient.ApiMocks.ImportApi.ImportSettingsResponse.Result,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
