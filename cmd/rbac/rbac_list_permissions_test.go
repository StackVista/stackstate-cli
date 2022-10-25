package rbac

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var SomePermissions = []string{
	"foo",
	"bar",
	"baz",
}

func TestPermissionsList(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListPermissionsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.PermissionsApi.GetPermissionsResponse.Result = stackstate_api.Permissions{
		Permissions: SomePermissions,
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.GetPermissionsCalls, 1)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.PermissionsApi.GetPermissionsCalls, 2)

	expected := []map[string]interface{}{
		{
			"permissions": SomePermissions,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
