package stackpack

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func setupStackPackUploadCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackUploadCommand(&cli.Deps)
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

func TestUploadStackPackPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPackUploadCmd()
	file := createTempFile()

	stackpack := sts.StackPack{
		Name:        "test",
		DisplayName: "display test",
		Version:     "1.0.0",
	}
	cli.MockClient.ApiMocks.StackpackApi.StackpackUploadResponse.Result = stackpack

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name())

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"name", "display name", "version"},
			Data:   [][]interface{}{{"test", "display test", "1.0.0"}},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestUploadStackPackPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackUploadCmd()
	file := createTempFile()

	stackpack := sts.StackPack{
		Name:        "test",
		DisplayName: "display test",
		Version:     "1.0.0",
	}
	cli.MockClient.ApiMocks.StackpackApi.StackpackUploadResponse.Result = stackpack

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--json")

	assert.Equal(
		t,
		[]interface{}{map[string]interface{}{
			"uploaded-stackpack": stackpack,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
