package settings

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const LongDescription = `Edit settings in StackState Templated JSON.

The edit command allows you to directly edit any StackState type you can retrieve via the API. It will open
the editor defined by your VISUAL, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for
Windows.
When '--unlock' is specified, the CLI will always unlock the settings node when editing it.
This might introduce changes that prevent the originating StackPack from upgrading correctly. Any changes you make are not the responsibility of the StackPack developer.
`

type EditArgs struct {
	Ids              []int64
	NodeTypes        []string
	AllowReferences  []string
	UnlockedStrategy string
	Timeout          int64
	Unlock           bool
}

func SettingsEditCommand(cli *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit settings in STJ format",
		Long:  LongDescription,
		RunE:  cli.CmdRunEWithApi(RunSettingsEditCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "List of ids to edit")
	cmd.Flags().StringSliceVar(&args.NodeTypes, TypeNameFlag, nil, "List of types to edit")
	cmd.Flags().StringSliceVar(&args.AllowReferences, AllowReferencesFlag, nil, "White list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"Strategy to use when encountering unlocked settings when applying settings to a namespace"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")))
	cmd.Flags().Int64VarP(&args.Timeout, TimeoutFlag, "t", 0, "Timeout in seconds")
	stscobra.MarkMutexFlags(cmd, []string{IdsFlag, TypeNameFlag}, "filter", true)
	cmd.Flags().BoolVar(&args.Unlock, UnlockFlag, false, "Unlocks the settings node before saving if needed")

	return cmd
}

func RunSettingsEditCommand(args *EditArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		if err := stscobra.CheckMutuallyExclusiveFlags(cmd, []string{IdsFlag, TypeNameFlag}, true); err != nil {
			return common.NewCLIArgParseError(err)
		}

		orig, resp, err := DoExport(cli.Context, api, args.Ids, "", args.NodeTypes, args.AllowReferences)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		c, err := cli.Editor.Edit("settings", ".json", strings.NewReader(orig))
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if strings.Compare(orig, string(c)) == 0 {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"updated-settings": []string{}, "message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}

			return nil
		}

		if args.Unlock {
			// NOTE Uses the wildcard type "_" to unlock all the nodes identified with ids.
			for _, id := range args.Ids {
				_, resp, err := api.NodeApi.Unlock(cli.Context, "_", id).Execute()
				if err != nil {
					return common.NewResponseError(err, resp)
				}
			}
		}

		nodes, resp, err := doImport(cli.Context, api, string(c), "", args.UnlockedStrategy, args.Timeout)
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
