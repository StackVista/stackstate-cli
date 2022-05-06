package mutex_flags

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGettingAndSettingMutexFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().Int("ra", 1, "")
	cmd.Flags().Int("rb", 2, "")
	cmd.Flags().Int("c", 3, "")
	cmd.Flags().Int("a", 4, "")
	cmd.Flags().Int("b", 5, "")

	MarkMutexFlags(cmd, []string{"ra", "rb"}, "rints", true)
	MarkMutexFlags(cmd, []string{"a", "b"}, "ints", true)
	assert.Equal(t, []string{"ints", "rints"}, GetAllMutexNames(cmd, true))
	assert.Equal(t, []string{"ra", "rb"}, GetAllFlagNamesOfMutex(cmd, "rints"))
}
