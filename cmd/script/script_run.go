package script

import (
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type ScriptRunArgs struct {
	Script          string
	ArgumentsScript string
	TimeoutMs       int32
	ScriptFile      string
}

const (
	ScriptFlag          = "script"
	TimeoutFlag         = "timeout"
	ArgumentsScriptFlag = "arguments-script"
)

func ScriptRunCommand(cli *di.Deps) *cobra.Command {
	args := &ScriptRunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a STSL script",
		Long:  "Run an STSL script.",
		Example: "# run a script from file\n" +
			"sts script run --file \"path/to/my.script\"\n" +
			"\n" +
			"# run a script with variables provided by an arguments-script\n" +
			"sts script run --script \"x+y\" --arguments-script \"[x: 1, y: 2]\"",
		RunE: cli.CmdRunEWithApi(RunScriptRunCommand(args)),
	}

	cmd.Flags().StringVar(&args.Script, ScriptFlag, "", "a script to run")
	cmd.Flags().StringVar(&args.ArgumentsScript, ArgumentsScriptFlag, "",
		"an extra script that generates arguments to be used as variables when the main script is executed, return format: java.util.Map",
	)
	cmd.Flags().Int32VarP(&args.TimeoutMs, TimeoutFlag, "t", 0, "timeout in milli-seconds for script execution")
	common.AddFileFlagVar(cmd, &args.ScriptFile, "path to a file containing the script to run")
	mutex_flags.MarkMutexFlags(cmd, []string{ScriptFlag, common.FileFlag}, "input", true)

	return cmd
}

func RunScriptRunCommand(args *ScriptRunArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		script := args.Script

		err := mutex_flags.CheckMutuallyExclusiveFlags(cmd, []string{ScriptFlag, common.FileFlag}, true)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}

		if args.ScriptFile != "" {
			b, err := os.ReadFile(args.ScriptFile)
			if err != nil {
				return common.NewReadFileError(err, args.ScriptFile)
			}
			script = string(b)
		}

		// execute script
		scriptRequest := stackstate_api.ExecuteScriptRequest{
			TimeoutMs:       util.Int32NilP(args.TimeoutMs),
			Script:          script,
			ArgumentsScript: util.StringNilP(args.ArgumentsScript),
		}

		scriptResponse, resp, err := api.ScriptingApi.
			ScriptExecute(cli.Context).
			ExecuteScriptRequest(scriptRequest).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		// print response
		value := scriptResponse.Result["value"]
		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"result": scriptResponse.Result,
			})
		} else {
			if value == nil {
				cli.Printer.Success("script executed (no response)")
			} else {
				cli.Printer.PrintStruct(value)
			}
		}

		return nil
	}
}
