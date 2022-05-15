package stackpack

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func StackpackUpgradeCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade a StackPack",
		Long:  "Upgrade SpackPack instance",
		RunE:  cli.CmdRunEWithApi(RunStackpackUpgradeCommand),
	}
	common.AddRequiredNameFlag(cmd, "name of the StackPack")
	cmd.Flags().String(UnlockedStrategyFlag, "", "unlocked-strategy flag must be one of {fail,overwrite,skip}")
	cmd.MarkFlagRequired(UnlockedStrategyFlag) //nolint:errcheck
	return cmd
}

func RunStackpackUpgradeCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(common.NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	unlockStrategy, err := cmd.Flags().GetString(UnlockedStrategyFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	if !isValidStrategy(unlockStrategy) {
		return common.NewCLIArgParseError(errors.New("only one of {fail,overwrite,skip} is accepted"))
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	stack, err := findStackName(stackpackList, name)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	if !stack.HasNextVersion() {
		return common.NewCLIArgParseError(fmt.Errorf("stackpack %s cannot be upgraded at this moment", name))
	}
	_, resp, err = api.StackpackApi.UpgradeStackPack(cli.Context, name).Unlocked(unlockStrategy).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"upgrade": fmt.Sprintf("Successfully triggered upgrade of StackPack %s", name),
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Successfully triggered upgrade of StackPack %s", name))
	}

	return nil
}

func isValidStrategy(name string) bool {
	return name == "fail" || name == "overwrite" || name == "skip"
}

func findStackName(stacks []stackstate_api.Sstackpack, name string) (stackstate_api.Sstackpack, error) {
	for _, v := range stacks {
		if v.GetName() == name {
			return v, nil
		}
	}
	return stackstate_api.Sstackpack{}, fmt.Errorf("stackpack %s does not exist", name)
}
