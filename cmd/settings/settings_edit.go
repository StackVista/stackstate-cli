package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
	"k8s.io/kubectl/pkg/cmd/util/editor"
)

type EditArgs struct {
	Ids              []int64
	Namespace        string
	NodeTypes        []string
	AllowReferences  []string
	UnlockedStrategy string
	Timeout          int64
}

func SettingsEditCommand(cli *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit settings in STJ format",
		Long:  "Edit settings in StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunSettingsEditCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "List of ids to edit")
	cmd.Flags().StringVar(&args.Namespace, Namespace, "", "Namespace to edit -- WARNING: This will overwrite all settings in the namespace")
	cmd.Flags().StringSliceVar(&args.NodeTypes, TypeNameFlag, nil, "List of types to edit")
	cmd.Flags().StringSliceVar(&args.AllowReferences, AllowReferencesFlag, nil, "White list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"Strategy to use when encountering unlocked settings when applying settings to a namespace"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")))
	cmd.Flags().Int64VarP(&args.Timeout, TimeoutFlag, "t", 0, "Timeout in seconds")
	stscobra.MarkMutexFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, "filter", true)
	return cmd
}

func RunSettingsEditCommand(args *EditArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		if err := stscobra.CheckMutuallyExclusiveFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, true); err != nil {
			return common.NewCLIArgParseError(err)
		}

		orig, resp, err := doExport(cli.Context, api, args.Ids, args.Namespace, args.NodeTypes, args.AllowReferences)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		e := editor.NewDefaultEditor([]string{"VISUAL", "EDITOR"})

		c, f, err := e.LaunchTempFile("setting-", ".json", resp.Body)
		if err != nil {
			os.Remove(f)
			return common.NewResponseError(err, resp)
		}

		defer os.Remove(f)

		if strings.Compare(orig, string(c)) == 0 {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"updated-settings": []string{}, "message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}

			return nil
		}

		nodes, resp, err := doImport(cli.Context, api, string(c), args.Namespace, args.UnlockedStrategy, args.Timeout)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"updated-settings": nodes,
			})
		} else {
			if len(nodes) == 0 {
				cli.Printer.PrintWarn("Nothing was updated.")
				return nil
			}

			tableData := make([][]interface{}, 0)
			for _, node := range nodes {
				tableData = append(tableData, []interface{}{node["_type"], node["id"], node["identifier"], node["name"]})
			}

			cli.Printer.Success(fmt.Sprintf("Updated <bold>%d</> setting node(s).\n", len(nodes)))
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
