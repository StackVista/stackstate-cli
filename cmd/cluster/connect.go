package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/kops"
)

func ConnectCommand(opts *ClusterOpts) *cobra.Command {
	var admin bool

	cmd := &cobra.Command{
		Use:   "connect <cluster>",
		Short: "Connect to a KOPS cluster using a role",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("need to pass cluster name as an argument")
			}

			return kops.Connect(cmd.Context(), args[0], admin, opts.Profile)
		},
	}

	cmd.Flags().BoolVarP(&admin, "admin", "a", false, "Whether to connect using the admin role")

	return cmd
}
