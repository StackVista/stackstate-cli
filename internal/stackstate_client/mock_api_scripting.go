package stackstate_client

import (
	"context"
	"net/http"
)

type MockScriptingApiService struct {
	ReturnFromScriptExecuteExecute ScriptExecuteExecuteReturnValues
	ExecuteScriptRequests          *[]ExecuteScriptRequest
}

func NewMockScriptingApiService() MockScriptingApiService {
	reqs := make([]ExecuteScriptRequest, 0)
	return MockScriptingApiService{
		ReturnFromScriptExecuteExecute: ScriptExecuteExecuteReturnValues{
			ScriptResponse: map[string]interface{}{"result": "hello test"},
			HttpResponse:   &http.Response{StatusCode: 200},
			Error:          nil,
		},
		ExecuteScriptRequests: &reqs,
	}
}

type ScriptExecuteExecuteReturnValues struct {
	ScriptResponse map[string]interface{}
	HttpResponse   *http.Response
	Error          error
}

func (mock MockScriptingApiService) ScriptExecute(ctx context.Context) ApiScriptExecuteRequest {
	return ApiScriptExecuteRequest{
		ApiService: mock,
	}
}

func (mock MockScriptingApiService) ScriptExecuteExecute(r ApiScriptExecuteRequest) (map[string]interface{}, *http.Response, error) {
	*mock.ExecuteScriptRequests = append(*mock.ExecuteScriptRequests, *r.executeScriptRequest)
	return mock.ReturnFromScriptExecuteExecute.ScriptResponse, mock.ReturnFromScriptExecuteExecute.HttpResponse, mock.ReturnFromScriptExecuteExecute.Error
}
