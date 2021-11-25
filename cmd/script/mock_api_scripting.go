package script

import (
	"context"
	"net/http"

	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

type MockScriptingApiService struct{}

func (mock MockScriptingApiService) ScriptExecute(ctx context.Context) sts.ApiScriptExecuteRequest {
	return sts.ApiScriptExecuteRequest{
		ApiService: mock,
	}
}

func (mock MockScriptingApiService) ScriptExecuteExecute(r sts.ApiScriptExecuteRequest) (map[string]interface{}, *http.Response, error) {
	resp := http.Response{}
	scriptResponse := make(map[string]interface{})
	scriptResponse["result"] = "hello test"
	return scriptResponse, &resp, nil
}
