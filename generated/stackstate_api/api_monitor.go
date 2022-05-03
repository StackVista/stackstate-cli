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

type MonitorApi interface {

	/*
	DeleteMonitor Delete a monitor

	Deletes existing monitor

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param monitorId The identifier of a monitor
	 @return ApiDeleteMonitorRequest
	*/
	DeleteMonitor(ctx _context.Context, monitorId int64) ApiDeleteMonitorRequest

	// DeleteMonitorExecute executes the request
	DeleteMonitorExecute(r ApiDeleteMonitorRequest) (*_nethttp.Response, error)

	/*
	DryRunMonitor Dry run a monitor and show a result

	Performs a dry run of a monitor without topology state modification

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param monitorId The identifier of a monitor
	 @return ApiDryRunMonitorRequest
	*/
	DryRunMonitor(ctx _context.Context, monitorId int64) ApiDryRunMonitorRequest

	// DryRunMonitorExecute executes the request
	//  @return MonitorRunResult
	DryRunMonitorExecute(r ApiDryRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error)

	/*
	GetAllMonitors List monitors

	List all available monitor in the system

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetAllMonitorsRequest
	*/
	GetAllMonitors(ctx _context.Context) ApiGetAllMonitorsRequest

	// GetAllMonitorsExecute executes the request
	//  @return MonitorList
	GetAllMonitorsExecute(r ApiGetAllMonitorsRequest) (MonitorList, *_nethttp.Response, error)

	/*
	GetMonitor Get a monitor

	Returns a monitor full representation

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param monitorId The identifier of a monitor
	 @return ApiGetMonitorRequest
	*/
	GetMonitor(ctx _context.Context, monitorId int64) ApiGetMonitorRequest

	// GetMonitorExecute executes the request
	//  @return Monitor
	GetMonitorExecute(r ApiGetMonitorRequest) (Monitor, *_nethttp.Response, error)

	/*
	GetMonitorWithStatus Get a monitor with stream information

	Returns a monitor full representation with the stream status information

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param monitorId The identifier of a monitor
	 @return ApiGetMonitorWithStatusRequest
	*/
	GetMonitorWithStatus(ctx _context.Context, monitorId int64) ApiGetMonitorWithStatusRequest

	// GetMonitorWithStatusExecute executes the request
	//  @return MonitorStatus
	GetMonitorWithStatusExecute(r ApiGetMonitorWithStatusRequest) (MonitorStatus, *_nethttp.Response, error)

	/*
	RunMonitor Run a monitor

	Runs a monitor once

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param monitorId The identifier of a monitor
	 @return ApiRunMonitorRequest
	*/
	RunMonitor(ctx _context.Context, monitorId int64) ApiRunMonitorRequest

	// RunMonitorExecute executes the request
	//  @return MonitorRunResult
	RunMonitorExecute(r ApiRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error)
}


// MonitorApiService MonitorApi service
type MonitorApiService service

type ApiDeleteMonitorRequest struct {
	ctx _context.Context
	ApiService MonitorApi
	monitorId int64
}


func (r ApiDeleteMonitorRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteMonitorExecute(r)
}

/*
DeleteMonitor Delete a monitor

Deletes existing monitor

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param monitorId The identifier of a monitor
 @return ApiDeleteMonitorRequest
*/
func (a *MonitorApiService) DeleteMonitor(ctx _context.Context, monitorId int64) ApiDeleteMonitorRequest {
	return ApiDeleteMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

// Execute executes the request
func (a *MonitorApiService) DeleteMonitorExecute(r ApiDeleteMonitorRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.DeleteMonitor")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorApiError
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

type ApiDryRunMonitorRequest struct {
	ctx _context.Context
	ApiService MonitorApi
	monitorId int64
}


func (r ApiDryRunMonitorRequest) Execute() (MonitorRunResult, *_nethttp.Response, error) {
	return r.ApiService.DryRunMonitorExecute(r)
}

/*
DryRunMonitor Dry run a monitor and show a result

Performs a dry run of a monitor without topology state modification

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param monitorId The identifier of a monitor
 @return ApiDryRunMonitorRequest
*/
func (a *MonitorApiService) DryRunMonitor(ctx _context.Context, monitorId int64) ApiDryRunMonitorRequest {
	return ApiDryRunMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

// Execute executes the request
//  @return MonitorRunResult
func (a *MonitorApiService) DryRunMonitorExecute(r ApiDryRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MonitorRunResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.DryRunMonitor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/{monitorId}/dryRun"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorApiError
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

type ApiGetAllMonitorsRequest struct {
	ctx _context.Context
	ApiService MonitorApi
}


func (r ApiGetAllMonitorsRequest) Execute() (MonitorList, *_nethttp.Response, error) {
	return r.ApiService.GetAllMonitorsExecute(r)
}

/*
GetAllMonitors List monitors

List all available monitor in the system

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllMonitorsRequest
*/
func (a *MonitorApiService) GetAllMonitors(ctx _context.Context) ApiGetAllMonitorsRequest {
	return ApiGetAllMonitorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MonitorList
func (a *MonitorApiService) GetAllMonitorsExecute(r ApiGetAllMonitorsRequest) (MonitorList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MonitorList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.GetAllMonitors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitors"

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
			var v MonitorApiError
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

type ApiGetMonitorRequest struct {
	ctx _context.Context
	ApiService MonitorApi
	monitorId int64
}


func (r ApiGetMonitorRequest) Execute() (Monitor, *_nethttp.Response, error) {
	return r.ApiService.GetMonitorExecute(r)
}

/*
GetMonitor Get a monitor

Returns a monitor full representation

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param monitorId The identifier of a monitor
 @return ApiGetMonitorRequest
*/
func (a *MonitorApiService) GetMonitor(ctx _context.Context, monitorId int64) ApiGetMonitorRequest {
	return ApiGetMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

// Execute executes the request
//  @return Monitor
func (a *MonitorApiService) GetMonitorExecute(r ApiGetMonitorRequest) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Monitor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.GetMonitor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorApiError
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

type ApiGetMonitorWithStatusRequest struct {
	ctx _context.Context
	ApiService MonitorApi
	monitorId int64
}


func (r ApiGetMonitorWithStatusRequest) Execute() (MonitorStatus, *_nethttp.Response, error) {
	return r.ApiService.GetMonitorWithStatusExecute(r)
}

/*
GetMonitorWithStatus Get a monitor with stream information

Returns a monitor full representation with the stream status information

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param monitorId The identifier of a monitor
 @return ApiGetMonitorWithStatusRequest
*/
func (a *MonitorApiService) GetMonitorWithStatus(ctx _context.Context, monitorId int64) ApiGetMonitorWithStatusRequest {
	return ApiGetMonitorWithStatusRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

// Execute executes the request
//  @return MonitorStatus
func (a *MonitorApiService) GetMonitorWithStatusExecute(r ApiGetMonitorWithStatusRequest) (MonitorStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MonitorStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.GetMonitorWithStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/{monitorId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorApiError
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

type ApiRunMonitorRequest struct {
	ctx _context.Context
	ApiService MonitorApi
	monitorId int64
}


func (r ApiRunMonitorRequest) Execute() (MonitorRunResult, *_nethttp.Response, error) {
	return r.ApiService.RunMonitorExecute(r)
}

/*
RunMonitor Run a monitor

Runs a monitor once

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param monitorId The identifier of a monitor
 @return ApiRunMonitorRequest
*/
func (a *MonitorApiService) RunMonitor(ctx _context.Context, monitorId int64) ApiRunMonitorRequest {
	return ApiRunMonitorRequest{
		ApiService: a,
		ctx: ctx,
		monitorId: monitorId,
	}
}

// Execute executes the request
//  @return MonitorRunResult
func (a *MonitorApiService) RunMonitorExecute(r ApiRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MonitorRunResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorApiService.RunMonitor")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/{monitorId}/run"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", _neturl.PathEscape(parameterToString(r.monitorId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorApiError
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


type MonitorApiMock struct {
	DeleteMonitorCalls *[]DeleteMonitorCall
	DeleteMonitorResponse DeleteMonitorMockResponse
	DryRunMonitorCalls *[]DryRunMonitorCall
	DryRunMonitorResponse DryRunMonitorMockResponse
	GetAllMonitorsCalls *[]GetAllMonitorsCall
	GetAllMonitorsResponse GetAllMonitorsMockResponse
	GetMonitorCalls *[]GetMonitorCall
	GetMonitorResponse GetMonitorMockResponse
	GetMonitorWithStatusCalls *[]GetMonitorWithStatusCall
	GetMonitorWithStatusResponse GetMonitorWithStatusMockResponse
	RunMonitorCalls *[]RunMonitorCall
	RunMonitorResponse RunMonitorMockResponse
}	

func NewMonitorApiMock() MonitorApiMock {
	xDeleteMonitorCalls := make([]DeleteMonitorCall, 0)
	xDryRunMonitorCalls := make([]DryRunMonitorCall, 0)
	xGetAllMonitorsCalls := make([]GetAllMonitorsCall, 0)
	xGetMonitorCalls := make([]GetMonitorCall, 0)
	xGetMonitorWithStatusCalls := make([]GetMonitorWithStatusCall, 0)
	xRunMonitorCalls := make([]RunMonitorCall, 0)
	return MonitorApiMock {
		DeleteMonitorCalls: &xDeleteMonitorCalls,
		DryRunMonitorCalls: &xDryRunMonitorCalls,
		GetAllMonitorsCalls: &xGetAllMonitorsCalls,
		GetMonitorCalls: &xGetMonitorCalls,
		GetMonitorWithStatusCalls: &xGetMonitorWithStatusCalls,
		RunMonitorCalls: &xRunMonitorCalls,
	}
}

type DeleteMonitorMockResponse struct {
	
	Response *_nethttp.Response
	Error error
}

type DeleteMonitorCall struct {
	PmonitorId int64
}


func (mock MonitorApiMock) DeleteMonitor(ctx _context.Context, monitorId int64) ApiDeleteMonitorRequest {
	return ApiDeleteMonitorRequest{
		ApiService: mock,
		ctx: ctx,
		monitorId: monitorId,
	}
}

func (mock MonitorApiMock) DeleteMonitorExecute(r ApiDeleteMonitorRequest) (*_nethttp.Response, error) {
	p := DeleteMonitorCall {
			PmonitorId: r.monitorId,
	}
	*mock.DeleteMonitorCalls = append(*mock.DeleteMonitorCalls, p)
	return mock.DeleteMonitorResponse.Response, mock.DeleteMonitorResponse.Error
}

type DryRunMonitorMockResponse struct {
	Result MonitorRunResult
	Response *_nethttp.Response
	Error error
}

type DryRunMonitorCall struct {
	PmonitorId int64
}


func (mock MonitorApiMock) DryRunMonitor(ctx _context.Context, monitorId int64) ApiDryRunMonitorRequest {
	return ApiDryRunMonitorRequest{
		ApiService: mock,
		ctx: ctx,
		monitorId: monitorId,
	}
}

func (mock MonitorApiMock) DryRunMonitorExecute(r ApiDryRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error) {
	p := DryRunMonitorCall {
			PmonitorId: r.monitorId,
	}
	*mock.DryRunMonitorCalls = append(*mock.DryRunMonitorCalls, p)
	return mock.DryRunMonitorResponse.Result, mock.DryRunMonitorResponse.Response, mock.DryRunMonitorResponse.Error
}

type GetAllMonitorsMockResponse struct {
	Result MonitorList
	Response *_nethttp.Response
	Error error
}

type GetAllMonitorsCall struct {
}


func (mock MonitorApiMock) GetAllMonitors(ctx _context.Context) ApiGetAllMonitorsRequest {
	return ApiGetAllMonitorsRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock MonitorApiMock) GetAllMonitorsExecute(r ApiGetAllMonitorsRequest) (MonitorList, *_nethttp.Response, error) {
	p := GetAllMonitorsCall {
	}
	*mock.GetAllMonitorsCalls = append(*mock.GetAllMonitorsCalls, p)
	return mock.GetAllMonitorsResponse.Result, mock.GetAllMonitorsResponse.Response, mock.GetAllMonitorsResponse.Error
}

type GetMonitorMockResponse struct {
	Result Monitor
	Response *_nethttp.Response
	Error error
}

type GetMonitorCall struct {
	PmonitorId int64
}


func (mock MonitorApiMock) GetMonitor(ctx _context.Context, monitorId int64) ApiGetMonitorRequest {
	return ApiGetMonitorRequest{
		ApiService: mock,
		ctx: ctx,
		monitorId: monitorId,
	}
}

func (mock MonitorApiMock) GetMonitorExecute(r ApiGetMonitorRequest) (Monitor, *_nethttp.Response, error) {
	p := GetMonitorCall {
			PmonitorId: r.monitorId,
	}
	*mock.GetMonitorCalls = append(*mock.GetMonitorCalls, p)
	return mock.GetMonitorResponse.Result, mock.GetMonitorResponse.Response, mock.GetMonitorResponse.Error
}

type GetMonitorWithStatusMockResponse struct {
	Result MonitorStatus
	Response *_nethttp.Response
	Error error
}

type GetMonitorWithStatusCall struct {
	PmonitorId int64
}


func (mock MonitorApiMock) GetMonitorWithStatus(ctx _context.Context, monitorId int64) ApiGetMonitorWithStatusRequest {
	return ApiGetMonitorWithStatusRequest{
		ApiService: mock,
		ctx: ctx,
		monitorId: monitorId,
	}
}

func (mock MonitorApiMock) GetMonitorWithStatusExecute(r ApiGetMonitorWithStatusRequest) (MonitorStatus, *_nethttp.Response, error) {
	p := GetMonitorWithStatusCall {
			PmonitorId: r.monitorId,
	}
	*mock.GetMonitorWithStatusCalls = append(*mock.GetMonitorWithStatusCalls, p)
	return mock.GetMonitorWithStatusResponse.Result, mock.GetMonitorWithStatusResponse.Response, mock.GetMonitorWithStatusResponse.Error
}

type RunMonitorMockResponse struct {
	Result MonitorRunResult
	Response *_nethttp.Response
	Error error
}

type RunMonitorCall struct {
	PmonitorId int64
}


func (mock MonitorApiMock) RunMonitor(ctx _context.Context, monitorId int64) ApiRunMonitorRequest {
	return ApiRunMonitorRequest{
		ApiService: mock,
		ctx: ctx,
		monitorId: monitorId,
	}
}

func (mock MonitorApiMock) RunMonitorExecute(r ApiRunMonitorRequest) (MonitorRunResult, *_nethttp.Response, error) {
	p := RunMonitorCall {
			PmonitorId: r.monitorId,
	}
	*mock.RunMonitorCalls = append(*mock.RunMonitorCalls, p)
	return mock.RunMonitorResponse.Result, mock.RunMonitorResponse.Response, mock.RunMonitorResponse.Error
}


