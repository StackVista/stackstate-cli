package monitor

import (
	"testing"

	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestMonitorEnableById(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been enabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorEnableByIdJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	expectedJsonCalls := []map[string]interface{}{{
		"monitor": &monitor,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestMonitorEnableByIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been enabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorEnableByIdentifierJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	expectedJsonCalls := []map[string]interface{}{{
		"monitor": &monitor,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
