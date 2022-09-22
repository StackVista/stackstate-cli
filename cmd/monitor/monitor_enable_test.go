package monitor

import (
	"testing"

	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestMonitorEnableById(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
}

func TestMonitorEnableByIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEnableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1")

	checkMonitorPatchStatusCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls, sts.MONITORSTATUSVALUE_ENABLED.Ptr())
}
