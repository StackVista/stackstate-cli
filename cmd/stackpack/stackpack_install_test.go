package stackpack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupInstallCommandFn() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := StackpackInstallCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestStackpackInstallPrintsToTable(t *testing.T) {
	name := "zabbix"
	provisionID := int64(1234)
	provisionName := "provision_name"
	provisionStatus := "PROVISIONING"
	provisionVersion := "3.2.0"
	timeMilli := int64(1438167001716)
	expectedUpdateTime := time.UnixMilli(timeMilli)
	mockProvisionResponse := stackstate_client.ProvisionResponse{
		Id:                  &provisionID,
		Name:                &provisionName,
		Status:              &provisionStatus,
		StackPackVersion:    &provisionVersion,
		LastUpdateTimestamp: &timeMilli,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(mockProvisionResponse); err != nil {
		t.Fatal(err)
	}
	cli, cmd := setupInstallCommandFn()
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Response = &http.Response{Body: ioutil.NopCloser(buf)}
	cli.MockClient.ApiMocks.StackpackApi.ProvisionDetailsResponse.Result = mockProvisionResponse
	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "install", "--name", "zabbix",
		"--parameter", "zabbix_instance_name=test_name",
		"--parameter", "zabbix_instance_url=test_url",
	)

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"id", "name", "status", "version", "last updated"},
			Data:                [][]interface{}{{&provisionID, &provisionName, &provisionStatus, &provisionVersion, expectedUpdateTime}},
			StructData:          mockProvisionResponse,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + name},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
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
