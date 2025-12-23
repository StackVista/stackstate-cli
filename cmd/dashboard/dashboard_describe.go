package dashboard

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/util"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type DescribeArgs struct {
	ID         int64
	Identifier string
	FilePath   string
}

func DashboardDescribeCommand(cli *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a dashboard in YAML format",
		Long:  "Describe a dashboard in StackState Templated YAML.",
		RunE:  cli.CmdRunEWithApi(RunDashboardDescribeCommand(args)),
	}

	cmd.Flags().Int64Var(&args.ID, "id", 0, "ID of the dashboard")
	cmd.Flags().StringVar(&args.Identifier, "identifier", "", "Identifier (URN) of the dashboard")
	common.AddFileFlagVar(cmd, &args.FilePath, "Path to the output file")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDashboardDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		dashboardIdOrUrn, err := ResolveDashboardIdOrUrn(args.ID, args.Identifier)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}

		// Get the full dashboard data
		dashboard, resp, err := api.DashboardsApi.GetDashboard(cli.Context, dashboardIdOrUrn).LoadFullDashboard(true).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		yamlData, err := yaml.Marshal(dashboard.DashboardReadFullSchema)
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to marshal dashboard: %v", err))
		}
		data := string(yamlData)

		if args.FilePath != "" {
			if err := util.WriteFile(args.FilePath, []byte(data)); err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"filepath": args.FilePath,
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("dashboard exported to: %s", args.FilePath))
			}

			return nil
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"data":   data,
					"format": "json",
				})
			} else {
				cli.Printer.PrintLn(data)
			}
			return nil
		}
	}
}
