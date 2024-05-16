package ingestionapikey

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func TestIngestionApiKeyList(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)
	key1desc := "main key"

	cli.MockClient.ApiMocks.IngestionApiKeyAPI.GetIngestionApiKeysResponse.Result = []stackstate_api.IngestionApiKey{
		{
			Id:          1,
			Name:        "key1",
			Description: &key1desc,
			Expiration:  int64p(1590105600000),
		},
		{
			Id:          2,
			Name:        "key2",
			Description: nil,
			Expiration:  nil,
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	tableData := []printer.TableData{
		{
			Header: []string{"id", "name", "expiration", "description"},
			Data: [][]interface{}{
				{int64(1), "key1", "2020-05-22", &key1desc},
				{int64(2), "key2", "", (*string)(nil)},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "ingestion api keys"},
		},
	}

	assert.Equal(t, tableData, *cli.MockPrinter.TableCalls)
}
