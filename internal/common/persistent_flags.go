package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	VerboseFlag      = "verbose"
	VerboseFlagShort = "v"
	URLFlag          = "url"
	APITokenFlag     = "api-token"
	NoColorFlag      = "no-color"
	OutputFlag       = "output"
	OutputFlagShort  = "o"
)

var AllowedOutputs = []string{JSONOutput.String(), TextOutput.String()}

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(URLFlag, "", "specify the URL of the StackState server")
	cmd.PersistentFlags().String(APITokenFlag, "", "specify the API token of the StackState server")
	cmd.PersistentFlags().CountP(VerboseFlag, VerboseFlagShort, "print verbose logging to the terminal to track what the CLI is doing (use multiple times to increase verbosity)")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "disable color when printing to the terminal")
	pflags.EnumP(cmd.PersistentFlags(), OutputFlag, OutputFlagShort, "text", AllowedOutputs, fmt.Sprintf("specify the output format (must be { %s })", strings.Join(AllowedOutputs, " | ")))
}
