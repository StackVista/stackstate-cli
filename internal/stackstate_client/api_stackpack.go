/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"
)

// Linger please
var (
	_ _context.Context
)

type StackpackApi interface {

	/*
		StackpackList StackPack API

		list of stackpack

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiStackpackListRequest
	*/
	StackpackList(ctx _context.Context) ApiStackpackListRequest

	// StackpackListExecute executes the request
	//  @return []Sstackpack
	StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *_nethttp.Response, error)

	/*
		StackpackUpload StackPack API

		upload a StackPack

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiStackpackUploadRequest
	*/
	StackpackUpload(ctx _context.Context) ApiStackpackUploadRequest

	// StackpackUploadExecute executes the request
	//  @return StackPack
	StackpackUploadExecute(r ApiStackpackUploadRequest) (StackPack, *_nethttp.Response, error)
}

// StackpackApiService StackpackApi service
type StackpackApiService service

type ApiStackpackListRequest struct {
	ctx        _context.Context
	ApiService StackpackApi
}

func (r ApiStackpackListRequest) Execute() ([]Sstackpack, *_nethttp.Response, error) {
	return r.ApiService.StackpackListExecute(r)
}

/*
StackpackList StackPack API

list of stackpack

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStackpackListRequest
*/
func (a *StackpackApiService) StackpackList(ctx _context.Context) ApiStackpackListRequest {
	return ApiStackpackListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []Sstackpack
func (a *StackpackApiService) StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Sstackpack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.StackpackList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v []string
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

type ApiStackpackUploadRequest struct {
	ctx        _context.Context
	ApiService StackpackApi
	stackPack  **os.File
}

func (r ApiStackpackUploadRequest) StackPack(stackPack *os.File) ApiStackpackUploadRequest {
	r.stackPack = &stackPack
	return r
}

func (r ApiStackpackUploadRequest) Execute() (StackPack, *_nethttp.Response, error) {
	return r.ApiService.StackpackUploadExecute(r)
}

/*
StackpackUpload StackPack API

upload a StackPack

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStackpackUploadRequest
*/
func (a *StackpackApiService) StackpackUpload(ctx _context.Context) ApiStackpackUploadRequest {
	return ApiStackpackUploadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return StackPack
func (a *StackpackApiService) StackpackUploadExecute(r ApiStackpackUploadRequest) (StackPack, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StackPack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.StackpackUpload")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "stackPack"
	var localVarFile *os.File
	if r.stackPack != nil {
		localVarFile = *r.stackPack
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v []string
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

// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------

type StackpackApiMock struct {
	StackpackListCalls      *[]StackpackListCall
	StackpackListResponse   StackpackListMockResponse
	StackpackUploadCalls    *[]StackpackUploadCall
	StackpackUploadResponse StackpackUploadMockResponse
}

func NewStackpackApiMock() StackpackApiMock {
	xStackpackListCalls := make([]StackpackListCall, 0)
	xStackpackUploadCalls := make([]StackpackUploadCall, 0)
	return StackpackApiMock{
		StackpackListCalls:   &xStackpackListCalls,
		StackpackUploadCalls: &xStackpackUploadCalls,
	}
}

type StackpackListMockResponse struct {
	Result   []Sstackpack
	Response *_nethttp.Response
	Error    error
}

type StackpackListCall struct {
}

func (mock StackpackApiMock) StackpackList(ctx _context.Context) ApiStackpackListRequest {
	return ApiStackpackListRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock StackpackApiMock) StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *_nethttp.Response, error) {
	p := StackpackListCall{}
	*mock.StackpackListCalls = append(*mock.StackpackListCalls, p)
	return mock.StackpackListResponse.Result, mock.StackpackListResponse.Response, mock.StackpackListResponse.Error
}

type StackpackUploadMockResponse struct {
	Result   StackPack
	Response *_nethttp.Response
	Error    error
}

type StackpackUploadCall struct {
	PstackPack **os.File
}

func (mock StackpackApiMock) StackpackUpload(ctx _context.Context) ApiStackpackUploadRequest {
	return ApiStackpackUploadRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock StackpackApiMock) StackpackUploadExecute(r ApiStackpackUploadRequest) (StackPack, *_nethttp.Response, error) {
	p := StackpackUploadCall{
		PstackPack: r.stackPack,
	}
	*mock.StackpackUploadCalls = append(*mock.StackpackUploadCalls, p)
	return mock.StackpackUploadResponse.Result, mock.StackpackUploadResponse.Response, mock.StackpackUploadResponse.Error
}
