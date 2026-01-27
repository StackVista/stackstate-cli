package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/usersession"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func UserSessionCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-session",
		Short: "Inspect the current user session",
		Long:  "Inspect the current user session. View session information and available roles.",
	}
	cmd.AddCommand(usersession.GetUserSessionRolesCommand(deps))
	return cmd
}
