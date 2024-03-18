package ingestionapikey

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestIngestApiKeyGenerate(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyResponse.Result = stackstate_api.GeneratedIngestionApiKeyResponse{
		Name:        "test-token",
		ApiKey:      "test-token-key",
		Expiration:  int64p(1590105600000),
		Description: stringp("test-token-description"),
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--description", "test-token-description", "--expiration", "2020-05-22")

	checkCreateCall(t, cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyCalls, "test-token", stringp("test-token-description"), int64p(1590105600000))
	assert.Equal(t, []string{"Ingestion API Key generated: \x1b[37mtest-token-key\x1b[0m\n"}, *cli.MockPrinter.SuccessCalls)
}

func TestIngestApiKeyGenerateOnlyRequriedFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyResponse.Result = stackstate_api.GeneratedIngestionApiKeyResponse{
		Name:   "test-token2",
		ApiKey: "test-token2-key",
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token2")

	checkCreateCall(t, cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyCalls, "test-token2", nil, nil)
	assert.Equal(t, []string{"Ingestion API Key generated: \x1b[37mtest-token2-key\x1b[0m\n"}, *cli.MockPrinter.SuccessCalls)
}

func TestIngestApiKeyGenerateJSON(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateCommand(&cli.Deps)

	r := &stackstate_api.GeneratedIngestionApiKeyResponse{
		Name:        "test-token",
		ApiKey:      "test-token-key",
		Expiration:  int64p(1590105600000),
		Description: stringp("test-token-description"),
	}

	cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyResponse.Result = *r

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--name", "test-token", "--description", "test-token-description", "--expiration", "2020-05-22", "-o", "json")

	checkCreateCall(t, cli.MockClient.ApiMocks.IngestionApiKeyApi.GenerateIngestionApiKeyCalls, "test-token", stringp("test-token-description"), int64p(1590105600000))
	assert.Equal(t,
		[]map[string]interface{}{{
			"ingestion-api-key": r,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func int64p(i int64) *int64 {
	return &i
}

func stringp(i string) *string {
	return &i
}

func checkCreateCall(t *testing.T, calls *[]stackstate_api.GenerateIngestionApiKeyCall, name string, description *string, expiration *int64) {
	assert.Len(t, *calls, 1)

	call := (*calls)[0]
	assert.Equal(t, name, call.PgenerateIngestionApiKeyRequest.Name)

	if description != nil {
		assert.Equal(t, *description, *call.PgenerateIngestionApiKeyRequest.Description)
	} else {
		assert.Nil(t, call.PgenerateIngestionApiKeyRequest.Description)
	}

	if expiration != nil {
		assert.Equal(t, *expiration, *call.PgenerateIngestionApiKeyRequest.Expiration)
	} else {
		assert.Nil(t, call.PgenerateIngestionApiKeyRequest.Expiration)
	}
}
