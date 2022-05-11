package stackpack

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	parameterName   = "zabbix"
	stepName        = "zabbix_instance_name"
	stepDisplayName = "Zabbix Instance Name"
	stepType        = "Text"
)

func setupStackpackListParametersFn() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackListParameterCommand(&cli.Deps)
	mockResponse := []stackstate_api.Sstackpack{
		{
			Name: &parameterName,
			Steps: []stackstate_api.StackPackStep{
				{
					Name:    &stepName,
					Display: &stepDisplayName,
					Value: &stackstate_api.StackPackStepValue{
						Type: &stepType,
					},
				},
			},
		},
	}
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = mockResponse
	return &cli, cmd
}

func TestStackpackListParameterPrintToTable(t *testing.T) {
	cli, cmd := setupStackpackListParametersFn()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-parameters", "--name", parameterName)
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "type"},
			Data:                [][]interface{}{{&stepName, &stepDisplayName, &stepType}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("StackPack %s does not exist", parameterName)},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListParameterPrintToJson(t *testing.T) {
	cli, cmd := setupStackpackListParametersFn()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-parameters", "--name", parameterName, "--json")
	expectedJsonCalls := []map[string]interface{}{{
		"parameters": []stackstate_api.StackPackStep{
			{
				Name:    &stepName,
				Display: &stepDisplayName,
				Value: &stackstate_api.StackPackStepValue{
					Type: &stepType,
				},
			},
		},
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
