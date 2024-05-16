package stackpack

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func setupStackPackUploadCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUploadCommand(&cli.Deps)
	return &cli, cmd
}

func createTempFile() *os.File {
	file, err := os.CreateTemp(os.TempDir(), "test_")
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
	cli, cmd := setupStackPackUploadCmd(t)
	file := createTempFile()

	name := "test"
	displayName := "display test"
	version := "1.0.0"
	stackpack := sts.StackPack{
		Name:        name,
		DisplayName: displayName,
		Version:     version,
	}
	cli.MockClient.ApiMocks.StackpackAPI.StackPackUploadResponse.Result = stackpack

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
	cli, cmd := setupStackPackUploadCmd(t)
	file := createTempFile()

	stackpack := &sts.StackPack{
		Name:        "test",
		DisplayName: "display test",
		Version:     "1.0.0",
	}
	cli.MockClient.ApiMocks.StackpackAPI.StackPackUploadResponse.Result = *stackpack

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "-o", "json")

	assert.Equal(
		t,
		[]map[string]interface{}{{
			"uploaded-stackpack": stackpack,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
