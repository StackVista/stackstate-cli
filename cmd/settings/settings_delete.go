package settings

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	Ids     []int64
	Timeout int64
}

func SettingsDeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete settings by ID",
		Long:  "Delete one or more settings by their IDs. This operation cannot be undone.",
		Example: `# delete a setting
sts settings delete --ids 123456789

# delete multiple settings
sts settings delete --ids 123,456,789`,
		RunE: cli.CmdRunEWithApi(RunSettingsDeleteCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "The IDs of nodes to delete")
	cmd.MarkFlagRequired(IdsFlag) //nolint:errcheck

	cmd.Flags().Int64VarP(&args.Timeout, TimeoutFlag, TimeoutFlagShort, DefaultTimeout, TimeoutUsage)

	return cmd
}

func DeleteNodes(cli *di.Deps, api *stackstate_api.APIClient, ids []int64, nodeType string, timeout int64) common.CLIError {
	for _, id := range ids {
		resp, err := api.NodeApi.Delete(cli.Context, nodeType, id).TimeoutSeconds(timeout).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
	}
	return nil
}

func RunSettingsDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		// NOTE Uses the wildcard type "_" to delete all the nodes as we don't want to fetch each one to figure out the type.
		err := DeleteNodes(cli, api, args.Ids, WildCardType, args.Timeout)
		if err != nil {
			return err
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted": args.Ids,
			})
		} else {
			ids := make([]string, len(args.Ids))
			for i, id := range args.Ids {
				ids[i] = fmt.Sprintf("%d", id)
			}
			cli.Printer.Successf("The following nodes were deleted: %s", strings.Join(ids, ", "))
		}

		return nil
	}
}
