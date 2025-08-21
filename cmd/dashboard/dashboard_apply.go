package dashboard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

// ApplyArgs contains arguments for dashboard apply command
type ApplyArgs struct {
	File string
}

func DashboardApplyCommand(cli *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Create or edit a dashboard from JSON",
		Long:  "Create or edit a dashboard from JSON file.",
		RunE:  cli.CmdRunEWithApi(RunDashboardApplyCommand(args)),
	}

	common.AddRequiredFileFlagVar(cmd, &args.File, "Path to a .json file with the dashboard definition")

	return cmd
}

func RunDashboardApplyCommand(args *ApplyArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		fileBytes, err := os.ReadFile(args.File)
		if err != nil {
			return common.NewReadFileError(err, args.File)
		}

		// Determine file type by extension
		ext := strings.ToLower(filepath.Ext(args.File))
		if ext != ".json" {
			return common.NewCLIArgParseError(fmt.Errorf("unsupported file type: %s. Only .json files are supported", ext))
		}

		return applyJSONDashboard(cli, api, fileBytes)
	}
}

// applyJSONDashboard processes JSON dashboard file and determines create vs update operation
func applyJSONDashboard(cli *di.Deps, api *stackstate_api.APIClient, fileBytes []byte) common.CLIError {
	// Parse the JSON to determine if it's a create or update operation
	var dashboardData map[string]interface{}
	if err := json.Unmarshal(fileBytes, &dashboardData); err != nil {
		return common.NewCLIArgParseError(fmt.Errorf("failed to parse JSON: %v", err))
	}

	// Check if it has an ID field (indicates update operation)
	if idField, hasId := dashboardData["id"]; hasId {
		// Update existing dashboard
		dashboardId := fmt.Sprintf("%.0f", idField.(float64))
		return updateDashboard(cli, api, dashboardId, dashboardData)
	} else {
		// Create new dashboard
		return createDashboard(cli, api, fileBytes)
	}
}

// createDashboard creates a new dashboard from JSON schema
func createDashboard(cli *di.Deps, api *stackstate_api.APIClient, fileBytes []byte) common.CLIError {
	var writeSchema stackstate_api.DashboardWriteSchema
	if err := json.Unmarshal(fileBytes, &writeSchema); err != nil {
		return common.NewCLIArgParseError(fmt.Errorf("failed to parse JSON as DashboardWriteSchema: %v", err))
	}

	// Validate required fields
	if writeSchema.Name == "" {
		return common.NewCLIArgParseError(fmt.Errorf("dashboard name is required"))
	}

	// Create new dashboard
	dashboard, resp, err := api.DashboardsApi.CreateDashboard(cli.Context).DashboardWriteSchema(writeSchema).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"dashboard": dashboard,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Dashboard created successfully! ID: %d, Name: %s", dashboard.GetId(), dashboard.GetName()))
	}

	return nil
}

// updateDashboard patches an existing dashboard with new data
func updateDashboard(cli *di.Deps, api *stackstate_api.APIClient, dashboardId string, dashboardData map[string]interface{}) common.CLIError {
	// Create patch schema from the JSON data
	patchSchema := stackstate_api.NewDashboardPatchSchema()

	if name, ok := dashboardData["name"].(string); ok && name != "" {
		patchSchema.SetName(name)
	}
	if description, ok := dashboardData["description"].(string); ok {
		patchSchema.SetDescription(description)
	}
	if scopeStr, ok := dashboardData["scope"].(string); ok {
		if scope, err := stackstate_api.NewDashboardScopeFromValue(scopeStr); err == nil {
			patchSchema.SetScope(*scope)
		}
	}
	if dashboardContent, ok := dashboardData["dashboard"]; ok {
		// Convert dashboard content to PersesDashboard
		dashboardBytes, err := json.Marshal(dashboardContent)
		if err == nil {
			var persesDashboard stackstate_api.PersesDashboard
			if err := json.Unmarshal(dashboardBytes, &persesDashboard); err == nil {
				patchSchema.SetDashboard(persesDashboard)
			}
		}
	}

	// Update existing dashboard
	dashboard, resp, err := api.DashboardsApi.PatchDashboard(cli.Context, dashboardId).DashboardPatchSchema(*patchSchema).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"dashboard": dashboard,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Dashboard updated successfully! ID: %d, Name: %s", dashboard.GetId(), dashboard.GetName()))
	}

	return nil
}
