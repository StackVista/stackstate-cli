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
	"strings"
)

type IngestionApiKeyApi interface {

	/*
		DeleteIngestionApiKey Delete Ingestion Api Key

		Deleted token can't be used by sources, so all ingestion pipelines for that key will fail

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ingestionApiKeyId The identifier of a key
		@return ApiDeleteIngestionApiKeyRequest
	*/
	DeleteIngestionApiKey(ctx context.Context, ingestionApiKeyId int64) ApiDeleteIngestionApiKeyRequest

	// DeleteIngestionApiKeyExecute executes the request
	DeleteIngestionApiKeyExecute(r ApiDeleteIngestionApiKeyRequest) (*http.Response, error)

	/*
		GenerateIngestionApiKey Generate a new Ingestion Api Key

		Generates token and then returns it in the response, the token can't be obtained any more after that

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGenerateIngestionApiKeyRequest
	*/
	GenerateIngestionApiKey(ctx context.Context) ApiGenerateIngestionApiKeyRequest

	// GenerateIngestionApiKeyExecute executes the request
	//  @return GeneratedIngestionApiKeyResponse
	GenerateIngestionApiKeyExecute(r ApiGenerateIngestionApiKeyRequest) (*GeneratedIngestionApiKeyResponse, *http.Response, error)

	/*
		GetIngestionApiKeys List Ingestion Api Keys

		Returns only metadata without token itself

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetIngestionApiKeysRequest
	*/
	GetIngestionApiKeys(ctx context.Context) ApiGetIngestionApiKeysRequest

	// GetIngestionApiKeysExecute executes the request
	//  @return []IngestionApiKey
	GetIngestionApiKeysExecute(r ApiGetIngestionApiKeysRequest) ([]IngestionApiKey, *http.Response, error)
}

// IngestionApiKeyApiService IngestionApiKeyApi service
type IngestionApiKeyApiService service

type ApiDeleteIngestionApiKeyRequest struct {
	ctx               context.Context
	ApiService        IngestionApiKeyApi
	ingestionApiKeyId int64
}

func (r ApiDeleteIngestionApiKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIngestionApiKeyExecute(r)
}

/*
DeleteIngestionApiKey Delete Ingestion Api Key

Deleted token can't be used by sources, so all ingestion pipelines for that key will fail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ingestionApiKeyId The identifier of a key
 @return ApiDeleteIngestionApiKeyRequest
*/
func (a *IngestionApiKeyApiService) DeleteIngestionApiKey(ctx context.Context, ingestionApiKeyId int64) ApiDeleteIngestionApiKeyRequest {
	return ApiDeleteIngestionApiKeyRequest{
		ApiService:        a,
		ctx:               ctx,
		ingestionApiKeyId: ingestionApiKeyId,
	}
}

// Execute executes the request
func (a *IngestionApiKeyApiService) DeleteIngestionApiKeyExecute(r ApiDeleteIngestionApiKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IngestionApiKeyApiService.DeleteIngestionApiKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/ingestion/api_keys/{ingestionApiKeyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ingestionApiKeyId"+"}", url.PathEscape(parameterToString(r.ingestionApiKeyId, "")), -1)

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

type ApiGenerateIngestionApiKeyRequest struct {
	ctx                            context.Context
	ApiService                     IngestionApiKeyApi
	generateIngestionApiKeyRequest *GenerateIngestionApiKeyRequest
}

func (r ApiGenerateIngestionApiKeyRequest) GenerateIngestionApiKeyRequest(generateIngestionApiKeyRequest GenerateIngestionApiKeyRequest) ApiGenerateIngestionApiKeyRequest {
	r.generateIngestionApiKeyRequest = &generateIngestionApiKeyRequest
	return r
}

func (r ApiGenerateIngestionApiKeyRequest) Execute() (*GeneratedIngestionApiKeyResponse, *http.Response, error) {
	return r.ApiService.GenerateIngestionApiKeyExecute(r)
}

/*
GenerateIngestionApiKey Generate a new Ingestion Api Key

Generates token and then returns it in the response, the token can't be obtained any more after that

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateIngestionApiKeyRequest
*/
func (a *IngestionApiKeyApiService) GenerateIngestionApiKey(ctx context.Context) ApiGenerateIngestionApiKeyRequest {
	return ApiGenerateIngestionApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GeneratedIngestionApiKeyResponse
func (a *IngestionApiKeyApiService) GenerateIngestionApiKeyExecute(r ApiGenerateIngestionApiKeyRequest) (*GeneratedIngestionApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeneratedIngestionApiKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IngestionApiKeyApiService.GenerateIngestionApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/ingestion/api_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.generateIngestionApiKeyRequest == nil {
		return localVarReturnValue, nil, reportError("generateIngestionApiKeyRequest is required and must be specified")
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
	localVarPostBody = r.generateIngestionApiKeyRequest
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
			var v IngestionApiKeyCreateError
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

type ApiGetIngestionApiKeysRequest struct {
	ctx        context.Context
	ApiService IngestionApiKeyApi
}

func (r ApiGetIngestionApiKeysRequest) Execute() ([]IngestionApiKey, *http.Response, error) {
	return r.ApiService.GetIngestionApiKeysExecute(r)
}

/*
GetIngestionApiKeys List Ingestion Api Keys

Returns only metadata without token itself

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIngestionApiKeysRequest
*/
func (a *IngestionApiKeyApiService) GetIngestionApiKeys(ctx context.Context) ApiGetIngestionApiKeysRequest {
	return ApiGetIngestionApiKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []IngestionApiKey
func (a *IngestionApiKeyApiService) GetIngestionApiKeysExecute(r ApiGetIngestionApiKeysRequest) ([]IngestionApiKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []IngestionApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IngestionApiKeyApiService.GetIngestionApiKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/ingestion/api_keys"

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

type IngestionApiKeyApiMock struct {
	DeleteIngestionApiKeyCalls      *[]DeleteIngestionApiKeyCall
	DeleteIngestionApiKeyResponse   DeleteIngestionApiKeyMockResponse
	GenerateIngestionApiKeyCalls    *[]GenerateIngestionApiKeyCall
	GenerateIngestionApiKeyResponse GenerateIngestionApiKeyMockResponse
	GetIngestionApiKeysCalls        *[]GetIngestionApiKeysCall
	GetIngestionApiKeysResponse     GetIngestionApiKeysMockResponse
}

func NewIngestionApiKeyApiMock() IngestionApiKeyApiMock {
	xDeleteIngestionApiKeyCalls := make([]DeleteIngestionApiKeyCall, 0)
	xGenerateIngestionApiKeyCalls := make([]GenerateIngestionApiKeyCall, 0)
	xGetIngestionApiKeysCalls := make([]GetIngestionApiKeysCall, 0)
	return IngestionApiKeyApiMock{
		DeleteIngestionApiKeyCalls:   &xDeleteIngestionApiKeyCalls,
		GenerateIngestionApiKeyCalls: &xGenerateIngestionApiKeyCalls,
		GetIngestionApiKeysCalls:     &xGetIngestionApiKeysCalls,
	}
}

type DeleteIngestionApiKeyMockResponse struct {
	Response *http.Response
	Error    error
}

type DeleteIngestionApiKeyCall struct {
	PingestionApiKeyId int64
}

func (mock IngestionApiKeyApiMock) DeleteIngestionApiKey(ctx context.Context, ingestionApiKeyId int64) ApiDeleteIngestionApiKeyRequest {
	return ApiDeleteIngestionApiKeyRequest{
		ApiService:        mock,
		ctx:               ctx,
		ingestionApiKeyId: ingestionApiKeyId,
	}
}

func (mock IngestionApiKeyApiMock) DeleteIngestionApiKeyExecute(r ApiDeleteIngestionApiKeyRequest) (*http.Response, error) {
	p := DeleteIngestionApiKeyCall{
		PingestionApiKeyId: r.ingestionApiKeyId,
	}
	*mock.DeleteIngestionApiKeyCalls = append(*mock.DeleteIngestionApiKeyCalls, p)
	return mock.DeleteIngestionApiKeyResponse.Response, mock.DeleteIngestionApiKeyResponse.Error
}

type GenerateIngestionApiKeyMockResponse struct {
	Result   GeneratedIngestionApiKeyResponse
	Response *http.Response
	Error    error
}

type GenerateIngestionApiKeyCall struct {
	PgenerateIngestionApiKeyRequest *GenerateIngestionApiKeyRequest
}

func (mock IngestionApiKeyApiMock) GenerateIngestionApiKey(ctx context.Context) ApiGenerateIngestionApiKeyRequest {
	return ApiGenerateIngestionApiKeyRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock IngestionApiKeyApiMock) GenerateIngestionApiKeyExecute(r ApiGenerateIngestionApiKeyRequest) (*GeneratedIngestionApiKeyResponse, *http.Response, error) {
	p := GenerateIngestionApiKeyCall{
		PgenerateIngestionApiKeyRequest: r.generateIngestionApiKeyRequest,
	}
	*mock.GenerateIngestionApiKeyCalls = append(*mock.GenerateIngestionApiKeyCalls, p)
	return &mock.GenerateIngestionApiKeyResponse.Result, mock.GenerateIngestionApiKeyResponse.Response, mock.GenerateIngestionApiKeyResponse.Error
}

type GetIngestionApiKeysMockResponse struct {
	Result   []IngestionApiKey
	Response *http.Response
	Error    error
}

type GetIngestionApiKeysCall struct {
}

func (mock IngestionApiKeyApiMock) GetIngestionApiKeys(ctx context.Context) ApiGetIngestionApiKeysRequest {
	return ApiGetIngestionApiKeysRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock IngestionApiKeyApiMock) GetIngestionApiKeysExecute(r ApiGetIngestionApiKeysRequest) ([]IngestionApiKey, *http.Response, error) {
	p := GetIngestionApiKeysCall{}
	*mock.GetIngestionApiKeysCalls = append(*mock.GetIngestionApiKeysCalls, p)
	return mock.GetIngestionApiKeysResponse.Result, mock.GetIngestionApiKeysResponse.Response, mock.GetIngestionApiKeysResponse.Error
}
