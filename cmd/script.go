package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "StackState scripting related commands.",
	}
	cmd.AddCommand(ScriptExecuteCommand())

	return cmd
}

func ScriptExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute a STSL script.",
		Run:   RunScriptExecuteCommand,
	}

	cmd.Flags().StringP("arguments-script", "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in milli-seconds.")

	return cmd
}

func RunScriptExecuteCommand(cmd *cobra.Command, args []string) {
	var argumentsScript *string
	a, _ := cmd.Flags().GetString("arguments-script")
	if a != "" {
		argumentsScript = &a
	}
	var timeoutMs *int32
	t, _ := cmd.Flags().GetInt("timeout")
	if t != 0 {
		t32 := int32(t)
		timeoutMs = &t32
	}
	verbose, _ := cmd.Flags().GetCount("verbose")

	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         "https://master.preprod.stackstate.io/api/",
		Description: "",
		Variables:   nil,
	}

	configuration.DefaultHeader["Authorization"] = "ApiToken ihm6EevUJ0lfKDx8ccKzNLxss3wyk1jk"
	client := stackstate_client.NewAPIClient(configuration)
	scriptExecute := client.ScriptingApi.ScriptExecute(cmd.Context())
	scriptRequest := stackstate_client.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          args[0],
		ArgumentsScript: argumentsScript,
	}

	if verbose > 0 {
		scriptRequestStr, _ := json.Marshal(scriptRequest)
		fmt.Println("Executing script request:", string(scriptRequestStr))
	}

	resp, _, err := scriptExecute.ExecuteScriptRequest(scriptRequest).Execute()
	if err != nil {
		fmt.Println("Got err ", err)
		return
	}

	fmt.Println(resp["result"])
}
