package scale

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/scaler"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
)

func UpCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up [<namespace>]",
		Args:  cobra.ExactArgs(1),
		Short: "Scale up a namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := kubernetes.NewClient()
			if err != nil {
				return err
			}

			s, err := scaler.NewScaler(cmd.Context(), cfg.Scale, c)
			if err != nil {
				return err
			}

			nsName := args[0]
			b, err := kubernetes.NamespaceExists(cmd.Context(), c, nsName)
			if err != nil {
				return err
			}

			if !b {
				return fmt.Errorf("namespace '%s' does not exist", nsName)
			}

			return s.ScaleUpNamespace(cmd.Context(), nsName)
		},
	}

	return cmd
}
