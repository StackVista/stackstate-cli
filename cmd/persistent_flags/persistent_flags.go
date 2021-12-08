package persistent_flags

import (
	"github.com/spf13/cobra"
)

const (
	VerboseFlag  = "verbose"
	ApiUrlFlag   = "api-url"
	ApiTokenFlag = "api-token"
	NoColorFlag  = "no-color"
	OutputFlag   = "output"
)

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(ApiUrlFlag, "", "StackState API URL.")
	cmd.PersistentFlags().String(ApiTokenFlag, "", "StackState API Token.")
	cmd.PersistentFlags().Bool(VerboseFlag, false, "Print verbose logging to see what the CLI is doing.")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "Print to terminal without color.")
	cmd.PersistentFlags().StringP(OutputFlag, "o", "Auto", "Format output as: JSON, YAML or Auto.")
}
