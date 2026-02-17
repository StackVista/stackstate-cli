package settings

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const WildCardType = "_"

type UnlockArgs struct {
	Ids []int64
}

func SettingsUnlockCommand(cli *di.Deps) *cobra.Command {
	args := &UnlockArgs{}
	cmd := &cobra.Command{
		Use:   "unlock",
		Short: "Unlock settings locked by StackPacks",
		Long:  `Unlock settings locked by StackPacks to allow editing. Warning: Unlocking settings may prevent StackPacks from upgrading correctly.`,
		Example: `# unlock a setting by ID
sts settings unlock --ids 123456789

# unlock multiple settings
sts settings unlock --ids 123,456,789`,
		RunE: cli.CmdRunEWithApi(RunSettingsUnlockCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "The IDs of nodes to unlock")
	cmd.MarkFlagRequired(IdsFlag) //nolint:errcheck

	return cmd
}

func UnlockNodes(cli *di.Deps, api *stackstate_api.APIClient, ids []int64, nodeType string) common.CLIError {
	for _, id := range ids {
		_, resp, err := api.NodeApi.Unlock(cli.Context, nodeType, id).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
	}
	return nil
}

func RunSettingsUnlockCommand(args *UnlockArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		// NOTE Uses the wildcard type "_" to unlock all the nodes identified with ids.
		err := UnlockNodes(cli, api, args.Ids, WildCardType)
		if err != nil {
			return err
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"unlocked": args.Ids,
			})
		} else {
			ids := make([]string, len(args.Ids))
			for i, id := range args.Ids {
				ids[i] = fmt.Sprintf("%d", id)
			}
			cli.Printer.Successf("The following nodes were unlocked: %s", strings.Join(ids, ", "))
		}

		return nil
	}
}
