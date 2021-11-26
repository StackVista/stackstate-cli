package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSuccessFromYaml(t *testing.T) {
	confYaml := `
api-url: https://my.stackstate.com/api
api-token: BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M
`

	conf, err := readConfFromFile(t, confYaml)

	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

////------------- Fixture code -----------

func readConfFromFile(t *testing.T, confYaml string) (Conf, error) {
	return ReadConfWithPaths([]string{createTmpConfigFile(t, confYaml)})
}

// returns the the test path where the config file is stored
func createTmpConfigFile(t *testing.T, confYaml string) string {
	testDir := t.TempDir()
	configFilePath := testDir + "/" + ViperConfigName + "." + ViperConfigType
	err := os.WriteFile(configFilePath, []byte(confYaml), 0644)
	if err != nil {
		t.Fatal(err)
	}
	return testDir
}
