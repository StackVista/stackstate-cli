package topic

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	topicName             = "sts_correlated_connections"
	mockTopicListResponse = []stackstate_api.Topic{
		{
			Name: topicName,
		},
	}
)

func setupTopicListCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := TopicListCommand(&cli.Deps)
	return &cli, cmd
}

func TestTopicListPrintToTable(t *testing.T) {
	cli, cmd := setupTopicListCmd()
	cli.MockClient.ApiMocks.TopicApi.ListTopicsResponse.Result = mockTopicListResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name"},
			Data:                [][]interface{}{{topicName}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "Topics"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestTopicListPrintToJson(t *testing.T) {
	cli, cmd := setupTopicListCmd()
	cli.MockClient.ApiMocks.TopicApi.ListTopicsResponse.Result = mockTopicListResponse
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "-o", "json")
	expectedJsonCall := []map[string]interface{}{{
		"topics": [][]interface{}{{topicName}},
	}}
	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
}
