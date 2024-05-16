package monitor

import (
	"testing"

	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestMonitorDisableById(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_DISABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been disabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorDisableByIdJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_DISABLED.Ptr())
	expectedJsonCalls := []map[string]interface{}{{
		"monitor": &monitor,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestMonitorDisableByIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_DISABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been disabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorDisableByIdentifierJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_DISABLED.Ptr())
	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorAPI.PatchMonitorCalls, sts.MONITORSTATUSVALUE_DISABLED.Ptr())
	expectedJsonCalls := []map[string]interface{}{{
		"monitor": &monitor,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func checkMonitorPatchStatusCall(t *testing.T, calls *[]sts.PatchMonitorCall, status *sts.MonitorStatusValue) {
	assert.Len(t, *calls, 1)

	call := (*calls)[0]
	assert.Equal(t, "1", call.PmonitorIdOrUrn)
	assert.Equal(t, status, call.PmonitorPatch.Status)
	assert.Nil(t, call.PmonitorPatch.Name)
	assert.Nil(t, call.PmonitorPatch.Description)
	assert.Nil(t, call.PmonitorPatch.Tags)
	assert.Nil(t, call.PmonitorPatch.Identifier)
	assert.Nil(t, call.PmonitorPatch.IntervalSeconds)
	assert.Nil(t, call.PmonitorPatch.RemediationHint)
}
