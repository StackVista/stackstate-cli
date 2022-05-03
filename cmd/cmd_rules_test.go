package cmd

import (
	"fmt"
	"io/ioutil"
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
			// test verb
			verbName := strings.ReplaceAll(verbCmd.Name(), "-", "_")
			verbFile := fmt.Sprintf("%s/%s_%s.go", nounCmd.Name(), nounCmd.Name(), verbName)
			exists, _ := util.DoesFileExist(verbFile)
			assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not.", verbFile))

			verbFileTest := fmt.Sprintf("%s/%s_%s_test.go", nounCmd.Name(), nounCmd.Name(), verbName)
			exists, _ = util.DoesFileExist(verbFileTest)
			assert.True(t, exists, fmt.Sprintf("file %s should exist, but does not. Please write tests!", verbFileTest))
		}
	}
}

func TestVerbCommandCheckForJsonOuput(t *testing.T) {
	root := setupCmd()
	for _, nounCmd := range root.Commands() {
		for _, verbCmd := range nounCmd.Commands() {
			verbName := strings.ReplaceAll(verbCmd.Name(), "-", "_")
			verCmdGoFile := fmt.Sprintf("%s/%s_%s.go", nounCmd.Name(), nounCmd.Name(), verbName)
			verbCmdGoCode, err := ioutil.ReadFile(verCmdGoFile)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(verbCmdGoCode), "if cli.IsJson {") {
				t.Fatalf("%s does not check whether to print to json!", verCmdGoFile)
			}
			if !strings.Contains(string(verbCmdGoCode), "cli.Printer.PrintJson(") {
				t.Fatalf("%s does not print json!", verCmdGoFile)
			}
		}
	}
}
