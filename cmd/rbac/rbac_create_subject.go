package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type CreateSubjectArgs struct {
	Subject string
	Scope   string
}

func CreateSubjectCommand(deps *di.Deps) *cobra.Command {
	args := &CreateSubjectArgs{}
	cmd := &cobra.Command{
		Use:   "create-subject",
		Short: "Create a new security subject",
		Long:  "Create a new security subject.",
		RunE:  deps.CmdRunEWithApi(RunCreateSubjectCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Scope, Scope, DefaultScope, ScopeUsage)

	return cmd
}

func RunCreateSubjectCommand(args *CreateSubjectArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		subject := stackstate_api.NewCreateSubject(args.Scope, DefaultSTQLVersion)
		resp, err := api.SubjectAPI.CreateSubject(cli.Context, args.Subject).
			CreateSubject(*subject).
			Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"created-subject": args.Subject,
			})
		} else {
			cli.Printer.Successf("Created subject '%s'", args.Subject)
		}

		return nil
	}
}
