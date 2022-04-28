package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func TestNounsAndVerbsExistAsFilesInDirectories(t *testing.T) {
	root := setupCmd()
	for _, nounCmd := range root.Commands() {
		// test noun
		nounFile := fmt.Sprintf("%s.go", nounCmd.Name())
		exists, _ := util.DoesFileExist(nounFile)
		assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not.", nounFile))

		for _, verbCmd := range nounCmd.Commands() {
			// test noun
			verbName := strings.Replace(verbCmd.Name(), "-", "_", -1)
			verbFile := fmt.Sprintf("%s/%s_%s.go", nounCmd.Name(), nounCmd.Name(), verbName)
			exists, _ := util.DoesFileExist(verbFile)
			assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not.", verbFile))
		}
	}
}
