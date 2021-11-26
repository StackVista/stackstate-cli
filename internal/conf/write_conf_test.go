package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSuccess(t *testing.T) {
	confIn := Conf{
		ApiUrl:   "https://my.stackstate.com/api",
		ApiToken: "BSOPSIY6Z3TuSmNIFzqPZyUMilggP9_M",
	}

	path := t.TempDir()
	err := WriteConfTo(confIn, path)
	if err != nil {
		t.Fatal(err)
	}

	confOut, err := ReadConfWithPaths(NewCommand(), []string{path})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, confIn, confOut)
}
