package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	Set1         = []string{"foo", "bar", "baz"}
	Resource1    = "some-view"
	Permissions1 = map[string][]string{
		Resource1: Set1,
	}

	Set2         = []string{"foo", "fai", "fum"}
	Resource2    = "system"
	Permissions2 = map[string][]string{
		Resource2: Set2,
	}

	AllPermissions = map[string][]string{
		Resource1: Set1,
		Resource2: Set2,
	}
	SubjectHandle = "some-handle"
	Description   = stackstate_api.NewPermissionDescription(SubjectHandle, AllPermissions)
)

func TestPermissionsDescribeJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	calls := *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls
	assert.Equal(t, SubjectHandle, calls[0].Psubject)

	expected := []map[string]interface{}{
		{
			"subject":     SubjectHandle,
			"permissions": (Permissions)(AllPermissions),
		},
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestPermissionsDescribeTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle)

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"permission", "resource"},
			Data: [][]interface{}{
				{"foo", Resource1},
				{"bar", Resource1},
				{"baz", Resource1},
				{"foo", Resource2},
				{"fai", Resource2},
				{"fum", Resource2},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestPermissionsDescribeFilterResource(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource1)

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"permission", "resource"},
			Data: [][]interface{}{
				{"foo", Resource1},
				{"bar", Resource1},
				{"baz", Resource1},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestPermissionsDescribeFilterResourceJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource2, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []map[string]interface{}{
		{
			"subject":     SubjectHandle,
			"permissions": (Permissions)(Permissions2),
		},
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestPermissionsDescribeFilterPermission(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--permission", "foo")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"permission", "resource"},
			Data: [][]interface{}{
				{"foo", Resource1},
				{"foo", Resource2},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestPermissionsDescribeFilterPermissionJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--permission", "bar", "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []map[string]interface{}{
		{
			"subject": SubjectHandle,
			"permissions": (Permissions)(map[string][]string{
				Resource1: {"bar"},
			}),
		},
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestPermissionsDescribeFilterResourceAndPermission(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource1, "--permission", "foo")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"permission", "resource"},
			Data: [][]interface{}{
				{"foo", Resource1},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestPermissionsDescribeFilterNothingRemains(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource2, "--permission", "bar")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		{
			Header:              []string{"permission", "resource"},
			Data:                [][]interface{}{},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}
