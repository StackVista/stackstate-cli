package rbac

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

type DescribePermissionsArgs struct {
	Subject    string
	Permission string
	Resource   string
	Source     string
}

func DescribePermissionsCommand(deps *di.Deps) *cobra.Command {
	args := &DescribePermissionsArgs{}
	cmd := &cobra.Command{
		Use:   "describe-permissions",
		Short: "Show permissions granted to a subject",
		Long:  "Show all permissions granted to a subject. Can filter by permission type or resource.",
		Example: `# show all permissions for a subject
sts rbac describe-permissions --subject my-team

# filter by permission type
sts rbac describe-permissions --subject my-team --permission access-view`,
		RunE: deps.CmdRunEWithApi(RunDescribePermissionsCommand(args)),
	}

	cmd.Flags().StringVar(&args.Subject, Subject, "", SubjectUsage)
	cmd.MarkFlagRequired(Subject) //nolint:errcheck

	cmd.Flags().StringVar(&args.Permission, Permission, "", PermissionDescribeUsage)
	cmd.Flags().StringVar(&args.Resource, Resource, "", ResourceDescribeUsage)
	pflags.EnumVar(cmd.Flags(), &args.Source, Source, "", SourceChoices, SourceUsage)

	return cmd
}

func RunDescribePermissionsCommand(args *DescribePermissionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		description, resp, err := describePermissions(cli, api, args.Subject, args.Permission, args.Resource, args.Source).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		sourceStrings := make([]string, 0)
		if description.FromSources != nil {
			for _, s := range description.FromSources {
				sourceStrings = append(sourceStrings, string(s))
			}
		} else {
			sourceStrings = nil
		}

		if cli.IsJson() {
			if sourceStrings != nil {
				cli.Printer.PrintJson(map[string]interface{}{
					"subject":     description.SubjectHandle,
					"permissions": description.Permissions,
					"sources":     sourceStrings,
				})
			} else {
				cli.Printer.PrintJson(map[string]interface{}{
					"subject":     description.SubjectHandle,
					"permissions": description.Permissions,
				})
			}
		} else {
			printPermissionsTable(cli, sourceStrings, description.Permissions)
		}

		return nil
	}
}

func describePermissions(cli *di.Deps, api *stackstate_api.APIClient, subject string, permission string, resource string, source string) stackstate_api.ApiDescribePermissionsRequest {
	request := api.PermissionsApi.DescribePermissions(cli.Context, subject)
	if permission != "" {
		request = request.Permission(permission)
	}
	if resource != "" {
		request = request.Resource(resource)
	}
	if source != "" {
		request = request.Source(stackstate_api.SubjectSource(capitalizeFirst(source)))
	}
	return request
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func printPermissionsTable(cli *di.Deps, sources []string, permissionsList map[string][]string) {
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

	if sources != nil {
		cli.Printer.PrintLn(fmt.Sprintf("Got subject from the following subject sources: %s", strings.Join(sources, ", ")))
		cli.Printer.PrintLn("")
	}

	cli.Printer.Table(printer.TableData{
		Header:              []string{"permission", "resource"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
	})
}
