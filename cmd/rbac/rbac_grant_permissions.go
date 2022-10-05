package rbac

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type GrantPermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
}

func GrantPermissionsCommand(deps *di.Deps) *cobra.Command {
	const (
		PermissionUsage = "The permission to grant"
		ResourceUsage   = "The resource to grant the permission to (e.g. \"system\" or a view name)"
		DefaultResource = "system"
	)

	args := &GrantPermissionsArgs{}
	cmd := &cobra.Command{
		Use:   "grant",
		Short: "Grant a permission to a subject",
		Long:  "Grant a permission to a subject.",
		RunE:  deps.CmdRunEWithApi(RunGrantPermissionsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionUsage)
	cmd.MarkFlagRequired(Permission) //nolint:errcheck

	cmd.Flags().StringVar(&args.Resource, Resource, DefaultResource, ResourceUsage)

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

		description, descrResp, descrErr := api.PermissionsApi.DescribePermissions(cli.Context, args.Subject).Execute()

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
			printPermissionsTable(cli, description.Permissions)
		}

		return nil
	}
}
