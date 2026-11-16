package stackpack

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

const (
	testStackPackName        = "test-stackpack"
	successfulResponseResult = "successful"
)

//nolint:funlen
func TestOperationWaiter_WaitForCompletion_Success(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	configID := int64(12345)
	timestamp := int64(1438167001716)

	// Mock successful completion
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: testStackPackName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &configID,
					Status:              StatusInstalled,
					StackPackVersion:    "1.0.0",
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
				},
			},
		},
	}

	waiter := NewOperationWaiter(&cli.Deps, api)
	options := WaitOptions{
		StackPackName: testStackPackName,
		Timeout:       5 * time.Second,
		PollInterval:  10 * time.Millisecond,
	}

	err := waiter.WaitForCompletion(options)

	assert.NoError(t, err)
	assert.True(t, len(*cli.MockClient.ApiMocks.StackpackApi.StackPackListCalls) >= 1)
}

//nolint:funlen
func TestOperationWaiter_WaitForCompletion_Error(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	configID := int64(12345)
	timestamp := int64(1438167001716)
	errorMessage := "Object is missing required member 'function'"

	// Mock error response
	stackPackError := stackstate_api.StackPackError{
		Retryable: false,
		Error: map[string]interface{}{
			"message": errorMessage,
		},
	}

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: testStackPackName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &configID,
					Status:              StatusError,
					StackPackVersion:    "1.0.0",
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
					Error:               &stackPackError,
				},
			},
		},
	}

	waiter := NewOperationWaiter(&cli.Deps, api)
	options := WaitOptions{
		StackPackName: testStackPackName,
		Timeout:       5 * time.Second,
		PollInterval:  10 * time.Millisecond,
	}

	err := waiter.WaitForCompletion(options)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Configuration %d failed: %s", configID, errorMessage))
	assert.Contains(t, err.Error(), testStackPackName)
}

//nolint:funlen
func TestOperationWaiter_WaitForCompletion_Timeout(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	configID := int64(12345)
	timestamp := int64(1438167001716)

	// Mock provisioning state that never completes
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: testStackPackName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &configID,
					Status:              StatusProvisioning,
					StackPackVersion:    "1.0.0",
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
				},
			},
		},
	}

	waiter := NewOperationWaiter(&cli.Deps, api)
	options := WaitOptions{
		StackPackName: testStackPackName,
		Timeout:       50 * time.Millisecond,
		PollInterval:  10 * time.Millisecond,
	}

	err := waiter.WaitForCompletion(options)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout waiting for stackpack")
	assert.Contains(t, err.Error(), testStackPackName)
}

func TestOperationWaiter_WaitForCompletion_StackpackNotFound(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	stackpackName := "nonexistent-stackpack"

	// Mock empty stackpack list
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{}

	waiter := NewOperationWaiter(&cli.Deps, api)
	options := WaitOptions{
		StackPackName: stackpackName,
		Timeout:       5 * time.Second,
		PollInterval:  10 * time.Millisecond,
	}

	err := waiter.WaitForCompletion(options)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "stackpack 'nonexistent-stackpack' not found")
}

//nolint:funlen
func TestOperationWaiter_WaitForCompletion_MultipleConfigurations(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	configID1 := int64(12345)
	configID2 := int64(67890)
	timestamp := int64(1438167001716)

	// Mock multiple configurations - one installed, one provisioning
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: testStackPackName,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &configID1,
					Status:              StatusInstalled,
					StackPackVersion:    "1.0.0",
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
				},
				{
					Id:                  &configID2,
					Status:              StatusProvisioning,
					StackPackVersion:    "1.0.0",
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
				},
			},
		},
	}

	waiter := NewOperationWaiter(&cli.Deps, api)
	options := WaitOptions{
		StackPackName: testStackPackName,
		Timeout:       50 * time.Millisecond,
		PollInterval:  10 * time.Millisecond,
	}

	err := waiter.WaitForCompletion(options)

	// Should timeout because one config is still provisioning
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout waiting for stackpack")
}

func TestFindStackPackByName_Found(t *testing.T) {
	stackpackName := "zabbix"
	stacks := []stackstate_api.FullStackPack{
		{Name: "mysql"},
		{Name: stackpackName},
		{Name: "redis"},
	}

	result, err := findStackPackByName(stacks, stackpackName)

	assert.NoError(t, err)
	assert.Equal(t, stackpackName, result.GetName())
}

func TestFindStackPackByName_NotFound(t *testing.T) {
	stacks := []stackstate_api.FullStackPack{
		{Name: "mysql"},
		{Name: "redis"},
	}

	_, err := findStackPackByName(stacks, "nonexistent")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "stackpack nonexistent does not exist")
}

func TestNewOperationWaiter(t *testing.T) {
	cli := di.NewMockDeps(t)
	api, _, _ := cli.MockClient.Connect()

	waiter := NewOperationWaiter(&cli.Deps, api)

	assert.NotNil(t, waiter)
	assert.Equal(t, &cli.Deps, waiter.cli)
	assert.Equal(t, api, waiter.api)
}

// validateWaitFlags is a helper function to test wait flag functionality across commands.
// This reduces code duplication between install and upgrade test files.
func validateWaitFlags(t *testing.T, cmd *cobra.Command) {
	// Test that the wait flag exists and has correct defaults
	waitFlag := cmd.Flags().Lookup("wait")
	assert.NotNil(t, waitFlag, "wait flag should exist")
	assert.Equal(t, "false", waitFlag.DefValue, "wait flag default should be false")

	timeoutFlag := cmd.Flags().Lookup("timeout")
	assert.NotNil(t, timeoutFlag, "timeout flag should exist")
	assert.Equal(t, "1m0s", timeoutFlag.DefValue, "timeout flag default should be 1m0s")

	// Verify the flags can be set
	err := cmd.ParseFlags([]string{"--wait", "--timeout", "30s"})
	assert.NoError(t, err)

	waitValue, err := cmd.Flags().GetBool("wait")
	assert.NoError(t, err)
	assert.True(t, waitValue)

	timeoutValue, err := cmd.Flags().GetDuration("timeout")
	assert.NoError(t, err)
	assert.Equal(t, 30*time.Second, timeoutValue)
}
