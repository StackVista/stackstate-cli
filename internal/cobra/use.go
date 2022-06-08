package cobra

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func AddRequiredFlagsToUseString(root *cobra.Command) {
	ForAllCmd(root, func(cmd *cobra.Command) {
		required := ""
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if IsSingleRequiredFlag(flag) {
				required += " " + fmt.Sprintf("--%s %s", flag.Name, strings.ToUpper(flag.Name))
			}
		})

		mutexFlags := GetAllMutexNames(cmd)
		for mutex := range mutexFlags[true] {
			mutexFlagUses := make([]string, 0)
			for _, mutexFlag := range GetAllFlagsOfMutex(cmd, mutex) {
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
