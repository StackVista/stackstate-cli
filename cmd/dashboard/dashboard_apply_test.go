package dashboard

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setDashboardApplyCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardApplyCommand(&cli.Deps)
	return &cli, cmd
}

func createTestApplyResult() sts.DashboardReadFullSchema {
	return sts.DashboardReadFullSchema{
		Id:          1234,
		Name:        "applied-dashboard",
		Identifier:  "urn:custom:dashboard:applied-dashboard",
		Description: "Dashboard created via apply",
		Scope:       sts.DASHBOARDSCOPE_PUBLIC_DASHBOARD,
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
}

func TestShouldApplyDashboardCreate(t *testing.T) {
	// Create a temporary file with dashboard JSON
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	dashboardYAML := `name: applied-dashboard
description: Dashboard created via apply
scope: publicDashboard
dashboard:
  spec:
    layouts:
      - kind: Grid
        spec:
          items: []
    panels: {}
`

	_, err = file.WriteString(dashboardYAML)
	assert.Nil(t, err)
	file.Close()

	cli, cmd := setDashboardApplyCmd(t)
	expectedResult := createTestApplyResult()
	cli.MockClient.ApiMocks.DashboardsApi.CreateDashboardResponse.Result = expectedResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name())

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.CreateDashboardCalls, 1)
	createCall := (*cli.MockClient.ApiMocks.DashboardsApi.CreateDashboardCalls)[0]
	assert.Equal(t, "applied-dashboard", createCall.PdashboardWriteSchema.Name)
	assert.Equal(t, "Dashboard created via apply", createCall.PdashboardWriteSchema.Description)
	assert.Equal(t, sts.DASHBOARDSCOPE_PUBLIC_DASHBOARD, createCall.PdashboardWriteSchema.Scope)
	assert.Equal(t, "Dashboard created successfully! ID: 1234, Name: applied-dashboard", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldApplyDashboardUpdate(t *testing.T) {
	// Create a temporary file with dashboard update JSON (includes ID)
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	updateYAML := `id: 1234
name: updated-dashboard
description: Updated dashboard description
scope: privateDashboard
`

	_, err = file.WriteString(updateYAML)
	assert.Nil(t, err)
	file.Close()

	cli, cmd := setDashboardApplyCmd(t)
	expectedResult := createTestApplyResult()
	expectedResult.Name = "updated-dashboard"
	expectedResult.Description = "Updated dashboard description"
	expectedResult.Scope = sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD
	cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Result = expectedResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name())

	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 1)
	patchCall := (*cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls)[0]
	assert.Equal(t, "1234", patchCall.PdashboardIdOrUrn)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Name)
	assert.Equal(t, "updated-dashboard", *patchCall.PdashboardPatchSchema.Name)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Description)
	assert.Equal(t, "Updated dashboard description", *patchCall.PdashboardPatchSchema.Description)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Scope)
	assert.Equal(t, sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD, *patchCall.PdashboardPatchSchema.Scope)
	assert.Equal(t, "Dashboard updated successfully! ID: 1234, Name: updated-dashboard", (*cli.MockPrinter.SuccessCalls)[0])
}

func TestShouldApplyDashboardWithJson(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	dashboardYAML := `name: yaml-output-dashboard
description: Dashboard for JSON output test
scope: publicDashboard
dashboard:
  spec:
    layouts: []
    panels: {}
`

	_, err = file.WriteString(dashboardYAML)
	assert.Nil(t, err)
	file.Close()

	cli, cmd := setDashboardApplyCmd(t)
	expectedResult := createTestApplyResult()
	cli.MockClient.ApiMocks.DashboardsApi.CreateDashboardResponse.Result = expectedResult

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", file.Name(), "--output", "json")

	expectedJsonCall := []map[string]interface{}{
		{
			"dashboard": &expectedResult,
		},
	}
	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestApplyDashboardInvalidFileType(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	cli, cmd := setDashboardApplyCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--file", file.Name())

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unsupported file type: .txt. Only .yaml files are supported")
}

func TestApplyDashboardMissingFile(t *testing.T) {
	cli, cmd := setDashboardApplyCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--file", "/nonexistent/file.yaml")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot read file")
}

func TestApplyDashboardInvalidJSON(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	invalidYAML := `name: test
invalid yaml
`

	_, err = file.WriteString(invalidYAML)
	assert.Nil(t, err)
	file.Close()

	cli, cmd := setDashboardApplyCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--file", file.Name())

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to parse YAML")
}

func TestApplyDashboardMissingName(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_dashboard_*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	invalidDashboard := `description: Dashboard without name
scope: publicDashboard
`

	_, err = file.WriteString(invalidDashboard)
	assert.Nil(t, err)
	file.Close()

	cli, cmd := setDashboardApplyCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--file", file.Name())

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "dashboard name is required")
}

func TestApplyDashboardMissingFileFlag(t *testing.T) {
	cli, cmd := setDashboardApplyCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"file\" not set")
}
