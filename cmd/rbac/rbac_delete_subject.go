package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteSubjectArgs struct {
	Subject string
}

func DeleteSubjectCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteSubjectArgs{}
	cmd := &cobra.Command{
		Use:   "delete-subject",
		Short: "Delete a security subject",
		Long:  "Delete a security subject from StackState.",
		RunE:  deps.CmdRunEWithApi(RunDeleteSubjectCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	return cmd
}

func RunDeleteSubjectCommand(args *DeleteSubjectArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		resp, err := api.SubjectApi.DeleteSubject(cli.Context, args.Subject).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-subject": args.Subject,
			})
		} else {
			cli.Printer.Successf("Deleted subject '%s'", args.Subject)
		}

		return nil
	}
}
