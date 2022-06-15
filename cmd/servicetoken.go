package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/servicetoken"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ServiceTokenCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-token",
		Short: "Manage service tokens",
		Long:  "Manage service tokens.",
	}
	cmd.AddCommand(servicetoken.CreateCommand(deps))
	cmd.AddCommand(servicetoken.DeleteCommand(deps))
	cmd.AddCommand(servicetoken.ListCommand(deps))
	return cmd
}
