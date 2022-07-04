package health

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type ClearErrorArgs struct {
	Urn string
}

func HealthClearErrorCommand(cli *di.Deps) *cobra.Command {
	args := &ClearErrorArgs{}
	cmd := &cobra.Command{
		Use:   "clear-error",
		Short: "Clear errors from a stream",
		Long:  "Clear errors from a health synchronization stream. More info at https://l.stackstate.com/cli-health-synchronization.",
		RunE:  cli.CmdRunEWithApi(RunHealthClearErrorCommand(args)),
	}
	common.AddRequiredUrnFlagVar(cmd, &args.Urn, "Urn of the stream")
	return cmd
}

func RunHealthClearErrorCommand(args *ClearErrorArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		resp, err := api.HealthSynchronizationApi.PostHealthSynchronizationStreamClearErrors(cli.Context, args.Urn).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"stream-error-clear": args.Urn,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Stream error clear: %s", args.Urn))
		}

		return nil
	}
}
