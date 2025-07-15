package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/usersession"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func UserSessionCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-session",
		Short: "Inspect user session",
		Long:  "Inspect user session.",
	}
	cmd.AddCommand(usersession.GetUserSessionRolesCommand(deps))
	return cmd
}
