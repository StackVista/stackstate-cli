package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func TestAllCommandsAreCompatible(t *testing.T) {
	root := AllCommands(&config.Config{})
	assert.NotNil(t, root)
}
