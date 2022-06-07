package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type SetArgs struct {
	Name string
}

func SetCommand(cli *di.Deps) *cobra.Command {
	args := &SetArgs{}
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set the current context",
		Long:  "Set the current context.",
		RunE:  cli.CmdRunE(RunContextSetCommand(args, cli)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the context")

	return cmd
}

func RunContextSetCommand(args *SetArgs, cli *di.Deps) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		cfg := cli.StsConfig

		ctx := cfg.GetContext(args.Name)
		if ctx == nil {
			return common.NewNotFoundError(fmt.Errorf("context with name '%s' not found", args.Name))
		}

		cfg.CurrentContext = args.Name

		if err := config.WriteConfig(cli.ConfigPath, cfg); err != nil {
			return common.NewWriteFileError(err, cli.ConfigPath)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"current-context": args.Name,
			})
		} else {
			// No output
		}

		return nil
	}
}
