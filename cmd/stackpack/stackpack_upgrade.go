package stackpack

import (
	"fmt"
	"strings"
	"time"

	"github.com/stackvista/stackstate-cli/pkg/pflags"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

var (
	UnlockedStrategyChoices = []string{"fail", "skip", "overwrite"}
)

type UpgradeArgs struct {
	TypeName         string
	UnlockedStrategy string
	Wait             bool
	Timeout          time.Duration
}

func StackpackUpgradeCommand(cli *di.Deps) *cobra.Command {
	args := &UpgradeArgs{}
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade a StackPack to the next version",
		Long:  "Upgrade all instances of a StackPack to the next available version. Check 'sts stackpack list' to see available upgrades.",
		Example: `# upgrade a StackPack
sts stackpack upgrade --name kubernetes

# upgrade and wait for completion
sts stackpack upgrade --name kubernetes --wait`,
		RunE: cli.CmdRunEWithApi(RunStackpackUpgradeCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.TypeName, "Name of the StackPack")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"",
		UnlockedStrategyChoices,
		"Strategy use to upgrade StackPack instance"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")),
	)
	cmd.Flags().BoolVar(&args.Wait, "wait", false, "Wait for upgrade to complete")
	cmd.Flags().DurationVar(&args.Timeout, "timeout", DefaultTimeout, "Timeout for waiting")
	return cmd
}
func RunStackpackUpgradeCommand(args *UpgradeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackPackList, resp, err := api.StackpackApi.StackPackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		stack, err := findStackName(stackPackList, args.TypeName)
		if err != nil {
			return common.NewNotFoundError(err)
		}
		if !stack.HasNextVersion() {
			return common.NewRuntimeError(fmt.Errorf("stackpack %s cannot be upgraded at this moment", args.TypeName))
		}
		_, resp, err = api.StackpackApi.UpgradeStackPack(cli.Context, args.TypeName).Unlocked(args.UnlockedStrategy).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		// Wait functionality: monitor upgrade until completion
		if args.Wait {
			if !cli.IsJson() {
				cli.Printer.PrintLn("Waiting for upgrade to complete...")
			}

			// Use OperationWaiter to poll until all configurations are upgraded
			waiter := NewOperationWaiter(cli, api)
			waitErr := waiter.WaitForCompletion(WaitOptions{
				StackPackName: args.TypeName,
				Timeout:       args.Timeout,
				PollInterval:  DefaultPollInterval,
			})
			if waitErr != nil {
				// Use NewRuntimeError to avoid showing usage on operation failures
				return common.NewRuntimeError(waitErr)
			}

			// Re-fetch final status for display after successful completion
			stackPackList, cliErr := fetchAllStackPacks(cli, api)
			if cliErr != nil {
				return cliErr
			}

			finalStackPack, err := findStackPackByName(stackPackList, args.TypeName)
			if err != nil {
				return common.NewNotFoundError(err)
			}

			// Display final status
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"stackpack":       finalStackPack,
					"status":          "completed",
					"current-version": finalStackPack.GetVersion(),
				})
			} else {
				cli.Printer.Success("StackPack upgrade completed successfully")

				// Show configurations status
				data := make([][]interface{}, 0)
				for _, config := range finalStackPack.GetConfigurations() {
					lastUpdateTime := time.UnixMilli(config.GetLastUpdateTimestamp())
					data = append(data, []interface{}{
						config.GetId(),
						finalStackPack.GetName(),
						config.GetStatus(),
						config.GetStackPackVersion(),
						lastUpdateTime,
					})
				}

				cli.Printer.Table(
					printer.TableData{
						Header:              []string{"id", "name", "status", "version", "last updated"},
						Data:                data,
						MissingTableDataMsg: printer.NotFoundMsg{Types: "configurations for " + args.TypeName},
					},
				)
			}
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"success":         true,
					"current-version": stack.GetVersion(),
					"next-version":    stack.NextVersion.GetVersion(),
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("Successfully triggered upgrade from %s to %s", stack.GetVersion(), stack.NextVersion.GetVersion()))
			}
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
