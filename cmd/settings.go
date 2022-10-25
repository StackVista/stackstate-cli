package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/settings"
	"github.com/stackvista/stackstate-cli/internal/di"
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
	cmd.AddCommand(settings.SettingsEditCommand(cli))
	cmd.AddCommand(settings.SettingsUnlockCommand(cli))
	cmd.AddCommand(settings.SettingsDeleteCommand(cli))
	return cmd
}
