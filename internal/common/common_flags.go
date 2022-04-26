package common

import (
	"github.com/spf13/cobra"
)

const (
	JsonFlag = "json"
)

func AddJsonFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(JsonFlag, false, "print as json")
}
