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

type LayoutApi interface {

	/*
		GetAllLayouts List layout hints

		List all available layout hints in the system

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetAllLayoutsRequest
	*/
	GetAllLayouts(ctx context.Context) ApiGetAllLayoutsRequest

	// GetAllLayoutsExecute executes the request
	//  @return LayoutList
	GetAllLayoutsExecute(r ApiGetAllLayoutsRequest) (*LayoutList, *http.Response, error)
}

// LayoutApiService LayoutApi service
type LayoutApiService service

type ApiGetAllLayoutsRequest struct {
	ctx        context.Context
	ApiService LayoutApi
}

func (r ApiGetAllLayoutsRequest) Execute() (*LayoutList, *http.Response, error) {
	return r.ApiService.GetAllLayoutsExecute(r)
}

/*
GetAllLayouts List layout hints

List all available layout hints in the system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllLayoutsRequest
*/
func (a *LayoutApiService) GetAllLayouts(ctx context.Context) ApiGetAllLayoutsRequest {
	return ApiGetAllLayoutsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return LayoutList
func (a *LayoutApiService) GetAllLayoutsExecute(r ApiGetAllLayoutsRequest) (*LayoutList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LayoutList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LayoutApiService.GetAllLayouts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/layouts"

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v LayoutApiError
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

type LayoutApiMock struct {
	GetAllLayoutsCalls    *[]GetAllLayoutsCall
	GetAllLayoutsResponse GetAllLayoutsMockResponse
}

func NewLayoutApiMock() LayoutApiMock {
	xGetAllLayoutsCalls := make([]GetAllLayoutsCall, 0)
	return LayoutApiMock{
		GetAllLayoutsCalls: &xGetAllLayoutsCalls,
	}
}

type GetAllLayoutsMockResponse struct {
	Result   LayoutList
	Response *http.Response
	Error    error
}

type GetAllLayoutsCall struct {
}

func (mock LayoutApiMock) GetAllLayouts(ctx context.Context) ApiGetAllLayoutsRequest {
	return ApiGetAllLayoutsRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock LayoutApiMock) GetAllLayoutsExecute(r ApiGetAllLayoutsRequest) (*LayoutList, *http.Response, error) {
	p := GetAllLayoutsCall{}
	*mock.GetAllLayoutsCalls = append(*mock.GetAllLayoutsCalls, p)
	return &mock.GetAllLayoutsResponse.Result, mock.GetAllLayoutsResponse.Response, mock.GetAllLayoutsResponse.Error
}
