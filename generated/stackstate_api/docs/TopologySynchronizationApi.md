# \TopologySynchronizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopologySynchronizationStreamById**](TopologySynchronizationAPI.md#GetTopologySynchronizationStreamById) | **Get** /synchronization/topology/streams/sync | Overview of a specific Topology Stream, queried by node id or sync identifier
[**GetTopologySynchronizationStreamStatusById**](TopologySynchronizationAPI.md#GetTopologySynchronizationStreamStatusById) | **Get** /synchronization/topology/streams/status | Metrics of a specific Topology Stream, queried by node id
[**GetTopologySynchronizationStreams**](TopologySynchronizationAPI.md#GetTopologySynchronizationStreams) | **Get** /synchronization/topology/streams | Overview of the topology synchronization streams
[**PostTopologySynchronizationStreamClearErrors**](TopologySynchronizationAPI.md#PostTopologySynchronizationStreamClearErrors) | **Post** /synchronization/topology/streams/clearErrors | Clear all the errors related to a specific sync



## GetTopologySynchronizationStreamById

> TopologyStreamListItemWithErrorDetails GetTopologySynchronizationStreamById(ctx).Identifier(identifier).IdentifierType(identifierType).Execute()

Overview of a specific Topology Stream, queried by node id or sync identifier



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
    identifierType := openapiclient.IdentifierType("NodeId") // IdentifierType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologySynchronizationAPI.GetTopologySynchronizationStreamById(context.Background()).Identifier(identifier).IdentifierType(identifierType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationAPI.GetTopologySynchronizationStreamById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreamById`: TopologyStreamListItemWithErrorDetails
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationAPI.GetTopologySynchronizationStreamById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologySynchronizationStreamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** |  | 
 **identifierType** | [**IdentifierType**](IdentifierType.md) |  | 

### Return type

[**TopologyStreamListItemWithErrorDetails**](TopologyStreamListItemWithErrorDetails.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologySynchronizationStreamStatusById

> TopologyStreamMetrics GetTopologySynchronizationStreamStatusById(ctx).Identifier(identifier).Execute()

Metrics of a specific Topology Stream, queried by node id



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
    resp, r, err := apiClient.TopologySynchronizationAPI.GetTopologySynchronizationStreamStatusById(context.Background()).Identifier(identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationAPI.GetTopologySynchronizationStreamStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreamStatusById`: TopologyStreamMetrics
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationAPI.GetTopologySynchronizationStreamStatusById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologySynchronizationStreamStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** |  | 

### Return type

[**TopologyStreamMetrics**](TopologyStreamMetrics.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologySynchronizationStreams

> TopologyStreamList GetTopologySynchronizationStreams(ctx).Execute()

Overview of the topology synchronization streams



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
    resp, r, err := apiClient.TopologySynchronizationAPI.GetTopologySynchronizationStreams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationAPI.GetTopologySynchronizationStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreams`: TopologyStreamList
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationAPI.GetTopologySynchronizationStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologySynchronizationStreamsRequest struct via the builder pattern


### Return type

[**TopologyStreamList**](TopologyStreamList.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTopologySynchronizationStreamClearErrors

> PostTopologySynchronizationStreamClearErrors(ctx).Identifier(identifier).IdentifierType(identifierType).Execute()

Clear all the errors related to a specific sync



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
    identifierType := openapiclient.IdentifierType("NodeId") // IdentifierType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologySynchronizationAPI.PostTopologySynchronizationStreamClearErrors(context.Background()).Identifier(identifier).IdentifierType(identifierType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationAPI.PostTopologySynchronizationStreamClearErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTopologySynchronizationStreamClearErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** |  | 
 **identifierType** | [**IdentifierType**](IdentifierType.md) |  | 

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

