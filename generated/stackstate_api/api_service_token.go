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
	"strings"
)

type ServiceTokenApi interface {

	/*
		CreateNewServiceToken Create new service token

		Create new service token.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateNewServiceTokenRequest
	*/
	CreateNewServiceToken(ctx context.Context) ApiCreateNewServiceTokenRequest

	// CreateNewServiceTokenExecute executes the request
	//  @return ServiceTokenCreatedResponse
	CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (*ServiceTokenCreatedResponse, *http.Response, error)

	/*
		DeleteServiceToken Delete service token

		Delete service token.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceTokenId The identifier of a service token
		@return ApiDeleteServiceTokenRequest
	*/
	DeleteServiceToken(ctx context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest

	// DeleteServiceTokenExecute executes the request
	DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*http.Response, error)

	/*
		GetServiceTokens Get service tokens

		Get service tokens.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetServiceTokensRequest
	*/
	GetServiceTokens(ctx context.Context) ApiGetServiceTokensRequest

	// GetServiceTokensExecute executes the request
	//  @return []ServiceToken
	GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *http.Response, error)
}

// ServiceTokenApiService ServiceTokenApi service
type ServiceTokenApiService service

type ApiCreateNewServiceTokenRequest struct {
	ctx                    context.Context
	ApiService             ServiceTokenApi
	newServiceTokenRequest *NewServiceTokenRequest
}

func (r ApiCreateNewServiceTokenRequest) NewServiceTokenRequest(newServiceTokenRequest NewServiceTokenRequest) ApiCreateNewServiceTokenRequest {
	r.newServiceTokenRequest = &newServiceTokenRequest
	return r
}

func (r ApiCreateNewServiceTokenRequest) Execute() (*ServiceTokenCreatedResponse, *http.Response, error) {
	return r.ApiService.CreateNewServiceTokenExecute(r)
}

/*
CreateNewServiceToken Create new service token

Create new service token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateNewServiceTokenRequest
*/
func (a *ServiceTokenApiService) CreateNewServiceToken(ctx context.Context) ApiCreateNewServiceTokenRequest {
	return ApiCreateNewServiceTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ServiceTokenCreatedResponse
func (a *ServiceTokenApiService) CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (*ServiceTokenCreatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServiceTokenCreatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.CreateNewServiceToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteServiceTokenRequest struct {
	ctx            context.Context
	ApiService     ServiceTokenApi
	serviceTokenId int64
}

func (r ApiDeleteServiceTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceTokenExecute(r)
}

/*
DeleteServiceToken Delete service token

Delete service token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceTokenId The identifier of a service token
	@return ApiDeleteServiceTokenRequest
*/
func (a *ServiceTokenApiService) DeleteServiceToken(ctx context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest {
	return ApiDeleteServiceTokenRequest{
		ApiService:     a,
		ctx:            ctx,
		serviceTokenId: serviceTokenId,
	}
}

// Execute executes the request
func (a *ServiceTokenApiService) DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.DeleteServiceToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/tokens/{serviceTokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceTokenId"+"}", url.PathEscape(parameterToString(r.serviceTokenId, "")), -1)

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

type ApiGetServiceTokensRequest struct {
	ctx        context.Context
	ApiService ServiceTokenApi
}

func (r ApiGetServiceTokensRequest) Execute() ([]ServiceToken, *http.Response, error) {
	return r.ApiService.GetServiceTokensExecute(r)
}

/*
GetServiceTokens Get service tokens

Get service tokens.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetServiceTokensRequest
*/
func (a *ServiceTokenApiService) GetServiceTokens(ctx context.Context) ApiGetServiceTokensRequest {
	return ApiGetServiceTokensRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ServiceToken
func (a *ServiceTokenApiService) GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ServiceToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTokenApiService.GetServiceTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/tokens"

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

type ServiceTokenApiMock struct {
	CreateNewServiceTokenCalls    *[]CreateNewServiceTokenCall
	CreateNewServiceTokenResponse CreateNewServiceTokenMockResponse
	DeleteServiceTokenCalls       *[]DeleteServiceTokenCall
	DeleteServiceTokenResponse    DeleteServiceTokenMockResponse
	GetServiceTokensCalls         *[]GetServiceTokensCall
	GetServiceTokensResponse      GetServiceTokensMockResponse
}

func NewServiceTokenApiMock() ServiceTokenApiMock {
	xCreateNewServiceTokenCalls := make([]CreateNewServiceTokenCall, 0)
	xDeleteServiceTokenCalls := make([]DeleteServiceTokenCall, 0)
	xGetServiceTokensCalls := make([]GetServiceTokensCall, 0)
	return ServiceTokenApiMock{
		CreateNewServiceTokenCalls: &xCreateNewServiceTokenCalls,
		DeleteServiceTokenCalls:    &xDeleteServiceTokenCalls,
		GetServiceTokensCalls:      &xGetServiceTokensCalls,
	}
}

type CreateNewServiceTokenMockResponse struct {
	Result   ServiceTokenCreatedResponse
	Response *http.Response
	Error    error
}

type CreateNewServiceTokenCall struct {
	PnewServiceTokenRequest *NewServiceTokenRequest
}

func (mock ServiceTokenApiMock) CreateNewServiceToken(ctx context.Context) ApiCreateNewServiceTokenRequest {
	return ApiCreateNewServiceTokenRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock ServiceTokenApiMock) CreateNewServiceTokenExecute(r ApiCreateNewServiceTokenRequest) (*ServiceTokenCreatedResponse, *http.Response, error) {
	p := CreateNewServiceTokenCall{
		PnewServiceTokenRequest: r.newServiceTokenRequest,
	}
	*mock.CreateNewServiceTokenCalls = append(*mock.CreateNewServiceTokenCalls, p)
	return &mock.CreateNewServiceTokenResponse.Result, mock.CreateNewServiceTokenResponse.Response, mock.CreateNewServiceTokenResponse.Error
}

type DeleteServiceTokenMockResponse struct {
	Response *http.Response
	Error    error
}

type DeleteServiceTokenCall struct {
	PserviceTokenId int64
}

func (mock ServiceTokenApiMock) DeleteServiceToken(ctx context.Context, serviceTokenId int64) ApiDeleteServiceTokenRequest {
	return ApiDeleteServiceTokenRequest{
		ApiService:     mock,
		ctx:            ctx,
		serviceTokenId: serviceTokenId,
	}
}

func (mock ServiceTokenApiMock) DeleteServiceTokenExecute(r ApiDeleteServiceTokenRequest) (*http.Response, error) {
	p := DeleteServiceTokenCall{
		PserviceTokenId: r.serviceTokenId,
	}
	*mock.DeleteServiceTokenCalls = append(*mock.DeleteServiceTokenCalls, p)
	return mock.DeleteServiceTokenResponse.Response, mock.DeleteServiceTokenResponse.Error
}

type GetServiceTokensMockResponse struct {
	Result   []ServiceToken
	Response *http.Response
	Error    error
}

type GetServiceTokensCall struct {
}

func (mock ServiceTokenApiMock) GetServiceTokens(ctx context.Context) ApiGetServiceTokensRequest {
	return ApiGetServiceTokensRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock ServiceTokenApiMock) GetServiceTokensExecute(r ApiGetServiceTokensRequest) ([]ServiceToken, *http.Response, error) {
	p := GetServiceTokensCall{}
	*mock.GetServiceTokensCalls = append(*mock.GetServiceTokensCalls, p)
	return mock.GetServiceTokensResponse.Result, mock.GetServiceTokensResponse.Response, mock.GetServiceTokensResponse.Error
}
