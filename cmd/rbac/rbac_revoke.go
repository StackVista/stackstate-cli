package rbac

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type RevokePermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
}

func RevokePermissionsCommand(deps *di.Deps) *cobra.Command {
	args := &RevokePermissionsArgs{}
	cmd := &cobra.Command{
		Use:   "revoke --subject SUBJECT --permission PERMISSION",
		Short: "Revoke a permission from a subject",
		Long:  "Revoke a previously granted permission from a subject on a resource.",
		Example: `# revoke access-view permission from all resources
sts rbac revoke --subject my-team --permission access-view

# revoke permission on a specific resource
sts rbac revoke --subject my-team --permission access-view --resource "label = 'production'"`,
		RunE: deps.CmdRunEWithApi(RunRevokePermissionsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionRevokeUsage)
	cmd.MarkFlagRequired(Permission) //nolint:errcheck

	cmd.Flags().StringVar(&args.Resource, Resource, DefaultResource, ResourceRevokeUsage)

	return cmd
}

func RunRevokePermissionsCommand(args *RevokePermissionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		revokeResp, revokeErr := api.PermissionsApi.RevokePermissions(cli.Context, args.Subject).
			Resource(args.Resource).
			Permission(args.Permission).
			Execute()

		if revokeErr != nil {
			return common.NewResponseError(revokeErr, revokeResp)
		}

		description, descrResp, descrErr := describePermissions(cli, api, args.Subject, args.Permission, "", "").Execute()

		if descrErr != nil {
			return common.NewResponseError(descrErr, descrResp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"subject":     description.SubjectHandle,
				"permissions": description.Permissions,
			})
		} else {
			cli.Printer.Successf("Revoked permission '%s' on '%s' from subject '%s'", args.Permission, args.Resource, args.Subject)
			printPermissionsTable(cli, nil, description.Permissions)
		}

		return nil
	}
}
