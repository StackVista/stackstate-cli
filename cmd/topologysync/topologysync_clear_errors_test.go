package topologysync

import (
	"fmt"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestTopologySyncClearErrors(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ClearErrorsCommand(&cli.Deps)

	id := 23
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", fmt.Sprintf("%d", id))

	expected := []string{
		fmt.Sprintf("Errors cleared from topology synchronization: %d", id),
	}
	assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
}

func TestTopologyClearErrorsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ClearErrorsCommand(&cli.Deps)

	id := 5

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", fmt.Sprintf("%d", id), "-o", "json")

	expected := []map[string]interface{}{
		{
			"cleared": int64(id),
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
