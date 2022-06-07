package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestShowContextJSON(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := ShowCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--output", "json")
	assert.NoError(t, err)

	calls := *cli.MockPrinter.PrintJsonCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, "foo", calls[0]["name"])
	assert.Equal(t, cli.StsConfig.GetCurrentContext(), calls[0]["current-context"])
}

func TestShowContextText(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := ShowCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--output", "text")
	assert.NoError(t, err)

	calls := *cli.MockPrinter.PrintLnCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, "foo", calls[0])
}
