package rbac

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

const (
	PermissionUsage = "Filter the permissions by permission name"
	ResourceUsage   = "Filter the permissions by a resource identifier (e.g. system or a view name)"
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

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionUsage)
	cmd.Flags().StringVar(&args.Resource, Resource, "", ResourceUsage)

	return cmd
}

func RunDescribePermissionsCommand(args *DescribePermissionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		description, resp, err := api.PermissionsApi.DescribePermissions(cli.Context, args.Subject).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		filtered := filterPermissions(description.Permissions, args)

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"subject":     description.SubjectHandle,
				"permissions": filtered,
			})
		} else {
			data := make([][]interface{}, 0)
			for resource, permissions := range filtered {
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

		return nil
	}
}

func filterPermissions(permissionsList Permissions, args *DescribePermissionsArgs) Permissions {
	filteredResources := make(Permissions, 0)
	if args.Resource != "" {
		for resource, permissions := range permissionsList {
			if resource == args.Resource {
				filteredResources[resource] = permissions
			}
		}
	} else {
		filteredResources = permissionsList
	}

	filteredPermissions := make(Permissions, 0)
	if args.Permission != "" {
		for resource, permissions := range filteredResources {
			for _, permission := range permissions {
				if permission == args.Permission {
					filteredPermissions[resource] = []string{permission}
				}
			}
		}
	} else {
		filteredPermissions = filteredResources
	}

	return filteredPermissions
}
