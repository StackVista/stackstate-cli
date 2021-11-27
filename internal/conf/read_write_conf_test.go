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

type WriteTests struct{}
type ReadTests struct{}

/*
Unfortuntately the tests in these suites interfere and this work around is
necessary, because OS environment variables are shared.

If you run these tests indiviudally they will succeed, but if you
run them all then they will fail, because Go's test runner runs these tests
concurrently.

This work around is the best I could do. I also tried using a Mutex, but to
no avail.
*/
func TestWriteReadRunner(t *testing.T) {
	t.Run("write_conf", func(t *testing.T) {
		tests := WriteTests{}
		tests.TestWriteSuccess(t)
	})
	t.Run("read_conf", func(t *testing.T) {
		tests := ReadTests{}
		tests.TestMissingConf(t)
		tests.TestYamlParseError(t)
		tests.TestValidationError(t)
		tests.TestLoadSuccessFromYaml(t)
		tests.TestLoadSuccessFromMinimumRequiredEnvs(t)
		tests.TestLoadSuccessFromMinimumFlags(t)
		tests.TestNoColorOnTermIsDumb(t)
	})
}

// executed by TestWriteReadRunner
func (p WriteTests) TestWriteSuccess(t *testing.T) {
	confIn := Conf{
		ApiUrl:   "https://write.stackstate.com/api",
		ApiToken: "BSOPSIYZ4TuSzNIFzqPZyUMilggP9_M",
	}

	path := t.TempDir()
	err := WriteConfTo(confIn, path)
	if err != nil {
		t.Fatal(err)
	}

	confOut, err := readConfWithPaths(NewCommand(), []string{path})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, confIn, confOut)
}

// executed by TestWriteReadRunner
func (p ReadTests) TestMissingConf(t *testing.T) {
	_, err := readConfWithPaths(NewCommand(), []string{})
	assert.NotNil(t, err)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, MissingConfError{}, err.(ReadConfError).RootCause)
}

// executed by TestWriteReadRunner
func (p ReadTests) TestYamlParseError(t *testing.T) {
	confYaml := `XXXX`
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err)
	assert.IsType(t, viper.ConfigParseError{}, err.(ReadConfError).RootCause)
}

// executed by TestWriteReadRunner
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

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromYaml(t *testing.T) {
	conf, err := readConfFromFile(t, MinimalConfYaml)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromMinimumRequiredEnvs(t *testing.T) {
	envs := strings.Split(strings.ReplaceAll(MinimumRequiredEnvVars, " ", ""), ",")
	defer os.Unsetenv(envs[0])
	defer os.Unsetenv(envs[1])
	os.Setenv(envs[0], "https://my.stackstate.com/api")
	os.Setenv(envs[1], "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M")

	conf, err := readConfWithPaths(NewCommand(), []string{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl)
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken)
}

// executed by TestWriteReadRunner
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

// executed by TestWriteReadRunner
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
