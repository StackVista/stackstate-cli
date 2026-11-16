package dashboard

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setDashboardDeleteCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardDeleteCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldDeleteDashboard(t *testing.T) {
	cli, cmd := setDashboardDeleteCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "123")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls, 1)
	deleteCall := (*cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls)[0]
	assert.Equal(t, "123", deleteCall.PdashboardIdOrUrn)
	assert.Equal(t, "Dashboard deleted: 123", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldDeleteDashboardWithIdentifier(t *testing.T) {
	cli, cmd := setDashboardDeleteCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "urn:custom:dashboard:test")

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls, 1)
	deleteCall := (*cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls)[0]
	assert.Equal(t, "urn:custom:dashboard:test", deleteCall.PdashboardIdOrUrn)
	assert.Equal(t, "Dashboard deleted: urn:custom:dashboard:test", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldDeleteDashboardWithJson(t *testing.T) {
	cli, cmd := setDashboardDeleteCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "123", "--output", "json")

	expectedJsonCall := []map[string]interface{}{
		{
			"deleted-dashboard-identifier": "123",
		},
	}
	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestDeleteDashboardMissingIdAndIdentifier(t *testing.T) {
	cli, cmd := setDashboardDeleteCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "one of the required flags {id | identifier} not set")
}

func TestDeleteDashboardCallsApiWithCorrectParams(t *testing.T) {
	cli, cmd := setDashboardDeleteCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "456")

	// Verify the API was called with the correct parameters
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls, 1)
	deleteCall := (*cli.MockClient.ApiMocks.DashboardsApi.DeleteDashboardCalls)[0]
	assert.Equal(t, "456", deleteCall.PdashboardIdOrUrn)

	// Verify no other dashboard API calls were made
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls, 0)
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CloneDashboardCalls, 0)
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardsCalls, 0)
}
