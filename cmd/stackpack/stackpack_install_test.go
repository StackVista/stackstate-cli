package stackpack

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	name                  = "zabbix"
	provisionID           = int64(1234)
	provisionName         = "provision_name"
	provisionStatus       = "PROVISIONING"
	provisionVersion      = "3.2.0"
	timeMilli             = int64(1438167001716)
	expectedUpdateTime    = time.UnixMilli(timeMilli)
	mockProvisionResponse = &stackstate_api.ProvisionResponse{
		Id:                  &provisionID,
		Name:                &provisionName,
		Status:              &provisionStatus,
		StackPackVersion:    &provisionVersion,
		LastUpdateTimestamp: &timeMilli,
	}
)

func setupStackPackInstallCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackInstallCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Result = *mockProvisionResponse
	return &cli, cmd
}

func TestStackpackInstallPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix",
		"--parameter", "zabbix_instance_name=test_name",
		"--parameter", "zabbix_instance_url=test_url",
	)

	assert.Equal(t,
		stackstate_api.ProvisionDetailsCall{
			PstackName: "zabbix",
			PrequestBody: &map[string]string{
				"zabbix_instance_name": "test_name",
				"zabbix_instance_url":  "test_url",
			},
		},
		(*cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsCalls)[0],
	)
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"id", "name", "status", "version", "last updated"},
			Data:                [][]interface{}{{&provisionID, &provisionName, &provisionStatus, &provisionVersion, expectedUpdateTime}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + name},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestStackpackInstallComplexParameters(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix",
		"--parameter", "one=foo,two=bar",
		"--parameter", "i=am=complex",
		"--parameter", "empty=",
		"--parameter", "comma=1,2",
	)

	assert.Equal(t,
		stackstate_api.ProvisionDetailsCall{
			PstackName: "zabbix",
			PrequestBody: &map[string]string{
				"one":   "foo",
				"two":   "bar",
				"empty": "",
				"i":     "am=complex",
				"comma": "1,2",
			},
		},
		(*cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsCalls)[0],
	)
}

func TestStackpackInstallPrintsToJson(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix", "--json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"instance": mockProvisionResponse},
		},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
