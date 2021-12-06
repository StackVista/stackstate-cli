package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

// These variables will be set on build by the ldlflags.
// See the `ldflags` in .goreleaser.yml
var (
	Version string
	Commit  string
	Date    string
)

func VersionCommand(cli *di.Deps) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display the version number",
		RunE:  di.CmdRunEWithDeps(cli, RunVersionCommand),
	}
}

func RunVersionCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	cli.Printer.PrintStruct(map[string]interface{}{
		"version": Version,
		"commit":  Commit,
		"date":    Date,
	})
	return nil
}
