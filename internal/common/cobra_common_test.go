package common

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
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
	mutex_flags.MarkMutexFlags(cmd, []string{"c", "db"}, "mutex1", true)
	mutex_flags.MarkMutexFlags(cmd, []string{"ef", "ge"}, "mutex2", true)

	AddRequiredFlagsToCmd(cmd)

	assert.Equal(t, "test --a INT --b STRING { --c BOOL | --db INTSLICE } { --ef STRING | --ge STRING }", cmd.Use)
}
