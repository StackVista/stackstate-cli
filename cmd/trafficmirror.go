package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/trafficmirror"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func TrafficMirrorCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "traffic-mirror",
		Short: "Operations for traffic mirror site",
	}

	cmd.AddCommand(trafficmirror.GetCommand(cfg))
	cmd.AddCommand(trafficmirror.PutCommand(cfg))
	cmd.AddCommand(trafficmirror.DeleteCommand(cfg))

	cmd.PersistentFlags().StringP("url", "u", "", "The URL of the Traffic Mirror instance")
	cmd.PersistentFlags().StringP("username", "n", "", "The username to access the Traffic Mirror instance")
	cmd.PersistentFlags().StringP("password", "p", "", "The password to access the Traffic Mirror instance")

	return cmd
}
