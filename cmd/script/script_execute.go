package script

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	FileFlag            = "file"
	TimeoutFlag         = "timeout"
	ArgumentsScriptFlag = "arguments-script"
)

func ScriptExecuteCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Short: "Execute a STSL script.",
		RunE:  di.CmdRunEWithDeps(cli, RunScriptExecuteCommand),
	}

	cmd.Flags().StringP(ArgumentsScriptFlag, "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP(TimeoutFlag, "t", 0, "Timeout in milli-seconds.")
	cmd.Flags().StringP(FileFlag, "f", "", "Read the script from file.")

	return cmd
}

func RunScriptExecuteCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	var script string
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	if file != "" {
		b, err := os.ReadFile(file)
		if err != nil {
			return common.NewCLIError(err)
		}
		script = string(b)
	} else if len(args) > 0 {
		script = args[0]
	} else {
		return common.NewCLIArgParseError(fmt.Errorf("missing <script> argument"))
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

	scriptRequest := sts.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          script,
		ArgumentsScript: argumentsScript,
	}

	cli.Printer.StartSpinner(common.AwaitingServer)
	defer cli.Printer.StopSpinner()
	scriptResponse, resp, err := cli.Client.ScriptingApi.
		ScriptExecute(cli.Context).
		ExecuteScriptRequest(scriptRequest).
		Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.PrintStruct(scriptResponse["result"])

	return nil
}
