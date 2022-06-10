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
		RunE:  cli.CmdRunEWithConfig(RunContextSetCommand(args, cli)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the context")

	return cmd
}

func RunContextSetCommand(args *SetArgs, cli *di.Deps) func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
		_, err := cfg.GetContext(args.Name)
		if err != nil {
			return common.NewNotFoundError(err)
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
			cli.Printer.PrintLn(fmt.Sprintf("Current context set to %s", args.Name))
			// No output
		}

		return nil
	}
}
