# \RelationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelationCheck**](RelationApi.md#GetRelationCheck) | **Get** /relations/{relationId}/checks/{checkId} | Get the check from the relation
[**GetRelationChecks**](RelationApi.md#GetRelationChecks) | **Get** /relations/{relationId}/checks | List all checks from the relation
[**GetRelationStream**](RelationApi.md#GetRelationStream) | **Get** /relations/{relationId}/streams/{streamId} | Get the stream from the relation
[**GetRelationStreams**](RelationApi.md#GetRelationStreams) | **Get** /relations/{relationId}/streams | List all streams from the relation



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

