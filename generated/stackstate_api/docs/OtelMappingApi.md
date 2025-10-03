# \OtelMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOtelComponentMappingStatus**](OtelMappingApi.md#GetOtelComponentMappingStatus) | **Get** /otel-component-mappings/{identifier}/status | Get the status of an otel component mapping synchronization.
[**GetOtelComponentMappings**](OtelMappingApi.md#GetOtelComponentMappings) | **Get** /otel-component-mappings | Get all otel component mappings.
[**GetOtelRelationMappingStatus**](OtelMappingApi.md#GetOtelRelationMappingStatus) | **Get** /otel-relation-mappings/{identifier}/status | Get the status of an otel relation mapping synchronization.
[**GetOtelRelationMappings**](OtelMappingApi.md#GetOtelRelationMappings) | **Get** /otel-relation-mappings | Get all otel relation mappings.



## GetOtelComponentMappingStatus

> OtelMappingStatus GetOtelComponentMappingStatus(ctx, identifier).Execute()

Get the status of an otel component mapping synchronization.

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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtelMappingApi.GetOtelComponentMappingStatus(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelComponentMappingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelComponentMappingStatus`: OtelMappingStatus
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelComponentMappingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelComponentMappingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtelMappingStatus**](OtelMappingStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOtelComponentMappings

> []OtelMappingItem GetOtelComponentMappings(ctx).Execute()

Get all otel component mappings.

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
    resp, r, err := apiClient.OtelMappingApi.GetOtelComponentMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelComponentMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelComponentMappings`: []OtelMappingItem
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelComponentMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelComponentMappingsRequest struct via the builder pattern


### Return type

[**[]OtelMappingItem**](OtelMappingItem.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOtelRelationMappingStatus

> OtelMappingStatus GetOtelRelationMappingStatus(ctx, identifier).Execute()

Get the status of an otel relation mapping synchronization.

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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtelMappingApi.GetOtelRelationMappingStatus(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelRelationMappingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelRelationMappingStatus`: OtelMappingStatus
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelRelationMappingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelRelationMappingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtelMappingStatus**](OtelMappingStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOtelRelationMappings

> []OtelMappingItem GetOtelRelationMappings(ctx).Execute()

Get all otel relation mappings.

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
    resp, r, err := apiClient.OtelMappingApi.GetOtelRelationMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelRelationMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelRelationMappings`: []OtelMappingItem
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelRelationMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelRelationMappingsRequest struct via the builder pattern


### Return type

[**[]OtelMappingItem**](OtelMappingItem.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

