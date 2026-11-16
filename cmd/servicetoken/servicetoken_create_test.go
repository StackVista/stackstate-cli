package servicetoken

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestServiceTokenCreate(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenResponse.Result = stackstate_api.ServiceTokenCreatedResponse{
		Name:       "test-token",
		Token:      "test-token-token",
		Expiration: int64p(100000000),
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--roles", "test-role", "--expiration", "2020-05-22")

	checkCreateCall(t, cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenCalls, "test-token", []string{"test-role"}, int64p(1590105600000))
	assert.Equal(t, []string{"Service token created: \x1b[37mtest-token-token\x1b[0m\n"}, *cli.MockPrinter.SuccessCalls)
}

func TestServiceTokenCreateExpirationOptional(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenResponse.Result = stackstate_api.ServiceTokenCreatedResponse{
		Name:       "test-token",
		Token:      "test-token-token",
		Expiration: int64p(100000000),
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--roles", "test-role")

	checkCreateCall(t, cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenCalls, "test-token", []string{"test-role"}, nil)
	assert.Equal(t, []string{"Service token created: \x1b[37mtest-token-token\x1b[0m\n"}, *cli.MockPrinter.SuccessCalls)
}

func TestServiceTokenCreateMultipleRoles(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenResponse.Result = stackstate_api.ServiceTokenCreatedResponse{
		Name:       "test-token",
		Token:      "test-token-token",
		Expiration: int64p(100000000),
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--roles", "test-role,another-role", "--roles", "third-role")

	checkCreateCall(t, cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenCalls, "test-token", []string{"test-role", "another-role", "third-role"}, nil)
	assert.Equal(t, []string{"Service token created: \x1b[37mtest-token-token\x1b[0m\n"}, *cli.MockPrinter.SuccessCalls)
}

func TestServiceTokenCreateJSON(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	r := &stackstate_api.ServiceTokenCreatedResponse{
		Name:  "test-token",
		Token: "test-token-token",
	}

	cli.MockClient.ApiMocks.ServiceTokenApi.CreateNewServiceTokenResponse.Result = *r

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--roles", "test-role", "-o", "json")

	assert.Equal(t,
		[]map[string]interface{}{{
			"service-token": r,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func int64p(i int64) *int64 {
	return &i
}

func checkCreateCall(t *testing.T, calls *[]stackstate_api.CreateNewServiceTokenCall, name string, roles []string, expiration *int64) {
	assert.Len(t, *calls, 1)

	call := (*calls)[0]
	assert.Equal(t, name, call.PnewServiceTokenRequest.Name)
	assert.Len(t, call.PnewServiceTokenRequest.Roles, len(roles))

	for _, role := range roles {
		assert.Contains(t, call.PnewServiceTokenRequest.Roles, role)
	}

	if expiration != nil {
		assert.Equal(t, *expiration, *call.PnewServiceTokenRequest.ExpiryDate)
	} else {
		assert.Nil(t, call.PnewServiceTokenRequest.ExpiryDate)
	}
}
