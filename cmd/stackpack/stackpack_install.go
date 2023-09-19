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
}

func StackpackInstallCommand(cli *di.Deps) *cobra.Command {
	args := &InstallArgs{
		Params: make(map[string]string),
	}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install a StackPack",
		Long: "Install an instance of a StackPack. " +
			"Be aware that each StackPack has a different set of parameters. " +
			"Run \"sts stackpack list-parameters --name NAME\" to list all of them.",
		Example: "# install an example StackPack with a full name and URL parameter\n" +
			"sts stackpack install --name example -p \"full_name=First Last\" -p URL=https://stackstate.com",
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

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"instance": instance,
			})
		} else {
			lastUpdateTime := time.UnixMilli(instance.GetLastUpdateTimestamp())

			cli.Printer.Success("StackPack instance installed")
			cli.Printer.Table(
				printer.TableData{
					Header:              []string{"id", "name", "status", "version", "last updated"},
					Data:                [][]interface{}{{instance.Id, instance.Name, instance.Status, instance.StackPackVersion, lastUpdateTime}},
					MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + args.Name},
				},
			)
		}

		return nil
	}
}
