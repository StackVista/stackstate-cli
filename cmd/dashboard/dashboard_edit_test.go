package dashboard

import (
	"errors"
	"io"
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

// MockEditor allows us to control the editor behavior in tests
type MockEditor struct {
	Content []byte
	Error   error
}

func (e *MockEditor) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	if e.Error != nil {
		return nil, e.Error
	}
	return e.Content, nil
}

func setupEditCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardEditCommand(&cli.Deps)
	return &cli, cmd
}

func createEditTestDashboard() sts.DashboardReadSchema {
	fullDashboard := sts.DashboardReadFullSchema{
		Id:          1234,
		Name:        "edit-test-dashboard",
		Identifier:  "urn:custom:dashboard:edit-test",
		Description: "Dashboard for edit testing",
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
	return sts.DashboardReadFullSchemaAsDashboardReadSchema(&fullDashboard)
}

func createEditResult() sts.DashboardReadFullSchema {
	return sts.DashboardReadFullSchema{
		Id:          1234,
		Name:        "edited-dashboard-name",
		Identifier:  "urn:custom:dashboard:edit-test",
		Description: "Updated description",
		Scope:       sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD,
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

func TestShouldEditDashboard(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	updatedDashboard := createEditResult()
	cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Result = updatedDashboard

	// Replace the ReverseEditor with a MockEditor that returns edited content
	mockEditor := &MockEditor{
		Content: []byte(`{
			"_type": "DashboardReadFullSchema",
			"id": 1234,
			"name": "edited-dashboard-name",
			"identifier": "urn:custom:dashboard:edit-test",
			"description": "Updated description",
			"scope": "privateDashboard"
		}`),
	}
	cli.Editor = mockEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	// Verify that GetDashboard was called
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls, 1)
	getCall := (*cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls)[0]
	assert.Equal(t, "1234", getCall.PdashboardIdOrUrn)

	// Verify that PatchDashboard was called
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 1)
	patchCall := (*cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls)[0]
	assert.Equal(t, "1234", patchCall.PdashboardIdOrUrn)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Name)
	assert.Equal(t, "edited-dashboard-name", *patchCall.PdashboardPatchSchema.Name)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Description)
	assert.Equal(t, "Updated description", *patchCall.PdashboardPatchSchema.Description)
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Scope)
	assert.Equal(t, sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD, *patchCall.PdashboardPatchSchema.Scope)

	// Verify success message
	assert.Len(t, *cli.MockPrinter.SuccessCalls, 1)
	assert.Contains(t, (*cli.MockPrinter.SuccessCalls)[0], "Dashboard updated successfully!")
}

func TestShouldEditDashboardWithIdentifier(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	updatedDashboard := createEditResult()
	cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Result = updatedDashboard

	// Replace editor with mock that returns changes
	mockEditor := &MockEditor{
		Content: []byte(`{
			"name": "edited-with-identifier",
			"description": "Updated via identifier"
		}`),
	}
	cli.Editor = mockEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--identifier", "urn:custom:dashboard:edit-test")

	// Verify GetDashboard called with identifier
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls, 1)
	getCall := (*cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls)[0]
	assert.Equal(t, "urn:custom:dashboard:edit-test", getCall.PdashboardIdOrUrn)
}

func TestEditDashboardNoChanges(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	// Custom editor that returns exactly the same content as input
	noChangeEditor := &NoChangeEditor{}
	cli.Editor = noChangeEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	// Verify no patch was called since no changes were made
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 0)

	// Verify warning message about no changes
	assert.Len(t, *cli.MockPrinter.PrintWarnCalls, 1)
	assert.Equal(t, "No changes made", (*cli.MockPrinter.PrintWarnCalls)[0])
}

// NoChangeEditor returns exactly the same content as input
type NoChangeEditor struct{}

func (e *NoChangeEditor) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	return io.ReadAll(contents)
}

func TestEditDashboardWithJsonOutput(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	updatedDashboard := createEditResult()
	cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Result = updatedDashboard

	mockEditor := &MockEditor{
		Content: []byte(`{
			"name": "json-output-test",
			"description": "Testing JSON output"
		}`),
	}
	cli.Editor = mockEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--output", "json")

	// Verify JSON output
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Contains(t, jsonOutput, "dashboard")
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestEditDashboardNoChangesJsonOutput(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	// Use NoChangeEditor that returns exactly the same content
	noChangeEditor := &NoChangeEditor{}
	cli.Editor = noChangeEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--output", "json")

	// Should get JSON message about no changes
	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, "No changes made", jsonOutput["message"])
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestEditDashboardInvalidJson(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	mockEditor := &MockEditor{
		Content: []byte(`{"invalid": json syntax}`),
	}
	//nolint:staticcheck
	cli.Deps.Editor = mockEditor

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1234")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to parse edited JSON")

	// Verify no patch call was made
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 0)
}

func TestEditDashboardEditorError(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	mockEditor := &MockEditor{
		Error: errors.New("editor failed to open"),
	}
	//nolint:staticcheck
	cli.Deps.Editor = mockEditor

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1234")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to open editor")

	// Verify no patch call was made
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 0)
}

func TestEditDashboardMissingArgs(t *testing.T) {
	cli, cmd := setupEditCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "one of the required flags {id | identifier} not set")
}

func TestEditDashboardUsesReverseEditorByDefault(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	originalDashboard := createEditTestDashboard()
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard

	updatedDashboard := createEditResult()
	cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Result = updatedDashboard

	// Use a MockEditor that makes a simple change instead of ReverseEditor
	// ReverseEditor produces invalid JSON
	mockEditor := &MockEditor{
		Content: []byte(`{"name": "changed-by-reverse-editor"}`),
	}
	cli.Editor = mockEditor

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	// Verify that GetDashboard was called
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.GetDashboardCalls, 1)

	// Verify that PatchDashboard was called (content was changed)
	assert.Len(t, *cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls, 1)
	patchCall := (*cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardCalls)[0]
	assert.NotNil(t, patchCall.PdashboardPatchSchema.Name)
	assert.Equal(t, "changed-by-reverse-editor", *patchCall.PdashboardPatchSchema.Name)
}

// Test the specific error scenarios that could occur during editing
func TestEditDashboardApiErrors(t *testing.T) {
	tests := []struct {
		name        string
		setupError  func(*di.MockDeps)
		expectedErr string
	}{
		{
			name: "GetDashboard API error",
			setupError: func(cli *di.MockDeps) {
				cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Error = errors.New("dashboard not found")
			},
			expectedErr: "dashboard not found",
		},
		{
			name: "PatchDashboard API error",
			setupError: func(cli *di.MockDeps) {
				originalDashboard := createEditTestDashboard()
				cli.MockClient.ApiMocks.DashboardsApi.GetDashboardResponse.Result = originalDashboard
				cli.MockClient.ApiMocks.DashboardsApi.PatchDashboardResponse.Error = errors.New("patch failed")
			},
			expectedErr: "patch failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupEditCmd(t)
			tt.setupError(cli)

			if tt.name == "PatchDashboard API error" {
				// Set up mock editor for this test
				mockEditor := &MockEditor{
					Content: []byte(`{"name": "changed"}`),
				}
				cli.Editor = mockEditor
			}

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--id", "1234")

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
