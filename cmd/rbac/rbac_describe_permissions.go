package rbac

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type DescribePermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
}

type Permissions map[string][]string

func DescribePermissionsCommand(deps *di.Deps) *cobra.Command {
	args := &DescribePermissionsArgs{}
	cmd := &cobra.Command{
		Use:   "describe-permissions",
		Short: "Show permissions of a subject",
		Long:  "Show permissions of a subject.",
		RunE:  deps.CmdRunEWithApi(RunDescribePermissionsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionDescribeUsage)
	cmd.Flags().StringVar(&args.Resource, Resource, "", ResourceDescribeUsage)

	return cmd
}

func RunDescribePermissionsCommand(args *DescribePermissionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		request := api.PermissionsApi.DescribePermissions(cli.Context, args.Subject)
		if args.Permission != "" {
			request = request.Permission(args.Permission)
		}
		if args.Resource != "" {
			request = request.Resource(args.Resource)
		}

		description, resp, err := request.Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"subject":     description.SubjectHandle,
				"permissions": description.Permissions,
			})
		} else {
			printPermissionsTable(cli, description.Permissions)
		}

		return nil
	}
}

func printPermissionsTable(cli *di.Deps, permissionsList Permissions) {
	data := make([][]interface{}, 0)
	for resource, permissions := range permissionsList {
		for _, permission := range permissions {
			data = append(data, []interface{}{permission, resource})
		}
	}

	cli.Printer.Table(printer.TableData{
		Header:              []string{"permission", "resource"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
	})
}
