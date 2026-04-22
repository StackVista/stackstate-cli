package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupStackPackDeleteVersionCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDeleteVersionCommand(&cli.Deps)
	return &cli, cmd
}

func TestStackpackDeleteVersionPrintToConsole(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-version", "--name", "kubernetes", "--stackpack-version", "1.0.0")

	expectedSuccessMessage := []string{"Successfully deleted version 1.0.0 of StackPack kubernetes"}
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedSuccessMessage, *cli.MockPrinter.SuccessCalls)

	assert.Equal(t, 1, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionCalls))
	call := (*cli.MockClient.ApiMocks.StackpackApi.StackPackDeleteVersionCalls)[0]
	assert.Equal(t, "kubernetes", call.PstackPackName)
	assert.Equal(t, "1.0.0", call.Pversion)
}

func TestStackpackDeleteVersionPrintToJson(t *testing.T) {
	cli, cmd := setupStackPackDeleteVersionCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete-version", "--name", "kubernetes", "--stackpack-version", "1.0.0", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"deleted": "1.0.0",
	}}
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
