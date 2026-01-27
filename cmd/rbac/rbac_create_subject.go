//nolint:dupl
package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type CreateSubjectArgs struct {
	Subject    string
	Scope      string
	CreateOnly bool
}

func CreateSubjectCommand(deps *di.Deps) *cobra.Command {
	args := &CreateSubjectArgs{}
	cmd := &cobra.Command{
		Use:   "create-subject --subject SUBJECT",
		Short: "Create a new security subject",
		Long:  "Create a new security subject for RBAC. Subjects can be users or groups that can be granted permissions.",
		Example: `# create a new subject
sts rbac create-subject --subject my-team

# create a subject with a scope query
sts rbac create-subject --subject my-team --scope "label = 'production'"`,
		RunE: deps.CmdRunEWithApi(RunCreateSubjectCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Scope, Scope, "", ScopeUsage)
	cmd.Flags().BoolVar(&args.CreateOnly, CreateOnly, false, CreateOnlyUsage)

	return cmd
}

func RunCreateSubjectCommand(args *CreateSubjectArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		subject := stackstate_api.NewCreateSubject()
		if args.Scope != "" {
			subject.SetQuery(args.Scope)
			subject.SetVersion("0.0.1")
		}

		resp, err := api.SubjectApi.CreateSubject(cli.Context, args.Subject).
			CreateSubject(*subject).
			CreateOnly(args.CreateOnly).
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
