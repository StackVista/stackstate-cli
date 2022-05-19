package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func AddRequiredFlagsToCmd(root *cobra.Command) {
	util.ForAllCmd(root, func(cmd *cobra.Command) {
		required := ""
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if util.IsSingleRequiredFlag(flag) {
				required += " " + fmt.Sprintf("--%s %s", flag.Name, strings.ToUpper(flag.Name))
			}
		})

		requiredMutexFlags := mutex_flags.GetAllMutexNames(cmd, true)
		for _, mutex := range requiredMutexFlags {
			mutexFlagUses := make([]string, 0)
			for _, mutexFlag := range mutex_flags.GetAllFlagsOfMutex(cmd, mutex) {
				mutexFlagUses = append(mutexFlagUses, fmt.Sprintf("--%s %s", mutexFlag.Name, strings.ToUpper(mutexFlag.Name)))
			}
			required += " { " + strings.Join(mutexFlagUses, " | ") + " }"
		}

		required = strings.Trim(required, " ")
		if required != "" {
			cmd.Use += " " + required
		}
	})
}
