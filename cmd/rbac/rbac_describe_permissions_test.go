package rbac

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
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
	ExpectedTable = printer.TableData{
		Header: []string{"permission", "resource"},
		Data: [][]interface{}{
			{"bar", Resource1},
			{"baz", Resource1},
			{"foo", Resource1},
			{"fai", Resource2},
			{"foo", Resource2},
			{"fum", Resource2},
		},
		MissingTableDataMsg: printer.NotFoundMsg{Types: "matching permissions"},
	}
	ExpectedJson = map[string]interface{}{
		"subject":     SubjectHandle,
		"permissions": AllPermissions,
	}
)

func TestPermissionsDescribeJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls, 1)

	calls := *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls
	assert.Equal(t, SubjectHandle, calls[0].Psubject)

	expected := []map[string]interface{}{
		ExpectedJson,
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestPermissionsDescribeTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle)

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls, 1)

	expected := []printer.TableData{
		ExpectedTable,
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestPermissionsDescribeFilterResource(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource1)

	calls := *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SubjectHandle, calls[0].Psubject)
	assert.Nil(t, calls[0].Ppermission)
	assert.Equal(t, Resource1, *calls[0].Presource)
}

func TestPermissionsDescribeFilterPermission(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--permission", "foo")

	calls := *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SubjectHandle, calls[0].Psubject)
	assert.Equal(t, "foo", *calls[0].Ppermission)
	assert.Nil(t, calls[0].Presource)
}

func TestPermissionsDescribeFilterResourceAndPermission(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--resource", Resource1, "--permission", "foo")

	calls := *cli.MockClient.ApiMocks.PermissionsAPI.DescribePermissionsCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SubjectHandle, calls[0].Psubject)
	assert.Equal(t, "foo", *calls[0].Ppermission)
	assert.Equal(t, Resource1, *calls[0].Presource)
}
