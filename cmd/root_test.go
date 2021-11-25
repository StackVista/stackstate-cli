package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestAllCommandsAreCompatible(t *testing.T) {
	root := AllCommands(&di.Deps{})
	assert.NotNil(t, root)
}
