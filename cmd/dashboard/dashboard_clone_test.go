package dashboard

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	clonedDashboardId         = int64(123)
	clonedDashboardName       = "cloned-dashboard"
	clonedDashboardIdentifier = "urn:custom:dashboard:cloned-dashboard"

	cloneResult = sts.DashboardReadFullSchema{
		Id:          clonedDashboardId,
		Name:        clonedDashboardName,
		Identifier:  clonedDashboardIdentifier,
		Description: "Cloned dashboard description",
		Scope:       sts.DASHBOARDSCOPE_PUBLIC_DASHBOARD,
		Dashboard: sts.PersesDashboard{
			Spec: &sts.PersesDashboardSpec{
				Layouts: []sts.PersesLayout{},
				Panels:  &map[string]sts.PersesPanel{},
			},
		},
	}
)

func setDashboardCloneCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardCloneCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldCloneDashboard(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 1)
	cloneCall := (*cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls)[0]
	assert.Equal(t, "1", cloneCall.PdashboardIdOrUrn)
	assert.Equal(t, "cloned-dashboard", cloneCall.PdashboardCloneSchema.Name)
	assert.Equal(t, "Dashboard cloned successfully! New dashboard ID: 123, Name: cloned-dashboard", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldCloneDashboardWithIdentifier(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "urn:custom:dashboard:original", "--name", "cloned-dashboard")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 1)
	cloneCall := (*cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls)[0]
	assert.Equal(t, "urn:custom:dashboard:original", cloneCall.PdashboardIdOrUrn)
}

func TestShouldCloneDashboardWithJson(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard", "--output", "json")

	expectedJsonCall := []map[string]interface{}{
		{
			"dashboard": &cloneResult,
		},
	}
	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestShouldCloneDashboardWithDescription(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard", "--description", "Custom description")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 1)
	cloneCall := (*cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls)[0]
	assert.NotNil(t, cloneCall.PdashboardCloneSchema.Description)
	assert.Equal(t, "Custom description", *cloneCall.PdashboardCloneSchema.Description)
}

func TestShouldCloneDashboardWithScope(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard", "--scope", "privateDashboard")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 1)
	cloneCall := (*cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls)[0]
	assert.NotNil(t, cloneCall.PdashboardCloneSchema.Scope)
	assert.Equal(t, sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD, *cloneCall.PdashboardCloneSchema.Scope)
}

func TestShouldCloneDashboardWithPublicScope(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardResponse.Result = cloneResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard", "--scope", "publicDashboard")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 1)
	cloneCall := (*cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls)[0]
	assert.NotNil(t, cloneCall.PdashboardCloneSchema.Scope)
	assert.Equal(t, sts.DASHBOARDSCOPE_PUBLIC_DASHBOARD, *cloneCall.PdashboardCloneSchema.Scope)
}

func TestShouldRejectInvalidScope(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1", "--name", "cloned-dashboard", "--scope", "invalid")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid scope: invalid. Must be 'publicDashboard' or 'privateDashboard'")
}

func TestCloneDashboardMissingName(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"name\" not set")
}

func TestCloneDashboardMissingIdAndIdentifier(t *testing.T) {
	cli, cmd := setDashboardCloneCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "cloned-dashboard")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "one of the required flags {id | identifier} not set")
}
