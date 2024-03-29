/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
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

type ApiTokenApi interface {

	/*
		GetCurrentUserApiTokens Get current user's API tokens

		Get all API token of the logged-in user.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetCurrentUserApiTokensRequest
	*/
	GetCurrentUserApiTokens(ctx context.Context) ApiGetCurrentUserApiTokensRequest

	// GetCurrentUserApiTokensExecute executes the request
	//  @return []ApiToken
	GetCurrentUserApiTokensExecute(r ApiGetCurrentUserApiTokensRequest) ([]ApiToken, *http.Response, error)
}

// ApiTokenApiService ApiTokenApi service
type ApiTokenApiService service

type ApiGetCurrentUserApiTokensRequest struct {
	ctx        context.Context
	ApiService ApiTokenApi
}

func (r ApiGetCurrentUserApiTokensRequest) Execute() ([]ApiToken, *http.Response, error) {
	return r.ApiService.GetCurrentUserApiTokensExecute(r)
}

/*
GetCurrentUserApiTokens Get current user's API tokens

Get all API token of the logged-in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCurrentUserApiTokensRequest
*/
func (a *ApiTokenApiService) GetCurrentUserApiTokens(ctx context.Context) ApiGetCurrentUserApiTokensRequest {
	return ApiGetCurrentUserApiTokensRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []ApiToken
func (a *ApiTokenApiService) GetCurrentUserApiTokensExecute(r ApiGetCurrentUserApiTokensRequest) ([]ApiToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ApiToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiTokenApiService.GetCurrentUserApiTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/profile/tokens"

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
				localVarHeaderParams["X-API-ServiceBearer"] = key
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
			var v UserNotFoundError
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

type ApiTokenApiMock struct {
	GetCurrentUserApiTokensCalls    *[]GetCurrentUserApiTokensCall
	GetCurrentUserApiTokensResponse GetCurrentUserApiTokensMockResponse
}

func NewApiTokenApiMock() ApiTokenApiMock {
	xGetCurrentUserApiTokensCalls := make([]GetCurrentUserApiTokensCall, 0)
	return ApiTokenApiMock{
		GetCurrentUserApiTokensCalls: &xGetCurrentUserApiTokensCalls,
	}
}

type GetCurrentUserApiTokensMockResponse struct {
	Result   []ApiToken
	Response *http.Response
	Error    error
}

type GetCurrentUserApiTokensCall struct {
}

func (mock ApiTokenApiMock) GetCurrentUserApiTokens(ctx context.Context) ApiGetCurrentUserApiTokensRequest {
	return ApiGetCurrentUserApiTokensRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock ApiTokenApiMock) GetCurrentUserApiTokensExecute(r ApiGetCurrentUserApiTokensRequest) ([]ApiToken, *http.Response, error) {
	p := GetCurrentUserApiTokensCall{}
	*mock.GetCurrentUserApiTokensCalls = append(*mock.GetCurrentUserApiTokensCalls, p)
	return mock.GetCurrentUserApiTokensResponse.Result, mock.GetCurrentUserApiTokensResponse.Response, mock.GetCurrentUserApiTokensResponse.Error
}
