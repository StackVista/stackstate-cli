package stackpack

import (
	"fmt"
	"strings"

	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var (
	UnlockedStrategyChoices = []string{"fail", "skip", "overwrite"}
)

type UpgradeArgs struct {
	TypeName         string
	UnlockedStrategy string
}

func StackpackUpgradeCommand(cli *di.Deps) *cobra.Command {
	args := &UpgradeArgs{}
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade a StackPack",
		Long:  "Upgrade a StackPack instance to the next version.",
		RunE:  cli.CmdRunEWithApi(RunStackpackUpgradeCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.TypeName, "name of the StackPack")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"strategy use to upgrade StackPack instance"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")),
	)
	return cmd
}
func RunStackpackUpgradeCommand(args *UpgradeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		stack, err := findStackName(stackpackList, args.TypeName)
		if err != nil {
			return common.NewNotFoundError(err)
		}
		if !stack.HasNextVersion() {
			return common.NewNotFoundError(fmt.Errorf("stackpack %s cannot be upgraded at this moment", args.TypeName))
		}
		_, resp, err = api.StackpackApi.UpgradeStackPack(cli.Context, args.TypeName).Unlocked(args.UnlockedStrategy).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success":         true,
				"current-version": stack.GetVersion(),
				"next-version":    stack.NextVersion.GetVersion(),
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Successfully triggered upgrade from %s to %s", stack.GetVersion(), stack.NextVersion.GetVersion()))
		}

		return nil
	}
}

func findStackName(stacks []stackstate_api.Sstackpack, name string) (stackstate_api.Sstackpack, error) {
	for _, v := range stacks {
		if v.GetName() == name {
			return v, nil
		}
	}
	return stackstate_api.Sstackpack{}, fmt.Errorf("stackpack %s does not exist", name)
}
