package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

var (
	UnlockedStrategyChoices = []string{"fail", "skip", "overwrite"}
)

type ApplyArgs struct {
	Filepath         string
	Namespace        string
	UnlockedStrategy string
	Timeout          int64
}

func SettingsApplyCommand(cli *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "apply saved settings",
		Long:  "Apply saved settings with StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunSettingsApplyCommand(args)),
	}
	common.AddRequiredFileFlagVar(cmd, &args.Filepath, ".stj file to import")
	cmd.Flags().StringVar(&args.Namespace, NamespaceFlag, "", "name of the namespace to overwrite"+
		" - WARNING this will overwrite the entire namespace")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"strategy to use when encountering unlocked settings when applying settings to a namespace"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")))
	cmd.Flags().Int64VarP(&args.Timeout, TimeoutFlag, "t", 0, "timeout in seconds")

	return cmd
}

func RunSettingsApplyCommand(args *ApplyArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		fileBytes, err := os.ReadFile(args.Filepath)
		if err != nil {
			return common.NewReadFileError(err, args.Filepath)
		}

		request := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes))
		if args.Namespace != "" {
			request = request.Namespace(args.Namespace)
		}
		if args.UnlockedStrategy != "" {
			request = request.Unlocked(args.UnlockedStrategy)
		}
		if args.Timeout > 0 {
			request = request.TimeoutSeconds(args.Timeout)
		}

		nodes, resp, err := request.Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"applied-settings": nodes,
			})
		} else {
			if len(nodes) == 0 {
				cli.Printer.PrintWarn("Nothing was imported.")
				return nil
			}

			tableData := make([][]interface{}, 0)
			for _, node := range nodes {
				tableData = append(tableData, []interface{}{node["_type"], node["id"], node["identifier"], node["name"]})
			}

			cli.Printer.Success(fmt.Sprintf("Applied <bold>%d</> setting node(s).\n", len(nodes)))
			if len(nodes) > 0 {
				cli.Printer.Table(printer.TableData{
					Header: []string{"Type", "Id", "Identifier", "Name"},
					Data:   tableData,
				})
			}
		}

		return nil
	}
}
