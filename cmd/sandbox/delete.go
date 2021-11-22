package sandbox

import (
	color "github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/cli"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/sandbox"
)

func DeleteCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete <sandbox name>",
		Short:   "Deletes a sandbox that you created",
		Args:    cobra.ExactArgs(1),
		PreRunE: cli.FlagsRequired("username"),
		RunE: func(cmd *cobra.Command, args []string) error {
			sandboxName := args[0]
			sandboxer, err := sandbox.New()
			if err != nil {
				return err
			}

			if err := sandboxer.Delete(cmd.Context(), sandboxName, cfg.Sandbox.Username); err != nil {
				return err
			}

			cmd.Printf("Sandbox '%s' has been deleted.\n", color.Red(sandboxName))

			return nil
		},
	}

	return cmd
}
