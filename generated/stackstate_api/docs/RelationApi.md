# \RelationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationCheck**](RelationApi.md#CreateRelationCheck) | **Post** /relations/{relationId}/checks | Create the Checks on the relation
[**CreateRelationStream**](RelationApi.md#CreateRelationStream) | **Post** /relations/{relationId}/streams | Create the Stream on the relation
[**DeleteRelationCheck**](RelationApi.md#DeleteRelationCheck) | **Delete** /relations/{relationId}/checks/{checkId} | Delete the check from the relation
[**DeleteRelationStream**](RelationApi.md#DeleteRelationStream) | **Delete** /relations/{relationId}/streams/{streamId} | Delete the stream from the relation
[**GetRelationCheck**](RelationApi.md#GetRelationCheck) | **Get** /relations/{relationId}/checks/{checkId} | Get the check from the relation
[**GetRelationChecks**](RelationApi.md#GetRelationChecks) | **Get** /relations/{relationId}/checks | List all checks from the relation
[**GetRelationStream**](RelationApi.md#GetRelationStream) | **Get** /relations/{relationId}/streams/{streamId} | Get the stream from the relation
[**GetRelationStreams**](RelationApi.md#GetRelationStreams) | **Get** /relations/{relationId}/streams | List all streams from the relation
[**PutRelationCheck**](RelationApi.md#PutRelationCheck) | **Put** /relations/{relationId}/checks/{checkId} | Update the check on the relation
[**PutRelationStream**](RelationApi.md#PutRelationStream) | **Put** /relations/{relationId}/streams/{streamId} | Update the stream on the relation



## CreateRelationCheck

> Check CreateRelationCheck(ctx, relationId).Check(check).Execute()

Create the Checks on the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    check := *openapiclient.NewCheck("Type_example", []openapiclient.Argument{openapiclient.Argument{ArgumentAnomalyDirectionVal: openapiclient.NewArgumentAnomalyDirectionVal("Type_example", int64(123), openapiclient.AnomalyDirection("RISE"))}}, int64(123), "Name_example", false) // Check | Single Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.CreateRelationCheck(context.Background(), relationId).Check(check).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.CreateRelationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelationCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.CreateRelationCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationCheckRequest struct via the builder pattern


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


## CreateRelationStream

> TelemetryStreamDefinition CreateRelationStream(ctx, relationId).DataStream(dataStream).Execute()

Create the Stream on the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    dataStream := openapiclient.DataStream{EventStream: openapiclient.NewEventStream("Type_example", int64(123), openapiclient.DataType("METRICS"), "Name_example", *openapiclient.NewEventTelemetryQuery("Type_example", []openapiclient.TelemetryQueryCondition{*openapiclient.NewTelemetryQueryCondition("Key_example", "Value_example")}), false)} // DataStream | Single telemetry stream definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.CreateRelationStream(context.Background(), relationId).DataStream(dataStream).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.CreateRelationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelationStream`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.CreateRelationStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationStreamRequest struct via the builder pattern


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


## DeleteRelationCheck

> DeleteRelationCheck(ctx, relationId, checkId).Execute()

Delete the check from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    checkId := int64(789) // int64 | The Identifier of a Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.DeleteRelationCheck(context.Background(), relationId, checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.DeleteRelationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelationCheckRequest struct via the builder pattern


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


## DeleteRelationStream

> DeleteRelationStream(ctx, relationId, streamId).Execute()

Delete the stream from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    streamId := int64(789) // int64 | The Identifier of a stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.DeleteRelationStream(context.Background(), relationId, streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.DeleteRelationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelationStreamRequest struct via the builder pattern


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


## GetRelationCheck

> Check GetRelationCheck(ctx, relationId, checkId).QueryTime(queryTime).Execute()

Get the check from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode
    checkId := int64(789) // int64 | The Identifier of a Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.GetRelationCheck(context.Background(), relationId, checkId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetRelationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetRelationCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationCheckRequest struct via the builder pattern


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


## GetRelationChecks

> []Check GetRelationChecks(ctx, relationId).QueryTime(queryTime).Execute()

List all checks from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.GetRelationChecks(context.Background(), relationId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetRelationChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationChecks`: []Check
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetRelationChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationChecksRequest struct via the builder pattern


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


## GetRelationStream

> TelemetryStreamDefinition GetRelationStream(ctx, relationId, streamId).QueryTime(queryTime).Execute()

Get the stream from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode
    streamId := int64(789) // int64 | The Identifier of a stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.GetRelationStream(context.Background(), relationId, streamId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetRelationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationStream`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetRelationStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationStreamRequest struct via the builder pattern


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


## GetRelationStreams

> []TelemetryStreamDefinition GetRelationStreams(ctx, relationId).QueryTime(queryTime).Execute()

List all streams from the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    queryTime := int32(56) // int32 | A Data point for a query either point from timeline or 'now' if livemode

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.GetRelationStreams(context.Background(), relationId).QueryTime(queryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetRelationStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationStreams`: []TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetRelationStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationStreamsRequest struct via the builder pattern


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


## PutRelationCheck

> Check PutRelationCheck(ctx, relationId, checkId).Check(check).Execute()

Update the check on the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    checkId := int64(789) // int64 | The Identifier of a Check
    check := *openapiclient.NewCheck("Type_example", []openapiclient.Argument{openapiclient.Argument{ArgumentAnomalyDirectionVal: openapiclient.NewArgumentAnomalyDirectionVal("Type_example", int64(123), openapiclient.AnomalyDirection("RISE"))}}, int64(123), "Name_example", false) // Check | Single Check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.PutRelationCheck(context.Background(), relationId, checkId).Check(check).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.PutRelationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutRelationCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.PutRelationCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**checkId** | **int64** | The Identifier of a Check | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRelationCheckRequest struct via the builder pattern


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


## PutRelationStream

> TelemetryStreamDefinition PutRelationStream(ctx, relationId, streamId).DataStream(dataStream).Execute()

Update the stream on the relation



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
    relationId := int64(789) // int64 | The Identifier of a relation
    streamId := int64(789) // int64 | The Identifier of a stream
    dataStream := openapiclient.DataStream{EventStream: openapiclient.NewEventStream("Type_example", int64(123), openapiclient.DataType("METRICS"), "Name_example", *openapiclient.NewEventTelemetryQuery("Type_example", []openapiclient.TelemetryQueryCondition{*openapiclient.NewTelemetryQueryCondition("Key_example", "Value_example")}), false)} // DataStream | Single telemetry stream definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationApi.PutRelationStream(context.Background(), relationId, streamId).DataStream(dataStream).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.PutRelationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutRelationStream`: TelemetryStreamDefinition
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.PutRelationStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationId** | **int64** | The Identifier of a relation | 
**streamId** | **int64** | The Identifier of a stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRelationStreamRequest struct via the builder pattern


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

