package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type DescribeSubjectsArgs struct {
	Subject string
}

func DescribeSubjectsCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeSubjectsArgs{}
	cmd := &cobra.Command{
		Use:   "describe-subjects",
		Short: "List the subjects available in StackState",
		Long:  "List the subjects available in StackState.",
		RunE:  deps.CmdRunEWithApi(RunDescribeSubjectsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)

	return cmd
}

func RunDescribeSubjectsCommand(args *DescribeSubjectsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if args.Subject != "" {
			subject, resp, err := api.SubjectApi.GetSubject(cli.Context, args.Subject).Execute()

			if err != nil {
				return common.NewResponseError(err, resp)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"handle": subject.Handle,
				})
			} else {
				cli.Printer.Table(printer.TableData{
					Header: []string{"Subject"},
					Data: [][]interface{}{
						{
							subject.Handle,
						},
					},
					MissingTableDataMsg: printer.NotFoundMsg{Types: "matching subjects"},
				})
			}
		} else {
			subjects, resp, err := api.SubjectApi.ListSubjects(cli.Context).Execute()

			if err != nil {
				return common.NewResponseError(err, resp)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"subjects": subjects,
				})
			} else {
				data := make([][]interface{}, 0)

				for _, subject := range subjects {
					data = append(data, []interface{}{subject.Handle})
				}

				cli.Printer.Table(printer.TableData{
					Header:              []string{"Subject"},
					Data:                data,
					MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
				})
			}
		}
		return nil
	}
}
