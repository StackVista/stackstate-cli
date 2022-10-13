package topic

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	MessageContents     = map[string]interface{}{"test": "value"}
	MessageContentsJson = "{\"test\":\"value\"}"

	Message1 = stackstate_api.NewMessage("key", 0, 0, MessageContents)
	Message2 = stackstate_api.NewMessage("key", 0, 2, MessageContents)
	Message3 = stackstate_api.NewMessage("key", 0, 3, MessageContents)
	Message4 = stackstate_api.NewMessage("key", 1, 1, MessageContents)
	Message5 = stackstate_api.NewMessage("key", 1, 4, MessageContents)

	Part0MessagesList = []stackstate_api.Message{
		*Message1,
		*Message2,
		*Message3,
	}
	Part0TopicMessages = stackstate_api.NewMessages(Part0MessagesList)

	AllMessagesList = []stackstate_api.Message{
		*Message1,
		*Message2,
		*Message3,
		*Message4,
		*Message5,
	}
	AllMessagesListSorted = []stackstate_api.Message{
		*Message1,
		*Message4,
		*Message2,
		*Message3,
		*Message5,
	}
	AllTopicMessages = stackstate_api.NewMessages(AllMessagesList)

	ZeroMessagesList  = make([]stackstate_api.Message, 0)
	ZeroTopicMessages = stackstate_api.NewMessages(ZeroMessagesList)
)

func TestTopicDescribeJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *Part0TopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 1)

	expectedJson := []map[string]interface{}{
		{
			"messages": Part0MessagesList,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}

func TestTopicDescribeTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *Part0TopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test")

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 1)

	expectedTable := []printer.TableData{
		{
			Header: []string{"Key", "Partition", "Offset", "Message"},
			Data: [][]interface{}{
				{Message1.Key, Message1.Partition, Message1.Offset, MessageContentsJson},
				{Message2.Key, Message2.Partition, Message2.Offset, MessageContentsJson},
				{Message3.Key, Message3.Partition, Message3.Offset, MessageContentsJson},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "messages"},
		},
	}

	assert.Equal(t, expectedTable, *cli.MockPrinter.TableCalls)
}

func TestTopicDescribeTableSorting(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *AllTopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test")

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 1)

	// NOTE Sorted by offset ignoring the partition.
	expectedTable := []printer.TableData{
		{
			Header: []string{"Key", "Partition", "Offset", "Message"},
			Data: [][]interface{}{
				{Message1.Key, Message1.Partition, Message1.Offset, MessageContentsJson},
				{Message4.Key, Message4.Partition, Message4.Offset, MessageContentsJson},
				{Message2.Key, Message2.Partition, Message2.Offset, MessageContentsJson},
				{Message3.Key, Message3.Partition, Message3.Offset, MessageContentsJson},
				{Message5.Key, Message5.Partition, Message5.Offset, MessageContentsJson},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "messages"},
		},
	}

	assert.Equal(t, expectedTable, *cli.MockPrinter.TableCalls)
}

func TestTopicDescribeFileIO(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "topic_test_")
	if err != nil {
		panic(err)
	}
	filePath := file.Name()
	file.Close()

	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *Part0TopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "--file", filePath)

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 1)

	expectedStrings := []string{
		fmt.Sprintf("Messages saved to '%s'", filePath),
	}
	assert.Equal(t, expectedStrings, *cli.MockPrinter.SuccessCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "-o", "json", "--file", filePath)

	assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 2)
	expectedJson := []map[string]interface{}{
		{
			"filepath": filePath,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)

	expectedContents, err := os.ReadFile("messages.json")
	assert.Nil(t, err)
	written, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	// NOTE Trim the newlines.
	assert.Equal(t, strings.Trim(string(expectedContents), "\n"), string(written))
}

func TestTopicDescribePaginationDefaults(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *AllTopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test")

	calls := *cli.MockClient.ApiMocks.TopicApi.DescribeCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, "test", calls[0].Ptopic)
	assert.Equal(t, DefaultNumber, *calls[0].Plimit)
	assert.Equal(t, DefaultOffset, *calls[0].Poffset)
	assert.Nil(t, calls[0].Ppartition)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "--partition", "23")

	calls = *cli.MockClient.ApiMocks.TopicApi.DescribeCalls
	assert.Len(t, calls, 2)
	assert.Equal(t, int32(23), *calls[1].Ppartition)
}

type InvalidArgs struct {
	Name  string
	Value int32
}

func TestTopicDescribePaginationLimits(t *testing.T) {
	tests := []InvalidArgs{
		{Name: Offset, Value: -23},
		{Name: Number, Value: -5},
		{Name: Number, Value: 0},
		{Name: PageSize, Value: -13},
		{Name: PageSize, Value: 0},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cli := di.NewMockDeps(t)
			cmd := DescribeCommand(&cli.Deps)

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "test", fmt.Sprintf("--%s", test.Name), fmt.Sprintf("%d", test.Value))

			assert.Equal(t, argValueError(test.Name, test.Value), err)
			assert.Len(t, *cli.MockClient.ApiMocks.TopicApi.DescribeCalls, 0)
		})
	}
}

func TestTopicDescribePaginationMoreThanAvailable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *AllTopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "--nr", "10", "-o", "json")

	calls := *cli.MockClient.ApiMocks.TopicApi.DescribeCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, "test", calls[0].Ptopic)
	assert.Equal(t, int32(0), *calls[0].Poffset)
	assert.Equal(t, int32(10), *calls[0].Plimit)

	expectedJson := []map[string]interface{}{
		{
			"messages": AllMessagesListSorted,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}

func TestTopicDescribePaginationPagedRetrieval(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.TopicApi.DescribeResponse.Result = *Part0TopicMessages

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test", "--nr", "10", "--pagesize", "3", "-o", "json")

	calls := *cli.MockClient.ApiMocks.TopicApi.DescribeCalls
	assert.Len(t, calls, 4)
	assert.Equal(t, int32(0), *calls[0].Poffset)
	assert.Equal(t, int32(3), *calls[0].Plimit)

	assert.Equal(t, int32(3), *calls[1].Poffset)
	assert.Equal(t, int32(3), *calls[1].Plimit)

	assert.Equal(t, int32(6), *calls[2].Poffset)
	assert.Equal(t, int32(3), *calls[2].Plimit)

	assert.Equal(t, int32(9), *calls[3].Poffset)
	assert.Equal(t, int32(1), *calls[3].Plimit)

	expectedJson := []map[string]interface{}{
		{
			// NOTE The mock API returns more than limit on the last call.
			// NOTE The messages are then sorted by offset.
			"messages": []stackstate_api.Message{
				*Message1,
				*Message1,
				*Message1,
				*Message1,
				*Message2,
				*Message2,
				*Message2,
				*Message2,
				*Message3,
				*Message3,
				*Message3,
				*Message3,
			},
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}
