# \OtelMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOtelComponentMapping**](OtelMappingApi.md#DeleteOtelComponentMapping) | **Delete** /otel-component-mappings/{identifier} | Deletes an OTel Component Mapping.
[**DeleteOtelRelationMapping**](OtelMappingApi.md#DeleteOtelRelationMapping) | **Delete** /otel-relation-mappings/{identifier} | Deletes an OTel Relation Mapping.
[**GetOtelComponentMapping**](OtelMappingApi.md#GetOtelComponentMapping) | **Get** /otel-component-mappings/{identifier} | Get an OTel Component Mapping.
[**GetOtelComponentMappingStatus**](OtelMappingApi.md#GetOtelComponentMappingStatus) | **Get** /otel-component-mappings/{identifier}/status | Get the status of an otel component mapping synchronization.
[**GetOtelComponentMappings**](OtelMappingApi.md#GetOtelComponentMappings) | **Get** /otel-component-mappings | Get all otel component mappings.
[**GetOtelRelationMapping**](OtelMappingApi.md#GetOtelRelationMapping) | **Get** /otel-relation-mappings/{identifier} | Get an OTel Relation Mapping.
[**GetOtelRelationMappingStatus**](OtelMappingApi.md#GetOtelRelationMappingStatus) | **Get** /otel-relation-mappings/{identifier}/status | Get the status of an otel relation mapping synchronization.
[**GetOtelRelationMappings**](OtelMappingApi.md#GetOtelRelationMappings) | **Get** /otel-relation-mappings | Get all otel relation mappings.
[**UpsertOtelComponentMappings**](OtelMappingApi.md#UpsertOtelComponentMappings) | **Put** /otel-component-mappings | Upserts (creates/updates) an OTel Component Mappings.
[**UpsertOtelRelationMappings**](OtelMappingApi.md#UpsertOtelRelationMappings) | **Put** /otel-relation-mappings | Upserts (creates/updates) an OTel Relation Mappings.



## DeleteOtelComponentMapping

> DeleteOtelComponentMapping(ctx, identifier).Execute()

Deletes an OTel Component Mapping.

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
    resp, r, err := apiClient.OtelMappingApi.DeleteOtelComponentMapping(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.DeleteOtelComponentMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOtelComponentMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteOtelRelationMapping

> DeleteOtelRelationMapping(ctx, identifier).Execute()

Deletes an OTel Relation Mapping.

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
    resp, r, err := apiClient.OtelMappingApi.DeleteOtelRelationMapping(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.DeleteOtelRelationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOtelRelationMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOtelComponentMapping

> OtelComponentMapping GetOtelComponentMapping(ctx, identifier).Execute()

Get an OTel Component Mapping.

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
    resp, r, err := apiClient.OtelMappingApi.GetOtelComponentMapping(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelComponentMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelComponentMapping`: OtelComponentMapping
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelComponentMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelComponentMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtelComponentMapping**](OtelComponentMapping.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetOtelRelationMapping

> OtelRelationMapping GetOtelRelationMapping(ctx, identifier).Execute()

Get an OTel Relation Mapping.

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
    resp, r, err := apiClient.OtelMappingApi.GetOtelRelationMapping(context.Background(), identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.GetOtelRelationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOtelRelationMapping`: OtelRelationMapping
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.GetOtelRelationMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOtelRelationMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtelRelationMapping**](OtelRelationMapping.md)

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


## UpsertOtelComponentMappings

> OtelMappingItem UpsertOtelComponentMappings(ctx).UpsertOtelComponentMappingsRequest(upsertOtelComponentMappingsRequest).Execute()

Upserts (creates/updates) an OTel Component Mappings.

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
    upsertOtelComponentMappingsRequest := *openapiclient.NewUpsertOtelComponentMappingsRequest("Identifier_example", "Name_example", *openapiclient.NewOtelInput([]openapiclient.OtelInputSignal{openapiclient.OtelInputSignal("TRACES")}, *openapiclient.NewOtelInputResource()), *openapiclient.NewOtelComponentMappingOutput("Identifier_example", "Name_example", "TypeName_example", "LayerName_example", "DomainName_example"), int64(123)) // UpsertOtelComponentMappingsRequest | Otel Component Mapping to create/update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtelMappingApi.UpsertOtelComponentMappings(context.Background()).UpsertOtelComponentMappingsRequest(upsertOtelComponentMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.UpsertOtelComponentMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertOtelComponentMappings`: OtelMappingItem
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.UpsertOtelComponentMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertOtelComponentMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertOtelComponentMappingsRequest** | [**UpsertOtelComponentMappingsRequest**](UpsertOtelComponentMappingsRequest.md) | Otel Component Mapping to create/update | 

### Return type

[**OtelMappingItem**](OtelMappingItem.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertOtelRelationMappings

> OtelMappingItem UpsertOtelRelationMappings(ctx).UpsertOtelRelationMappingsRequest(upsertOtelRelationMappingsRequest).Execute()

Upserts (creates/updates) an OTel Relation Mappings.

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
    upsertOtelRelationMappingsRequest := *openapiclient.NewUpsertOtelRelationMappingsRequest("Identifier_example", "Name_example", *openapiclient.NewOtelInput([]openapiclient.OtelInputSignal{openapiclient.OtelInputSignal("TRACES")}, *openapiclient.NewOtelInputResource()), *openapiclient.NewOtelRelationMappingOutput("SourceId_example", "TargetId_example", "TypeName_example"), int64(123)) // UpsertOtelRelationMappingsRequest | Otel Relation Mapping to create/update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtelMappingApi.UpsertOtelRelationMappings(context.Background()).UpsertOtelRelationMappingsRequest(upsertOtelRelationMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtelMappingApi.UpsertOtelRelationMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertOtelRelationMappings`: OtelMappingItem
    fmt.Fprintf(os.Stdout, "Response from `OtelMappingApi.UpsertOtelRelationMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertOtelRelationMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertOtelRelationMappingsRequest** | [**UpsertOtelRelationMappingsRequest**](UpsertOtelRelationMappingsRequest.md) | Otel Relation Mapping to create/update | 

### Return type

[**OtelMappingItem**](OtelMappingItem.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

