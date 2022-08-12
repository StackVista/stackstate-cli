package context

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupSaveCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SaveCommand(&cli.Deps)

	return &cli, cmd
}

func TestSaveNewContext(t *testing.T) {
	cli, cmd := setupSaveCmd(t)
	setupConfig(t, cli)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "baz", "--url", "http://baz.com", "--api-token", "my-token")
	assert.NoError(t, err)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "baz", cfg.CurrentContext)
	assert.Len(t, cfg.Contexts, 4)

	validateContext(t, cfg, cfg.CurrentContext, "http://baz.com", "my-token", "", "", "/api")
}

func TestSaveExistingContext(t *testing.T) {
	cli, cmd := setupSaveCmd(t)
	setupConfig(t, cli)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar", "--url", "http://bar.com", "--service-token", "my-token")
	assert.NoError(t, err)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "bar", cfg.CurrentContext)
	assert.Len(t, cfg.Contexts, 3)
	validateContext(t, cfg, cfg.CurrentContext, "http://bar.com", "", "my-token", "", "/api")
}

func validateContext(t *testing.T, cfg *config.Config, name string, url string, apiToken, serviceToken, k8sSAToken string, apiPath string) {
	ctx, err := cfg.GetContext(name)
	assert.NoError(t, err)
	assert.Equal(t, url, ctx.Context.URL)
	assert.Equal(t, apiToken, ctx.Context.APIToken)
	assert.Equal(t, serviceToken, ctx.Context.ServiceToken)
	assert.Equal(t, k8sSAToken, ctx.Context.K8sSAToken)
	assert.Equal(t, apiPath, ctx.Context.APIPath)
}

func TestNoSaveOnMissingTokens(t *testing.T) {
	cli, cmd := setupSaveCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar", "--url", "http://my-bar.com")
	assert.Errorf(t, err, "missing required argument: --api-token")

	// Should not have written config file
	assert.NoFileExists(t, cli.ConfigPath)
}
