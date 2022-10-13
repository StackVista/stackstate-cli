package topic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	SomeTopicNames = []string{
		"foo",
		"bar",
		"baz",
	}
	SomeTopics = []stackstate_api.Topic{
		{Name: SomeTopicNames[0]},
		{Name: SomeTopicNames[1]},
		{Name: SomeTopicNames[2]},
	}
)

func TestTopicList(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.ListResponse.Result = SomeTopics

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.ListCalls, 1)

	expectedTable := []printer.TableData{
		{
			Header: []string{"Topic Name"},
			Data: [][]interface{}{
				{SomeTopicNames[0]},
				{SomeTopicNames[1]},
				{SomeTopicNames[2]},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "topics"},
		},
	}

	assert.Equal(t, expectedTable, *cli.MockPrinter.TableCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.ListCalls, 2)

	expectedJson := []map[string]interface{}{
		{
			"topics": SomeTopicNames,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}
