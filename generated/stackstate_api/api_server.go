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


type ServerApi interface {

	/*
	ServerInfo Get server info

	Get information of the StackState information, such as version, deployment mode, etc.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiServerInfoRequest
	*/
	ServerInfo(ctx context.Context) ApiServerInfoRequest

	// ServerInfoExecute executes the request
	//  @return ServerInfo
	ServerInfoExecute(r ApiServerInfoRequest) (*ServerInfo, *http.Response, error)
}

// ServerApiService ServerApi service
type ServerApiService service

type ApiServerInfoRequest struct {
	ctx context.Context
	ApiService ServerApi
}

func (r ApiServerInfoRequest) Execute() (*ServerInfo, *http.Response, error) {
	return r.ApiService.ServerInfoExecute(r)
}

/*
ServerInfo Get server info

Get information of the StackState information, such as version, deployment mode, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerInfoRequest
*/
func (a *ServerApiService) ServerInfo(ctx context.Context) ApiServerInfoRequest {
	return ApiServerInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServerInfo
func (a *ServerApiService) ServerInfoExecute(r ApiServerInfoRequest) (*ServerInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServerInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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


type ServerApiMock struct {
	ServerInfoCalls *[]ServerInfoCall
	ServerInfoResponse ServerInfoMockResponse
}	

func NewServerApiMock() ServerApiMock {
	xServerInfoCalls := make([]ServerInfoCall, 0)
	return ServerApiMock {
		ServerInfoCalls: &xServerInfoCalls,
	}
}

type ServerInfoMockResponse struct {
	Result ServerInfo
	Response *http.Response
	Error error
}

type ServerInfoCall struct {
}


func (mock ServerApiMock) ServerInfo(ctx context.Context) ApiServerInfoRequest {
	return ApiServerInfoRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock ServerApiMock) ServerInfoExecute(r ApiServerInfoRequest) (*ServerInfo, *http.Response, error) {
	p := ServerInfoCall {
	}
	*mock.ServerInfoCalls = append(*mock.ServerInfoCalls, p)
	return &mock.ServerInfoResponse.Result, mock.ServerInfoResponse.Response, mock.ServerInfoResponse.Error
}


