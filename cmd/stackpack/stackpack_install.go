package stackpack

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

type InstallArgs struct {
	Name             string
	UnlockedStrategy string
	Params           map[string]string
	Wait             bool
	Timeout          time.Duration
}

func StackpackInstallCommand(cli *di.Deps) *cobra.Command {
	args := &InstallArgs{
		Params: make(map[string]string),
	}
	cmd := &cobra.Command{
		Use:   "install --name NAME",
		Short: "Install a new instance of a StackPack",
		Long:  `Install a new instance of a StackPack. Use 'sts stackpack list-parameters --name NAME' to see required parameters.`,
		Example: `# install an example StackPack with parameters
sts stackpack install --name example -p "full_name=First Last" -p URL=https://stackstate.com

# install and wait for completion
sts stackpack install --name kubernetes -p cluster_name=production --wait`,
		RunE: cli.CmdRunEWithApi(RunStackpackInstallCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the StackPack")
	pflags.EnumVar(cmd.Flags(), &args.UnlockedStrategy,
		UnlockedStrategyFlag,
		"fail",
		UnlockedStrategyChoices,
		"Strategy use to upgrade StackPack instance"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")),
	)
	cmd.Flags().StringToStringVarP(&args.Params, ParameterFlag, "p", args.Params, "List of parameters of the form \"key=value\"")
	cmd.Flags().BoolVar(&args.Wait, "wait", false, "Wait for installation to complete")
	cmd.Flags().DurationVar(&args.Timeout, "timeout", DefaultTimeout, "Timeout for waiting")
	return cmd
}

func RunStackpackInstallCommand(args *InstallArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		instance, resp, err := api.StackpackApi.ProvisionDetails(cli.Context, args.Name).RequestBody(args.Params).Unlocked(args.UnlockedStrategy).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		// Wait functionality: monitor installation until completion
		if args.Wait {
			if !cli.IsJson() {
				cli.Printer.PrintLn("Waiting for installation to complete...")
			}

			// Use OperationWaiter to poll until all configurations are installed
			waiter := NewOperationWaiter(cli, api)
			waitErr := waiter.WaitForCompletion(WaitOptions{
				StackPackName: args.Name,
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

			finalStackPack, err := findStackPackByName(stackPackList, args.Name)
			if err != nil {
				return common.NewNotFoundError(err)
			}

			// Display final status
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"stackpack": finalStackPack,
					"status":    "completed",
				})
			} else {
				cli.Printer.Success("StackPack installation completed successfully")

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
						MissingTableDataMsg: printer.NotFoundMsg{Types: "configurations for " + args.Name},
					},
				)
			}
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"instance": instance,
				})
			} else {
				lastUpdateTime := time.UnixMilli(instance.GetLastUpdateTimestamp())

				cli.Printer.Success("StackPack instance installation triggered")
				cli.Printer.Table(
					printer.TableData{
						Header:              []string{"id", "name", "status", "version", "last updated"},
						Data:                [][]interface{}{{instance.Id, instance.Name, instance.Status, instance.StackPackVersion, lastUpdateTime}},
						MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + args.Name},
					},
				)
			}
		}

		return nil
	}
}
