package agent

import (
	"testing"
	"time"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	SomeAgentRegistrations = []stackstate_api.AgentRegistration{
		{
			AgentId:           "id1",
			Lease:             stackstate_api.AGENTLEASE_ACTIVE,
			LeaseUntilEpochMs: 0,
			RegisteredEpochMs: 0,
			AgentData: &stackstate_api.AgentData{
				Platform:      "platform",
				CoreCount:     4,
				MemoryBytes:   16 * 1024 * 1024,
				KernelVersion: "5.15",
			},
			NodeBudgetCount: 1,
		}, {
			AgentId:           "id2",
			Lease:             stackstate_api.AGENTLEASE_LIMITED,
			LeaseUntilEpochMs: 0,
			RegisteredEpochMs: 0,
			AgentData:         nil,
			NodeBudgetCount:   2,
		},
	}
)

func TestAgentList(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)

	cli.MockClient.ApiMocks.AgentRegistrationsApi.AllAgentRegistrationsResponse.Result = stackstate_api.AgentRegistrations{
		Agents: SomeAgentRegistrations,
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.AgentRegistrationsApi.AllAgentRegistrationsCalls, 1)

	nulTime := time.UnixMilli(0).Format("2006-01-02 15:04:05 MST")

	expectedTable := []printer.TableData{
		{
			Header: []string{"Host", "Lease", "Registered", "Last Lease", "CPUS", "Memory", "Platform", "Kernel", "# Nodes"},
			Data: [][]interface{}{
				{"id1", stackstate_api.AGENTLEASE_ACTIVE, nulTime, nulTime, "4", "16MB", "platform", "5.15", int32(1)},
				{"id2", stackstate_api.AGENTLEASE_LIMITED, nulTime, nulTime, "0", "0MB", "", "", int32(2)},
			},
		},
	}

	assert.Equal(t, expectedTable, *cli.MockPrinter.TableCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.AgentRegistrationsApi.AllAgentRegistrationsCalls, 2)

	expectedJson := []map[string]interface{}{
		{
			"agents": SomeAgentRegistrations,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}
