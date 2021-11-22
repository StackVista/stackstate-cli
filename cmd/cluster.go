package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/cluster"
)

func ClusterCommand() *cobra.Command {
	opts := &cluster.ClusterOpts{}
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Operate on KOPS clusters",
	}

	cmd.PersistentFlags().StringVarP(&opts.Profile, "profile", "p", "", "The AWS profile to use")

	cmd.AddCommand(cluster.ConnectCommand(opts))
	cmd.AddCommand(cluster.ListCommand(opts))

	return cmd
}
