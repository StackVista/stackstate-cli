package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type GrantPermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
}

func GrantPermissionsCommand(deps *di.Deps) *cobra.Command {
	args := &GrantPermissionsArgs{}
	cmd := &cobra.Command{
		Use:   "grant --subject SUBJECT --permission PERMISSION",
		Short: "Grant a permission to a subject",
		Long:  "Grant a permission to a subject on a resource. Use 'sts rbac list-permissions' to see available permissions.",
		Example: `# grant access-view permission to all resources
sts rbac grant --subject my-team --permission access-view

# grant permission on a specific resource
sts rbac grant --subject my-team --permission access-view --resource "label = 'production'"`,
		RunE: deps.CmdRunEWithApi(RunGrantPermissionsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionGrantUsage)
	cmd.MarkFlagRequired(Permission) //nolint:errcheck

	cmd.Flags().StringVar(&args.Resource, Resource, DefaultResource, ResourceGrantUsage)

	return cmd
}

func RunGrantPermissionsCommand(args *GrantPermissionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		permission := stackstate_api.NewGrantPermission(args.Permission, args.Resource)
		grantResp, grantErr := api.PermissionsApi.GrantPermissions(cli.Context, args.Subject).
			GrantPermission(*permission).
			Execute()

		if grantErr != nil {
			return common.NewResponseError(grantErr, grantResp)
		}

		description, descrResp, descrErr := describePermissions(cli, api, args.Subject, args.Permission, args.Resource, "").Execute()

		if descrErr != nil {
			return common.NewResponseError(descrErr, descrResp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"subject":     description.SubjectHandle,
				"permissions": description.Permissions,
			})
		} else {
			cli.Printer.Successf("Granted permission '%s' on '%s' to subject '%s'", args.Permission, args.Resource, args.Subject)
			printPermissionsTable(cli, nil, description.Permissions)
		}

		return nil
	}
}
