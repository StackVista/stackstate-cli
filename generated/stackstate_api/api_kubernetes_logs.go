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

type KubernetesLogsApi interface {

	/*
		GetKubernetesLogs Get Kubernetes logs

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetKubernetesLogsRequest
	*/
	GetKubernetesLogs(ctx context.Context) ApiGetKubernetesLogsRequest

	// GetKubernetesLogsExecute executes the request
	//  @return GetKubernetesLogsResult
	GetKubernetesLogsExecute(r ApiGetKubernetesLogsRequest) (*GetKubernetesLogsResult, *http.Response, error)

	/*
		GetKubernetesLogsAutocomplete Get Kubernetes logs autocomplete values

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetKubernetesLogsAutocompleteRequest
	*/
	GetKubernetesLogsAutocomplete(ctx context.Context) ApiGetKubernetesLogsAutocompleteRequest

	// GetKubernetesLogsAutocompleteExecute executes the request
	//  @return GetKubernetesLogsAutocompleteResult
	GetKubernetesLogsAutocompleteExecute(r ApiGetKubernetesLogsAutocompleteRequest) (*GetKubernetesLogsAutocompleteResult, *http.Response, error)

	/*
		GetKubernetesLogsHistogram Get Kubernetes logs histogram

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetKubernetesLogsHistogramRequest
	*/
	GetKubernetesLogsHistogram(ctx context.Context) ApiGetKubernetesLogsHistogramRequest

	// GetKubernetesLogsHistogramExecute executes the request
	//  @return GetKubernetesLogsHistogramResult
	GetKubernetesLogsHistogramExecute(r ApiGetKubernetesLogsHistogramRequest) (*GetKubernetesLogsHistogramResult, *http.Response, error)
}

// KubernetesLogsApiService KubernetesLogsApi service
type KubernetesLogsApiService service

type ApiGetKubernetesLogsRequest struct {
	ctx            context.Context
	ApiService     KubernetesLogsApi
	from           *int32
	to             *int32
	podUID         *string
	pageSize       *int32
	page           *int32
	query          *string
	containerNames *[]string
	direction      *LogsDirection
	levels         *[]LogLevel
}

// Logs initial timestamp.
func (r ApiGetKubernetesLogsRequest) From(from int32) ApiGetKubernetesLogsRequest {
	r.from = &from
	return r
}

// Logs last timestamp.
func (r ApiGetKubernetesLogsRequest) To(to int32) ApiGetKubernetesLogsRequest {
	r.to = &to
	return r
}

// Find only logs for the given pod UID.
func (r ApiGetKubernetesLogsRequest) PodUID(podUID string) ApiGetKubernetesLogsRequest {
	r.podUID = &podUID
	return r
}

// Maximum number of the log lines in the result.
func (r ApiGetKubernetesLogsRequest) PageSize(pageSize int32) ApiGetKubernetesLogsRequest {
	r.pageSize = &pageSize
	return r
}

// The page for which the log lines of pageSize must be returned.
func (r ApiGetKubernetesLogsRequest) Page(page int32) ApiGetKubernetesLogsRequest {
	r.page = &page
	return r
}

// Find only logs containing query text.
func (r ApiGetKubernetesLogsRequest) Query(query string) ApiGetKubernetesLogsRequest {
	r.query = &query
	return r
}

// Find only logs for the given container names.
func (r ApiGetKubernetesLogsRequest) ContainerNames(containerNames []string) ApiGetKubernetesLogsRequest {
	r.containerNames = &containerNames
	return r
}

// Fetch Oldest or Newest first.
func (r ApiGetKubernetesLogsRequest) Direction(direction LogsDirection) ApiGetKubernetesLogsRequest {
	r.direction = &direction
	return r
}

// Search a specific log level DEBUG, INFO, WARN, ERROR.
func (r ApiGetKubernetesLogsRequest) Levels(levels []LogLevel) ApiGetKubernetesLogsRequest {
	r.levels = &levels
	return r
}

func (r ApiGetKubernetesLogsRequest) Execute() (*GetKubernetesLogsResult, *http.Response, error) {
	return r.ApiService.GetKubernetesLogsExecute(r)
}

/*
GetKubernetesLogs Get Kubernetes logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKubernetesLogsRequest
*/
func (a *KubernetesLogsApiService) GetKubernetesLogs(ctx context.Context) ApiGetKubernetesLogsRequest {
	return ApiGetKubernetesLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GetKubernetesLogsResult
func (a *KubernetesLogsApiService) GetKubernetesLogsExecute(r ApiGetKubernetesLogsRequest) (*GetKubernetesLogsResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetKubernetesLogsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KubernetesLogsApiService.GetKubernetesLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/k8s/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.podUID == nil {
		return localVarReturnValue, nil, reportError("podUID is required and must be specified")
	}

	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	localVarQueryParams.Add("podUID", parameterToString(*r.podUID, ""))
	if r.containerNames != nil {
		localVarQueryParams.Add("containerNames", parameterToString(*r.containerNames, "csv"))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	if r.levels != nil {
		localVarQueryParams.Add("levels", parameterToString(*r.levels, "csv"))
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
			var v GetKubernetesLogsBadRequest
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

type ApiGetKubernetesLogsAutocompleteRequest struct {
	ctx        context.Context
	ApiService KubernetesLogsApi
	from       *int32
	to         *int32
	podUID     *string
}

// Logs initial timestamp.
func (r ApiGetKubernetesLogsAutocompleteRequest) From(from int32) ApiGetKubernetesLogsAutocompleteRequest {
	r.from = &from
	return r
}

// Logs last timestamp.
func (r ApiGetKubernetesLogsAutocompleteRequest) To(to int32) ApiGetKubernetesLogsAutocompleteRequest {
	r.to = &to
	return r
}

// Find only logs for the given pod UID.
func (r ApiGetKubernetesLogsAutocompleteRequest) PodUID(podUID string) ApiGetKubernetesLogsAutocompleteRequest {
	r.podUID = &podUID
	return r
}

func (r ApiGetKubernetesLogsAutocompleteRequest) Execute() (*GetKubernetesLogsAutocompleteResult, *http.Response, error) {
	return r.ApiService.GetKubernetesLogsAutocompleteExecute(r)
}

/*
GetKubernetesLogsAutocomplete Get Kubernetes logs autocomplete values

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKubernetesLogsAutocompleteRequest
*/
func (a *KubernetesLogsApiService) GetKubernetesLogsAutocomplete(ctx context.Context) ApiGetKubernetesLogsAutocompleteRequest {
	return ApiGetKubernetesLogsAutocompleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GetKubernetesLogsAutocompleteResult
func (a *KubernetesLogsApiService) GetKubernetesLogsAutocompleteExecute(r ApiGetKubernetesLogsAutocompleteRequest) (*GetKubernetesLogsAutocompleteResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetKubernetesLogsAutocompleteResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KubernetesLogsApiService.GetKubernetesLogsAutocomplete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/k8s/logs/autocomplete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.podUID == nil {
		return localVarReturnValue, nil, reportError("podUID is required and must be specified")
	}

	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	localVarQueryParams.Add("podUID", parameterToString(*r.podUID, ""))
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
			var v GetKubernetesLogsAutocompleteBadRequest
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

type ApiGetKubernetesLogsHistogramRequest struct {
	ctx            context.Context
	ApiService     KubernetesLogsApi
	from           *int32
	to             *int32
	podUID         *string
	bucketsCount   *int32
	query          *string
	containerNames *[]string
	levels         *[]LogLevel
}

// Logs initial timestamp.
func (r ApiGetKubernetesLogsHistogramRequest) From(from int32) ApiGetKubernetesLogsHistogramRequest {
	r.from = &from
	return r
}

// Logs last timestamp.
func (r ApiGetKubernetesLogsHistogramRequest) To(to int32) ApiGetKubernetesLogsHistogramRequest {
	r.to = &to
	return r
}

// Find only logs for the given pod UID.
func (r ApiGetKubernetesLogsHistogramRequest) PodUID(podUID string) ApiGetKubernetesLogsHistogramRequest {
	r.podUID = &podUID
	return r
}

// The number of histogram buckets.
func (r ApiGetKubernetesLogsHistogramRequest) BucketsCount(bucketsCount int32) ApiGetKubernetesLogsHistogramRequest {
	r.bucketsCount = &bucketsCount
	return r
}

// Find only logs containing query text.
func (r ApiGetKubernetesLogsHistogramRequest) Query(query string) ApiGetKubernetesLogsHistogramRequest {
	r.query = &query
	return r
}

// Find only logs for the given container names.
func (r ApiGetKubernetesLogsHistogramRequest) ContainerNames(containerNames []string) ApiGetKubernetesLogsHistogramRequest {
	r.containerNames = &containerNames
	return r
}

// Search a specific log level DEBUG, INFO, WARN, ERROR.
func (r ApiGetKubernetesLogsHistogramRequest) Levels(levels []LogLevel) ApiGetKubernetesLogsHistogramRequest {
	r.levels = &levels
	return r
}

func (r ApiGetKubernetesLogsHistogramRequest) Execute() (*GetKubernetesLogsHistogramResult, *http.Response, error) {
	return r.ApiService.GetKubernetesLogsHistogramExecute(r)
}

/*
GetKubernetesLogsHistogram Get Kubernetes logs histogram

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKubernetesLogsHistogramRequest
*/
func (a *KubernetesLogsApiService) GetKubernetesLogsHistogram(ctx context.Context) ApiGetKubernetesLogsHistogramRequest {
	return ApiGetKubernetesLogsHistogramRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GetKubernetesLogsHistogramResult
func (a *KubernetesLogsApiService) GetKubernetesLogsHistogramExecute(r ApiGetKubernetesLogsHistogramRequest) (*GetKubernetesLogsHistogramResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetKubernetesLogsHistogramResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KubernetesLogsApiService.GetKubernetesLogsHistogram")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/k8s/logs/histogram"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.podUID == nil {
		return localVarReturnValue, nil, reportError("podUID is required and must be specified")
	}
	if r.bucketsCount == nil {
		return localVarReturnValue, nil, reportError("bucketsCount is required and must be specified")
	}

	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	localVarQueryParams.Add("podUID", parameterToString(*r.podUID, ""))
	if r.containerNames != nil {
		localVarQueryParams.Add("containerNames", parameterToString(*r.containerNames, "csv"))
	}
	localVarQueryParams.Add("bucketsCount", parameterToString(*r.bucketsCount, ""))
	if r.levels != nil {
		localVarQueryParams.Add("levels", parameterToString(*r.levels, "csv"))
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
			var v GetKubernetesLogsHistogramBadRequest
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

// ---------------------------------------------
// ------------------ MOCKS --------------------
// ---------------------------------------------

type KubernetesLogsApiMock struct {
	GetKubernetesLogsCalls                *[]GetKubernetesLogsCall
	GetKubernetesLogsResponse             GetKubernetesLogsMockResponse
	GetKubernetesLogsAutocompleteCalls    *[]GetKubernetesLogsAutocompleteCall
	GetKubernetesLogsAutocompleteResponse GetKubernetesLogsAutocompleteMockResponse
	GetKubernetesLogsHistogramCalls       *[]GetKubernetesLogsHistogramCall
	GetKubernetesLogsHistogramResponse    GetKubernetesLogsHistogramMockResponse
}

func NewKubernetesLogsApiMock() KubernetesLogsApiMock {
	xGetKubernetesLogsCalls := make([]GetKubernetesLogsCall, 0)
	xGetKubernetesLogsAutocompleteCalls := make([]GetKubernetesLogsAutocompleteCall, 0)
	xGetKubernetesLogsHistogramCalls := make([]GetKubernetesLogsHistogramCall, 0)
	return KubernetesLogsApiMock{
		GetKubernetesLogsCalls:             &xGetKubernetesLogsCalls,
		GetKubernetesLogsAutocompleteCalls: &xGetKubernetesLogsAutocompleteCalls,
		GetKubernetesLogsHistogramCalls:    &xGetKubernetesLogsHistogramCalls,
	}
}

type GetKubernetesLogsMockResponse struct {
	Result   GetKubernetesLogsResult
	Response *http.Response
	Error    error
}

type GetKubernetesLogsCall struct {
	Pfrom           *int32
	Pto             *int32
	PpodUID         *string
	PpageSize       *int32
	Ppage           *int32
	Pquery          *string
	PcontainerNames *[]string
	Pdirection      *LogsDirection
	Plevels         *[]LogLevel
}

func (mock KubernetesLogsApiMock) GetKubernetesLogs(ctx context.Context) ApiGetKubernetesLogsRequest {
	return ApiGetKubernetesLogsRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock KubernetesLogsApiMock) GetKubernetesLogsExecute(r ApiGetKubernetesLogsRequest) (*GetKubernetesLogsResult, *http.Response, error) {
	p := GetKubernetesLogsCall{
		Pfrom:           r.from,
		Pto:             r.to,
		PpodUID:         r.podUID,
		PpageSize:       r.pageSize,
		Ppage:           r.page,
		Pquery:          r.query,
		PcontainerNames: r.containerNames,
		Pdirection:      r.direction,
		Plevels:         r.levels,
	}
	*mock.GetKubernetesLogsCalls = append(*mock.GetKubernetesLogsCalls, p)
	return &mock.GetKubernetesLogsResponse.Result, mock.GetKubernetesLogsResponse.Response, mock.GetKubernetesLogsResponse.Error
}

type GetKubernetesLogsAutocompleteMockResponse struct {
	Result   GetKubernetesLogsAutocompleteResult
	Response *http.Response
	Error    error
}

type GetKubernetesLogsAutocompleteCall struct {
	Pfrom   *int32
	Pto     *int32
	PpodUID *string
}

func (mock KubernetesLogsApiMock) GetKubernetesLogsAutocomplete(ctx context.Context) ApiGetKubernetesLogsAutocompleteRequest {
	return ApiGetKubernetesLogsAutocompleteRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock KubernetesLogsApiMock) GetKubernetesLogsAutocompleteExecute(r ApiGetKubernetesLogsAutocompleteRequest) (*GetKubernetesLogsAutocompleteResult, *http.Response, error) {
	p := GetKubernetesLogsAutocompleteCall{
		Pfrom:   r.from,
		Pto:     r.to,
		PpodUID: r.podUID,
	}
	*mock.GetKubernetesLogsAutocompleteCalls = append(*mock.GetKubernetesLogsAutocompleteCalls, p)
	return &mock.GetKubernetesLogsAutocompleteResponse.Result, mock.GetKubernetesLogsAutocompleteResponse.Response, mock.GetKubernetesLogsAutocompleteResponse.Error
}

type GetKubernetesLogsHistogramMockResponse struct {
	Result   GetKubernetesLogsHistogramResult
	Response *http.Response
	Error    error
}

type GetKubernetesLogsHistogramCall struct {
	Pfrom           *int32
	Pto             *int32
	PpodUID         *string
	PbucketsCount   *int32
	Pquery          *string
	PcontainerNames *[]string
	Plevels         *[]LogLevel
}

func (mock KubernetesLogsApiMock) GetKubernetesLogsHistogram(ctx context.Context) ApiGetKubernetesLogsHistogramRequest {
	return ApiGetKubernetesLogsHistogramRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock KubernetesLogsApiMock) GetKubernetesLogsHistogramExecute(r ApiGetKubernetesLogsHistogramRequest) (*GetKubernetesLogsHistogramResult, *http.Response, error) {
	p := GetKubernetesLogsHistogramCall{
		Pfrom:           r.from,
		Pto:             r.to,
		PpodUID:         r.podUID,
		PbucketsCount:   r.bucketsCount,
		Pquery:          r.query,
		PcontainerNames: r.containerNames,
		Plevels:         r.levels,
	}
	*mock.GetKubernetesLogsHistogramCalls = append(*mock.GetKubernetesLogsHistogramCalls, p)
	return &mock.GetKubernetesLogsHistogramResponse.Result, mock.GetKubernetesLogsHistogramResponse.Response, mock.GetKubernetesLogsHistogramResponse.Error
}
