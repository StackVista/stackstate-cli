package stackpack

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	// StackPack configuration status constants for wait operations
	StatusInstalled    = "INSTALLED"    // Configuration is successfully installed
	StatusProvisioning = "PROVISIONING" // Configuration is still being processed
	StatusError        = "ERROR"        // Configuration failed with errors

	// Default wait operation settings
	DefaultPollInterval = 5 * time.Second // How often to check status during wait
	DefaultTimeout      = 1 * time.Minute // Default timeout for wait operations
)

// OperationWaiter provides functionality to wait for StackPack operations to complete
// by polling the API and monitoring configuration status changes
type OperationWaiter struct {
	cli *di.Deps
	api *stackstate_api.APIClient
}

// WaitOptions configures how the wait operation should behave
type WaitOptions struct {
	StackPackName string        // Name of the StackPack to monitor
	Timeout       time.Duration // Maximum time to wait before giving up
	PollInterval  time.Duration // How often to check the status
}

func NewOperationWaiter(cli *di.Deps, api *stackstate_api.APIClient) *OperationWaiter {
	return &OperationWaiter{
		cli: cli,
		api: api,
	}
}

// WaitForCompletion polls the StackPack API until all configurations are installed or an error occurs.
// Returns nil on success, error on timeout or configuration failures.
func (w *OperationWaiter) WaitForCompletion(options WaitOptions) error {
	// Set up timeout context for the entire wait operation
	ctx, cancel := context.WithTimeout(w.cli.Context, options.Timeout)
	defer cancel()

	// Set up ticker for periodic polling
	ticker := time.NewTicker(options.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for stackpack '%s' operation to complete after %v", options.StackPackName, options.Timeout)
		case <-ticker.C:
			// Poll the API to check current status
			stackPackList, cliErr := fetchAllStackPacks(w.cli, w.api)
			if cliErr != nil {
				return fmt.Errorf("failed to check stackpack status: %v", cliErr)
			}

			stackPack, err := findStackPackByName(stackPackList, options.StackPackName)
			if err != nil {
				return fmt.Errorf("stackpack '%s' not found: %v", options.StackPackName, err)
			}

			// Check the status of all configurations for this StackPack
			allInstalled := true
			hasProvisioning := false
			var errorMessages []string

			for _, config := range stackPack.GetConfigurations() {
				status := config.GetStatus()
				switch status {
				case StatusError:
					// Extract detailed error message from the API response
					errorMsg := fmt.Sprintf("Configuration %d failed", config.GetId())
					if config.HasError() {
						stackPackError := config.GetError()
						apiError := stackPackError.GetError()
						if message, ok := apiError["message"]; ok {
							if msgStr, ok := message.(string); ok {
								errorMsg = fmt.Sprintf("Configuration %d failed: %s", config.GetId(), msgStr)
							}
						}
					}
					errorMessages = append(errorMessages, errorMsg)
				case StatusProvisioning:
					hasProvisioning = true
					allInstalled = false
				case StatusInstalled:
					// Continue checking other configs
				default:
					// Unknown status, treat as still in progress
					allInstalled = false
				}
			}

			// Return immediately if any configuration has failed
			if len(errorMessages) > 0 {
				return fmt.Errorf("stackpack '%s' failed:\n%s", options.StackPackName, strings.Join(errorMessages, "\n"))
			}

			// Success: all configurations are installed and none are provisioning
			if allInstalled && !hasProvisioning {
				return nil
			}

			// Continue polling - some configurations are still in progress
		}
	}
}

func findStackPackByName(stacks []stackstate_api.FullStackPack, name string) (stackstate_api.FullStackPack, error) {
	for _, v := range stacks {
		if v.GetName() == name {
			return v, nil
		}
	}
	return stackstate_api.FullStackPack{}, fmt.Errorf("stackpack %s does not exist", name)
}

// fetchAllStackPacks retrieves all StackPacks from the API and returns them sorted by name.
// This function was moved from stackpack_list.go to common.go for reuse in wait operations.
func fetchAllStackPacks(cli *di.Deps, api *stackstate_api.APIClient) ([]stackstate_api.FullStackPack, common.CLIError) {
	stackPackList, resp, err := api.StackpackApi.StackPackList(cli.Context).Execute()
	if err != nil {
		return nil, common.NewResponseError(err, resp)
	}

	// Sort by name for consistent ordering
	sort.SliceStable(stackPackList, func(i, j int) bool {
		return stackPackList[i].Name < stackPackList[j].Name
	})
	return stackPackList, nil
}
