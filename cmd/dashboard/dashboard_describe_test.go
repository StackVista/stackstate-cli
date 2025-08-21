package dashboard

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupDescribeCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardDescribeCommand(&cli.Deps)
	return &cli, cmd
}

func createTestDashboard() sts.DashboardReadSchema {
	fullDashboard := sts.DashboardReadFullSchema{
		Id:          firstDashboardId,
		Name:        firstDashboardName,
		Identifier:  firstDashboardIdentifier,
		Description: firstDashboardDescription,
		Scope:       firstDashboardScope,
		Dashboard: sts.PersesDashboard{
			Spec: &sts.PersesDashboardSpec{
				Layouts: []sts.PersesLayout{
					{
						Kind: "Grid",
						Spec: sts.PersesLayoutSpec{
							Items: []sts.PersesGridItem{},
						},
					},
				},
				Panels: &map[string]sts.PersesPanel{},
			},
		},
	}
	return sts.DashboardReadFullSchemaAsDashboardReadSchema(&fullDashboard)
}

func TestDashboardDescribe(t *testing.T) {
	cli, cmd := setupDescribeCmd(t)
	expectedDashboard := createTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = expectedDashboard

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1")
	assert.Nil(t, err)

	// Verify that the command printed the dashboard JSON
	assert.Len(t, *cli.MockPrinter.PrintLnCalls, 1)
	printedOutput := (*cli.MockPrinter.PrintLnCalls)[0]
	assert.Contains(t, printedOutput, `"id": 1`)
	assert.Contains(t, printedOutput, `"name": "aDashboard"`)
	assert.Contains(t, printedOutput, `"identifier": "urn:custom:dashboard:aDashboard"`)
}

func TestDashboardDescribeWithIdentifier(t *testing.T) {
	cli, cmd := setupDescribeCmd(t)
	expectedDashboard := createTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = expectedDashboard

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--identifier", "urn:custom:dashboard:aDashboard")
	assert.Nil(t, err)

	// Verify the API was called with the correct identifier
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls, 1)
	assert.Equal(t, "urn:custom:dashboard:aDashboard", (*cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls)[0].PdashboardIdOrUrn)
}

func TestDashboardDescribeJson(t *testing.T) {
	cli, cmd := setupDescribeCmd(t)
	expectedDashboard := createTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = expectedDashboard

	di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1", "-o", "json") //nolint:errcheck

	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Contains(t, jsonOutput, "data")
	assert.Contains(t, jsonOutput, "format")
	assert.Equal(t, "json", jsonOutput["format"])
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestDashboardDescribeToFile(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	cli, cmd := setupDescribeCmd(t)
	expectedDashboard := createTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = expectedDashboard

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1", "--file", file.Name())
	assert.Nil(t, err)

	// Verify success message was printed
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], "dashboard exported to:")
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], file.Name())

	// Verify file contents
	body, err := os.ReadFile(file.Name())
	assert.Nil(t, err)
	assert.Contains(t, string(body), `"id": 1`)
	assert.Contains(t, string(body), `"name": "aDashboard"`)
}

func TestDashboardDescribeToFileJson(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	cli, cmd := setupDescribeCmd(t)
	expectedDashboard := createTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = expectedDashboard

	di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1", "--file", file.Name(), "-o", "json") //nolint:errcheck

	// Verify JSON response with filepath
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, file.Name(), jsonOutput["filepath"])
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestDashboardDescribeMissingArgs(t *testing.T) {
	cli, cmd := setupDescribeCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "one of the required flags {id | identifier} not set")
}
