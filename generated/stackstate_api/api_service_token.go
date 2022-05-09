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
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ServiceTokenApi interface {

	/*
	CreateNewServiceToken Create new service token

	Create new service token.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiCreateNewServiceTokenRequest
	*/
	CreateNewServiceToken(ctx _context.Context) ApiCreateNewServiceTokenRequest

	// CreateNewServiceTokenExecute executes the request
	//  @return ServiceTokenCreatedResponse
	CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (ServiceTokenCreatedResponse, *_nethttp.Response, error)

	/*
	DeleteServiceToken Delete service token

	Delete service token.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param serviceTokenId The identifier of a service token
	 @return ApiDeleteServiceTokenRequest
	*/
	DeleteServiceToken(ctx _context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest

	// DeleteServiceTokenExecute executes the request
	DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*_nethttp.Response, error)

	/*
	GetServiceTokens Get service tokens

	Get service tokens.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetServiceTokensRequest
	*/
	GetServiceTokens(ctx _context.Context) ApiGetServiceTokensRequest

	// GetServiceTokensExecute executes the request
	//  @return []ServiceToken
	GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *_nethttp.Response, error)
}


// ServiceTokenApiService ServiceTokenApi service
type ServiceTokenApiService service

type ApiCreateNewServiceTokenRequest struct {
	ctx _context.Context
	ApiService ServiceTokenApi
	newServiceTokenRequest *NewServiceTokenRequest
}

func (r ApiCreateNewServiceTokenRequest) NewServiceTokenRequest(newServiceTokenRequest NewServiceTokenRequest) ApiCreateNewServiceTokenRequest {
	r.newServiceTokenRequest = &newServiceTokenRequest
	return r
}

func (r ApiCreateNewServiceTokenRequest) Execute() (ServiceTokenCreatedResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateNewServiceTokenExecute(r)
}

/*
CreateNewServiceToken Create new service token

Create new service token.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNewServiceTokenRequest
*/
func (a *ServiceTokenApiService) CreateNewServiceToken(ctx _context.Context) ApiCreateNewServiceTokenRequest {
	return ApiCreateNewServiceTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceTokenCreatedResponse
func (a *ServiceTokenApiService) CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (ServiceTokenCreatedResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceTokenCreatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.CreateNewServiceToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.newServiceTokenRequest == nil {
		return localVarReturnValue, nil, reportError("newServiceTokenRequest is required and must be specified")
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
	localVarPostBody = r.newServiceTokenRequest
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ServiceTokenCreateError
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteServiceTokenRequest struct {
	ctx _context.Context
	ApiService ServiceTokenApi
	serviceTokenId int64
}


func (r ApiDeleteServiceTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteServiceTokenExecute(r)
}

/*
DeleteServiceToken Delete service token

Delete service token.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceTokenId The identifier of a service token
 @return ApiDeleteServiceTokenRequest
*/
func (a *ServiceTokenApiService) DeleteServiceToken(ctx _context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest {
	return ApiDeleteServiceTokenRequest{
		ApiService: a,
		ctx: ctx,
		serviceTokenId: serviceTokenId,
	}
}

// Execute executes the request
func (a *ServiceTokenApiService) DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.DeleteServiceToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/token/{serviceTokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceTokenId"+"}", _neturl.PathEscape(parameterToString(r.serviceTokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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

type ApiGetServiceTokensRequest struct {
	ctx _context.Context
	ApiService ServiceTokenApi
}


func (r ApiGetServiceTokensRequest) Execute() ([]ServiceToken, *_nethttp.Response, error) {
	return r.ApiService.GetServiceTokensExecute(r)
}

/*
GetServiceTokens Get service tokens

Get service tokens.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetServiceTokensRequest
*/
func (a *ServiceTokenApiService) GetServiceTokens(ctx _context.Context) ApiGetServiceTokensRequest {
	return ApiGetServiceTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ServiceToken
func (a *ServiceTokenApiService) GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ServiceToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.GetServiceTokens")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
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


type ServiceTokenApiMock struct {
	CreateNewServiceTokenCalls *[]CreateNewServiceTokenCall
	CreateNewServiceTokenResponse CreateNewServiceTokenMockResponse
	DeleteServiceTokenCalls *[]DeleteServiceTokenCall
	DeleteServiceTokenResponse DeleteServiceTokenMockResponse
	GetServiceTokensCalls *[]GetServiceTokensCall
	GetServiceTokensResponse GetServiceTokensMockResponse
}	

func NewServiceTokenApiMock() ServiceTokenApiMock {
	xCreateNewServiceTokenCalls := make([]CreateNewServiceTokenCall, 0)
	xDeleteServiceTokenCalls := make([]DeleteServiceTokenCall, 0)
	xGetServiceTokensCalls := make([]GetServiceTokensCall, 0)
	return ServiceTokenApiMock {
		CreateNewServiceTokenCalls: &xCreateNewServiceTokenCalls,
		DeleteServiceTokenCalls: &xDeleteServiceTokenCalls,
		GetServiceTokensCalls: &xGetServiceTokensCalls,
	}
}

type CreateNewServiceTokenMockResponse struct {
	Result ServiceTokenCreatedResponse
	Response *_nethttp.Response
	Error error
}

type CreateNewServiceTokenCall struct {
	PnewServiceTokenRequest *NewServiceTokenRequest
}


func (mock ServiceTokenApiMock) CreateNewServiceToken(ctx _context.Context) ApiCreateNewServiceTokenRequest {
	return ApiCreateNewServiceTokenRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock ServiceTokenApiMock) CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (ServiceTokenCreatedResponse, *_nethttp.Response, error) {
	p := CreateNewServiceTokenCall {
			PnewServiceTokenRequest: r.newServiceTokenRequest,
	}
	*mock.CreateNewServiceTokenCalls = append(*mock.CreateNewServiceTokenCalls, p)
	return mock.CreateNewServiceTokenResponse.Result, mock.CreateNewServiceTokenResponse.Response, mock.CreateNewServiceTokenResponse.Error
}

type DeleteServiceTokenMockResponse struct {
	
	Response *_nethttp.Response
	Error error
}

type DeleteServiceTokenCall struct {
	PserviceTokenId int64
}


func (mock ServiceTokenApiMock) DeleteServiceToken(ctx _context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest {
	return ApiDeleteServiceTokenRequest{
		ApiService: mock,
		ctx: ctx,
		serviceTokenId: serviceTokenId,
	}
}

func (mock ServiceTokenApiMock) DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*_nethttp.Response, error) {
	p := DeleteServiceTokenCall {
			PserviceTokenId: r.serviceTokenId,
	}
	*mock.DeleteServiceTokenCalls = append(*mock.DeleteServiceTokenCalls, p)
	return mock.DeleteServiceTokenResponse.Response, mock.DeleteServiceTokenResponse.Error
}

type GetServiceTokensMockResponse struct {
	Result []ServiceToken
	Response *_nethttp.Response
	Error error
}

type GetServiceTokensCall struct {
}


func (mock ServiceTokenApiMock) GetServiceTokens(ctx _context.Context) ApiGetServiceTokensRequest {
	return ApiGetServiceTokensRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock ServiceTokenApiMock) GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *_nethttp.Response, error) {
	p := GetServiceTokensCall {
	}
	*mock.GetServiceTokensCalls = append(*mock.GetServiceTokensCalls, p)
	return mock.GetServiceTokensResponse.Result, mock.GetServiceTokensResponse.Response, mock.GetServiceTokensResponse.Error
}

