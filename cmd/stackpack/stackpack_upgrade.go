package stackpack

import (
	"fmt"
	"strings"

	"github.com/stackvista/stackstate-cli/pkg/pflags"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
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
		Short: "Upgrade a StackPack",
		Long:  "Upgrade a StackPack instance to the next version.",
		RunE:  cli.CmdRunEWithApi(RunStackpackUpgradeCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.TypeName, "Name of the StackPack")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"Strategy use to upgrade StackPack instance"+
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
		stackPackList, resp, err := api.StackpackAPI.StackPackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		stack, err := findStackName(stackPackList, args.TypeName)
		if err != nil {
			return common.NewNotFoundError(err)
		}
		if !stack.HasNextVersion() {
			return common.NewNotFoundError(fmt.Errorf("stackpack %s cannot be upgraded at this moment", args.TypeName))
		}
		_, resp, err = api.StackpackAPI.UpgradeStackPack(cli.Context, args.TypeName).Unlocked(args.UnlockedStrategy).Execute()
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

func findStackName(stacks []stackstate_api.FullStackPack, name string) (stackstate_api.FullStackPack, error) {
	for _, v := range stacks {
		if v.GetName() == name {
			return v, nil
		}
	}
	return stackstate_api.FullStackPack{}, fmt.Errorf("stackpack %s does not exist", name)
}
