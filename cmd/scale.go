package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/scale"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func ScaleCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scale",
		Short: "Scale a namespace either up or down",
	}

	cmd.AddCommand(scale.UpCommand(cfg))
	cmd.AddCommand(scale.DownCommand(cfg))

	return cmd
}
