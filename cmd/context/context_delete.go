package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	Name string
}

func DeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a context",
		Long:  "Delete a context from the local config file.",
		RunE:  cli.CmdRunEWithConfig(RunContextDeleteCommand(args)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the context")

	return cmd
}

func RunContextDeleteCommand(args *DeleteArgs) func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
		if args.Name == cfg.CurrentContext {
			return common.NewCLIArgParseError(fmt.Errorf("cannot delete the current context (%s)", args.Name))
		}

		found := -1

		for i, c := range cfg.Contexts {
			if c.Name == args.Name {
				found = i
				break
			}
		}

		if found == -1 {
			return common.NewNotFoundError(fmt.Errorf("context with name '%s' not found", args.Name))
		}

		cfg.Contexts = append(cfg.Contexts[:found], cfg.Contexts[found+1:]...)

		if err := config.WriteConfig(cli.ConfigPath, cfg); err != nil {
			return common.NewWriteFileError(err, cli.ConfigPath)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted context": args.Name,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Deleted context '%s'\n", args.Name))
		}

		return nil
	}
}
