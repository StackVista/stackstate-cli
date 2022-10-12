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

type SubjectApi interface {

	/*
		CreateSubject Create a subject

		Create a new security subject

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subject
		@return ApiCreateSubjectRequest
	*/
	CreateSubject(ctx context.Context, subject string) ApiCreateSubjectRequest

	// CreateSubjectExecute executes the request
	CreateSubjectExecute(r ApiCreateSubjectRequest) (*http.Response, error)

	/*
		DeleteSubject Delete a subject

		Remove a security subject

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subject
		@return ApiDeleteSubjectRequest
	*/
	DeleteSubject(ctx context.Context, subject string) ApiDeleteSubjectRequest

	// DeleteSubjectExecute executes the request
	DeleteSubjectExecute(r ApiDeleteSubjectRequest) (*http.Response, error)

	/*
		GetSubject Get subject

		Describe a subject and its scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param subject
		@return ApiGetSubjectRequest
	*/
	GetSubject(ctx context.Context, subject string) ApiGetSubjectRequest

	// GetSubjectExecute executes the request
	//  @return SubjectConfig
	GetSubjectExecute(r ApiGetSubjectRequest) (*SubjectConfig, *http.Response, error)

	/*
		ListSubjects List subjects

		List all subjects and their scopes

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListSubjectsRequest
	*/
	ListSubjects(ctx context.Context) ApiListSubjectsRequest

	// ListSubjectsExecute executes the request
	//  @return []SubjectConfig
	ListSubjectsExecute(r ApiListSubjectsRequest) ([]SubjectConfig, *http.Response, error)
}

// SubjectApiService SubjectApi service
type SubjectApiService service

type ApiCreateSubjectRequest struct {
	ctx           context.Context
	ApiService    SubjectApi
	subject       string
	createSubject *CreateSubject
}

func (r ApiCreateSubjectRequest) CreateSubject(createSubject CreateSubject) ApiCreateSubjectRequest {
	r.createSubject = &createSubject
	return r
}

func (r ApiCreateSubjectRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateSubjectExecute(r)
}

/*
CreateSubject Create a subject

Create a new security subject

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiCreateSubjectRequest
*/
func (a *SubjectApiService) CreateSubject(ctx context.Context, subject string) ApiCreateSubjectRequest {
	return ApiCreateSubjectRequest{
		ApiService: a,
		ctx:        ctx,
		subject:    subject,
	}
}

// Execute executes the request
func (a *SubjectApiService) CreateSubjectExecute(r ApiCreateSubjectRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectApiService.CreateSubject")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/subjects/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSubject == nil {
		return nil, reportError("createSubject is required and must be specified")
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
	localVarPostBody = r.createSubject
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
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

type ApiDeleteSubjectRequest struct {
	ctx        context.Context
	ApiService SubjectApi
	subject    string
}

func (r ApiDeleteSubjectRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSubjectExecute(r)
}

/*
DeleteSubject Delete a subject

Remove a security subject

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiDeleteSubjectRequest
*/
func (a *SubjectApiService) DeleteSubject(ctx context.Context, subject string) ApiDeleteSubjectRequest {
	return ApiDeleteSubjectRequest{
		ApiService: a,
		ctx:        ctx,
		subject:    subject,
	}
}

// Execute executes the request
func (a *SubjectApiService) DeleteSubjectExecute(r ApiDeleteSubjectRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectApiService.DeleteSubject")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/subjects/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
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

type ApiGetSubjectRequest struct {
	ctx        context.Context
	ApiService SubjectApi
	subject    string
}

func (r ApiGetSubjectRequest) Execute() (*SubjectConfig, *http.Response, error) {
	return r.ApiService.GetSubjectExecute(r)
}

/*
GetSubject Get subject

Describe a subject and its scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiGetSubjectRequest
*/
func (a *SubjectApiService) GetSubject(ctx context.Context, subject string) ApiGetSubjectRequest {
	return ApiGetSubjectRequest{
		ApiService: a,
		ctx:        ctx,
		subject:    subject,
	}
}

// Execute executes the request
//
//	@return SubjectConfig
func (a *SubjectApiService) GetSubjectExecute(r ApiGetSubjectRequest) (*SubjectConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubjectConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectApiService.GetSubject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/subjects/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericErrorsResponse
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

type ApiListSubjectsRequest struct {
	ctx        context.Context
	ApiService SubjectApi
}

func (r ApiListSubjectsRequest) Execute() ([]SubjectConfig, *http.Response, error) {
	return r.ApiService.ListSubjectsExecute(r)
}

/*
ListSubjects List subjects

List all subjects and their scopes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListSubjectsRequest
*/
func (a *SubjectApiService) ListSubjects(ctx context.Context) ApiListSubjectsRequest {
	return ApiListSubjectsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []SubjectConfig
func (a *SubjectApiService) ListSubjectsExecute(r ApiListSubjectsRequest) ([]SubjectConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []SubjectConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubjectApiService.ListSubjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/subjects"

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

type SubjectApiMock struct {
	CreateSubjectCalls    *[]CreateSubjectCall
	CreateSubjectResponse CreateSubjectMockResponse
	DeleteSubjectCalls    *[]DeleteSubjectCall
	DeleteSubjectResponse DeleteSubjectMockResponse
	GetSubjectCalls       *[]GetSubjectCall
	GetSubjectResponse    GetSubjectMockResponse
	ListSubjectsCalls     *[]ListSubjectsCall
	ListSubjectsResponse  ListSubjectsMockResponse
}

func NewSubjectApiMock() SubjectApiMock {
	xCreateSubjectCalls := make([]CreateSubjectCall, 0)
	xDeleteSubjectCalls := make([]DeleteSubjectCall, 0)
	xGetSubjectCalls := make([]GetSubjectCall, 0)
	xListSubjectsCalls := make([]ListSubjectsCall, 0)
	return SubjectApiMock{
		CreateSubjectCalls: &xCreateSubjectCalls,
		DeleteSubjectCalls: &xDeleteSubjectCalls,
		GetSubjectCalls:    &xGetSubjectCalls,
		ListSubjectsCalls:  &xListSubjectsCalls,
	}
}

type CreateSubjectMockResponse struct {
	Response *http.Response
	Error    error
}

type CreateSubjectCall struct {
	Psubject       string
	PcreateSubject *CreateSubject
}

func (mock SubjectApiMock) CreateSubject(ctx context.Context, subject string) ApiCreateSubjectRequest {
	return ApiCreateSubjectRequest{
		ApiService: mock,
		ctx:        ctx,
		subject:    subject,
	}
}

func (mock SubjectApiMock) CreateSubjectExecute(r ApiCreateSubjectRequest) (*http.Response, error) {
	p := CreateSubjectCall{
		Psubject:       r.subject,
		PcreateSubject: r.createSubject,
	}
	*mock.CreateSubjectCalls = append(*mock.CreateSubjectCalls, p)
	return mock.CreateSubjectResponse.Response, mock.CreateSubjectResponse.Error
}

type DeleteSubjectMockResponse struct {
	Response *http.Response
	Error    error
}

type DeleteSubjectCall struct {
	Psubject string
}

func (mock SubjectApiMock) DeleteSubject(ctx context.Context, subject string) ApiDeleteSubjectRequest {
	return ApiDeleteSubjectRequest{
		ApiService: mock,
		ctx:        ctx,
		subject:    subject,
	}
}

func (mock SubjectApiMock) DeleteSubjectExecute(r ApiDeleteSubjectRequest) (*http.Response, error) {
	p := DeleteSubjectCall{
		Psubject: r.subject,
	}
	*mock.DeleteSubjectCalls = append(*mock.DeleteSubjectCalls, p)
	return mock.DeleteSubjectResponse.Response, mock.DeleteSubjectResponse.Error
}

type GetSubjectMockResponse struct {
	Result   SubjectConfig
	Response *http.Response
	Error    error
}

type GetSubjectCall struct {
	Psubject string
}

func (mock SubjectApiMock) GetSubject(ctx context.Context, subject string) ApiGetSubjectRequest {
	return ApiGetSubjectRequest{
		ApiService: mock,
		ctx:        ctx,
		subject:    subject,
	}
}

func (mock SubjectApiMock) GetSubjectExecute(r ApiGetSubjectRequest) (*SubjectConfig, *http.Response, error) {
	p := GetSubjectCall{
		Psubject: r.subject,
	}
	*mock.GetSubjectCalls = append(*mock.GetSubjectCalls, p)
	return &mock.GetSubjectResponse.Result, mock.GetSubjectResponse.Response, mock.GetSubjectResponse.Error
}

type ListSubjectsMockResponse struct {
	Result   []SubjectConfig
	Response *http.Response
	Error    error
}

type ListSubjectsCall struct {
}

func (mock SubjectApiMock) ListSubjects(ctx context.Context) ApiListSubjectsRequest {
	return ApiListSubjectsRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock SubjectApiMock) ListSubjectsExecute(r ApiListSubjectsRequest) ([]SubjectConfig, *http.Response, error) {
	p := ListSubjectsCall{}
	*mock.ListSubjectsCalls = append(*mock.ListSubjectsCalls, p)
	return mock.ListSubjectsResponse.Result, mock.ListSubjectsResponse.Response, mock.ListSubjectsResponse.Error
}