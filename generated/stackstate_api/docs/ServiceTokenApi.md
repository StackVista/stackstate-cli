# \ServiceTokenAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewServiceToken**](ServiceTokenAPI.md#CreateNewServiceToken) | **Post** /security/tokens | Create new service token
[**DeleteServiceToken**](ServiceTokenAPI.md#DeleteServiceToken) | **Delete** /security/tokens/{serviceTokenId} | Delete service token
[**GetServiceTokens**](ServiceTokenAPI.md#GetServiceTokens) | **Get** /security/tokens | Get service tokens



## CreateNewServiceToken

> ServiceTokenCreatedResponse CreateNewServiceToken(ctx).NewServiceTokenRequest(newServiceTokenRequest).Execute()

Create new service token



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
    newServiceTokenRequest := *openapiclient.NewNewServiceTokenRequest("Name_example", []string{"Roles_example"}) // NewServiceTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceTokenAPI.CreateNewServiceToken(context.Background()).NewServiceTokenRequest(newServiceTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenAPI.CreateNewServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewServiceToken`: ServiceTokenCreatedResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceTokenAPI.CreateNewServiceToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newServiceTokenRequest** | [**NewServiceTokenRequest**](NewServiceTokenRequest.md) |  | 

### Return type

[**ServiceTokenCreatedResponse**](ServiceTokenCreatedResponse.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceToken

> DeleteServiceToken(ctx, serviceTokenId).Execute()

Delete service token



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
    serviceTokenId := int64(789) // int64 | The identifier of a service token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceTokenAPI.DeleteServiceToken(context.Background(), serviceTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenAPI.DeleteServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceTokenId** | **int64** | The identifier of a service token | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTokens

> []ServiceToken GetServiceTokens(ctx).Execute()

Get service tokens



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
    resp, r, err := apiClient.ServiceTokenAPI.GetServiceTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenAPI.GetServiceTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceTokens`: []ServiceToken
    fmt.Fprintf(os.Stdout, "Response from `ServiceTokenAPI.GetServiceTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTokensRequest struct via the builder pattern


### Return type

[**[]ServiceToken**](ServiceToken.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

