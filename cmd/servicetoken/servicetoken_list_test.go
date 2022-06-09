package servicetoken

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func TestServiceTokenList(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)

	cli.MockClient.ApiMocks.ServiceTokenApi.GetServiceTokensResponse.Result = []stackstate_api.ServiceToken{
		{
			Id:   int64p(1),
			Name: "test",
			Roles: []string{
				"test-role",
				"another-role",
			},
			Expiration: int64p(1590105600000),
		},
		{
			Id:   int64p(2),
			Name: "test2",
			Roles: []string{
				"test-role",
			},
			Expiration: nil,
		},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	tableData := []printer.TableData{
		{
			Header: []string{"id", "name", "expiration", "roles"},
			Data: [][]interface{}{
				{int64(1), "test", "2020-05-22", []string{"test-role", "another-role"}},
				{int64(2), "test2", "", []string{"test-role"}},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "service tokens"},
		},
	}

	assert.Equal(t, tableData, *cli.MockPrinter.TableCalls)
}
