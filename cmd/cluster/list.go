package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/kops"
)

func ListCommand(opts *ClusterOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available KOPS clusters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clusters, err := kops.ListClusters(cmd.Context(), opts.Profile)
			if err != nil {
				return err
			}

			for _, c := range clusters {
				fmt.Println(c)
			}

			return nil
		},
	}

	return cmd
}
