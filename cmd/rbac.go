package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/rbac"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RbacCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rbac",
		Short: "Manage the StackState RBAC",
		Long:  "Manage the StackState rule-based access control configuration.",
	}

	cmd.AddCommand(rbac.CreateSubjectCommand(deps))
	cmd.AddCommand(rbac.ListPermissionsCommand(deps))
	cmd.AddCommand(rbac.DescribePermissionsCommand(deps))
	cmd.AddCommand(rbac.GrantPermissionsCommand(deps))
	cmd.AddCommand(rbac.RevokePermissionsCommand(deps))
	return cmd
}
