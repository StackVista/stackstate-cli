package conf

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	MinimalConfYaml = `
api.url: https://my.stackstate.com/api
api.token: BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M
`
)

type ReadTests struct{}

func (p ReadTests) TestMissingConf(t *testing.T) {
	_, err := readConfWithPaths(NewCommand(), []string{})
	assert.NotNil(t, err)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, MissingConfError{}, err.(ReadConfError).RootCause)
}

func (p ReadTests) TestYamlParseError(t *testing.T) {
	confYaml := `XXXX`
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, viper.ConfigParseError{}, err.(ReadConfError).RootCause)
}

func (p ReadTests) TestValidationError(t *testing.T) {
	confYaml := "api.token: 123871283"
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, ValidateConfError{}, err.(ReadConfError).RootCause)
	valErrs := err.(ReadConfError).RootCause.(ValidateConfError).ValidationErrors
	assert.Greater(t, len(valErrs), 0)
	assert.IsType(t, MissingFieldError{}, valErrs[0])
	assert.Equal(t, "api-url", valErrs[0].(MissingFieldError).FieldName)
}

func (p ReadTests) TestLoadSuccessFromYaml(t *testing.T) {
	conf, err := readConfFromFile(t, MinimalConfYaml)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

func (p ReadTests) TestLoadSuccessFromMinimumRequiredEnvs(t *testing.T) {
	envs := strings.Split(strings.ReplaceAll(MinimumRequiredEnvVars, " ", ""), ",")
	assert.Equal(t, 2, len(envs))
	t.Setenv(envs[0], "https://my.stackstate.com/api")
	t.Setenv(envs[1], "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M")

	conf, err := readConfWithPaths(NewCommand(), []string{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

func (p ReadTests) TestLoadSuccessFromMinimumFlags(t *testing.T) {
	cmd := NewCommand()
	flags := strings.Split(strings.ReplaceAll(MinimumRequiredFlags, " ", ""), ",")
	assert.Equal(t, 2, len(flags))
	cmd.Flags().String(flags[0], "", "")
	cmd.Flags().String(flags[1], "", "")
	cmd.Flags().Set(flags[0], "https://my.stackstate.com/api")
	cmd.Flags().Set(flags[1], "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M")

	conf, err := readConfWithPaths(cmd, []string{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

func (p ReadTests) TestNoColorOnTermIsDumb(t *testing.T) {
	t.Setenv("TERM", "dumb")

	conf, err := readConfFromFile(t, MinimalConfYaml)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, conf.NoColor)
}

////------------- Fixture code -----------

func readConfFromFile(t *testing.T, confYaml string) (Conf, error) {
	return readConfWithPaths(
		NewCommand(),
		[]string{createTmpConfigFile(t, confYaml)},
	)
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

func NewCommand() *cobra.Command {
	x := cobra.Command{}
	return &x
}
