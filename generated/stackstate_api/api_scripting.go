/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type ScriptingApi interface {

	/*
	ScriptExecute Execute script

	Execute a StackState Scripting Language or Template Language script with arbitrary arguments.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiScriptExecuteRequest
	*/
	ScriptExecute(ctx context.Context) ApiScriptExecuteRequest

	// ScriptExecuteExecute executes the request
	//  @return ExecuteScriptResponse
	ScriptExecuteExecute(r ApiScriptExecuteRequest) (*ExecuteScriptResponse, *http.Response, error)
}

// ScriptingApiService ScriptingApi service
type ScriptingApiService service

type ApiScriptExecuteRequest struct {
	ctx context.Context
	ApiService ScriptingApi
	executeScriptRequest *ExecuteScriptRequest
}

func (r ApiScriptExecuteRequest) ExecuteScriptRequest(executeScriptRequest ExecuteScriptRequest) ApiScriptExecuteRequest {
	r.executeScriptRequest = &executeScriptRequest
	return r
}

func (r ApiScriptExecuteRequest) Execute() (*ExecuteScriptResponse, *http.Response, error) {
	return r.ApiService.ScriptExecuteExecute(r)
}

/*
ScriptExecute Execute script

Execute a StackState Scripting Language or Template Language script with arbitrary arguments.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScriptExecuteRequest
*/
func (a *ScriptingApiService) ScriptExecute(ctx context.Context) ApiScriptExecuteRequest {
	return ApiScriptExecuteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExecuteScriptResponse
func (a *ScriptingApiService) ScriptExecuteExecute(r ApiScriptExecuteRequest) (*ExecuteScriptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExecuteScriptResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScriptingApiService.ScriptExecute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/script/execute"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.executeScriptRequest == nil {
		return localVarReturnValue, nil, reportError("executeScriptRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.executeScriptRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ServiceBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Service-Bearer"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ServiceToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExecuteScriptError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 408 {
			var v ExecuteScriptTimeoutError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}


// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------


type ScriptingApiMock struct {
	ScriptExecuteCalls *[]ScriptExecuteCall
	ScriptExecuteResponse ScriptExecuteMockResponse
}	

func NewScriptingApiMock() ScriptingApiMock {
	xScriptExecuteCalls := make([]ScriptExecuteCall, 0)
	return ScriptingApiMock {
		ScriptExecuteCalls: &xScriptExecuteCalls,
	}
}

type ScriptExecuteMockResponse struct {
	Result ExecuteScriptResponse
	Response *http.Response
	Error error
}

type ScriptExecuteCall struct {
	PexecuteScriptRequest *ExecuteScriptRequest
}


func (mock ScriptingApiMock) ScriptExecute(ctx context.Context) ApiScriptExecuteRequest {
	return ApiScriptExecuteRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock ScriptingApiMock) ScriptExecuteExecute(r ApiScriptExecuteRequest) (*ExecuteScriptResponse, *http.Response, error) {
	p := ScriptExecuteCall {
			PexecuteScriptRequest: r.executeScriptRequest,
	}
	*mock.ScriptExecuteCalls = append(*mock.ScriptExecuteCalls, p)
	return &mock.ScriptExecuteResponse.Result, mock.ScriptExecuteResponse.Response, mock.ScriptExecuteResponse.Error
}


