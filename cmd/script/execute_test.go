package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func TestExecute(t *testing.T) {
	client := sts.APIClient{}
	client.ScriptingApi = MockScriptingApiService{}
	assert.Equal(t, true, false)
}
