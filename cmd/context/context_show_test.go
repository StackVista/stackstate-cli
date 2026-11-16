package context

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
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

	assert.NoError(t, err)
	assert.Equal(t, &config.StsContext{
		URL:      "http://foo.com",
		APIToken: "apiToken",
		APIPath:  "/api",
	}, calls[0]["current-context"])
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
