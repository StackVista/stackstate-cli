package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupStackPackDeleteVersionsCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDeleteVersionsCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsResponse.Result = stackstate_api.DeleteVersionsResult{
		Deleted:      []string{"1.0.0", "1.1.0"},
		SkippedInUse: []string{"1.2.0"},
	}
	return &cli, cmd
}

func TestStackpackDeleteVersionsWithTo(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--to", "1.5.0")

	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, []string{"Deleted versions: 1.0.0, 1.1.0"}, *cli.MockPrinter.SuccessCalls)

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls))
	call := (*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls)[0]
	assert.Equal(t, "kubernetes", call.PstackPackName)
	assert.Nil(t, call.Pfrom)
	assert.Equal(t, "1.5.0", *call.Pto)
	assert.Nil(t, call.Pall)
	assert.Nil(t, call.Pdev)
}

func TestStackpackDeleteVersionsWithRange(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--from", "1.0.0", "--to", "1.5.0")

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls))
	call := (*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls)[0]
	assert.Equal(t, "kubernetes", call.PstackPackName)
	assert.Equal(t, "1.0.0", *call.Pfrom)
	assert.Equal(t, "1.5.0", *call.Pto)
	assert.Nil(t, call.Pall)
	assert.Nil(t, call.Pdev)
}

func TestStackpackDeleteVersionsWithAll(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--all")

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls))
	call := (*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls)[0]
	assert.Equal(t, "kubernetes", call.PstackPackName)
	assert.Nil(t, call.Pfrom)
	assert.Nil(t, call.Pto)
	assert.Equal(t, true, *call.Pall)
	assert.Nil(t, call.Pdev)
}

func TestStackpackDeleteVersionsWithAllAndDevOnly(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--all", "--dev-only")

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls))
	call := (*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionsCalls)[0]
	assert.Equal(t, "kubernetes", call.PstackPackName)
	assert.Nil(t, call.Pfrom)
	assert.Nil(t, call.Pto)
	assert.Equal(t, true, *call.Pall)
	assert.Equal(t, true, *call.Pdev)
}

func TestStackpackDeleteVersionsPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--to", "2.0.0", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"deleted":      []string{"1.0.0", "1.1.0"},
		"skippedInUse": []string{"1.2.0"},
	}}
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}

func TestStackpackDeleteVersionsNoFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDeleteVersionsCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes")
	assert.Error(t, err)
}

func TestStackpackDeleteVersionsAllWithFrom(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDeleteVersionsCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--all", "--from", "1.0.0")
	assert.Error(t, err)
}

func TestStackpackDeleteVersionsAllWithTo(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDeleteVersionsCommand(&cli.Deps)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "delete-versions", "--name", "kubernetes", "--all", "--to", "1.5.0")
	assert.Error(t, err)
}
