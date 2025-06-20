# \UserAuthorizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAuthorizationFor**](UserAuthorizationApi.md#GetUserAuthorizationFor) | **Get** /user/authorization/for | Is the current user authorized for the provided permission and resource



## GetUserAuthorizationFor

> GetUserAuthorizationFor(ctx).Permission(permission).ResourceName(resourceName).Execute()

Is the current user authorized for the provided permission and resource



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
    permission := "permission_example" // string | 
    resourceName := "resourceName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAuthorizationApi.GetUserAuthorizationFor(context.Background()).Permission(permission).ResourceName(resourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthorizationApi.GetUserAuthorizationFor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAuthorizationForRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string** |  | 
 **resourceName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

