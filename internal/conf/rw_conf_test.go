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
url: https://my.stackstate.com/api
api-token: BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M
`
)

type WriteTests struct{}
type ReadTests struct{}

/*
Unfortuntately the tests in these suites interfere and this work around is
necessary, because OS environment variables are shared.

If you run these tests individually they will succeed, but if you
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
		tests.TestYamlParseError(t)
		tests.TestMissingURLFieldValidationError(t)
		tests.TestIncorrectURLFieldValidationError(t)
		tests.TestLoadSuccessFromYaml(t)
		tests.TestLoadSuccessFromMinimumRequiredEnvs(t)
		tests.TestLoadSuccessFromMinimumFlags(t)
	})
}

// executed by TestWriteReadRunner
func (p WriteTests) TestWriteSuccess(t *testing.T) {
	confIn := Conf{
		URL:      "https://write.stackstate.com/",
		ApiToken: "BSOPSIYZ4TuSzNIFzqPZyUMilggP9_M",
		ApiPath:  "test",
	}

	path := t.TempDir()
	err := WriteConfTo(confIn, path+"/"+ViperConfigName+"."+ViperConfigType)
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
func (p ReadTests) TestYamlParseError(t *testing.T) {
	confYaml := `XXXX`
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err, "TestYamlParseError")
	assert.IsType(t, viper.ConfigParseError{}, err.(ReadConfError).RootCause, "TestYamlParseError")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestMissingURLFieldValidationError(t *testing.T) {
	confYaml := "api-token: 123871283"
	_, err := readConfFromFile(t, confYaml)
	assert.IsType(t, ReadConfError{}, err, "TestValidationError")
	assert.IsType(t, ValidateConfError{}, err.(ReadConfError).RootCause, "TestValidationError")
	valErrs := err.(ReadConfError).RootCause.(ValidateConfError).ValidationErrors
	assert.Greater(t, len(valErrs), 0, "TestValidationError")
	assert.IsType(t, MissingFieldError{}, valErrs[0], "TestValidationError")
	assert.Equal(t, "url", valErrs[0].(MissingFieldError).FieldName, "TestValidationError")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestIncorrectURLFieldValidationError(t *testing.T) {
	confYaml := "api-token: 123871283\nurl: test"
	_, err := readConfFromFile(t, confYaml)
	assert.Equal(t, "could not load StackState CLI config\nValidation error: URL test must start with \"https://\" or \"http://\"", err.Error(), "TestIncorrectURLFieldValidationError")
}

// executed by TestWriteReadRunner
func (p ReadTests) TestLoadSuccessFromYaml(t *testing.T) {
	conf, err := readConfFromFile(t, MinimalConfYaml)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://my.stackstate.com/api", conf.URL, "TestLoadSuccessFromYaml")
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
	assert.Equal(t, "https://my.stackstate.com/api", conf.URL, "TestLoadSuccessFromMinimumRequiredEnvs")
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken, "TestLoadSuccessFromMinimumRequiredEnvs")
}

// executed by TestWriteReadRunner
//nolint:golint,errcheck
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
	assert.Equal(t, "https://my.stackstate.com/api", conf.URL, "TestLoadSuccessFromMinimumFlags")
	assert.Equal(t, "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M", conf.ApiToken, "TestLoadSuccessFromMinimumFlags")
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
	err := os.WriteFile(configFilePath, []byte(confYaml), 0600)
	if err != nil {
		t.Fatal(err)
	}
	return testDir
}

func newCmd() *cobra.Command {
	return &cobra.Command{}
}
