package graph

import (
	"testing"

	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestDeleteExpiredData(t *testing.T) {
	for n, e := range responses {
		t.Run(n, func(t *testing.T) {
			result := stackstate_admin_api.NewRemovalProgress(*stackstate_admin_api.NewRemovalProgressProgress(n))

			cli := di.NewMockDeps(t)
			cmd := DeleteExpiredDataCommand(&cli.Deps)

			cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataResponse.Result = *result

			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

			assert.Len(t, *cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataCalls, 1)

			expected := []string{
				e,
			}

			assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
		})
	}
}

func TestDeleteExpiredDataDefaultResponse(t *testing.T) {
	result := stackstate_admin_api.NewRemovalProgress(*stackstate_admin_api.NewRemovalProgressProgress("Some Other Type"))

	cli := di.NewMockDeps(t)
	cmd := DeleteExpiredDataCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataResponse.Result = *result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataCalls, 1)

	expected := []string{
		"Command executed successfully.",
	}

	assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
}

func TestDeleteExpiredDataJSON(t *testing.T) {
	for n := range responses {
		t.Run(n, func(t *testing.T) {
			progress := stackstate_admin_api.NewRemovalProgressProgress(n)
			result := stackstate_admin_api.NewRemovalProgress(*progress)

			cli := di.NewMockDeps(t)
			cmd := DeleteExpiredDataCommand(&cli.Deps)

			cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataResponse.Result = *result

			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

			assert.Len(t, *cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataCalls, 1)

			expected := []map[string]interface{}{
				{
					"progress": *progress,
				},
			}

			assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
		})
	}
}

func TestDeleteExpiredDataImmediately(t *testing.T) {
	result := stackstate_admin_api.NewRemovalProgress(*stackstate_admin_api.NewRemovalProgressProgress("RemovalInProgress"))

	cli := di.NewMockDeps(t)
	cmd := DeleteExpiredDataCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataResponse.Result = *result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--immediate")

	calls := *cli.MockClient.ApiMocks.RetentionAPI.RemoveExpiredDataCalls

	assert.Len(t, calls, 1)
	assert.Equal(t, true, *calls[0].PexpireImmediatelyAndRestart)
}
