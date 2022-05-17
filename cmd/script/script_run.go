package script

import (
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

const (
	ScriptFlag          = "script"
	TimeoutFlag         = "timeout"
	ArgumentsScriptFlag = "arguments-script"
)

func ScriptRunCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a STSL script",
		Long:  "Run an STSL script.",
		Example: "# run a script from file\n" +
			"sts script run --file \"path/to/my.script\"\n" +
			"\n" +
			"# run a script with variables provided by an arguments-script\n" +
			"sts script run --script \"x+y\" --arguments-script \"[x: 1, y: 2]\"",
		RunE: cli.CmdRunEWithApi(RunScriptRunCommand),
	}

	cmd.Flags().String(ScriptFlag, "", "a script to run")
	cmd.Flags().String(ArgumentsScriptFlag, "",
		"an extra script that generates arguments to be used as variables when the main script is executed, return format: java.util.Map",
	)
	cmd.Flags().IntP(TimeoutFlag, "t", 0, "timeout in milli-seconds for script execution")
	common.AddFileFlag(cmd, "path to a file containing the script to run")
	mutex_flags.MarkMutexFlags(cmd, []string{ScriptFlag, common.FileFlag}, "input", true)

	return cmd
}

func RunScriptRunCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	var script string

	script, err := cmd.Flags().GetString(ScriptFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	filepath, err := cmd.Flags().GetString(common.FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	err = mutex_flags.CheckMutuallyExclusiveFlags(cmd, []string{ScriptFlag, common.FileFlag}, true)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	if filepath != "" {
		b, err := os.ReadFile(filepath)
		if err != nil {
			return common.NewReadFileError(err, filepath)
		}
		script = string(b)
	}

	var argumentsScript *string
	a, _ := cmd.Flags().GetString(ArgumentsScriptFlag)
	if a != "" {
		argumentsScript = &a
	}
	var timeoutMs *int32
	t, _ := cmd.Flags().GetInt(TimeoutFlag)
	if t != 0 {
		t32 := int32(t)
		timeoutMs = &t32
	}

	// execute script
	scriptRequest := stackstate_api.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          script,
		ArgumentsScript: argumentsScript,
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
