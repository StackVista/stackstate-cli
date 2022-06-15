package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/settings"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func SettingsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Manage settings",
		Long:  "View, manage, export and import settings.",
	}
	cmd.AddCommand(settings.SettingsApplyCommand(cli))
	cmd.AddCommand(settings.SettingsListTypesCommand(cli))
	cmd.AddCommand(settings.SettingsListCommand(cli))
	cmd.AddCommand(settings.SettingsDescribeCommand(cli))

	return cmd
}
