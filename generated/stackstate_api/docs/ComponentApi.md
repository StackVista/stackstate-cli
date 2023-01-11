# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentCheck**](ComponentApi.md#GetComponentCheck) | **Get** /components/{componentId}/checks/{checkId} | Get the check from the component
[**GetComponentChecks**](ComponentApi.md#GetComponentChecks) | **Get** /components/{componentId}/checks | List all checks from the component
[**GetComponentStream**](ComponentApi.md#GetComponentStream) | **Get** /components/{componentId}/streams/{streamId} | Get the stream from the component
[**GetComponentStreams**](ComponentApi.md#GetComponentStreams) | **Get** /components/{componentId}/streams | List all streams from the component
[**GetLatestMetrics**](ComponentApi.md#GetLatestMetrics) | **Get** /components/{componentId}/streams/latest | Get the latests metrics



## GetComponentCheck

> Check GetComponentCheck(ctx, componentId, checkId).QueryTime(queryTime).Execute()

Get the check from the component



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
    componentId := int64(789) // int64 | The Identifier of a component
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode
    checkId := int64(789) // int64 | The Identifier of a Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentCheck(context.Background(), componentId, checkId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 


### Return type

[**Check**](Check.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentChecks

> []Check GetComponentChecks(ctx, componentId).QueryTime(queryTime).Execute()

List all checks from the component



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
    componentId := int64(789) // int64 | The Identifier of a component
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentChecks(context.Background(), componentId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentChecks`: []Check
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 

### Return type

[**[]Check**](Check.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentStream

> TelemetryStreamDefinition GetComponentStream(ctx, componentId, streamId).QueryTime(queryTime).Execute()

Get the stream from the component



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
    componentId := int64(789) // int64 | The Identifier of a component
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode
    streamId := int64(789) // int64 | The Identifier of a stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentStream(context.Background(), componentId, streamId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentStream`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 


### Return type

[**TelemetryStreamDefinition**](TelemetryStreamDefinition.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentStreams

> []TelemetryStreamDefinition GetComponentStreams(ctx, componentId).QueryTime(queryTime).Execute()

List all streams from the component



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
    componentId := int64(789) // int64 | The Identifier of a component
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentStreams(context.Background(), componentId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentStreams`: []TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 

### Return type

[**[]TelemetryStreamDefinition**](TelemetryStreamDefinition.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestMetrics

> TelemetryLatestSnapshotsResponse GetLatestMetrics(ctx, componentId).QueryTime(queryTime).StreamIds(streamIds).Execute()

Get the latests metrics



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
    componentId := int64(789) // int64 | The Identifier of a component
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode
    streamIds := []int64{int64(123)} // []int64 | Ids of streams to query for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetLatestMetrics(context.Background(), componentId).QueryTime(queryTime).StreamIds(streamIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetLatestMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestMetrics`: TelemetryLatestSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetLatestMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 
 **streamIds** | **[]int64** | Ids of streams to query for | 

### Return type

[**TelemetryLatestSnapshotsResponse**](TelemetryLatestSnapshotsResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

