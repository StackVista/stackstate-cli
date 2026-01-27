package health

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type ClearErrorArgs struct {
	Urn string
}

func HealthClearErrorCommand(cli *di.Deps) *cobra.Command {
	args := &ClearErrorArgs{}
	cmd := &cobra.Command{
		Use:   "clear-error --urn URN",
		Short: "Clear errors from a health synchronization stream",
		Long: `Clear errors from a health synchronization stream, allowing it to resume normal operation. Use this after resolving the underlying cause of synchronization errors.

More info: https://l.stackstate.com/cli-health-synchronization.`,
		Example: `# clear errors from a health stream
sts health clear-error --urn urn:health:my-stream`,
		RunE: cli.CmdRunEWithApi(RunHealthClearErrorCommand(args)),
	}
	common.AddRequiredUrnFlagVar(cmd, &args.Urn, "URN of the health synchronization stream")
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
			msg := fmt.Sprintf("Stream error clear: %s", args.Urn)
			cli.Printer.Success(msg)
		}

		return nil
	}
}
