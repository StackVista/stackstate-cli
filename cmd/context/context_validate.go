package context

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type ValidateArgs struct {
	Name string
}

func ValidateCommand(cli *di.Deps) *cobra.Command {
	args := &ValidateArgs{}
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate that a context can connect to the server",
		Long:  "Validate a connection context by attempting to connect to the SUSE Observability server. Validates the current context if no name is specified.",
		Example: `# validate the current context
sts context validate

# validate a specific context
sts context validate --name production`,
		RunE: cli.CmdRunEWithConfig(RunValidateCommand(args)),
	}

	common.AddNameFlagVar(cmd, &args.Name, "Name of the context")

	return cmd
}

func RunValidateCommand(args *ValidateArgs) func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
		ctxName := cfg.CurrentContext
		if args.Name != "" {
			ctxName = args.Name
		}

		ctx, err := cfg.GetContext(ctxName)
		if err != nil {
			return common.NewNotFoundError(err)
		}

		serverInfo, cerr := ValidateContext(cli, cmd, ctx.Context)
		if cerr != nil {
			return cerr
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"context":     ctxName,
				"server-info": serverInfo,
				"connected":   true,
			})
		} else {
			cli.Printer.Successf("Context %s is valid.\n", ctxName)
		}

		return nil
	}
}
