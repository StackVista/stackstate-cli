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
)


type TopologySynchronizationApi interface {

	/*
	GetTopologySynchronizationStreamById Overview of a specific Topology Stream, queried by node id or sync identifier

	Overview of a specific Topology Stream, queried by node id or sync identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTopologySynchronizationStreamByIdRequest
	*/
	GetTopologySynchronizationStreamById(ctx context.Context) ApiGetTopologySynchronizationStreamByIdRequest

	// GetTopologySynchronizationStreamByIdExecute executes the request
	//  @return TopologyStreamListItemWithErrorDetails
	GetTopologySynchronizationStreamByIdExecute(r ApiGetTopologySynchronizationStreamByIdRequest) (*TopologyStreamListItemWithErrorDetails, *http.Response, error)

	/*
	GetTopologySynchronizationStreamStatusById Metrics of a specific Topology Stream, queried by node id

	Metrics of a specific Topology Stream, queried by node id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTopologySynchronizationStreamStatusByIdRequest
	*/
	GetTopologySynchronizationStreamStatusById(ctx context.Context) ApiGetTopologySynchronizationStreamStatusByIdRequest

	// GetTopologySynchronizationStreamStatusByIdExecute executes the request
	//  @return TopologyStreamMetrics
	GetTopologySynchronizationStreamStatusByIdExecute(r ApiGetTopologySynchronizationStreamStatusByIdRequest) (*TopologyStreamMetrics, *http.Response, error)

	/*
	GetTopologySynchronizationStreams Overview of the topology synchronization streams

	Overview of the topology synchronization streams

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTopologySynchronizationStreamsRequest
	*/
	GetTopologySynchronizationStreams(ctx context.Context) ApiGetTopologySynchronizationStreamsRequest

	// GetTopologySynchronizationStreamsExecute executes the request
	//  @return TopologyStreamList
	GetTopologySynchronizationStreamsExecute(r ApiGetTopologySynchronizationStreamsRequest) (*TopologyStreamList, *http.Response, error)

	/*
	PostTopologySynchronizationStreamClearErrors Clear all the errors related to a specific sync

	Clear all the errors related to a specific sync

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostTopologySynchronizationStreamClearErrorsRequest
	*/
	PostTopologySynchronizationStreamClearErrors(ctx context.Context) ApiPostTopologySynchronizationStreamClearErrorsRequest

	// PostTopologySynchronizationStreamClearErrorsExecute executes the request
	PostTopologySynchronizationStreamClearErrorsExecute(r ApiPostTopologySynchronizationStreamClearErrorsRequest) (*http.Response, error)
}

// TopologySynchronizationApiService TopologySynchronizationApi service
type TopologySynchronizationApiService service

type ApiGetTopologySynchronizationStreamByIdRequest struct {
	ctx context.Context
	ApiService TopologySynchronizationApi
	identifier *string
	identifierType *IdentifierType
}

func (r ApiGetTopologySynchronizationStreamByIdRequest) Identifier(identifier string) ApiGetTopologySynchronizationStreamByIdRequest {
	r.identifier = &identifier
	return r
}

func (r ApiGetTopologySynchronizationStreamByIdRequest) IdentifierType(identifierType IdentifierType) ApiGetTopologySynchronizationStreamByIdRequest {
	r.identifierType = &identifierType
	return r
}

func (r ApiGetTopologySynchronizationStreamByIdRequest) Execute() (*TopologyStreamListItemWithErrorDetails, *http.Response, error) {
	return r.ApiService.GetTopologySynchronizationStreamByIdExecute(r)
}

/*
GetTopologySynchronizationStreamById Overview of a specific Topology Stream, queried by node id or sync identifier

Overview of a specific Topology Stream, queried by node id or sync identifier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTopologySynchronizationStreamByIdRequest
*/
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreamById(ctx context.Context) ApiGetTopologySynchronizationStreamByIdRequest {
	return ApiGetTopologySynchronizationStreamByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TopologyStreamListItemWithErrorDetails
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreamByIdExecute(r ApiGetTopologySynchronizationStreamByIdRequest) (*TopologyStreamListItemWithErrorDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TopologyStreamListItemWithErrorDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySynchronizationApiService.GetTopologySynchronizationStreamById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization/topology/streams/sync"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identifier == nil {
		return localVarReturnValue, nil, reportError("identifier is required and must be specified")
	}
	if r.identifierType == nil {
		return localVarReturnValue, nil, reportError("identifierType is required and must be specified")
	}

	localVarQueryParams.Add("identifier", parameterToString(*r.identifier, ""))
	localVarQueryParams.Add("identifierType", parameterToString(*r.identifierType, ""))
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
				localVarHeaderParams["X-API-Service-Bearer"] = key
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
			var v InvalidSyncIdentifier
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v TopologySyncError
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

type ApiGetTopologySynchronizationStreamStatusByIdRequest struct {
	ctx context.Context
	ApiService TopologySynchronizationApi
	identifier *string
}

func (r ApiGetTopologySynchronizationStreamStatusByIdRequest) Identifier(identifier string) ApiGetTopologySynchronizationStreamStatusByIdRequest {
	r.identifier = &identifier
	return r
}

func (r ApiGetTopologySynchronizationStreamStatusByIdRequest) Execute() (*TopologyStreamMetrics, *http.Response, error) {
	return r.ApiService.GetTopologySynchronizationStreamStatusByIdExecute(r)
}

/*
GetTopologySynchronizationStreamStatusById Metrics of a specific Topology Stream, queried by node id

Metrics of a specific Topology Stream, queried by node id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTopologySynchronizationStreamStatusByIdRequest
*/
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreamStatusById(ctx context.Context) ApiGetTopologySynchronizationStreamStatusByIdRequest {
	return ApiGetTopologySynchronizationStreamStatusByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TopologyStreamMetrics
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreamStatusByIdExecute(r ApiGetTopologySynchronizationStreamStatusByIdRequest) (*TopologyStreamMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TopologyStreamMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySynchronizationApiService.GetTopologySynchronizationStreamStatusById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization/topology/streams/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identifier == nil {
		return localVarReturnValue, nil, reportError("identifier is required and must be specified")
	}

	localVarQueryParams.Add("identifier", parameterToString(*r.identifier, ""))
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
				localVarHeaderParams["X-API-Service-Bearer"] = key
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
			var v InvalidSyncIdentifier
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v TopologySyncError
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

type ApiGetTopologySynchronizationStreamsRequest struct {
	ctx context.Context
	ApiService TopologySynchronizationApi
}

func (r ApiGetTopologySynchronizationStreamsRequest) Execute() (*TopologyStreamList, *http.Response, error) {
	return r.ApiService.GetTopologySynchronizationStreamsExecute(r)
}

/*
GetTopologySynchronizationStreams Overview of the topology synchronization streams

Overview of the topology synchronization streams

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTopologySynchronizationStreamsRequest
*/
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreams(ctx context.Context) ApiGetTopologySynchronizationStreamsRequest {
	return ApiGetTopologySynchronizationStreamsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TopologyStreamList
func (a *TopologySynchronizationApiService) GetTopologySynchronizationStreamsExecute(r ApiGetTopologySynchronizationStreamsRequest) (*TopologyStreamList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TopologyStreamList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySynchronizationApiService.GetTopologySynchronizationStreams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization/topology/streams"

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
				localVarHeaderParams["X-API-Service-Bearer"] = key
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

type ApiPostTopologySynchronizationStreamClearErrorsRequest struct {
	ctx context.Context
	ApiService TopologySynchronizationApi
	identifier *string
	identifierType *IdentifierType
}

func (r ApiPostTopologySynchronizationStreamClearErrorsRequest) Identifier(identifier string) ApiPostTopologySynchronizationStreamClearErrorsRequest {
	r.identifier = &identifier
	return r
}

func (r ApiPostTopologySynchronizationStreamClearErrorsRequest) IdentifierType(identifierType IdentifierType) ApiPostTopologySynchronizationStreamClearErrorsRequest {
	r.identifierType = &identifierType
	return r
}

func (r ApiPostTopologySynchronizationStreamClearErrorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostTopologySynchronizationStreamClearErrorsExecute(r)
}

/*
PostTopologySynchronizationStreamClearErrors Clear all the errors related to a specific sync

Clear all the errors related to a specific sync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostTopologySynchronizationStreamClearErrorsRequest
*/
func (a *TopologySynchronizationApiService) PostTopologySynchronizationStreamClearErrors(ctx context.Context) ApiPostTopologySynchronizationStreamClearErrorsRequest {
	return ApiPostTopologySynchronizationStreamClearErrorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TopologySynchronizationApiService) PostTopologySynchronizationStreamClearErrorsExecute(r ApiPostTopologySynchronizationStreamClearErrorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySynchronizationApiService.PostTopologySynchronizationStreamClearErrors")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization/topology/streams/clearErrors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identifier == nil {
		return nil, reportError("identifier is required and must be specified")
	}
	if r.identifierType == nil {
		return nil, reportError("identifierType is required and must be specified")
	}

	localVarQueryParams.Add("identifier", parameterToString(*r.identifier, ""))
	localVarQueryParams.Add("identifierType", parameterToString(*r.identifierType, ""))
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
				localVarHeaderParams["X-API-Service-Bearer"] = key
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
			var v InvalidSyncIdentifier
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v TopologySyncError
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


type TopologySynchronizationApiMock struct {
	GetTopologySynchronizationStreamByIdCalls *[]GetTopologySynchronizationStreamByIdCall
	GetTopologySynchronizationStreamByIdResponse GetTopologySynchronizationStreamByIdMockResponse
	GetTopologySynchronizationStreamStatusByIdCalls *[]GetTopologySynchronizationStreamStatusByIdCall
	GetTopologySynchronizationStreamStatusByIdResponse GetTopologySynchronizationStreamStatusByIdMockResponse
	GetTopologySynchronizationStreamsCalls *[]GetTopologySynchronizationStreamsCall
	GetTopologySynchronizationStreamsResponse GetTopologySynchronizationStreamsMockResponse
	PostTopologySynchronizationStreamClearErrorsCalls *[]PostTopologySynchronizationStreamClearErrorsCall
	PostTopologySynchronizationStreamClearErrorsResponse PostTopologySynchronizationStreamClearErrorsMockResponse
}	

func NewTopologySynchronizationApiMock() TopologySynchronizationApiMock {
	xGetTopologySynchronizationStreamByIdCalls := make([]GetTopologySynchronizationStreamByIdCall, 0)
	xGetTopologySynchronizationStreamStatusByIdCalls := make([]GetTopologySynchronizationStreamStatusByIdCall, 0)
	xGetTopologySynchronizationStreamsCalls := make([]GetTopologySynchronizationStreamsCall, 0)
	xPostTopologySynchronizationStreamClearErrorsCalls := make([]PostTopologySynchronizationStreamClearErrorsCall, 0)
	return TopologySynchronizationApiMock {
		GetTopologySynchronizationStreamByIdCalls: &xGetTopologySynchronizationStreamByIdCalls,
		GetTopologySynchronizationStreamStatusByIdCalls: &xGetTopologySynchronizationStreamStatusByIdCalls,
		GetTopologySynchronizationStreamsCalls: &xGetTopologySynchronizationStreamsCalls,
		PostTopologySynchronizationStreamClearErrorsCalls: &xPostTopologySynchronizationStreamClearErrorsCalls,
	}
}

type GetTopologySynchronizationStreamByIdMockResponse struct {
	Result TopologyStreamListItemWithErrorDetails
	Response *http.Response
	Error error
}

type GetTopologySynchronizationStreamByIdCall struct {
	Pidentifier *string
	PidentifierType *IdentifierType
}


func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreamById(ctx context.Context) ApiGetTopologySynchronizationStreamByIdRequest {
	return ApiGetTopologySynchronizationStreamByIdRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreamByIdExecute(r ApiGetTopologySynchronizationStreamByIdRequest) (*TopologyStreamListItemWithErrorDetails, *http.Response, error) {
	p := GetTopologySynchronizationStreamByIdCall {
			Pidentifier: r.identifier,
			PidentifierType: r.identifierType,
	}
	*mock.GetTopologySynchronizationStreamByIdCalls = append(*mock.GetTopologySynchronizationStreamByIdCalls, p)
	return &mock.GetTopologySynchronizationStreamByIdResponse.Result, mock.GetTopologySynchronizationStreamByIdResponse.Response, mock.GetTopologySynchronizationStreamByIdResponse.Error
}

type GetTopologySynchronizationStreamStatusByIdMockResponse struct {
	Result TopologyStreamMetrics
	Response *http.Response
	Error error
}

type GetTopologySynchronizationStreamStatusByIdCall struct {
	Pidentifier *string
}


func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreamStatusById(ctx context.Context) ApiGetTopologySynchronizationStreamStatusByIdRequest {
	return ApiGetTopologySynchronizationStreamStatusByIdRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreamStatusByIdExecute(r ApiGetTopologySynchronizationStreamStatusByIdRequest) (*TopologyStreamMetrics, *http.Response, error) {
	p := GetTopologySynchronizationStreamStatusByIdCall {
			Pidentifier: r.identifier,
	}
	*mock.GetTopologySynchronizationStreamStatusByIdCalls = append(*mock.GetTopologySynchronizationStreamStatusByIdCalls, p)
	return &mock.GetTopologySynchronizationStreamStatusByIdResponse.Result, mock.GetTopologySynchronizationStreamStatusByIdResponse.Response, mock.GetTopologySynchronizationStreamStatusByIdResponse.Error
}

type GetTopologySynchronizationStreamsMockResponse struct {
	Result TopologyStreamList
	Response *http.Response
	Error error
}

type GetTopologySynchronizationStreamsCall struct {
}


func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreams(ctx context.Context) ApiGetTopologySynchronizationStreamsRequest {
	return ApiGetTopologySynchronizationStreamsRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock TopologySynchronizationApiMock) GetTopologySynchronizationStreamsExecute(r ApiGetTopologySynchronizationStreamsRequest) (*TopologyStreamList, *http.Response, error) {
	p := GetTopologySynchronizationStreamsCall {
	}
	*mock.GetTopologySynchronizationStreamsCalls = append(*mock.GetTopologySynchronizationStreamsCalls, p)
	return &mock.GetTopologySynchronizationStreamsResponse.Result, mock.GetTopologySynchronizationStreamsResponse.Response, mock.GetTopologySynchronizationStreamsResponse.Error
}

type PostTopologySynchronizationStreamClearErrorsMockResponse struct {
	
	Response *http.Response
	Error error
}

type PostTopologySynchronizationStreamClearErrorsCall struct {
	Pidentifier *string
	PidentifierType *IdentifierType
}


func (mock TopologySynchronizationApiMock) PostTopologySynchronizationStreamClearErrors(ctx context.Context) ApiPostTopologySynchronizationStreamClearErrorsRequest {
	return ApiPostTopologySynchronizationStreamClearErrorsRequest{
		ApiService: mock,
		ctx: ctx,
	}
}

func (mock TopologySynchronizationApiMock) PostTopologySynchronizationStreamClearErrorsExecute(r ApiPostTopologySynchronizationStreamClearErrorsRequest) (*http.Response, error) {
	p := PostTopologySynchronizationStreamClearErrorsCall {
			Pidentifier: r.identifier,
			PidentifierType: r.identifierType,
	}
	*mock.PostTopologySynchronizationStreamClearErrorsCalls = append(*mock.PostTopologySynchronizationStreamClearErrorsCalls, p)
	return mock.PostTopologySynchronizationStreamClearErrorsResponse.Response, mock.PostTopologySynchronizationStreamClearErrorsResponse.Error
}


