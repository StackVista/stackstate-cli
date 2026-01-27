package topologysync

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type ClearErrorsArgs struct {
	ID         int64
	Identifier string
}

func ClearErrorsCommand(deps *di.Deps) *cobra.Command {
	args := &ClearErrorsArgs{}
	cmd := &cobra.Command{
		Use:   "clear-errors {--id ID | --identifier URN}",
		Short: "Clear errors from a topology synchronization",
		Long:  "Clear all errors from a topology synchronization. Use this after resolving the underlying data issues.",
		Example: `# clear errors by ID
sts topology-sync clear-errors --id 123456789

# clear errors by identifier
sts topology-sync clear-errors --identifier urn:stackpack:kubernetes:sync`,
		RunE: deps.CmdRunEWithApi(RunClearErrorsCommand(args)),
	}
	cmd.Flags().Int64VarP(&args.ID, common.IDFlag, common.IDFlagShort, 0, "The ID of a topology synchronization")
	cmd.Flags().StringVar(&args.Identifier, common.IdentifierFlag, "", "The identifier of a topology synchronization")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunClearErrorsCommand(args *ClearErrorsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		idType, id := IdOrIdentifier(args.ID, args.Identifier)
		resp, err := api.TopologySynchronizationApi.
			PostTopologySynchronizationStreamClearErrors(cli.Context).
			IdentifierType(idType).
			Identifier(id).
			Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"cleared": id,
			})
		} else {
			cli.Printer.Successf("Errors cleared from topology synchronization: %s", id)
		}
		return nil
	}
}
