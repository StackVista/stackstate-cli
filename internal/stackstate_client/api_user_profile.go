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
)

// Linger please
var (
	_ _context.Context
)

type UserProfileApi interface {

	/*
	GetCurrentUserProfile Get current user profile

	Get current user profile.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiGetCurrentUserProfileRequest
	*/
	GetCurrentUserProfile(ctx _context.Context) ApiGetCurrentUserProfileRequest

	// GetCurrentUserProfileExecute executes the request
	//  @return UserProfile
	GetCurrentUserProfileExecute(r ApiGetCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error)

	/*
	SaveCurrentUserProfile Save current user profile

	Save current user profile.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiSaveCurrentUserProfileRequest
	*/
	SaveCurrentUserProfile(ctx _context.Context) ApiSaveCurrentUserProfileRequest

	// SaveCurrentUserProfileExecute executes the request
	//  @return UserProfile
	SaveCurrentUserProfileExecute(r ApiSaveCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error)
}

type UserProfileApiMock struct {
	GetCurrentUserProfileCalls []ApiGetCurrentUserProfileRequest
	GetCurrentUserProfileResponse GetCurrentUserProfileMockResponse
	SaveCurrentUserProfileCalls []ApiSaveCurrentUserProfileRequest
	SaveCurrentUserProfileResponse SaveCurrentUserProfileMockResponse

}	

type GetCurrentUserProfileMockResponse struct {
	A UserProfile
	B *_nethttp.Response
	C error
}

func (a *UserProfileApiMock) GetCurrentUserProfile(ctx _context.Context) ApiGetCurrentUserProfileRequest {
	return ApiGetCurrentUserProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

func (a *UserProfileApiMock) GetCurrentUserProfileExecute(r ApiGetCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error) {
	a.GetCurrentUserProfileCalls = append(a.GetCurrentUserProfileCalls, r)
	return a.GetCurrentUserProfileResponse.A, a.GetCurrentUserProfileResponse.B, a.GetCurrentUserProfileResponse.C
}

type SaveCurrentUserProfileMockResponse struct {
	A UserProfile
	B *_nethttp.Response
	C error
}

func (a *UserProfileApiMock) SaveCurrentUserProfile(ctx _context.Context) ApiSaveCurrentUserProfileRequest {
	return ApiSaveCurrentUserProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

func (a *UserProfileApiMock) SaveCurrentUserProfileExecute(r ApiSaveCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error) {
	a.SaveCurrentUserProfileCalls = append(a.SaveCurrentUserProfileCalls, r)
	return a.SaveCurrentUserProfileResponse.A, a.SaveCurrentUserProfileResponse.B, a.SaveCurrentUserProfileResponse.C
}




// UserProfileApiService UserProfileApi service
type UserProfileApiService service

type ApiGetCurrentUserProfileRequest struct {
	ctx _context.Context
	ApiService UserProfileApi
}


func (r ApiGetCurrentUserProfileRequest) Execute() (UserProfile, *_nethttp.Response, error) {
	return r.ApiService.GetCurrentUserProfileExecute(r)
}

/*
GetCurrentUserProfile Get current user profile

Get current user profile.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCurrentUserProfileRequest
*/
func (a *UserProfileApiService) GetCurrentUserProfile(ctx _context.Context) ApiGetCurrentUserProfileRequest {
	return ApiGetCurrentUserProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserProfile
func (a *UserProfileApiService) GetCurrentUserProfileExecute(r ApiGetCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserProfileApiService.GetCurrentUserProfile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/profile"

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
			var v UserNotFoundError
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

type ApiSaveCurrentUserProfileRequest struct {
	ctx _context.Context
	ApiService UserProfileApi
	userProfile *UserProfile
}

func (r ApiSaveCurrentUserProfileRequest) UserProfile(userProfile UserProfile) ApiSaveCurrentUserProfileRequest {
	r.userProfile = &userProfile
	return r
}

func (r ApiSaveCurrentUserProfileRequest) Execute() (UserProfile, *_nethttp.Response, error) {
	return r.ApiService.SaveCurrentUserProfileExecute(r)
}

/*
SaveCurrentUserProfile Save current user profile

Save current user profile.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSaveCurrentUserProfileRequest
*/
func (a *UserProfileApiService) SaveCurrentUserProfile(ctx _context.Context) ApiSaveCurrentUserProfileRequest {
	return ApiSaveCurrentUserProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserProfile
func (a *UserProfileApiService) SaveCurrentUserProfileExecute(r ApiSaveCurrentUserProfileRequest) (UserProfile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserProfileApiService.SaveCurrentUserProfile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/profile"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.userProfile == nil {
		return localVarReturnValue, nil, reportError("userProfile is required and must be specified")
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
	localVarPostBody = r.userProfile
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
			var v UserProfileSaveError
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
