package cobra

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func ForAllCmd(parent *cobra.Command, fn func(*cobra.Command)) {
	fn(parent)
	for _, child := range parent.Commands() {
		fn(child)
		if child.HasSubCommands() {
			ForAllCmd(child, fn)
		}
	}
}

func ForAllFlags(parent *cobra.Command, fn func(*cobra.Command, *pflag.Flag)) {
	ForAllCmd(parent, func(cmd *cobra.Command) {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fn(cmd, flag)
		})
	})
}

func IsFlagRequired(flag *pflag.Flag) bool {
	return IsSingleRequiredFlag(flag) || len(flag.Annotations["rmutex"]) > 0
}

func IsSingleRequiredFlag(flag *pflag.Flag) bool {
	return len(flag.Annotations[cobra.BashCompOneRequiredFlag]) > 0
}
