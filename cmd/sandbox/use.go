package sandbox

import (
	"fmt"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	sandboxpkg "github.com/stackvista/sandbox-operator/pkg/sandbox"
	"gitlab.com/stackvista/stackstate-cli2/internal/cli"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/sandbox"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
)

func SwitchCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use [<sandbox name>]",
		Short:   "Switch to using a different sandbox on the Kubernetes cluster",
		Args:    cobra.MaximumNArgs(1),
		PreRunE: cli.FlagsRequired("username"),
		RunE: func(cmd *cobra.Command, args []string) error {
			boxer, err := sandbox.New()
			if err != nil {
				return err
			}

			boxes, err := boxer.ListMySandboxes(cmd.Context(), cfg.Sandbox.Username)
			if err != nil {
				return err
			}

			namespaceName := ""
			if len(args) > 0 {
				namespaceName = args[0]

				found := false
				for _, n := range boxes {
					found = found || n == namespaceName
				}

				if !found {
					return fmt.Errorf("sandbox %s is unknown", color.Red(namespaceName))
				}
			} else {
				sel := &promptui.Select{
					Label: "Choose sandbox",
					Items: boxes,
				}

				_, v, err := sel.Run()
				if err != nil {
					return err
				}

				namespaceName = v
			}

			box, err := boxer.Get(cmd.Context(), namespaceName)
			if err != nil {
				return err
			}

			if err := kubernetes.SwitchNamespace(sandboxpkg.SandboxName(box)); err != nil {
				return err
			}

			cmd.Printf("You are now connected to sandbox '%s' on Kubernetes cluster %s\n",
				color.Green(namespaceName),
				color.Yellow(boxer.GetKubernetesHost()))

			return nil
		},
	}

	return cmd
}
