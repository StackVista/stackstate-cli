# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComponentCheck**](ComponentApi.md#CreateComponentCheck) | **Post** /components/{componentId}/checks | Create the Checks on the component
[**CreateComponentStreams**](ComponentApi.md#CreateComponentStreams) | **Post** /components/{componentId}/streams | Create the Stream on the component
[**DeleteComponentCheck**](ComponentApi.md#DeleteComponentCheck) | **Delete** /components/{componentId}/checks/{checkId} | Delete the check from the component
[**DeleteComponentStream**](ComponentApi.md#DeleteComponentStream) | **Delete** /components/{componentId}/streams/{streamId} | Delete the stream from the component
[**GetComponentCheck**](ComponentApi.md#GetComponentCheck) | **Get** /components/{componentId}/checks/{checkId} | Get the check from the component
[**GetComponentChecks**](ComponentApi.md#GetComponentChecks) | **Get** /components/{componentId}/checks | List all checks from the component
[**GetComponentStream**](ComponentApi.md#GetComponentStream) | **Get** /components/{componentId}/streams/{streamId} | Get the stream from the component
[**GetComponentStreams**](ComponentApi.md#GetComponentStreams) | **Get** /components/{componentId}/streams | List all streams from the component
[**GetLatestMetrics**](ComponentApi.md#GetLatestMetrics) | **Get** /components/{componentId}/streams/latest | Get the latests metrics
[**UpdateComponentCheck**](ComponentApi.md#UpdateComponentCheck) | **Put** /components/{componentId}/checks/{checkId} | Update the check on the component
[**UpdateComponentStream**](ComponentApi.md#UpdateComponentStream) | **Put** /components/{componentId}/streams/{streamId} | Update the stream on the component



## CreateComponentCheck

> Check CreateComponentCheck(ctx, componentId).Check(check).Execute()

Create the Checks on the component



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
    check := *openapiclient.NewCheck("Type_example", []openapiclient.Argument{openapiclient.Argument{ArgumentAnomalyDirectionVal: openapiclient.NewArgumentAnomalyDirectionVal("Type_example", int64(123), openapiclient.AnomalyDirection("RISE"))}}, int64(123), "Name_example", false) // Check | Single Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.CreateComponentCheck(context.Background(), componentId).Check(check).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.CreateComponentCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComponentCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.CreateComponentCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateComponentCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **check** | [**Check**](Check.md) | Single Check | 

### Return type

[**Check**](Check.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComponentStreams

> TelemetryStreamDefinition CreateComponentStreams(ctx, componentId).DataStream(dataStream).Execute()

Create the Stream on the component



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
    dataStream := openapiclient.DataStream{EventStream: openapiclient.NewEventStream("Type_example", int64(123), openapiclient.DataType("METRICS"), "Name_example", *openapiclient.NewEventTelemetryQuery("Type_example", []openapiclient.TelemetryQueryCondition{*openapiclient.NewTelemetryQueryCondition("Key_example", "Value_example")}), false)} // DataStream | Single telemetry stream definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.CreateComponentStreams(context.Background(), componentId).DataStream(dataStream).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.CreateComponentStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComponentStreams`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.CreateComponentStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateComponentStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataStream** | [**DataStream**](DataStream.md) | Single telemetry stream definition | 

### Return type

[**TelemetryStreamDefinition**](TelemetryStreamDefinition.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComponentCheck

> DeleteComponentCheck(ctx, componentId, checkId).Execute()

Delete the check from the component



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
    checkId := int64(789) // int64 | The Identifier of a Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.DeleteComponentCheck(context.Background(), componentId, checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.DeleteComponentCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentCheckRequest struct via the builder pattern


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


## DeleteComponentStream

> DeleteComponentStream(ctx, componentId, streamId).Execute()

Delete the stream from the component



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
    streamId := int64(789) // int64 | The Identifier of a stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.DeleteComponentStream(context.Background(), componentId, streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.DeleteComponentStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentStreamRequest struct via the builder pattern


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


## UpdateComponentCheck

> Check UpdateComponentCheck(ctx, componentId, checkId).Check(check).Execute()

Update the check on the component



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
    checkId := int64(789) // int64 | The Identifier of a Check
    check := *openapiclient.NewCheck("Type_example", []openapiclient.Argument{openapiclient.Argument{ArgumentAnomalyDirectionVal: openapiclient.NewArgumentAnomalyDirectionVal("Type_example", int64(123), openapiclient.AnomalyDirection("RISE"))}}, int64(123), "Name_example", false) // Check | Single Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.UpdateComponentCheck(context.Background(), componentId, checkId).Check(check).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.UpdateComponentCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComponentCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.UpdateComponentCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **check** | [**Check**](Check.md) | Single Check | 

### Return type

[**Check**](Check.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponentStream

> TelemetryStreamDefinition UpdateComponentStream(ctx, componentId, streamId).DataStream(dataStream).Execute()

Update the stream on the component



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
    streamId := int64(789) // int64 | The Identifier of a stream
    dataStream := openapiclient.DataStream{EventStream: openapiclient.NewEventStream("Type_example", int64(123), openapiclient.DataType("METRICS"), "Name_example", *openapiclient.NewEventTelemetryQuery("Type_example", []openapiclient.TelemetryQueryCondition{*openapiclient.NewTelemetryQueryCondition("Key_example", "Value_example")}), false)} // DataStream | Single telemetry stream definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.UpdateComponentStream(context.Background(), componentId, streamId).DataStream(dataStream).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.UpdateComponentStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComponentStream`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.UpdateComponentStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataStream** | [**DataStream**](DataStream.md) | Single telemetry stream definition | 

### Return type

[**TelemetryStreamDefinition**](TelemetryStreamDefinition.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

