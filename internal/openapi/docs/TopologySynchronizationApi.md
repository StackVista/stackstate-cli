# \TopologySynchronizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopologySynchronizationStreamById**](TopologySynchronizationApi.md#GetTopologySynchronizationStreamById) | **Get** /synchronization/topology/streams/sync | Overview of a specific Topology Stream, queried by node id or sync identifier
[**GetTopologySynchronizationStreamStatusById**](TopologySynchronizationApi.md#GetTopologySynchronizationStreamStatusById) | **Get** /synchronization/topology/streams/status | Metrics of a specific Topology Stream, queried by node id
[**GetTopologySynchronizationStreams**](TopologySynchronizationApi.md#GetTopologySynchronizationStreams) | **Get** /synchronization/topology/streams | Overview of the topology synchronization streams



## GetTopologySynchronizationStreamById

> TopologyStreamListItem GetTopologySynchronizationStreamById(ctx).Identifier(identifier).IdentifierType(identifierType).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySynchronizationApi.GetTopologySynchronizationStreamById(context.Background()).Identifier(identifier).IdentifierType(identifierType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationApi.GetTopologySynchronizationStreamById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreamById`: TopologyStreamListItem
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationApi.GetTopologySynchronizationStreamById`: %v\n", resp)
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

[**TopologyStreamListItem**](TopologyStreamListItem.md)

### Authorization

No authorization required

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySynchronizationApi.GetTopologySynchronizationStreamStatusById(context.Background()).Identifier(identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationApi.GetTopologySynchronizationStreamStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreamStatusById`: TopologyStreamMetrics
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationApi.GetTopologySynchronizationStreamStatusById`: %v\n", resp)
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

No authorization required

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySynchronizationApi.GetTopologySynchronizationStreams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySynchronizationApi.GetTopologySynchronizationStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologySynchronizationStreams`: TopologyStreamList
    fmt.Fprintf(os.Stdout, "Response from `TopologySynchronizationApi.GetTopologySynchronizationStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologySynchronizationStreamsRequest struct via the builder pattern


### Return type

[**TopologyStreamList**](TopologyStreamList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

