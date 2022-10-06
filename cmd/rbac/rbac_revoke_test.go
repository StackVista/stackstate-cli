package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func TestRevokePermissions(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RevokePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "--permission", SomePermission)

	revokeCalls := *cli.MockClient.ApiMocks.PermissionsApi.RevokePermissionsCalls
	assert.Len(t, revokeCalls, 1)
	assert.Equal(t, SomeSubject, revokeCalls[0].Psubject)
	assert.Equal(t, SomePermission, *revokeCalls[0].Ppermission)
	assert.Equal(t, "system", *revokeCalls[0].Presource)

	describeCalls := *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls
	assert.Len(t, describeCalls, 1)
	assert.Equal(t, SomeSubject, describeCalls[0].Psubject)
	assert.Equal(t, SomePermission, *describeCalls[0].Ppermission)
	assert.Equal(t, "system", *describeCalls[0].Presource)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "--permission", SomePermission, "--resource", SomeResource)

	revokeCalls = *cli.MockClient.ApiMocks.PermissionsApi.RevokePermissionsCalls
	assert.Len(t, revokeCalls, 2)
	assert.Equal(t, SomeSubject, revokeCalls[1].Psubject)
	assert.Equal(t, SomePermission, *revokeCalls[1].Ppermission)
	assert.Equal(t, SomeResource, *revokeCalls[1].Presource)

	describeCalls = *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls
	assert.Len(t, describeCalls, 2)
	assert.Equal(t, SomeSubject, describeCalls[1].Psubject)
	assert.Equal(t, SomePermission, *describeCalls[1].Ppermission)
	assert.Equal(t, SomeResource, *describeCalls[1].Presource)

	expectedResult := []printer.TableData{
		ExpectedTable,
		ExpectedTable,
	}

	assert.Equal(t, expectedResult, *cli.MockPrinter.TableCalls)
}

func TestRevokePermissionsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RevokePermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--permission", SomePermission, "-o", "json")

	expected := []map[string]interface{}{
		ExpectedJson,
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
