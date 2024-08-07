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

type AgentLeasesApi interface {

	/*
		AgentCheckLease Check the lease of an agent.

		Checks the lease of an agent and might register it if it does not exist yet.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param agentId The identifier of an agent
		@return ApiAgentCheckLeaseRequest
	*/
	AgentCheckLease(ctx context.Context, agentId string) ApiAgentCheckLeaseRequest

	// AgentCheckLeaseExecute executes the request
	//  @return AgentRegistration
	AgentCheckLeaseExecute(r ApiAgentCheckLeaseRequest) (*AgentRegistration, *http.Response, error)
}

// AgentLeasesApiService AgentLeasesApi service
type AgentLeasesApiService service

type ApiAgentCheckLeaseRequest struct {
	ctx               context.Context
	ApiService        AgentLeasesApi
	agentId           string
	checkLeaseRequest *CheckLeaseRequest
}

func (r ApiAgentCheckLeaseRequest) CheckLeaseRequest(checkLeaseRequest CheckLeaseRequest) ApiAgentCheckLeaseRequest {
	r.checkLeaseRequest = &checkLeaseRequest
	return r
}

func (r ApiAgentCheckLeaseRequest) Execute() (*AgentRegistration, *http.Response, error) {
	return r.ApiService.AgentCheckLeaseExecute(r)
}

/*
AgentCheckLease Check the lease of an agent.

Checks the lease of an agent and might register it if it does not exist yet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentId The identifier of an agent
 @return ApiAgentCheckLeaseRequest
*/
func (a *AgentLeasesApiService) AgentCheckLease(ctx context.Context, agentId string) ApiAgentCheckLeaseRequest {
	return ApiAgentCheckLeaseRequest{
		ApiService: a,
		ctx:        ctx,
		agentId:    agentId,
	}
}

// Execute executes the request
//  @return AgentRegistration
func (a *AgentLeasesApiService) AgentCheckLeaseExecute(r ApiAgentCheckLeaseRequest) (*AgentRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentLeasesApiService.AgentCheckLease")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agents/{agentId}/checkLease"
	localVarPath = strings.Replace(localVarPath, "{"+"agentId"+"}", url.PathEscape(parameterToString(r.agentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkLeaseRequest == nil {
		return localVarReturnValue, nil, reportError("checkLeaseRequest is required and must be specified")
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
	localVarPostBody = r.checkLeaseRequest
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v AgentRegistration
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

type AgentLeasesApiMock struct {
	AgentCheckLeaseCalls    *[]AgentCheckLeaseCall
	AgentCheckLeaseResponse AgentCheckLeaseMockResponse
}

func NewAgentLeasesApiMock() AgentLeasesApiMock {
	xAgentCheckLeaseCalls := make([]AgentCheckLeaseCall, 0)
	return AgentLeasesApiMock{
		AgentCheckLeaseCalls: &xAgentCheckLeaseCalls,
	}
}

type AgentCheckLeaseMockResponse struct {
	Result   AgentRegistration
	Response *http.Response
	Error    error
}

type AgentCheckLeaseCall struct {
	PagentId           string
	PcheckLeaseRequest *CheckLeaseRequest
}

func (mock AgentLeasesApiMock) AgentCheckLease(ctx context.Context, agentId string) ApiAgentCheckLeaseRequest {
	return ApiAgentCheckLeaseRequest{
		ApiService: mock,
		ctx:        ctx,
		agentId:    agentId,
	}
}

func (mock AgentLeasesApiMock) AgentCheckLeaseExecute(r ApiAgentCheckLeaseRequest) (*AgentRegistration, *http.Response, error) {
	p := AgentCheckLeaseCall{
		PagentId:           r.agentId,
		PcheckLeaseRequest: r.checkLeaseRequest,
	}
	*mock.AgentCheckLeaseCalls = append(*mock.AgentCheckLeaseCalls, p)
	return &mock.AgentCheckLeaseResponse.Result, mock.AgentCheckLeaseResponse.Response, mock.AgentCheckLeaseResponse.Error
}
