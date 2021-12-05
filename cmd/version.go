package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var (
	Version string
	Commit  string
	Date    string
)

func VersionCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display the version number",
		Run: func(_ *cobra.Command, _ []string) {
			cli.Printer.PrintStruct(map[string]interface{}{
				"version": Version,
				"commit":  Commit,
				"date":    Date,
			})
		},
	}

	return cmd
}
