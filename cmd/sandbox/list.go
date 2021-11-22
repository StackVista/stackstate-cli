package sandbox

import (
	color "github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/cli"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/sandbox"
)

func ListCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists your Kubernetes Sandboxes",
		PreRunE: cli.FlagsRequired("username"),
		RunE: func(cmd *cobra.Command, args []string) error {
			sandboxer, err := sandbox.New()
			if err != nil {
				return err
			}

			boxes, err := sandboxer.ListMySandboxes(cmd.Context(), cfg.Sandbox.Username)
			if err != nil {
				return err
			}

			if len(boxes) > 0 {
				cmd.Printf("You have %d sandboxes running on %s:\n", color.Green(len(boxes)), color.Green(sandboxer.GetKubernetesHost()))
				for _, r := range boxes {
					cmd.Printf("- %s\n", r)
				}
			} else {
				cmd.Printf("You have %d sandboxes running on %s\n", color.Red(len(boxes)), color.Green(sandboxer.GetKubernetesHost()))
			}
			return nil
		},
	}

	return cmd
}
