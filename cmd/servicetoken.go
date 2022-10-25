package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/servicetoken"
	"github.com/stackvista/stackstate-cli/internal/di"
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
