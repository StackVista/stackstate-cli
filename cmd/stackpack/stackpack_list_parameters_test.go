package stackpack

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	parameterName   = "zabbix"
	stepName        = "zabbix_instance_name"
	stepDisplayName = "Zabbix Instance Name"
	stepType        = "Text"

	firstStepName = "abbix_instance_name"
)

func setupStackpackListParametersFn(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
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
				{
					Name:    &firstStepName,
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
	cli, cmd := setupStackpackListParametersFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-parameters", "--name", parameterName)
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"name", "display name", "type"},
			Data:                [][]interface{}{{&firstStepName, &stepDisplayName, &stepType}, {&stepName, &stepDisplayName, &stepType}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("StackPack \"%s\"", parameterName)},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackListParameterPrintToJson(t *testing.T) {
	cli, cmd := setupStackpackListParametersFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list-parameters", "--name", parameterName, "-o", "json")
	expectedJsonCalls := []map[string]interface{}{{
		"parameters": []stackstate_api.StackPackStep{
			{
				Name:    &firstStepName,
				Display: &stepDisplayName,
				Value: &stackstate_api.StackPackStepValue{
					Type: &stepType,
				},
			}, {
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
