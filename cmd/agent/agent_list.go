package agent

import (
	"cmp"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"slices"
	"time"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all registered agents",
		Long:  "List all registered agents",
		RunE:  deps.CmdRunEWithApi(RunListCommand),
	}

	return cmd
}

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
		if agent.Lease == stackstate_api.AGENTLEASE_ACTIVE {
			active++
		} else if agent.Lease == stackstate_api.AGENTLEASE_LIMITED {
			limited++
		} else {
			stale++
		}
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"agents": agentList,
		})
	} else {
		data := make([][]interface{}, len(agentList))
		currentTime := time.Now()

		for i, agent := range agentList {
			data[i] = []interface{}{agent.AgentId, agent.Lease, currentTime.Sub(time.UnixMilli(agent.RegisteredEpochMs)).String(), time.UnixMilli(agent.GetLeaseUntilEpochMs()).String()}
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Host", "Lease", "Age", "Last Lease"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "topics"},
		})

		cli.Printer.PrintLn("")
		cli.Printer.PrintLn(fmt.Sprintf("Totals: %d active, %d limited, %d stale", active, limited, stale))
	}

	return nil
}
