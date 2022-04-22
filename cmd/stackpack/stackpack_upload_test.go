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
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupStackPackUploadCmd() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := StackpackUploadCommand(&mockCli.Deps)

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

func TestUploadStackPackSuccess(t *testing.T) {
	cli, cmd := setupStackPackUploadCmd()
	file := createTempFile()

	name := "test"
	displayName := "display test"
	version := "1.0.0"
	stackpack := sts.Stackpack{
		Name:        &name,
		DisplayName: &displayName,
		Version:     &version,
	}
	cli.MockClient.ApiMocks.StackpackApi.StackpackUploadResponse.Result = stackpack

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--file", file.Name())

	expectedTableCall := []printer.TableCall{
		{
			Header:     []string{"name", "display name", "version"},
			Data:       [][]string{{"test", "display test", "1.0.0"}},
			StructData: stackpack,
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
