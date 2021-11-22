package scale

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/scaler"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
)

func DownCommand(cfg *config.Config) *cobra.Command {
	var DownNow bool
	var All bool

	cmd := &cobra.Command{
		Use:   "down [<namespace>]",
		Short: "Scale down the namespace",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 && DownNow {
				return fmt.Errorf("can only supply --now when a namespace is given")
			} else if len(args) == 0 && !All {
				return fmt.Errorf("should explicitly supply '--all' if you want to scan all namespaces for scaling down")
			}

			c, err := kubernetes.NewClient()
			if err != nil {
				return err
			}

			s, err := scaler.NewScaler(cmd.Context(), cfg.Scale, c)
			if err != nil {
				return err
			}

			if len(args) > 0 {
				nsName := args[0]
				b, err := kubernetes.NamespaceExists(cmd.Context(), c, nsName)
				if err != nil {
					return err
				}

				if !b {
					return fmt.Errorf("namespace '%s' does not exist", nsName)
				}

				return s.ScaleDownNamespace(cmd.Context(), nsName, DownNow)
			}

			return s.ScaleDownAll(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&DownNow, "now", false, "Scale down the namespace now, only valid when a namespace is given.")
	cmd.Flags().BoolVarP(&All, "all", "a", false, "Check all namespaces to see whether they need to be scaled down.")

	return cmd
}
