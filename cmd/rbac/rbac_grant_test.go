package rbac

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	SomePermission = "acess-view"
	SomeResource   = "view-x"
)

func TestGrantPermissions(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := GrantPermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "--permission", SomePermission)

	grantCalls := *cli.MockClient.ApiMocks.PermissionsApi.GrantPermissionsCalls
	assert.Len(t, grantCalls, 1)
	assert.Equal(t, SomeSubject, grantCalls[0].Psubject)

	expectedGrant := stackstate_api.NewGrantPermission(SomePermission, "system")
	assert.Equal(t, expectedGrant, grantCalls[0].PgrantPermission)

	describeCalls := *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls
	assert.Len(t, describeCalls, 1)
	assert.Equal(t, SomeSubject, describeCalls[0].Psubject)
	assert.Equal(t, expectedGrant.Permission, *describeCalls[0].Ppermission)
	assert.Equal(t, expectedGrant.ResourceName, *describeCalls[0].Presource)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "--permission", SomePermission, "--resource", SomeResource)

	grantCalls = *cli.MockClient.ApiMocks.PermissionsApi.GrantPermissionsCalls
	assert.Len(t, grantCalls, 2)
	assert.Equal(t, SomeSubject, grantCalls[1].Psubject)

	expectedGrant = stackstate_api.NewGrantPermission(SomePermission, SomeResource)
	assert.Equal(t, expectedGrant, grantCalls[1].PgrantPermission)

	describeCalls = *cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsCalls
	assert.Len(t, describeCalls, 2)
	assert.Equal(t, SomeSubject, describeCalls[1].Psubject)
	assert.Equal(t, expectedGrant.Permission, *describeCalls[1].Ppermission)
	assert.Equal(t, expectedGrant.ResourceName, *describeCalls[1].Presource)

	expectedResult := []printer.TableData{
		ExpectedTable,
		ExpectedTable,
	}

	assert.Equal(t, expectedResult, *cli.MockPrinter.TableCalls)
}

func TestGrantPermissionsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := GrantPermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.DescribePermissionsResponse.Result = *Description

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectHandle, "--permission", SomePermission, "-o", "json")

	expected := []map[string]interface{}{
		ExpectedJson,
	}
	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
