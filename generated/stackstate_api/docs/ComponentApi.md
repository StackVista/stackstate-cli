# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentsComponentIdStreamsLatestGet**](ComponentApi.md#ComponentsComponentIdStreamsLatestGet) | **Get** /components/{componentId}/streams/latest | Get the latests metrics



## ComponentsComponentIdStreamsLatestGet

> TelemetryLatestSnapshotsResponse ComponentsComponentIdStreamsLatestGet(ctx, componentId).QueryTime(queryTime).StreamIds(streamIds).Execute()

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
    resp, r, err := apiClient.ComponentApi.ComponentsComponentIdStreamsLatestGet(context.Background(), componentId).QueryTime(queryTime).StreamIds(streamIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentsComponentIdStreamsLatestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentsComponentIdStreamsLatestGet`: TelemetryLatestSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentsComponentIdStreamsLatestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **int64** | The Identifier of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentsComponentIdStreamsLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryTime** | **int32** | A Data point for a query either point from timeline or &#39;now&#39; if livemode | 
 **streamIds** | **[]int64** | Ids of streams to query for | 

### Return type

[**TelemetryLatestSnapshotsResponse**](TelemetryLatestSnapshotsResponse.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

