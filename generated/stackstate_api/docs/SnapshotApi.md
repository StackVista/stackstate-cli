# \SnapshotApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuerySnapshot**](SnapshotApi.md#QuerySnapshot) | **Post** /snapshot | Query topology snapshot



## QuerySnapshot

> QuerySnapshotResult QuerySnapshot(ctx).ViewSnapshotRequest(viewSnapshotRequest).TimeoutMs(timeoutMs).Execute()

Query topology snapshot

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
    viewSnapshotRequest := *openapiclient.NewViewSnapshotRequest("Type_example", "Query_example", "0.0.1", *openapiclient.NewQueryMetadata(false, false, int64(123), false, false, false, false, false, false, false)) // ViewSnapshotRequest | Request body for querying a topology snapshot
    timeoutMs := int64(789) // int64 | Query timeout in milliseconds (optional) (default to 30000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotApi.QuerySnapshot(context.Background()).ViewSnapshotRequest(viewSnapshotRequest).TimeoutMs(timeoutMs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.QuerySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySnapshot`: QuerySnapshotResult
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.QuerySnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuerySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewSnapshotRequest** | [**ViewSnapshotRequest**](ViewSnapshotRequest.md) | Request body for querying a topology snapshot | 
 **timeoutMs** | **int64** | Query timeout in milliseconds | [default to 30000]

### Return type

[**QuerySnapshotResult**](QuerySnapshotResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

