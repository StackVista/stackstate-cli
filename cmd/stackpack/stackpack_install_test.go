package stackpack

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	FailStrategy = "fail"
)

var (
	name                  = "zabbix"
	provisionID           = int64(1234)
	provisionName         = "provision_name"
	provisionStatus       = "PROVISIONING"
	provisionVersion      = "3.2.0"
	timeMilli             = int64(1438167001716)
	expectedUpdateTime    = time.UnixMilli(timeMilli)
	mockProvisionResponse = &stackstate_api.ProvisionResponse{
		Id:                  &provisionID,
		Name:                &provisionName,
		Status:              &provisionStatus,
		StackPackVersion:    &provisionVersion,
		LastUpdateTimestamp: &timeMilli,
	}
)

func setupStackPackInstallCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackInstallCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Result = *mockProvisionResponse
	return &cli, cmd
}

func TestStackpackInstallPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix",
		"--parameter", "zabbix_instance_name=test_name",
		"--parameter", "zabbix_instance_url=test_url",
	)

	strategyFlag := FailStrategy
	assert.Equal(t,
		stackstate_api.ProvisionDetailsCall{
			PstackPackName: "zabbix",
			Punlocked:      &strategyFlag,
			PrequestBody: &map[string]string{
				"zabbix_instance_name": "test_name",
				"zabbix_instance_url":  "test_url",
			},
		},
		(*cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsCalls)[0],
	)
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"id", "name", "status", "version", "last updated"},
			Data:                [][]interface{}{{&provisionID, &provisionName, &provisionStatus, &provisionVersion, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + name},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackInstallComplexParameters(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix",
		"--parameter", "one=foo,two=bar",
		"--parameter", "i=am=complex",
		"--parameter", "empty=",
		"--parameter", "comma=1,2",
	)

	strategyFlag := FailStrategy
	assert.Equal(t,
		stackstate_api.ProvisionDetailsCall{
			PstackPackName: "zabbix",
			Punlocked:      &strategyFlag,
			PrequestBody: &map[string]string{
				"one":   "foo",
				"two":   "bar",
				"empty": "",
				"i":     "am=complex",
				"comma": "1,2",
			},
		},
		(*cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsCalls)[0],
	)
}

func TestStackpackInstallNoParameters(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix")

	strategyFlag := FailStrategy
	assert.Equal(t,
		stackstate_api.ProvisionDetailsCall{
			PstackPackName: "zabbix",
			Punlocked:      &strategyFlag,
			PrequestBody:   &map[string]string{},
		},
		(*cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsCalls)[0],
	)
}

func TestStackpackInstallPrintsToJson(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix", "-o", "json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"instance": mockProvisionResponse},
		},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestStackpackInstallHasWaitFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackInstallCommand(&cli.Deps)

	// Test that the wait and timeout flags exist
	waitFlag := cmd.Flags().Lookup("wait")
	assert.NotNil(t, waitFlag, "wait flag should exist")
	assert.Equal(t, "false", waitFlag.DefValue, "wait flag default should be false")

	timeoutFlag := cmd.Flags().Lookup("timeout")
	assert.NotNil(t, timeoutFlag, "timeout flag should exist")
	assert.Equal(t, "1m0s", timeoutFlag.DefValue, "timeout flag default should be 1m0s")
}

func TestStackpackInstallWithWaitError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackInstallCommand(&cli.Deps)

	validateWaitFlags(t, cmd)
}

func TestStackpackInstallWithWaitTimeout(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackInstallCommand(&cli.Deps)

	configID := int64(12345)
	timestamp := int64(1438167001716)

	// Setup provision response
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Result = *mockProvisionResponse

	// Setup stackpack list response with provisioning state (never completes)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: name,
			Configurations: []stackstate_api.StackPackConfiguration{
				{
					Id:                  &configID,
					Status:              StatusProvisioning,
					StackPackVersion:    provisionVersion,
					LastUpdateTimestamp: &timestamp,
					Config:              map[string]interface{}{},
				},
			},
		},
	}

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "install", "--name", name, "--wait", "--timeout", "50ms")

	// Should return timeout error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout waiting for stackpack")
}

func TestStackpackInstallWithWaitJsonFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackInstallCommand(&cli.Deps)

	// Test that wait and timeout flags can be parsed with other flags
	err := cmd.ParseFlags([]string{"--name", "test", "--wait", "--timeout", "100ms"})
	assert.NoError(t, err)

	// Verify flag values
	waitValue, err := cmd.Flags().GetBool("wait")
	assert.NoError(t, err)
	assert.True(t, waitValue)

	timeoutValue, err := cmd.Flags().GetDuration("timeout")
	assert.NoError(t, err)
	assert.Equal(t, 100*time.Millisecond, timeoutValue)

	nameValue, err := cmd.Flags().GetString("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", nameValue)
}
