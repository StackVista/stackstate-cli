# \UserSessionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSessionAssumedRole**](UserSessionApi.md#GetUserSessionAssumedRole) | **Get** /user/session/assumedRole | Get the assumed a role for the current session
[**GetUserSessionAvailableRoles**](UserSessionApi.md#GetUserSessionAvailableRoles) | **Get** /user/session/availableRoles | Get a list of available roles for this session
[**SaveUserSessionAssumedRole**](UserSessionApi.md#SaveUserSessionAssumedRole) | **Put** /user/session/assumedRole | Set the assumed role for the current session



## GetUserSessionAssumedRole

> Role GetUserSessionAssumedRole(ctx).Execute()

Get the assumed a role for the current session



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
    resp, r, err := apiClient.UserSessionApi.GetUserSessionAssumedRole(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionApi.GetUserSessionAssumedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSessionAssumedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `UserSessionApi.GetUserSessionAssumedRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionAssumedRoleRequest struct via the builder pattern


### Return type

[**Role**](Role.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSessionAvailableRoles

> Roles GetUserSessionAvailableRoles(ctx).Execute()

Get a list of available roles for this session



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
    resp, r, err := apiClient.UserSessionApi.GetUserSessionAvailableRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionApi.GetUserSessionAvailableRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSessionAvailableRoles`: Roles
    fmt.Fprintf(os.Stdout, "Response from `UserSessionApi.GetUserSessionAvailableRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionAvailableRolesRequest struct via the builder pattern


### Return type

[**Roles**](Roles.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUserSessionAssumedRole

> Role SaveUserSessionAssumedRole(ctx).Role(role).Execute()

Set the assumed role for the current session



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
    role := *openapiclient.NewRole() // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSessionApi.SaveUserSessionAssumedRole(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionApi.SaveUserSessionAssumedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveUserSessionAssumedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `UserSessionApi.SaveUserSessionAssumedRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveUserSessionAssumedRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

