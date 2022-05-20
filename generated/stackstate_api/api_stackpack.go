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
	"os"
	"strings"
)

type StackpackApi interface {

	/*
		ConfirmManualSteps Confirm manual steps

		Confirm manual steps of the stackpack

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stackpackName
		@param stackpackInstanceId
		@return ApiConfirmManualStepsRequest
	*/
	ConfirmManualSteps(ctx context.Context, stackpackName string, stackpackInstanceId int64) ApiConfirmManualStepsRequest

	// ConfirmManualStepsExecute executes the request
	//  @return string
	ConfirmManualStepsExecute(r ApiConfirmManualStepsRequest) (string, *http.Response, error)

	/*
		ProvisionDetails Provision API

		Provision details

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stackName
		@return ApiProvisionDetailsRequest
	*/
	ProvisionDetails(ctx context.Context, stackName string) ApiProvisionDetailsRequest

	// ProvisionDetailsExecute executes the request
	//  @return ProvisionResponse
	ProvisionDetailsExecute(r ApiProvisionDetailsRequest) (*ProvisionResponse, *http.Response, error)

	/*
		ProvisionUninstall Provision API

		Provision details

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stackName
		@param stackId
		@return ApiProvisionUninstallRequest
	*/
	ProvisionUninstall(ctx context.Context, stackName string, stackId int64) ApiProvisionUninstallRequest

	// ProvisionUninstallExecute executes the request
	//  @return string
	ProvisionUninstallExecute(r ApiProvisionUninstallRequest) (string, *http.Response, error)

	/*
		StackpackList StackPack API

		list of stackpack

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiStackpackListRequest
	*/
	StackpackList(ctx context.Context) ApiStackpackListRequest

	// StackpackListExecute executes the request
	//  @return []Sstackpack
	StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *http.Response, error)

	/*
		StackpackUpload StackPack API

		upload a StackPack

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiStackpackUploadRequest
	*/
	StackpackUpload(ctx context.Context) ApiStackpackUploadRequest

	// StackpackUploadExecute executes the request
	//  @return StackPack
	StackpackUploadExecute(r ApiStackpackUploadRequest) (*StackPack, *http.Response, error)

	/*
		UpgradeStackPack Upgrade API

		Upgrade stackpack

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stackName
		@return ApiUpgradeStackPackRequest
	*/
	UpgradeStackPack(ctx context.Context, stackName string) ApiUpgradeStackPackRequest

	// UpgradeStackPackExecute executes the request
	//  @return string
	UpgradeStackPackExecute(r ApiUpgradeStackPackRequest) (string, *http.Response, error)
}

// StackpackApiService StackpackApi service
type StackpackApiService service

type ApiConfirmManualStepsRequest struct {
	ctx                 context.Context
	ApiService          StackpackApi
	stackpackName       string
	stackpackInstanceId int64
}

func (r ApiConfirmManualStepsRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ConfirmManualStepsExecute(r)
}

/*
ConfirmManualSteps Confirm manual steps

Confirm manual steps of the stackpack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stackpackName
 @param stackpackInstanceId
 @return ApiConfirmManualStepsRequest
*/
func (a *StackpackApiService) ConfirmManualSteps(ctx context.Context, stackpackName string, stackpackInstanceId int64) ApiConfirmManualStepsRequest {
	return ApiConfirmManualStepsRequest{
		ApiService:          a,
		ctx:                 ctx,
		stackpackName:       stackpackName,
		stackpackInstanceId: stackpackInstanceId,
	}
}

// Execute executes the request
//  @return string
func (a *StackpackApiService) ConfirmManualStepsExecute(r ApiConfirmManualStepsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.ConfirmManualSteps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack/{stackpackName}/confirm-manual-steps/{stackpackInstanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stackpackName"+"}", url.PathEscape(parameterToString(r.stackpackName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stackpackInstanceId"+"}", url.PathEscape(parameterToString(r.stackpackInstanceId, "")), -1)

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

type ApiProvisionDetailsRequest struct {
	ctx         context.Context
	ApiService  StackpackApi
	stackName   string
	requestBody *map[string]string
}

func (r ApiProvisionDetailsRequest) RequestBody(requestBody map[string]string) ApiProvisionDetailsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiProvisionDetailsRequest) Execute() (*ProvisionResponse, *http.Response, error) {
	return r.ApiService.ProvisionDetailsExecute(r)
}

/*
ProvisionDetails Provision API

Provision details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stackName
 @return ApiProvisionDetailsRequest
*/
func (a *StackpackApiService) ProvisionDetails(ctx context.Context, stackName string) ApiProvisionDetailsRequest {
	return ApiProvisionDetailsRequest{
		ApiService: a,
		ctx:        ctx,
		stackName:  stackName,
	}
}

// Execute executes the request
//  @return ProvisionResponse
func (a *StackpackApiService) ProvisionDetailsExecute(r ApiProvisionDetailsRequest) (*ProvisionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProvisionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.ProvisionDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack/{stackName}/provision"
	localVarPath = strings.Replace(localVarPath, "{"+"stackName"+"}", url.PathEscape(parameterToString(r.stackName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.requestBody
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

type ApiProvisionUninstallRequest struct {
	ctx        context.Context
	ApiService StackpackApi
	stackName  string
	stackId    int64
}

func (r ApiProvisionUninstallRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ProvisionUninstallExecute(r)
}

/*
ProvisionUninstall Provision API

Provision details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stackName
 @param stackId
 @return ApiProvisionUninstallRequest
*/
func (a *StackpackApiService) ProvisionUninstall(ctx context.Context, stackName string, stackId int64) ApiProvisionUninstallRequest {
	return ApiProvisionUninstallRequest{
		ApiService: a,
		ctx:        ctx,
		stackName:  stackName,
		stackId:    stackId,
	}
}

// Execute executes the request
//  @return string
func (a *StackpackApiService) ProvisionUninstallExecute(r ApiProvisionUninstallRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.ProvisionUninstall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack/{stackName}/deprovision/{stackId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stackName"+"}", url.PathEscape(parameterToString(r.stackName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stackId"+"}", url.PathEscape(parameterToString(r.stackId, "")), -1)

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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStackpackListRequest struct {
	ctx        context.Context
	ApiService StackpackApi
}

func (r ApiStackpackListRequest) Execute() ([]Sstackpack, *http.Response, error) {
	return r.ApiService.StackpackListExecute(r)
}

/*
StackpackList StackPack API

list of stackpack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStackpackListRequest
*/
func (a *StackpackApiService) StackpackList(ctx context.Context) ApiStackpackListRequest {
	return ApiStackpackListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []Sstackpack
func (a *StackpackApiService) StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sstackpack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.StackpackList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack"

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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStackpackUploadRequest struct {
	ctx        context.Context
	ApiService StackpackApi
	stackPack  **os.File
}

func (r ApiStackpackUploadRequest) StackPack(stackPack *os.File) ApiStackpackUploadRequest {
	r.stackPack = &stackPack
	return r
}

func (r ApiStackpackUploadRequest) Execute() (*StackPack, *http.Response, error) {
	return r.ApiService.StackpackUploadExecute(r)
}

/*
StackpackUpload StackPack API

upload a StackPack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStackpackUploadRequest
*/
func (a *StackpackApiService) StackpackUpload(ctx context.Context) ApiStackpackUploadRequest {
	return ApiStackpackUploadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return StackPack
func (a *StackpackApiService) StackpackUploadExecute(r ApiStackpackUploadRequest) (*StackPack, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StackPack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.StackpackUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	var stackPackLocalVarFormFileName string
	var stackPackLocalVarFileName string
	var stackPackLocalVarFileBytes []byte

	stackPackLocalVarFormFileName = "stackPack"

	var stackPackLocalVarFile *os.File
	if r.stackPack != nil {
		stackPackLocalVarFile = *r.stackPack
	}
	if stackPackLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(stackPackLocalVarFile)
		stackPackLocalVarFileBytes = fbs
		stackPackLocalVarFileName = stackPackLocalVarFile.Name()
		stackPackLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: stackPackLocalVarFileBytes, fileName: stackPackLocalVarFileName, formFileName: stackPackLocalVarFormFileName})
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpgradeStackPackRequest struct {
	ctx        context.Context
	ApiService StackpackApi
	stackName  string
	unlocked   *string
}

func (r ApiUpgradeStackPackRequest) Unlocked(unlocked string) ApiUpgradeStackPackRequest {
	r.unlocked = &unlocked
	return r
}

func (r ApiUpgradeStackPackRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.UpgradeStackPackExecute(r)
}

/*
UpgradeStackPack Upgrade API

Upgrade stackpack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stackName
 @return ApiUpgradeStackPackRequest
*/
func (a *StackpackApiService) UpgradeStackPack(ctx context.Context, stackName string) ApiUpgradeStackPackRequest {
	return ApiUpgradeStackPackRequest{
		ApiService: a,
		ctx:        ctx,
		stackName:  stackName,
	}
}

// Execute executes the request
//  @return string
func (a *StackpackApiService) UpgradeStackPackExecute(r ApiUpgradeStackPackRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackpackApiService.UpgradeStackPack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stackpack/{stackName}/upgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"stackName"+"}", url.PathEscape(parameterToString(r.stackName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unlocked == nil {
		return localVarReturnValue, nil, reportError("unlocked is required and must be specified")
	}

	localVarQueryParams.Add("unlocked", parameterToString(*r.unlocked, ""))
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

type StackpackApiMock struct {
	ConfirmManualStepsCalls    *[]ConfirmManualStepsCall
	ConfirmManualStepsResponse ConfirmManualStepsMockResponse
	ProvisionDetailsCalls      *[]ProvisionDetailsCall
	ProvisionDetailsResponse   ProvisionDetailsMockResponse
	ProvisionUninstallCalls    *[]ProvisionUninstallCall
	ProvisionUninstallResponse ProvisionUninstallMockResponse
	StackpackListCalls         *[]StackpackListCall
	StackpackListResponse      StackpackListMockResponse
	StackpackUploadCalls       *[]StackpackUploadCall
	StackpackUploadResponse    StackpackUploadMockResponse
	UpgradeStackPackCalls      *[]UpgradeStackPackCall
	UpgradeStackPackResponse   UpgradeStackPackMockResponse
}

func NewStackpackApiMock() StackpackApiMock {
	xConfirmManualStepsCalls := make([]ConfirmManualStepsCall, 0)
	xProvisionDetailsCalls := make([]ProvisionDetailsCall, 0)
	xProvisionUninstallCalls := make([]ProvisionUninstallCall, 0)
	xStackpackListCalls := make([]StackpackListCall, 0)
	xStackpackUploadCalls := make([]StackpackUploadCall, 0)
	xUpgradeStackPackCalls := make([]UpgradeStackPackCall, 0)
	return StackpackApiMock{
		ConfirmManualStepsCalls: &xConfirmManualStepsCalls,
		ProvisionDetailsCalls:   &xProvisionDetailsCalls,
		ProvisionUninstallCalls: &xProvisionUninstallCalls,
		StackpackListCalls:      &xStackpackListCalls,
		StackpackUploadCalls:    &xStackpackUploadCalls,
		UpgradeStackPackCalls:   &xUpgradeStackPackCalls,
	}
}

type ConfirmManualStepsMockResponse struct {
	Result   string
	Response *http.Response
	Error    error
}

type ConfirmManualStepsCall struct {
	PstackpackName       string
	PstackpackInstanceId int64
}

func (mock StackpackApiMock) ConfirmManualSteps(ctx context.Context, stackpackName string, stackpackInstanceId int64) ApiConfirmManualStepsRequest {
	return ApiConfirmManualStepsRequest{
		ApiService:          mock,
		ctx:                 ctx,
		stackpackName:       stackpackName,
		stackpackInstanceId: stackpackInstanceId,
	}
}

func (mock StackpackApiMock) ConfirmManualStepsExecute(r ApiConfirmManualStepsRequest) (string, *http.Response, error) {
	p := ConfirmManualStepsCall{
		PstackpackName:       r.stackpackName,
		PstackpackInstanceId: r.stackpackInstanceId,
	}
	*mock.ConfirmManualStepsCalls = append(*mock.ConfirmManualStepsCalls, p)
	return mock.ConfirmManualStepsResponse.Result, mock.ConfirmManualStepsResponse.Response, mock.ConfirmManualStepsResponse.Error
}

type ProvisionDetailsMockResponse struct {
	Result   ProvisionResponse
	Response *http.Response
	Error    error
}

type ProvisionDetailsCall struct {
	PstackName   string
	PrequestBody *map[string]string
}

func (mock StackpackApiMock) ProvisionDetails(ctx context.Context, stackName string) ApiProvisionDetailsRequest {
	return ApiProvisionDetailsRequest{
		ApiService: mock,
		ctx:        ctx,
		stackName:  stackName,
	}
}

func (mock StackpackApiMock) ProvisionDetailsExecute(r ApiProvisionDetailsRequest) (*ProvisionResponse, *http.Response, error) {
	p := ProvisionDetailsCall{
		PstackName:   r.stackName,
		PrequestBody: r.requestBody,
	}
	*mock.ProvisionDetailsCalls = append(*mock.ProvisionDetailsCalls, p)
	return &mock.ProvisionDetailsResponse.Result, mock.ProvisionDetailsResponse.Response, mock.ProvisionDetailsResponse.Error
}

type ProvisionUninstallMockResponse struct {
	Result   string
	Response *http.Response
	Error    error
}

type ProvisionUninstallCall struct {
	PstackName string
	PstackId   int64
}

func (mock StackpackApiMock) ProvisionUninstall(ctx context.Context, stackName string, stackId int64) ApiProvisionUninstallRequest {
	return ApiProvisionUninstallRequest{
		ApiService: mock,
		ctx:        ctx,
		stackName:  stackName,
		stackId:    stackId,
	}
}

func (mock StackpackApiMock) ProvisionUninstallExecute(r ApiProvisionUninstallRequest) (string, *http.Response, error) {
	p := ProvisionUninstallCall{
		PstackName: r.stackName,
		PstackId:   r.stackId,
	}
	*mock.ProvisionUninstallCalls = append(*mock.ProvisionUninstallCalls, p)
	return mock.ProvisionUninstallResponse.Result, mock.ProvisionUninstallResponse.Response, mock.ProvisionUninstallResponse.Error
}

type StackpackListMockResponse struct {
	Result   []Sstackpack
	Response *http.Response
	Error    error
}

type StackpackListCall struct {
}

func (mock StackpackApiMock) StackpackList(ctx context.Context) ApiStackpackListRequest {
	return ApiStackpackListRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock StackpackApiMock) StackpackListExecute(r ApiStackpackListRequest) ([]Sstackpack, *http.Response, error) {
	p := StackpackListCall{}
	*mock.StackpackListCalls = append(*mock.StackpackListCalls, p)
	return mock.StackpackListResponse.Result, mock.StackpackListResponse.Response, mock.StackpackListResponse.Error
}

type StackpackUploadMockResponse struct {
	Result   StackPack
	Response *http.Response
	Error    error
}

type StackpackUploadCall struct {
	PstackPack **os.File
}

func (mock StackpackApiMock) StackpackUpload(ctx context.Context) ApiStackpackUploadRequest {
	return ApiStackpackUploadRequest{
		ApiService: mock,
		ctx:        ctx,
	}
}

func (mock StackpackApiMock) StackpackUploadExecute(r ApiStackpackUploadRequest) (*StackPack, *http.Response, error) {
	p := StackpackUploadCall{
		PstackPack: r.stackPack,
	}
	*mock.StackpackUploadCalls = append(*mock.StackpackUploadCalls, p)
	return &mock.StackpackUploadResponse.Result, mock.StackpackUploadResponse.Response, mock.StackpackUploadResponse.Error
}

type UpgradeStackPackMockResponse struct {
	Result   string
	Response *http.Response
	Error    error
}

type UpgradeStackPackCall struct {
	PstackName string
	Punlocked  *string
}

func (mock StackpackApiMock) UpgradeStackPack(ctx context.Context, stackName string) ApiUpgradeStackPackRequest {
	return ApiUpgradeStackPackRequest{
		ApiService: mock,
		ctx:        ctx,
		stackName:  stackName,
	}
}

func (mock StackpackApiMock) UpgradeStackPackExecute(r ApiUpgradeStackPackRequest) (string, *http.Response, error) {
	p := UpgradeStackPackCall{
		PstackName: r.stackName,
		Punlocked:  r.unlocked,
	}
	*mock.UpgradeStackPackCalls = append(*mock.UpgradeStackPackCalls, p)
	return mock.UpgradeStackPackResponse.Result, mock.UpgradeStackPackResponse.Response, mock.UpgradeStackPackResponse.Error
}
