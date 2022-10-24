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


type PermissionsApi interface {

	/*
	DescribePermissions Describe permissions

	Describe permissions granted to a subject

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiDescribePermissionsRequest
	*/
	DescribePermissions(ctx context.Context, subject string) ApiDescribePermissionsRequest

	// DescribePermissionsExecute executes the request
	//  @return PermissionDescription
	DescribePermissionsExecute(r ApiDescribePermissionsRequest) (*PermissionDescription, *http.Response, error)

	/*
	GetPermissions List permissions

	Get a list of available permissions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPermissionsRequest
	*/
	GetPermissions(ctx context.Context) ApiGetPermissionsRequest

	// GetPermissionsExecute executes the request
	//  @return Permissions
	GetPermissionsExecute(r ApiGetPermissionsRequest) (*Permissions, *http.Response, error)

	/*
	GrantPermissions Grant permissions

	Grant permissions to a subject

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiGrantPermissionsRequest
	*/
	GrantPermissions(ctx context.Context, subject string) ApiGrantPermissionsRequest

	// GrantPermissionsExecute executes the request
	GrantPermissionsExecute(r ApiGrantPermissionsRequest) (*http.Response, error)

	/*
	RevokePermissions Revoke permissions

	Revoke permissions of a subject

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subject
	@return ApiRevokePermissionsRequest
	*/
	RevokePermissions(ctx context.Context, subject string) ApiRevokePermissionsRequest

	// RevokePermissionsExecute executes the request
	RevokePermissionsExecute(r ApiRevokePermissionsRequest) (*http.Response, error)
}

// PermissionsApiService PermissionsApi service
type PermissionsApiService service

type ApiDescribePermissionsRequest struct {
	ctx context.Context
	ApiService PermissionsApi
	subject string
	resource *string
	permission *string
}

func (r ApiDescribePermissionsRequest) Resource(resource string) ApiDescribePermissionsRequest {
	r.resource = &resource
	return r
}

func (r ApiDescribePermissionsRequest) Permission(permission string) ApiDescribePermissionsRequest {
	r.permission = &permission
	return r
}

func (r ApiDescribePermissionsRequest) Execute() (*PermissionDescription, *http.Response, error) {
	return r.ApiService.DescribePermissionsExecute(r)
}

/*
DescribePermissions Describe permissions

Describe permissions granted to a subject

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subject
 @return ApiDescribePermissionsRequest
*/
func (a *PermissionsApiService) DescribePermissions(ctx context.Context, subject string) ApiDescribePermissionsRequest {
	return ApiDescribePermissionsRequest{
		ApiService: a,
		ctx: ctx,
		subject: subject,
	}
}

// Execute executes the request
//  @return PermissionDescription
func (a *PermissionsApiService) DescribePermissionsExecute(r ApiDescribePermissionsRequest) (*PermissionDescription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PermissionDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.DescribePermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/permissions/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resource != nil {
		localVarQueryParams.Add("resource", parameterToString(*r.resource, ""))
	}
	if r.permission != nil {
		localVarQueryParams.Add("permission", parameterToString(*r.permission, ""))
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

type ApiGetPermissionsRequest struct {
	ctx context.Context
	ApiService PermissionsApi
}

func (r ApiGetPermissionsRequest) Execute() (*Permissions, *http.Response, error) {
	return r.ApiService.GetPermissionsExecute(r)
}

/*
GetPermissions List permissions

Get a list of available permissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPermissionsRequest
*/
func (a *PermissionsApiService) GetPermissions(ctx context.Context) ApiGetPermissionsRequest {
	return ApiGetPermissionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Permissions
func (a *PermissionsApiService) GetPermissionsExecute(r ApiGetPermissionsRequest) (*Permissions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Permissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.GetPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/permissions/list"

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

type ApiGrantPermissionsRequest struct {
	ctx context.Context
	ApiService PermissionsApi
	subject string
	grantPermission *GrantPermission
}

func (r ApiGrantPermissionsRequest) GrantPermission(grantPermission GrantPermission) ApiGrantPermissionsRequest {
	r.grantPermission = &grantPermission
	return r
}

func (r ApiGrantPermissionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GrantPermissionsExecute(r)
}

/*
GrantPermissions Grant permissions

Grant permissions to a subject

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subject
 @return ApiGrantPermissionsRequest
*/
func (a *PermissionsApiService) GrantPermissions(ctx context.Context, subject string) ApiGrantPermissionsRequest {
	return ApiGrantPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		subject: subject,
	}
}

// Execute executes the request
func (a *PermissionsApiService) GrantPermissionsExecute(r ApiGrantPermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.GrantPermissions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/permissions/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.grantPermission == nil {
		return nil, reportError("grantPermission is required and must be specified")
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
	localVarPostBody = r.grantPermission
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

type ApiRevokePermissionsRequest struct {
	ctx context.Context
	ApiService PermissionsApi
	subject string
	resource *string
	permission *string
}

func (r ApiRevokePermissionsRequest) Resource(resource string) ApiRevokePermissionsRequest {
	r.resource = &resource
	return r
}

func (r ApiRevokePermissionsRequest) Permission(permission string) ApiRevokePermissionsRequest {
	r.permission = &permission
	return r
}

func (r ApiRevokePermissionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokePermissionsExecute(r)
}

/*
RevokePermissions Revoke permissions

Revoke permissions of a subject

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subject
 @return ApiRevokePermissionsRequest
*/
func (a *PermissionsApiService) RevokePermissions(ctx context.Context, subject string) ApiRevokePermissionsRequest {
	return ApiRevokePermissionsRequest{
		ApiService: a,
		ctx: ctx,
		subject: subject,
	}
}

// Execute executes the request
func (a *PermissionsApiService) RevokePermissionsExecute(r ApiRevokePermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.RevokePermissions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/permissions/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resource != nil {
		localVarQueryParams.Add("resource", parameterToString(*r.resource, ""))
	}
	if r.permission != nil {
		localVarQueryParams.Add("permission", parameterToString(*r.permission, ""))
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


// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------


type PermissionsApiMock struct {
	DescribePermissionsCalls *[]DescribePermissionsCall
	DescribePermissionsResponse DescribePermissionsMockResponse
	GetPermissionsCalls *[]GetPermissionsCall
	GetPermissionsResponse GetPermissionsMockResponse
	GrantPermissionsCalls *[]GrantPermissionsCall
	GrantPermissionsResponse GrantPermissionsMockResponse
	RevokePermissionsCalls *[]RevokePermissionsCall
	RevokePermissionsResponse RevokePermissionsMockResponse
}	

func NewPermissionsApiMock() PermissionsApiMock {
	xDescribePermissionsCalls := make([]DescribePermissionsCall, 0)
	xGetPermissionsCalls := make([]GetPermissionsCall, 0)
	xGrantPermissionsCalls := make([]GrantPermissionsCall, 0)
	xRevokePermissionsCalls := make([]RevokePermissionsCall, 0)
	return PermissionsApiMock {
		DescribePermissionsCalls: &xDescribePermissionsCalls,
		GetPermissionsCalls: &xGetPermissionsCalls,
		GrantPermissionsCalls: &xGrantPermissionsCalls,
		RevokePermissionsCalls: &xRevokePermissionsCalls,
	}
}

type DescribePermissionsMockResponse struct {
	Result PermissionDescription
	Response *http.Response
	Error error
}

type DescribePermissionsCall struct {
	Psubject string
	Presource *string
	Ppermission *string
}


func (mock PermissionsApiMock) DescribePermissions(ctx context.Context, subject string) ApiDescribePermissionsRequest {
	return ApiDescribePermissionsRequest{
		ApiService: mock,
		ctx: ctx,
		subject: subject,
	}
}

func (mock PermissionsApiMock) DescribePermissionsExecute(r ApiDescribePermissionsRequest) (*PermissionDescription, *http.Response, error) {
	p := DescribePermissionsCall {
			Psubject: r.subject,
			Presource: r.resource,
			Ppermission: r.permission,
	}
	*mock.DescribePermissionsCalls = append(*mock.DescribePermissionsCalls, p)
	return &mock.DescribePermissionsResponse.Result, mock.DescribePermissionsResponse.Response, mock.DescribePermissionsResponse.Error
}

type GetPermissionsMockResponse struct {
	Result Permissions
	Response *http.Response
	Error error
}

type GetPermissionsCall struct {
}


func (mock PermissionsApiMock) GetPermissions(ctx context.Context) ApiGetPermissionsRequest {
	return ApiGetPermissionsRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock PermissionsApiMock) GetPermissionsExecute(r ApiGetPermissionsRequest) (*Permissions, *http.Response, error) {
	p := GetPermissionsCall {
	}
	*mock.GetPermissionsCalls = append(*mock.GetPermissionsCalls, p)
	return &mock.GetPermissionsResponse.Result, mock.GetPermissionsResponse.Response, mock.GetPermissionsResponse.Error
}

type GrantPermissionsMockResponse struct {
	
	Response *http.Response
	Error error
}

type GrantPermissionsCall struct {
	Psubject string
	PgrantPermission *GrantPermission
}


func (mock PermissionsApiMock) GrantPermissions(ctx context.Context, subject string) ApiGrantPermissionsRequest {
	return ApiGrantPermissionsRequest{
		ApiService: mock,
		ctx: ctx,
		subject: subject,
	}
}

func (mock PermissionsApiMock) GrantPermissionsExecute(r ApiGrantPermissionsRequest) (*http.Response, error) {
	p := GrantPermissionsCall {
			Psubject: r.subject,
			PgrantPermission: r.grantPermission,
	}
	*mock.GrantPermissionsCalls = append(*mock.GrantPermissionsCalls, p)
	return mock.GrantPermissionsResponse.Response, mock.GrantPermissionsResponse.Error
}

type RevokePermissionsMockResponse struct {
	
	Response *http.Response
	Error error
}

type RevokePermissionsCall struct {
	Psubject string
	Presource *string
	Ppermission *string
}


func (mock PermissionsApiMock) RevokePermissions(ctx context.Context, subject string) ApiRevokePermissionsRequest {
	return ApiRevokePermissionsRequest{
		ApiService: mock,
		ctx: ctx,
		subject: subject,
	}
}

func (mock PermissionsApiMock) RevokePermissionsExecute(r ApiRevokePermissionsRequest) (*http.Response, error) {
	p := RevokePermissionsCall {
			Psubject: r.subject,
			Presource: r.resource,
			Ppermission: r.permission,
	}
	*mock.RevokePermissionsCalls = append(*mock.RevokePermissionsCalls, p)
	return mock.RevokePermissionsResponse.Response, mock.RevokePermissionsResponse.Error
}


