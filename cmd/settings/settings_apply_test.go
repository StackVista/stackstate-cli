package settings

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type Filename = string

func setupCmdSettingsApply() (di.MockDeps, *cobra.Command) {

	mockCli := di.NewMockDeps()
	cmd := SettingsApplyCommand(&mockCli.Deps)

	return mockCli, cmd
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

func TestSetingsApplyFromFile(t *testing.T) {
	cli, cmd := setupCmdSettingsApply()
	file := createTempFile()
	defer os.Remove(file.Name())

	cli.MockClient.ApiMocks.ImportApi.ImportSettingsResponse.Result = []map[string]interface{}{
		{
			"_type":      "Layer",
			"id":         12345,
			"identifier": "urn:stackpack:test:layer:test",
			"name":       "test",
		},
	}

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
