package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/settings"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func SettingsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Manage settings.",
	}
	cmd.AddCommand(settings.SettingsApplyCommand(cli))

	return cmd
}
