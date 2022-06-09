package cobra

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestAddingRequiredFlagsToCmd(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().Int("a", 0, "")
	cmd.Flags().String("b", "", "")
	cmd.Flags().Bool("c", false, "")
	cmd.Flags().IntSliceP("db", "d", []int{}, "")
	cmd.Flags().String("e", "", "")
	cmd.Flags().String("ef", "", "")
	cmd.Flags().String("ge", "", "")
	cmd.MarkFlagRequired("a") //nolint:errcheck
	cmd.MarkFlagRequired("b") //nolint:errcheck
	MarkMutexFlags(cmd, []string{"c", "db"}, "mutex1", true)
	MarkMutexFlags(cmd, []string{"ef", "ge"}, "mutex2", true)

	AddRequiredFlagsToUseString(cmd)

	assert.Equal(t, "test --a A --b B { --c C | --db DB } { --ef EF | --ge GE }", cmd.Use)
}
