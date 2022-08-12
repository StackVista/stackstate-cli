package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func TestNounsAndVerbsExistAsFilesInDirectories(t *testing.T) {
	root := setupCmd(t)
	for _, nounCmd := range root.Commands() {
		// test noun
		nounName := strings.ReplaceAll(nounCmd.Name(), "-", "")
		nounFile := fmt.Sprintf("%s.go", nounName)
		exists, _ := util.DoesFileExist(nounFile)
		assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not.", nounFile))

		for _, verbCmd := range nounCmd.Commands() {
			// test verb
			verbName := strings.ReplaceAll(verbCmd.Name(), "-", "_")
			verbFile := fmt.Sprintf("%s/%s_%s.go", nounName, nounName, verbName)
			exists, _ := util.DoesFileExist(verbFile)
			assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not.", verbFile))

			verbFileTest := fmt.Sprintf("%s/%s_%s_test.go", nounName, nounName, verbName)
			exists, _ = util.DoesFileExist(verbFileTest)
			assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not. Please write tests!", verbFileTest))
		}
	}
}

func TestVerbCommandCheckForJsonOuput(t *testing.T) {
	root := setupCmd(t)
	for _, nounCmd := range root.Commands() {
		for _, verbCmd := range nounCmd.Commands() {
			nounName := strings.ReplaceAll(nounCmd.Name(), "-", "")
			verbName := strings.ReplaceAll(verbCmd.Name(), "-", "_")
			verCmdGoFile := fmt.Sprintf("%s/%s_%s.go", nounName, nounName, verbName)
			verbCmdGoCode, err := os.ReadFile(verCmdGoFile)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(verbCmdGoCode), "if cli.IsJson() {") {
				t.Errorf("%s does not check whether to print to json!", verCmdGoFile)
			}
			if !strings.Contains(string(verbCmdGoCode), "cli.Printer.PrintJson(") {
				t.Errorf("%s does not print json!", verCmdGoFile)
			}
		}
	}
}
