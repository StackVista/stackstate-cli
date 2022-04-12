package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/stackpack"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func StackPackCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stackpack",
		Short: "manage stackpacks",
	}
	cmd.AddCommand(stackpack.StackpackUploadCommand(cli))

	return cmd
}
