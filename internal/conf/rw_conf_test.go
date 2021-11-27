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
		tests.TestNoColorFlag(t)
		tests.TestNoColorEnv(t)
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

	confOut, err := readConfWithPaths(newCmd(), viper.New(), []string{path})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, confIn, confOut)
}

// executed by TestWriteReadRunner
func (p ReadTests) TestMissingConf(t *testing.T) {
	_, err := readConfWithPaths(newCmd(), viper.New(), []string{})
	assert.NotNil(t, err, "TestMissingConf")
	assert.IsType(t, ReadConfError{}, err, "TestMissingConf")
	assert.IsType(t, MissingConfError{}, err.(ReadConfError).RootCause, "TestMissingConf")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestYamlParseError(t *testing.T) {
	confYaml := `XXXX`
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err, "TestYamlParseError")
	assert.IsType(t, viper.ConfigParseError{}, err.(ReadConfError).RootCause, "TestYamlParseError")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestValidationError(t *testing.T) {
	confYaml := "api.token: 123871283"
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err, "TestValidationError")
	assert.IsType(t, ValidateConfError{}, err.(ReadConfError).RootCause, "TestValidationError")
	valErrs := err.(ReadConfError).RootCause.(ValidateConfError).ValidationErrors
	assert.Greater(t, len(valErrs), 0, "TestValidationError")
	assert.IsType(t, MissingFieldError{}, valErrs[0], "TestValidationError")
	assert.Equal(t, "api-url", valErrs[0].(MissingFieldError).FieldName, "TestValidationError")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromYaml(t *testing.T) {
	conf, err := readConfFromFile(t, MinimalConfYaml)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl, "TestLoadSuccessFromYaml")
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken, "TestLoadSuccessFromYaml")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromMinimumRequiredEnvs(t *testing.T) {
	envs := strings.Split(strings.ReplaceAll(MinimumRequiredEnvVars, " ", ""), ",")
	defer os.Unsetenv(envs[0])
	defer os.Unsetenv(envs[1])
	os.Setenv(envs[0], "https://my.stackstate.com/api")
	os.Setenv(envs[1], "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M")

	conf, err := readConfWithPaths(newCmd(), viper.New(), []string{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl, "TestLoadSuccessFromMinimumRequiredEnvs")
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken, "TestLoadSuccessFromMinimumRequiredEnvs")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromMinimumFlags(t *testing.T) {
	cmd := newCmd()
	flags := strings.Split(strings.ReplaceAll(MinimumRequiredFlags, " ", ""), ",")
	assert.Equal(t, 2, len(flags))
	cmd.Flags().String(flags[0], "", "")
	cmd.Flags().String(flags[1], "", "")
	cmd.Flags().Set(flags[0], "https://my.stackstate.com/api")
	cmd.Flags().Set(flags[1], "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M")

	conf, err := readConfWithPaths(cmd, viper.New(), []string{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.ApiUrl, "TestLoadSuccessFromMinimumFlags")
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken, "TestLoadSuccessFromMinimumFlags")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestNoColorOnTermIsDumb(t *testing.T) {
	term := os.Getenv("TERM")
	defer os.Setenv("TERM", term)
	os.Setenv("TERM", "Dumb")
	conf := readConfWithMinimal(t, newCmd())
	assert.Equal(t, true, conf.NoColor, "TestNoColorOnTermIsDumb")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestNoColorFlag(t *testing.T) {
	cmd := newCmd()
	cmd.Flags().Bool("no-color", false, "") // register flag
	cmd.Flags().Set("no-color", "true")
	conf := readConfWithMinimal(t, cmd)
	assert.Equal(t, true, conf.NoColor, "TestNoColorFlag")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestNoColorEnv(t *testing.T) {
	defer os.Unsetenv("STS_CLI_NO_COLOR")
	os.Setenv("STS_CLI_NO_COLOR", "true")
	conf := readConfWithMinimal(t, newCmd())
	assert.Equal(t, true, conf.NoColor)
}

////------------- Fixture code -----------

func readConfFromFile(t *testing.T, confYaml string) (Conf, error) {
	return readConfWithPaths(
		newCmd(),
		viper.New(),
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

func newCmd() *cobra.Command {
	return &cobra.Command{}
}

// Always reads the config succesfully. Handy for testing success paths.
func readConfWithMinimal(t *testing.T, cmd *cobra.Command) Conf {
	conf, err := readConfWithPaths(
		cmd,
		viper.New(),
		[]string{createTmpConfigFile(t, MinimalConfYaml)},
	)
	if err != nil {
		t.Fatal(err)
	}
	return conf
}
