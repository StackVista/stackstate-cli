package conf

import (
	"testing"
)

func TestRunner(t *testing.T) {

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
