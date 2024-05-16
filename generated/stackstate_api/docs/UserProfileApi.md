# \UserProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentUserProfile**](UserProfileAPI.md#GetCurrentUserProfile) | **Get** /user/profile | Get current user profile
[**SaveCurrentUserProfile**](UserProfileAPI.md#SaveCurrentUserProfile) | **Put** /user/profile | Save current user profile



## GetCurrentUserProfile

> UserProfile GetCurrentUserProfile(ctx).Execute()

Get current user profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProfileAPI.GetCurrentUserProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProfileAPI.GetCurrentUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserProfile`: UserProfile
    fmt.Fprintf(os.Stdout, "Response from `UserProfileAPI.GetCurrentUserProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserProfileRequest struct via the builder pattern


### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCurrentUserProfile

> UserProfile SaveCurrentUserProfile(ctx).UserProfile(userProfile).Execute()

Save current user profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userProfile := *openapiclient.NewUserProfile("Name_example") // UserProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProfileAPI.SaveCurrentUserProfile(context.Background()).UserProfile(userProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProfileAPI.SaveCurrentUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveCurrentUserProfile`: UserProfile
    fmt.Fprintf(os.Stdout, "Response from `UserProfileAPI.SaveCurrentUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveCurrentUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userProfile** | [**UserProfile**](UserProfile.md) |  | 

### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

