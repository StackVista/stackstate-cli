package conf

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMissingConf(t *testing.T) {
	_, err := ReadConfWithPaths([]string{})
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, MissingConfError{}, err.(ReadConfError).RootCause)
}

func TestYamlParseError(t *testing.T) {
	confYaml := `XXXX`
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, viper.ConfigParseError{}, err.(ReadConfError).RootCause)
}

func TestValidationError(t *testing.T) {
	confYaml := "api-token: 123871283"
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, ValidateConfError{}, err.(ReadConfError).RootCause)
	valErrs := err.(ReadConfError).RootCause.(ValidateConfError).ValidationErrors
	assert.Greater(t, len(valErrs), 0)
	assert.IsType(t, MissingFieldError{}, valErrs[0])
	assert.Equal(t, "api-url", valErrs[0].(MissingFieldError).FieldName)
}

func TestLoadSuccessFromYaml(t *testing.T) {
	confYaml := `
api-url: https://my.stackstate.com/api
api-token: BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M
`

	conf, err := readConfFromFile(t, confYaml)

	if err != nil {
		t.Fatal(err)
	}
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
