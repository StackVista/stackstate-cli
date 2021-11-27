package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type WriteTests struct{}

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
