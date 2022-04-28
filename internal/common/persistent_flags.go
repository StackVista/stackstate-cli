package common

import (
	"github.com/spf13/cobra"
)

const (
	VerboseFlag  = "verbose"
	APIURLFlag   = "api-url"
	APITokenFlag = "api-token"
	NoColorFlag  = "no-color"
	JsonFlag     = "json"
)

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(APIURLFlag, "", "specify the API URL of the StackState the CLI should connect to")
	cmd.PersistentFlags().String(APITokenFlag, "", "specify the API token of the StackState the CLI should connect to")
	cmd.PersistentFlags().Bool(VerboseFlag, false, "print verbose logging to the terminal to track what the CLI is doing")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "disable color when printing to the terminal")
	cmd.PersistentFlags().Bool(JsonFlag, false, "print as json")
}
