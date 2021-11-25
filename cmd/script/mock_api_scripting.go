package script

import (
	"context"
	"net/http"

	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

type MockScriptingApiService struct {
	ReturnFromScriptExecuteExecute ScriptExecuteExecuteReturnValues
}

func NewMockScriptingApiService() MockScriptingApiService {
	resp := http.Response{
		StatusCode: 200,
	}
	scriptResponse := make(map[string]interface{})
	scriptResponse["result"] = "hello test"

	return MockScriptingApiService{
		ReturnFromScriptExecuteExecute: ScriptExecuteExecuteReturnValues{
			ScriptResponse: scriptResponse,
			HttpResponse:   &resp,
			Error:          nil,
		},
	}
}

type ScriptExecuteExecuteReturnValues struct {
	ScriptResponse map[string]interface{}
	HttpResponse   *http.Response
	Error          error
}

func (mock MockScriptingApiService) ScriptExecute(ctx context.Context) sts.ApiScriptExecuteRequest {
	return sts.ApiScriptExecuteRequest{
		ApiService: mock,
	}
}

func (mock MockScriptingApiService) ScriptExecuteExecute(r sts.ApiScriptExecuteRequest) (map[string]interface{}, *http.Response, error) {
	return mock.ReturnFromScriptExecuteExecute.ScriptResponse, mock.ReturnFromScriptExecuteExecute.HttpResponse, mock.ReturnFromScriptExecuteExecute.Error
}
