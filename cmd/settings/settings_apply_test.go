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
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type Filename = string

func setupCmdSettingsApply() (di.MockDeps, *cobra.Command) {
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
	return cli, cmd
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
	cli, cmd := setupCmdSettingsApply()
	file := createTempFile()
	defer os.Remove(file.Name())

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--file", file.Name())

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pbody, "hello world")
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pnamespace)
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].PtimeoutSeconds)
	assert.Nil(t, (*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Punlocked)
	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"Type", "Id", "Identifier", "Name"},
			Data:       [][]string{{"Layer", "12345", "urn:stackpack:test:layer:test", "test"}},
			StructData: cli.MockClient.ApiMocks.ImportApi.ImportSettingsResponse.Result,
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestSetingsApplyNamespace(t *testing.T) {
	cli, cmd := setupCmdSettingsApply()
	file := createTempFile()
	defer os.Remove(file.Name())

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--file", file.Name(), "--namespace", "urn:test")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pnamespace, "urn:test")
}

func TestSetingsApplyWrongUnlockedStrategy(t *testing.T) {
	cli, cmd := setupCmdSettingsApply()
	file := createTempFile()
	defer os.Remove(file.Name())

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--file", file.Name(), "--unlocked-strategy", "woopz")

	assert.Equal(t, common.NewCLIArgParseError(
		fmt.Errorf("invalid 'unlocked-strategy' flag value 'woopz' (must be { fail | skip | overwrite })")),
		err,
	)
}

func TestSetingsApplyUnlockedStrategyFail(t *testing.T) {
	cli, cmd := setupCmdSettingsApply()
	file := createTempFile()
	defer os.Remove(file.Name())

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--file", file.Name(), "--unlocked-strategy", "fail")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Punlocked, "fail")
}
