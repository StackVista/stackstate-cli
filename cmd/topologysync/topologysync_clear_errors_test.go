package topologysync

import (
	"fmt"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestTopologySyncClearErrorsId(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ClearErrorsCommand(&cli.Deps)

	id := "23"

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-i", id)

	calls := *cli.MockClient.ApiMocks.TopologySynchronizationApi.PostTopologySynchronizationStreamClearErrorsCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, id, *calls[0].Pidentifier)
	assert.Equal(t, NodeIdType, *calls[0].PidentifierType)

	expected := []string{
		fmt.Sprintf("Errors cleared from topology synchronization: %s", id),
	}
	assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
}

func TestTopologySyncClearErrorsIdentifier(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ClearErrorsCommand(&cli.Deps)

	identifier := "urn:stackstate:5"

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", identifier)

	calls := *cli.MockClient.ApiMocks.TopologySynchronizationApi.PostTopologySynchronizationStreamClearErrorsCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, identifier, *calls[0].Pidentifier)
	assert.Equal(t, IdentifierType, *calls[0].PidentifierType)

	expected := []string{
		fmt.Sprintf("Errors cleared from topology synchronization: %s", identifier),
	}
	assert.Equal(t, expected, *cli.MockPrinter.SuccessCalls)
}

func TestTopologyClearErrorsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ClearErrorsCommand(&cli.Deps)

	id := "5"

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", id, "-o", "json")

	expected := []map[string]interface{}{
		{
			"cleared": id,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
