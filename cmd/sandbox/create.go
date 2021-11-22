package sandbox

import (
	"context"
	"time"

	color "github.com/logrusorgru/aurora/v3"
	devopsv1 "github.com/stackvista/sandbox-operator/apis/devops/v1"
	sandboxpkg "github.com/stackvista/sandbox-operator/pkg/sandbox"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/cli"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/sandbox"
	"gitlab.com/stackvista/stackstate-cli2/pkg/kubernetes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateCommand(cfg *config.Config) *cobra.Command {
	var expirationDate string
	cmd := &cobra.Command{
		Use:     "create <sandbox name>",
		Short:   "Create your personal Kubernetes Sandbox",
		Args:    cobra.ExactArgs(1),
		PreRunE: cli.FlagsRequired("slack-id", "username"),
		RunE: func(cmd *cobra.Command, args []string) error {
			sandboxName := args[0]
			expDate, err := ParseExpirationDate(cmd.Context(), expirationDate)
			if err != nil {
				return err
			}

			spec := &devopsv1.SandboxSpec{
				User:           cfg.Sandbox.Username,
				SlackId:        cfg.Sandbox.SlackID,
				ExpirationDate: expDate,
				ManualExpiry:   cfg.Sandbox.ManualExpiry,
			}

			sb, err := Create(cmd.Context(), sandboxName, spec)
			if err != nil {
				return err
			}

			cmd.Printf("Sandbox '%s' has been created.", color.Green(sb.Name))
			if expirationDate != "" {
				cmd.Printf(" It will expire on '%s'.", color.Yellow(spec.ExpirationDate.Time))
			}

			cmd.Println()

			createdNamespace := sandboxpkg.SandboxName(sb)
			if err := kubernetes.SwitchNamespace(createdNamespace); err != nil {
				return err
			}

			cmd.Printf("You are now using sandbox %s\n", color.Green(createdNamespace))
			return nil
		},
	}

	cmd.Flags().StringVarP(&expirationDate, "expiry", "e", "", "The expiration date of your Sandbox, in the format yyyy-MM-dd. If not set, the sandbox will be cleaned after a pre-determined timeout, typically 1 week.")
	cmd.Flags().StringP("slack-id", "s", "", "Your Slack ID, used to notify you of the lifecycle of your sandbox")
	cmd.Flags().BoolP("manual-expiry", "m", false, "If set, prevents your sandbox from automatically being cleaned")

	return cmd
}

func Create(ctx context.Context, name string, spec *devopsv1.SandboxSpec) (*devopsv1.Sandbox, error) {
	sandboxer, err := sandbox.New()
	if err != nil {
		return nil, err
	}

	return sandboxer.Create(ctx, name, spec)
}

func ParseExpirationDate(ctx context.Context, expirationDate string) (*v1.Time, error) {
	if expirationDate != "" {
		t, err := time.Parse("2006-01-02", expirationDate)
		if err != nil {
			return nil, err
		}

		return &v1.Time{
			Time: t,
		}, nil
	}

	return nil, nil
}
