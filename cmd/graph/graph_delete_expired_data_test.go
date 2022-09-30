package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	stackstate_admin_api "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestDeleteExpiredData(t *testing.T) {
	for n, e := range responses {
		t.Run(n, func(t *testing.T) {
			progress := stackstate_admin_api.NewRemovalProgressProgress(n)
			result := stackstate_admin_api.RemovalProgress{
				Progress: progress,
			}

			cli := di.NewMockDeps(t)
			cmd := DeleteExpiredDataCommand(&cli.Deps)

			cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataResponse.Result = result

			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

			assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataCalls, 1)

			expected := []string{
				e,
			}

			assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
		})
	}
}

func TestDeleteExpiredDataDefaultResponse(t *testing.T) {
	progress := stackstate_admin_api.NewRemovalProgressProgress("Some Other Type")
	result := stackstate_admin_api.RemovalProgress{
		Progress: progress,
	}

	cli := di.NewMockDeps(t)
	cmd := DeleteExpiredDataCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataCalls, 1)

	expected := []string{
		"Command executed successfully.",
	}

	assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
}

func TestDeleteExpiredDataJSON(t *testing.T) {
	for n := range responses {
		t.Run(n, func(t *testing.T) {
			progress := stackstate_admin_api.NewRemovalProgressProgress(n)
			result := stackstate_admin_api.RemovalProgress{
				Progress: progress,
			}

			cli := di.NewMockDeps(t)
			cmd := DeleteExpiredDataCommand(&cli.Deps)

			cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataResponse.Result = result

			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

			assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataCalls, 1)

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
	progress := stackstate_admin_api.NewRemovalProgressProgress("RemovalInProgress")
	result := stackstate_admin_api.RemovalProgress{
		Progress: progress,
	}

	cli := di.NewMockDeps(t)
	cmd := DeleteExpiredDataCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--immediate")

	calls := *cli.MockClient.ApiMocks.RetentionApi.RemoveExpiredDataCalls

	assert.Len(t, calls, 1)
	assert.Equal(t, true, *calls[0].PexpireImmediatelyAndRestart)
}
