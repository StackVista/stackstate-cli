package health

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	Urn string
}

func HealthDeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete --urn URN",
		Short: "Delete a health synchronization stream",
		Long:  "Delete a health synchronization stream and all its sub-streams. This removes all check states associated with the stream.",
		Example: `# delete a health stream
sts health delete --urn urn:health:my-stream`,
		RunE: cli.CmdRunEWithApi(RunHealthDeleteCommand(args)),
	}
	common.AddRequiredUrnFlagVar(cmd, &args.Urn, "URN of the health synchronization stream to delete")

	return cmd
}

func RunHealthDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		resp, err := api.HealthSynchronizationApi.DeleteHealthSynchronizationStream(cli.Context, args.Urn).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-stream-urn": args.Urn,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Stream deleted: %s", args.Urn))
		}
		return nil
	}
}
