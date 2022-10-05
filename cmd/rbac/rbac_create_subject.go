package rbac

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const STQLVersion = "0.0.1"

type CreateSubjectArgs struct {
	Subject string
	Scope   string
}

func CreateSubjectCommand(deps *di.Deps) *cobra.Command {
	const ScopeUsage = "The query in STQL that will be prepended to every topology element retrieved in StackState. " +
		"For example, if your scope is \"label = 'A'\", then all STQL executed in StackState" +
		" (e.g. Retrieving topology) will only return elements that have the label A"

	args := &CreateSubjectArgs{}
	cmd := &cobra.Command{
		Use:   "create-subject",
		Short: "Create a new security subject",
		Long:  "Create a new security subject.",
		RunE:  deps.CmdRunEWithApi(RunCreateSubjectCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Scope, Scope, ScopeDefault, ScopeUsage)

	return cmd
}

func RunCreateSubjectCommand(args *CreateSubjectArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		subject := stackstate_api.NewCreateSubject(args.Scope, STQLVersion)
		resp, err := api.SubjectApi.CreateSubject(cli.Context, args.Subject).
			CreateSubject(*subject).
			Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"create-subject": args.Subject,
			})
		} else {
			cli.Printer.Successf("Created subject '%s'", args.Subject)
		}

		return nil
	}
}
