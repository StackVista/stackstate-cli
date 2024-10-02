package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

var (
	UnlockedStrategyChoices = []string{"fail", "skip", "overwrite"}
	LockedStrategyChoices   = []string{"fail", "skip", "overwrite"}
)

type ApplyArgs struct {
	Filepath         string
	Namespace        string
	UnlockedStrategy string
	LockedStrategy   string
	Timeout          int64
}

func SettingsApplyCommand(cli *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply saved settings",
		Long:  "Apply saved settings with StackState Templated YAML.",
		RunE:  cli.CmdRunEWithApi(RunSettingsApplyCommand(args)),
	}
	common.AddRequiredFileFlagVar(cmd, &args.Filepath, "Path to a .sty or .stj file with the settings to import")
	cmd.Flags().StringVar(&args.Namespace, NamespaceFlag, "", "Name of the namespace to overwrite"+
		" - WARNING this will overwrite the entire namespace")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"Strategy to use when encountering unlocked settings when applying settings to a namespace"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")))
	pflags.EnumVar(cmd.Flags(), &args.LockedStrategy,
		LockedStrategyFlag,
		"",
		LockedStrategyChoices,
		"Strategy to use when encountering locked settings"+
			fmt.Sprintf(" (must be { %s })", strings.Join(LockedStrategyChoices, " | ")))
	cmd.Flags().Int64VarP(&args.Timeout, TimeoutFlag, TimeoutFlagShort, 0, TimeoutUsage)

	return cmd
}

func RunSettingsApplyCommand(args *ApplyArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		fileBytes, err := os.ReadFile(args.Filepath)
		if err != nil {
			return common.NewReadFileError(err, args.Filepath)
		}

		nodes, resp, err := doImport(cli.Context, api, string(fileBytes), args.Namespace, args.UnlockedStrategy, args.LockedStrategy, args.Timeout)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
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
