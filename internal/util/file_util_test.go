package util

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFileShouldTruncateContents(t *testing.T) {
	dir := t.TempDir()

	err := WriteFile(fmt.Sprintf("%s/file", dir), []byte("hello johnny"))
	assert.NoError(t, err)

	err = WriteFile(fmt.Sprintf("%s/file", dir), []byte("hello"))
	assert.NoError(t, err)

	contents, err := os.ReadFile(fmt.Sprintf("%s/file", dir))
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(contents))
}
