package rbac

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

type DescribeSubjectsArgs struct {
	Subject string
	Source  string
}

func DescribeSubjectsCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeSubjectsArgs{}
	cmd := &cobra.Command{
		Use:   "describe-subjects",
		Short: "List all subjects or describe a specific subject",
		Long:  "List all security subjects or describe a specific subject. Shows subject handle and source.",
		Example: `# list all subjects
sts rbac describe-subjects

# describe a specific subject
sts rbac describe-subjects --subject my-team`,
		RunE: deps.CmdRunEWithApi(RunDescribeSubjectsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	pflags.EnumVar(cmd.Flags(), &args.Source, Source, "", SourceChoices, SourceUsage)

	return cmd
}

func RunDescribeSubjectsCommand(args *DescribeSubjectsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		// Why do wre use the list api instead of the GetSubject api for retrieving a subject? The reason is that due to rbac
		// subjects can come from multiple sources. Instead of updating the GetSubject api in a breaking way, we have deprecated that
		// api and use the list api with a client side filter to get all subjects.
		subjects, resp, err := api.SubjectApi.ListSubjects(cli.Context).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		filteredSubjects := make([]stackstate_api.SubjectConfig, 0)

		for _, subject := range subjects {
			if (args.Subject == "" || subject.Handle == args.Subject) && (args.Source == "" || strings.EqualFold(string(subject.Source), args.Source)) {
				filteredSubjects = append(filteredSubjects, subject)
			}
		}

		if len(filteredSubjects) == 0 && args.Subject != "" {
			return common.NewNotFoundError(fmt.Errorf("could not find subject: '%s'", args.Subject))
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"subjects": filteredSubjects,
				})
			} else {
				data := make([][]interface{}, 0)

				for _, subject := range filteredSubjects {
					data = append(data, []interface{}{subject.Handle, subject.Source})
				}

				cli.Printer.Table(printer.TableData{
					Header:              []string{"Subject", "Source"},
					Data:                data,
					MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
				})
			}
		}
		return nil
	}
}
