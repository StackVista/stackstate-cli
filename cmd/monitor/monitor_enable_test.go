package monitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestMonitorEnableById(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been enabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorEnableByIdJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorApi.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
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

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	assert.Equal(t, []string{"Monitor 1 has been enabled"}, *cli.MockPrinter.SuccessCalls)
}

func TestMonitorEnableByIdentifierJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)
	cli.MockClient.ApiMocks.MonitorApi.PatchMonitorResponse.Result = monitor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1", "-o", "json")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
	expectedJsonCalls := []map[string]interface{}{{
		"monitor": &monitor,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
