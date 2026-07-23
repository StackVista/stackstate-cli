package stackpack

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type DowngradeArgs struct {
	TypeName string
	Version  string
	Wait     bool
	Timeout  time.Duration
}

func StackpackDowngradeCommand(cli *di.Deps) *cobra.Command {
	args := &DowngradeArgs{}
	cmd := &cobra.Command{
		Use:   "downgrade",
		Short: "Downgrade a StackPack to an older version",
		Long: "Downgrade all instances of a StackPack to an older version. " +
			"Only supported for StackPacks 2.0 and the target version must be lower than the currently installed version. " +
			"To move a StackPacks 1.0 instance to an older version, uninstall it and install the desired version instead.",
		Example: `# downgrade a StackPack to an older version
sts stackpack downgrade --name kubernetes --stackpack-version 1.2.3

# downgrade and wait for completion
sts stackpack downgrade --name kubernetes --stackpack-version 1.2.3 --wait`,
		RunE: cli.CmdRunEWithApi(RunStackpackDowngradeCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.TypeName, "Name of the StackPack")
	cmd.Flags().StringVar(&args.Version, StackpackVersionFlag, "", "Version to downgrade to")
	cmd.MarkFlagRequired(StackpackVersionFlag)
	cmd.Flags().BoolVar(&args.Wait, "wait", false, "Wait for downgrade to complete")
	cmd.Flags().DurationVar(&args.Timeout, "timeout", DefaultTimeout, "Timeout for waiting")
	return cmd
}

func RunStackpackDowngradeCommand(args *DowngradeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		_, resp, err := api.StackpackApi.DowngradeStackPack(cli.Context, args.TypeName).
			Version(args.Version).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if args.Wait {
			if !cli.IsJson() {
				cli.Printer.PrintLn("Waiting for downgrade to complete...")
			}

			waiter := NewOperationWaiter(cli, api)
			waitErr := waiter.WaitForCompletion(WaitOptions{
				StackPackName: args.TypeName,
				Timeout:       args.Timeout,
				PollInterval:  DefaultPollInterval,
			})
			if waitErr != nil {
				return common.NewRuntimeError(waitErr)
			}

			stackPackList, cliErr := fetchAllStackPacks(cli, api)
			if cliErr != nil {
				return cliErr
			}

			finalStackPack, err := findStackPackByName(stackPackList, args.TypeName)
			if err != nil {
				return common.NewNotFoundError(err)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"stackpack":       finalStackPack,
					"status":          "completed",
					"current-version": finalStackPack.GetVersion(),
				})
			} else {
				cli.Printer.Success("StackPack downgrade completed successfully")

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
					"success":        true,
					"target-version": args.Version,
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("Successfully triggered downgrade of %s to %s", args.TypeName, args.Version))
			}
		}

		return nil
	}
}
