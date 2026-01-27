package monitor

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/settings"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const MonitorNodeType = "Monitor"

type EditArgs struct {
	ID         int64
	Identifier string
	Unlock     bool
}

func MonitorEditCommand(cli *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit {--id ID | --identifier URN}",
		Short: "Edit a monitor interactively in your default editor",
		Long:  `Edit a monitor interactively. Opens the monitor definition in the editor defined by your VISUAL or EDITOR environment variable. Locked monitors (from StackPacks) require cloning first.`,
		Example: `# edit a monitor by ID
sts monitor edit --id 123456789

# edit a monitor by identifier
sts monitor edit --identifier urn:stackpack:my-monitor`,
		RunE: cli.CmdRunEWithApi(RunMonitorEditCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)
	cmd.Flags().BoolVar(&args.Unlock, Unlock, false, UnlockFlagUsage)
	err := cmd.Flags().MarkHidden(Unlock)
	if err != nil {
		return nil
	}
	return cmd
}

func RunMonitorEditCommand(args *EditArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		id := args.ID
		monitorName := ""

		if len(args.Identifier) > 0 {
			monitor, resp, err := api.MonitorApi.GetMonitor(cli.Context, args.Identifier).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}
			id = monitor.Id
			monitorName = monitor.Name
		}

		// test whether the node is locked to lead the user to a different flow
		locked, isLockedResponse, isLockedErr := api.NodeApi.Lock(cli.Context, "Monitor", args.ID).Execute()
		if isLockedErr != nil {
			return common.NewResponseError(isLockedErr, isLockedResponse)
		}

		if locked.NodeLocked != nil && !args.Unlock {
			// Print message for the user
			cli.Printer.PrintLn(fmt.Sprintf("The monitor %s that you are trying is locked (SUSE Observability specific), it cannot be edited", monitorName))
			cli.Printer.PrintLn("")
			cli.Printer.PrintLn("To change the behaviour of this monitor you need to follow the following steps:")
			cli.Printer.PrintLn(fmt.Sprintf("1. Clone the monitor:\n\tsts monitor clone -i %d --name <new-name>", id))
			cli.Printer.PrintLn("2. Edit this new monitor:\n\tsts monitor edit -i <monitor-clone-id>")
			cli.Printer.PrintLn("3. If the original monitor was disabled the cloned monitor will also be disable you can enable it using:\n\tsts monitor enable -i <monitor-clone-id>")
			cli.Printer.PrintLn("4. In order to test your edited monitor dry run it with:\n\tsts monitor run -i <monitor-clone-id>")
			cli.Printer.PrintLn(fmt.Sprintf("5. If needed disable the original monitor with:\n\tsts monitor disable -i %d", id))
		} else {
			export := stackstate_api.NewExport()
			export.NodesWithIds = []int64{id}

			orig, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*export).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}

			c, err := cli.Editor.Edit("monitor-", ".json", resp.Body)
			if err != nil {
				return common.NewResponseError(err, resp)
			}

			if strings.Compare(orig, string(c)) == 0 {
				if cli.IsJson() {
					cli.Printer.PrintJson(map[string]interface{}{"nodes": []string{}, "message": "No changes made"})
				} else {
					cli.Printer.PrintWarn("No changes made")
				}

				return nil
			}

			if args.Unlock {
				err := settings.UnlockNodes(cli, api, []int64{id}, MonitorNodeType)
				if err != nil {
					return err
				}
			}

			nodes, resp, err := api.ImportApi.ImportSettings(cli.Context).Body(string(c)).Execute()
			if err != nil {
				return common.NewResponseError(err, resp)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"nodes": nodes,
				})
			} else {
				tableData := make([][]interface{}, 0)
				for _, node := range nodes {
					tableData = append(tableData, []interface{}{node["_type"], node["id"], node["status"], node["identifier"], node["name"]})
				}

				cli.Printer.Success(fmt.Sprintf("Updated <bold>%d</> monitor(s).", len(nodes)))
				if len(nodes) > 0 {
					cli.Printer.Table(printer.TableData{
						Header: []string{"Type", "Id", "Status", "Identifier", "Name"},
						Data:   tableData,
					})
				}
			}
		}

		return nil
	}
}
