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

type UserAuthorizationApi interface {

	/*
		GetUserAuthorizationFor Is the current user authorized for the provided permission and resource

		Is the current user authorized for the provided permission and resource

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetUserAuthorizationForRequest
	*/
	GetUserAuthorizationFor(ctx context.Context) ApiGetUserAuthorizationForRequest

	// GetUserAuthorizationForExecute executes the request
	GetUserAuthorizationForExecute(r ApiGetUserAuthorizationForRequest) (*http.Response, error)
}

// UserAuthorizationApiService UserAuthorizationApi service
type UserAuthorizationApiService service

type ApiGetUserAuthorizationForRequest struct {
	ctx          context.Context
	ApiService   UserAuthorizationApi
	permission   *string
	resourceName *string
}

func (r ApiGetUserAuthorizationForRequest) Permission(permission string) ApiGetUserAuthorizationForRequest {
	r.permission = &permission
	return r
}

func (r ApiGetUserAuthorizationForRequest) ResourceName(resourceName string) ApiGetUserAuthorizationForRequest {
	r.resourceName = &resourceName
	return r
}

func (r ApiGetUserAuthorizationForRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetUserAuthorizationForExecute(r)
}

/*
GetUserAuthorizationFor Is the current user authorized for the provided permission and resource

Is the current user authorized for the provided permission and resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserAuthorizationForRequest
*/
func (a *UserAuthorizationApiService) GetUserAuthorizationFor(ctx context.Context) ApiGetUserAuthorizationForRequest {
	return ApiGetUserAuthorizationForRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserAuthorizationApiService) GetUserAuthorizationForExecute(r ApiGetUserAuthorizationForRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthorizationApiService.GetUserAuthorizationFor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/authorization/for"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permission == nil {
		return nil, reportError("permission is required and must be specified")
	}

	localVarQueryParams.Add("permission", parameterToString(*r.permission, ""))
	if r.resourceName != nil {
		localVarQueryParams.Add("resourceName", parameterToString(*r.resourceName, ""))
	}
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------

type UserAuthorizationApiMock struct {
	GetUserAuthorizationForCalls    *[]GetUserAuthorizationForCall
	GetUserAuthorizationForResponse GetUserAuthorizationForMockResponse
}

func NewUserAuthorizationApiMock() UserAuthorizationApiMock {
	xGetUserAuthorizationForCalls := make([]GetUserAuthorizationForCall, 0)
	return UserAuthorizationApiMock{
		GetUserAuthorizationForCalls: &xGetUserAuthorizationForCalls,
	}
}

type GetUserAuthorizationForMockResponse struct {
	Response *http.Response
	Error    error
}

type GetUserAuthorizationForCall struct {
	Ppermission   *string
	PresourceName *string
}

func (mock UserAuthorizationApiMock) GetUserAuthorizationFor(ctx context.Context) ApiGetUserAuthorizationForRequest {
	return ApiGetUserAuthorizationForRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock UserAuthorizationApiMock) GetUserAuthorizationForExecute(r ApiGetUserAuthorizationForRequest) (*http.Response, error) {
	p := GetUserAuthorizationForCall{
		Ppermission:   r.permission,
		PresourceName: r.resourceName,
	}
	*mock.GetUserAuthorizationForCalls = append(*mock.GetUserAuthorizationForCalls, p)
	return mock.GetUserAuthorizationForResponse.Response, mock.GetUserAuthorizationForResponse.Error
}
