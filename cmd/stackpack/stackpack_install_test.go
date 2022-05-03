package stackpack

import (
	"errors"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
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
	mockProvisionResponse = stackstate_client.ProvisionResponse{
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
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Result = mockProvisionResponse
	return &cli, cmd
}

func TestStackpackInstallPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPackInstallCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "install", "--name", "zabbix",
		"--parameter", "zabbix_instance_name=test_name",
		"--parameter", "zabbix_instance_url=test_url",
	)

	assert.Equal(t,
		stackstate_client.ProvisionDetailsCall{
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

func TestParseParameter(t *testing.T) {
	testCase := []struct {
		desc     string
		input    string
		expKey   string
		expValue string
		expErr   error
	}{
		{
			desc:     "should parses successfully",
			input:    "name=mr=jon",
			expKey:   "name",
			expValue: "mr=jon",
			expErr:   nil,
		},
		{
			desc:     "equal operator at the end",
			input:    "name=",
			expKey:   "",
			expValue: "",
			expErr:   errors.New("expected parameter format key=value"),
		},
		{
			desc:     "no key",
			input:    "=jon",
			expKey:   "",
			expValue: "",
			expErr:   errors.New("expected parameter format key=value"),
		},
		{
			desc:     "no value",
			input:    "name=",
			expKey:   "",
			expValue: "",
			expErr:   errors.New("expected parameter format key=value"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.desc, func(t *testing.T) {
			key, value, err := parseParameter(tc.input)
			assert.Equal(t, tc.expErr, err)
			assert.Equal(t, tc.expKey, key)
			assert.Equal(t, tc.expValue, value)
		})
	}
}
