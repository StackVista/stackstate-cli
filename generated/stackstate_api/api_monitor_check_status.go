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

type MonitorCheckStatusApi interface {

	/*
		GetMonitorCheckStatus Get a monitor check status

		Get a monitor check status by check state id

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id The id of a monitor check status
		@return ApiGetMonitorCheckStatusRequest
	*/
	GetMonitorCheckStatus(ctx context.Context, id int64) ApiGetMonitorCheckStatusRequest

	// GetMonitorCheckStatusExecute executes the request
	//  @return MonitorCheckStatus
	GetMonitorCheckStatusExecute(r ApiGetMonitorCheckStatusRequest) (*MonitorCheckStatus, *http.Response, error)

	/*
		GetMonitorCheckStatusHealthHistory Get a monitor check health history

		Get a monitor check status health history for a defined period of time by the check state id

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id The id of a monitor check status
		@return ApiGetMonitorCheckStatusHealthHistoryRequest
	*/
	GetMonitorCheckStatusHealthHistory(ctx context.Context, id int64) ApiGetMonitorCheckStatusHealthHistoryRequest

	// GetMonitorCheckStatusHealthHistoryExecute executes the request
	//  @return MonitorCheckStatusHealthHistory
	GetMonitorCheckStatusHealthHistoryExecute(r ApiGetMonitorCheckStatusHealthHistoryRequest) (*MonitorCheckStatusHealthHistory, *http.Response, error)

	/*
		GetMonitorCheckStatusRelatedFailures Get a monitor check related failures

		Get a monitor check status related failures by the check state id

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id The id of a monitor check status
		@return ApiGetMonitorCheckStatusRelatedFailuresRequest
	*/
	GetMonitorCheckStatusRelatedFailures(ctx context.Context, id int64) ApiGetMonitorCheckStatusRelatedFailuresRequest

	// GetMonitorCheckStatusRelatedFailuresExecute executes the request
	//  @return MonitorCheckStatusRelatedFailures
	GetMonitorCheckStatusRelatedFailuresExecute(r ApiGetMonitorCheckStatusRelatedFailuresRequest) (*MonitorCheckStatusRelatedFailures, *http.Response, error)
}

// MonitorCheckStatusApiService MonitorCheckStatusApi service
type MonitorCheckStatusApiService service

type ApiGetMonitorCheckStatusRequest struct {
	ctx          context.Context
	ApiService   MonitorCheckStatusApi
	id           int64
	topologyTime *int32
}

// A timestamp at which resources will be queried. If not given the resources are queried at current time.
func (r ApiGetMonitorCheckStatusRequest) TopologyTime(topologyTime int32) ApiGetMonitorCheckStatusRequest {
	r.topologyTime = &topologyTime
	return r
}

func (r ApiGetMonitorCheckStatusRequest) Execute() (*MonitorCheckStatus, *http.Response, error) {
	return r.ApiService.GetMonitorCheckStatusExecute(r)
}

/*
GetMonitorCheckStatus Get a monitor check status

Get a monitor check status by check state id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of a monitor check status
 @return ApiGetMonitorCheckStatusRequest
*/
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatus(ctx context.Context, id int64) ApiGetMonitorCheckStatusRequest {
	return ApiGetMonitorCheckStatusRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return MonitorCheckStatus
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatusExecute(r ApiGetMonitorCheckStatusRequest) (*MonitorCheckStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MonitorCheckStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorCheckStatusApiService.GetMonitorCheckStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/checkStatus/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.topologyTime != nil {
		localVarQueryParams.Add("topologyTime", parameterToString(*r.topologyTime, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v MonitorCheckStatusApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorCheckStatusNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorCheckStatusApiError
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

type ApiGetMonitorCheckStatusHealthHistoryRequest struct {
	ctx          context.Context
	ApiService   MonitorCheckStatusApi
	id           int64
	startTime    *int32
	endTime      *int32
	topologyTime *int32
}

// The start time of a time range to query resources.
func (r ApiGetMonitorCheckStatusHealthHistoryRequest) StartTime(startTime int32) ApiGetMonitorCheckStatusHealthHistoryRequest {
	r.startTime = &startTime
	return r
}

// The end time of a time range to query resources. If not given the endTime is set to current time.
func (r ApiGetMonitorCheckStatusHealthHistoryRequest) EndTime(endTime int32) ApiGetMonitorCheckStatusHealthHistoryRequest {
	r.endTime = &endTime
	return r
}

// A timestamp at which resources will be queried. If not given the resources are queried at current time.
func (r ApiGetMonitorCheckStatusHealthHistoryRequest) TopologyTime(topologyTime int32) ApiGetMonitorCheckStatusHealthHistoryRequest {
	r.topologyTime = &topologyTime
	return r
}

func (r ApiGetMonitorCheckStatusHealthHistoryRequest) Execute() (*MonitorCheckStatusHealthHistory, *http.Response, error) {
	return r.ApiService.GetMonitorCheckStatusHealthHistoryExecute(r)
}

/*
GetMonitorCheckStatusHealthHistory Get a monitor check health history

Get a monitor check status health history for a defined period of time by the check state id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of a monitor check status
 @return ApiGetMonitorCheckStatusHealthHistoryRequest
*/
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatusHealthHistory(ctx context.Context, id int64) ApiGetMonitorCheckStatusHealthHistoryRequest {
	return ApiGetMonitorCheckStatusHealthHistoryRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return MonitorCheckStatusHealthHistory
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatusHealthHistoryExecute(r ApiGetMonitorCheckStatusHealthHistoryRequest) (*MonitorCheckStatusHealthHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MonitorCheckStatusHealthHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorCheckStatusApiService.GetMonitorCheckStatusHealthHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/checkStatus/{id}/healthHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}

	localVarQueryParams.Add("startTime", parameterToString(*r.startTime, ""))
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.topologyTime != nil {
		localVarQueryParams.Add("topologyTime", parameterToString(*r.topologyTime, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v MonitorCheckStatusApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorCheckStatusApiError
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

type ApiGetMonitorCheckStatusRelatedFailuresRequest struct {
	ctx          context.Context
	ApiService   MonitorCheckStatusApi
	id           int64
	topologyTime *int32
}

// A timestamp at which resources will be queried. If not given the resources are queried at current time.
func (r ApiGetMonitorCheckStatusRelatedFailuresRequest) TopologyTime(topologyTime int32) ApiGetMonitorCheckStatusRelatedFailuresRequest {
	r.topologyTime = &topologyTime
	return r
}

func (r ApiGetMonitorCheckStatusRelatedFailuresRequest) Execute() (*MonitorCheckStatusRelatedFailures, *http.Response, error) {
	return r.ApiService.GetMonitorCheckStatusRelatedFailuresExecute(r)
}

/*
GetMonitorCheckStatusRelatedFailures Get a monitor check related failures

Get a monitor check status related failures by the check state id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of a monitor check status
 @return ApiGetMonitorCheckStatusRelatedFailuresRequest
*/
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatusRelatedFailures(ctx context.Context, id int64) ApiGetMonitorCheckStatusRelatedFailuresRequest {
	return ApiGetMonitorCheckStatusRelatedFailuresRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return MonitorCheckStatusRelatedFailures
func (a *MonitorCheckStatusApiService) GetMonitorCheckStatusRelatedFailuresExecute(r ApiGetMonitorCheckStatusRelatedFailuresRequest) (*MonitorCheckStatusRelatedFailures, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MonitorCheckStatusRelatedFailures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitorCheckStatusApiService.GetMonitorCheckStatusRelatedFailures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/monitor/checkStatus/{id}/relatedFailures"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.topologyTime != nil {
		localVarQueryParams.Add("topologyTime", parameterToString(*r.topologyTime, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v MonitorCheckStatusApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MonitorCheckStatusNotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MonitorCheckStatusApiError
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

type MonitorCheckStatusApiMock struct {
	GetMonitorCheckStatusCalls                   *[]GetMonitorCheckStatusCall
	GetMonitorCheckStatusResponse                GetMonitorCheckStatusMockResponse
	GetMonitorCheckStatusHealthHistoryCalls      *[]GetMonitorCheckStatusHealthHistoryCall
	GetMonitorCheckStatusHealthHistoryResponse   GetMonitorCheckStatusHealthHistoryMockResponse
	GetMonitorCheckStatusRelatedFailuresCalls    *[]GetMonitorCheckStatusRelatedFailuresCall
	GetMonitorCheckStatusRelatedFailuresResponse GetMonitorCheckStatusRelatedFailuresMockResponse
}

func NewMonitorCheckStatusApiMock() MonitorCheckStatusApiMock {
	xGetMonitorCheckStatusCalls := make([]GetMonitorCheckStatusCall, 0)
	xGetMonitorCheckStatusHealthHistoryCalls := make([]GetMonitorCheckStatusHealthHistoryCall, 0)
	xGetMonitorCheckStatusRelatedFailuresCalls := make([]GetMonitorCheckStatusRelatedFailuresCall, 0)
	return MonitorCheckStatusApiMock{
		GetMonitorCheckStatusCalls:                &xGetMonitorCheckStatusCalls,
		GetMonitorCheckStatusHealthHistoryCalls:   &xGetMonitorCheckStatusHealthHistoryCalls,
		GetMonitorCheckStatusRelatedFailuresCalls: &xGetMonitorCheckStatusRelatedFailuresCalls,
	}
}

type GetMonitorCheckStatusMockResponse struct {
	Result   MonitorCheckStatus
	Response *http.Response
	Error    error
}

type GetMonitorCheckStatusCall struct {
	Pid           int64
	PtopologyTime *int32
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatus(ctx context.Context, id int64) ApiGetMonitorCheckStatusRequest {
	return ApiGetMonitorCheckStatusRequest{
		ApiService: mock,
		ctx:        ctx,
		id:         id,
	}
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatusExecute(r ApiGetMonitorCheckStatusRequest) (*MonitorCheckStatus, *http.Response, error) {
	p := GetMonitorCheckStatusCall{
		Pid:           r.id,
		PtopologyTime: r.topologyTime,
	}
	*mock.GetMonitorCheckStatusCalls = append(*mock.GetMonitorCheckStatusCalls, p)
	return &mock.GetMonitorCheckStatusResponse.Result, mock.GetMonitorCheckStatusResponse.Response, mock.GetMonitorCheckStatusResponse.Error
}

type GetMonitorCheckStatusHealthHistoryMockResponse struct {
	Result   MonitorCheckStatusHealthHistory
	Response *http.Response
	Error    error
}

type GetMonitorCheckStatusHealthHistoryCall struct {
	Pid           int64
	PstartTime    *int32
	PendTime      *int32
	PtopologyTime *int32
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatusHealthHistory(ctx context.Context, id int64) ApiGetMonitorCheckStatusHealthHistoryRequest {
	return ApiGetMonitorCheckStatusHealthHistoryRequest{
		ApiService: mock,
		ctx:        ctx,
		id:         id,
	}
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatusHealthHistoryExecute(r ApiGetMonitorCheckStatusHealthHistoryRequest) (*MonitorCheckStatusHealthHistory, *http.Response, error) {
	p := GetMonitorCheckStatusHealthHistoryCall{
		Pid:           r.id,
		PstartTime:    r.startTime,
		PendTime:      r.endTime,
		PtopologyTime: r.topologyTime,
	}
	*mock.GetMonitorCheckStatusHealthHistoryCalls = append(*mock.GetMonitorCheckStatusHealthHistoryCalls, p)
	return &mock.GetMonitorCheckStatusHealthHistoryResponse.Result, mock.GetMonitorCheckStatusHealthHistoryResponse.Response, mock.GetMonitorCheckStatusHealthHistoryResponse.Error
}

type GetMonitorCheckStatusRelatedFailuresMockResponse struct {
	Result   MonitorCheckStatusRelatedFailures
	Response *http.Response
	Error    error
}

type GetMonitorCheckStatusRelatedFailuresCall struct {
	Pid           int64
	PtopologyTime *int32
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatusRelatedFailures(ctx context.Context, id int64) ApiGetMonitorCheckStatusRelatedFailuresRequest {
	return ApiGetMonitorCheckStatusRelatedFailuresRequest{
		ApiService: mock,
		ctx:        ctx,
		id:         id,
	}
}

func (mock MonitorCheckStatusApiMock) GetMonitorCheckStatusRelatedFailuresExecute(r ApiGetMonitorCheckStatusRelatedFailuresRequest) (*MonitorCheckStatusRelatedFailures, *http.Response, error) {
	p := GetMonitorCheckStatusRelatedFailuresCall{
		Pid:           r.id,
		PtopologyTime: r.topologyTime,
	}
	*mock.GetMonitorCheckStatusRelatedFailuresCalls = append(*mock.GetMonitorCheckStatusRelatedFailuresCalls, p)
	return &mock.GetMonitorCheckStatusRelatedFailuresResponse.Result, mock.GetMonitorCheckStatusRelatedFailuresResponse.Response, mock.GetMonitorCheckStatusRelatedFailuresResponse.Error
}
