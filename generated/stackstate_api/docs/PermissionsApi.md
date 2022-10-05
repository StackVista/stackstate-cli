# \PermissionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribePermissions**](PermissionsApi.md#DescribePermissions) | **Get** /security/permissions/{subject} | Describe permissions
[**GetPermissions**](PermissionsApi.md#GetPermissions) | **Get** /security/permissions/list | List permissions
[**GrantPermissions**](PermissionsApi.md#GrantPermissions) | **Post** /security/permissions/{subject} | Grant permissions
[**RevokePermissions**](PermissionsApi.md#RevokePermissions) | **Delete** /security/permissions/{subject} | Revoke permissions



## DescribePermissions

> PermissionDescription DescribePermissions(ctx, subject).Execute()

Describe permissions



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
    subject := "subject_example" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.DescribePermissions(context.Background(), subject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.DescribePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribePermissions`: PermissionDescription
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.DescribePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDescribePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionDescription**](PermissionDescription.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> Permissions GetPermissions(ctx).Execute()

List permissions



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
    resp, r, err := apiClient.PermissionsApi.GetPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: Permissions
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


### Return type

[**Permissions**](Permissions.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPermissions

> GrantPermissions(ctx, subject).GrantPermission(grantPermission).Execute()

Grant permissions



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
    subject := "subject_example" // string |
    grantPermission := *openapiclient.NewGrantPermission("Permission_example", "ResourceName_example") // GrantPermission |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.GrantPermissions(context.Background(), subject).GrantPermission(grantPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GrantPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantPermission** | [**GrantPermission**](GrantPermission.md) |  |

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePermissions

> PermissionDescription RevokePermissions(ctx, subject).RevokePermission(revokePermission).Execute()

Revoke permissions



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
    subject := "subject_example" // string |
    revokePermission := *openapiclient.NewRevokePermission("Permission_example", "Resource_example") // RevokePermission |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.RevokePermissions(context.Background(), subject).RevokePermission(revokePermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.RevokePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokePermissions`: PermissionDescription
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.RevokePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokePermission** | [**RevokePermission**](RevokePermission.md) |  |

### Return type

[**PermissionDescription**](PermissionDescription.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
