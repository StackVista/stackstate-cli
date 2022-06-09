package cobra

import (
	"fmt"
	"sort"
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
		reqMutexFlags := mutexFlags[true].ToSlice()
		sort.Strings(reqMutexFlags)

		for _, mutex := range reqMutexFlags {
			mutexFlagUses := make([]string, 0)
			for _, mutexFlag := range GetAllFlagsOfMutex(cmd, mutex) {
				mutexFlagUses = append(mutexFlagUses, fmt.Sprintf("--%s %s", mutexFlag.Name, strings.ToUpper(mutexFlag.Name)))
			}
			sort.Strings(mutexFlagUses)

			required += " { " + strings.Join(mutexFlagUses, " | ") + " }"
		}

		required = strings.Trim(required, " ")
		if required != "" {
			cmd.Use += " " + required
		}
	})
}
