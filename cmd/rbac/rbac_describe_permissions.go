package rbac

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type DescribePermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
}

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
		description, resp, err := describePermissions(cli, api, args.Subject, args.Permission, args.Resource).Execute()

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

func describePermissions(cli *di.Deps, api *stackstate_api.APIClient, subject string, permission string, resource string) stackstate_api.ApiDescribePermissionsRequest {
	request := api.PermissionsApi.DescribePermissions(cli.Context, subject)
	if permission != "" {
		request = request.Permission(permission)
	}
	if resource != "" {
		request = request.Resource(resource)
	}
	return request
}

func printPermissionsTable(cli *di.Deps, permissionsList map[string][]string) {
	keys := make([]string, len(permissionsList))
	for key := range permissionsList {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	data := make([][]interface{}, 0)
	for _, resource := range keys {
		permissions := permissionsList[resource]
		sort.SliceStable(permissions, func(i, j int) bool {
			return permissions[i] < permissions[j]
		})
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
