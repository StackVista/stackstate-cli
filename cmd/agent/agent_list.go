package agent

import (
	"cmp"
	"fmt"
	"slices"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all registered agents",
		Long:  "List all registered agents.",
		RunE:  deps.CmdRunEWithApi(RunListCommand),
	}

	return cmd
}

const toMB = 1024 * 1024

func RunListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	agents, resp, err := api.AgentRegistrationsApi.AllAgentRegistrations(cli.Context).Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	agentList := agents.Agents

	slices.SortFunc(agentList, func(a, b stackstate_api.AgentRegistration) int {
		if n := cmp.Compare(a.Lease, b.Lease); n != 0 {
			return n
		}
		// If leases are equal, order by registration moment
		return cmp.Compare(a.RegisteredEpochMs, b.RegisteredEpochMs)
	})

	var active = 0
	var limited = 0
	var stale = 0

	for _, agent := range agentList {
		switch agent.Lease {
		case stackstate_api.AGENTLEASE_ACTIVE:
			active += int(agent.NodeBudgetCount)
		case stackstate_api.AGENTLEASE_LIMITED:
			limited += int(agent.NodeBudgetCount)
		case stackstate_api.AGENTLEASE_STALE:
			stale += int(agent.NodeBudgetCount)
		default:
		}
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"agents": agentList,
		})
	} else {
		data := make([][]interface{}, len(agentList))

		for i, agent := range agentList {
			var info = agent.AgentData
			if info == nil {
				info = &stackstate_api.AgentData{Platform: "", CoreCount: 0, MemoryBytes: 0, KernelVersion: ""}
			}

			data[i] = []interface{}{
				agent.AgentId,
				agent.Lease,
				time.UnixMilli(agent.RegisteredEpochMs).Format("2006-01-02 15:04:05 MST"),
				time.UnixMilli(agent.LeaseUntilEpochMs).Format("2006-01-02 15:04:05 MST"),
				fmt.Sprintf("%v", info.CoreCount),
				fmt.Sprintf("%vMB", info.MemoryBytes/toMB),
				info.Platform,
				info.KernelVersion,
				agent.NodeBudgetCount,
			}
		}
		cli.Printer.Table(printer.TableData{
			Header: []string{"Host", "Lease", "Registered", "Last Lease", "CPUS", "Memory", "Platform", "Kernel", "# Nodes"},
			Data:   data,
		})

		cli.Printer.PrintLn("")
		cli.Printer.PrintLn(fmt.Sprintf("Totals: %d active, %d limited, %d stale", active, limited, stale))
	}

	return nil
}
