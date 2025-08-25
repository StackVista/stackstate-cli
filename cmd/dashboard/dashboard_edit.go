package dashboard

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const LongDescription = `Edit a dashboard.

The edit command allows you to directly edit any StackState Dashboard. It will open
the editor defined by your EDITOR environment variables.

The dashboard will be presented as JSON format for editing.
`

type EditArgs struct {
	ID         int64
	Identifier string
}

func DashboardEditCommand(cli *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a dashboard",
		Long:  LongDescription,
		RunE:  cli.CmdRunEWithApi(RunDashboardEditCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, "ID of the dashboard")
	common.AddIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the dashboard")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDashboardEditCommand(args *EditArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		dashboardIdOrUrn, err := ResolveDashboardIdOrUrn(args.ID, args.Identifier)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}

		// Get the current dashboard
		dashboard, resp, err := api.DashboardsApi.GetDashboard(cli.Context, dashboardIdOrUrn).LoadFullDashboard(true).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		// Convert dashboard to pretty JSON for editing
		originalJSON, err := json.MarshalIndent(dashboard, "", "  ")
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to marshal dashboard to JSON: %v", err))
		}

		// Open editor with the dashboard JSON
		editedContent, err := cli.Editor.Edit("dashboard-", ".json", strings.NewReader(string(originalJSON)))
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to open editor: %v", err))
		}

		// Check if any changes were made
		if strings.Compare(string(originalJSON), string(editedContent)) == 0 {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}
			return nil
		}

		// Parse the edited JSON
		var editedDashboard map[string]interface{}
		if err := json.Unmarshal(editedContent, &editedDashboard); err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to parse edited JSON: %v", err))
		}

		// Create patch schema from the edited JSON
		patchSchema := stackstate_api.NewDashboardPatchSchema()

		if name, ok := editedDashboard["name"].(string); ok && name != "" {
			patchSchema.SetName(name)
		}
		if description, ok := editedDashboard["description"].(string); ok {
			patchSchema.SetDescription(description)
		}
		if scopeStr, ok := editedDashboard["scope"].(string); ok {
			if scope, err := stackstate_api.NewDashboardScopeFromValue(scopeStr); err == nil {
				patchSchema.SetScope(*scope)
			}
		}
		if dashboardContent, ok := editedDashboard["dashboard"]; ok {
			// Convert dashboard content to PersesDashboard
			dashboardBytes, err := json.Marshal(dashboardContent)
			if err == nil {
				var persesDashboard stackstate_api.PersesDashboard
				if err := json.Unmarshal(dashboardBytes, &persesDashboard); err == nil {
					patchSchema.SetDashboard(persesDashboard)
				}
			}
		}

		// Apply the changes
		updatedDashboard, resp, err := api.DashboardsApi.PatchDashboard(cli.Context, dashboardIdOrUrn).DashboardPatchSchema(*patchSchema).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"dashboard": updatedDashboard,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Dashboard updated successfully! ID: %d, Name: %s", updatedDashboard.GetId(), updatedDashboard.GetName()))
		}

		return nil
	}
}
