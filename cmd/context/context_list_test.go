package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestContextListJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cli.StsConfig.CurrentContext = "foo"
	cli.StsConfig.Contexts = []*config.NamedContext{
		{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "/api")},
	}
	cmd := ListCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--output", "json")

	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	assert.Equal(t, cli.StsConfig.Contexts, (*cli.MockPrinter.PrintJsonCalls)[0]["contexts"])
}
