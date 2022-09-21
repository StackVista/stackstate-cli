package monitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestMonitorDisableById(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", "1")

	checkDisableCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls)
}

func TestMonitorDisableByIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := MonitorDisableCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "1")

	checkDisableCall(t, cli.MockClient.ApiMocks.MonitorApi.PatchMonitorCalls)
}

func checkDisableCall(t *testing.T, calls *[]sts.PatchMonitorCall) {
	assert.Len(t, *calls, 1)

	call := (*calls)[0]
	assert.Equal(t, "1", call.PmonitorIdOrUrn)
	assert.Equal(t, sts.MONITORSTATUSVALUE_DISABLED.Ptr(), call.PmonitorPatch.Status)
	assert.Nil(t, call.PmonitorPatch.Name)
	assert.Nil(t, call.PmonitorPatch.Description)
	assert.Nil(t, call.PmonitorPatch.Tags)
	assert.Nil(t, call.PmonitorPatch.Identifier)
	assert.Nil(t, call.PmonitorPatch.IntervalSeconds)
	assert.Nil(t, call.PmonitorPatch.RemediationHint)
}
