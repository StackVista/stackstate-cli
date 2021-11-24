package script

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute a STSL script.",
		RunE:  RunScriptExecuteCommand,
	}

	cmd.Flags().StringP("arguments-script", "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in milli-seconds.")

	return cmd
}

func RunScriptExecuteCommand(cmd *cobra.Command, args []string) error {
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

	configuration.DefaultHeader["Authorization"] = "ApiToken Sxeqe2hKAaTUgpQnIv3_ctkTYjsN8pz2"
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

	scriptResponse, resp, err := scriptExecute.ExecuteScriptRequest(scriptRequest).Execute()

	if resp.StatusCode == 401 {
		return fmt.Errorf("401 Unauthorized. Please check your API Token")
	}

	if err != nil {
		switch v := err.(type) {
		case stackstate_client.GenericOpenAPIError:
			return fmt.Errorf("%v. Error response: %+v", resp.Status, string(v.Body()))
		default:
			return fmt.Errorf("%v. %+v", resp.Status, v)
		}
	}

	fmt.Println(scriptResponse["result"])
	return nil
}
