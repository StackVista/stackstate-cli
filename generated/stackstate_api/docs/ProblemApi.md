# \ProblemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProblemCausingEvents**](ProblemApi.md#GetProblemCausingEvents) | **Get** /problems/{problemId}/causing-events | List possible events which led to the problem



## GetProblemCausingEvents

> GetCausingEventsResult GetProblemCausingEvents(ctx, problemId).TopologyTime(topologyTime).Limit(limit).Execute()

List possible events which led to the problem



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
    problemId := int64(789) // int64 | The problem id number.
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried for current time. (optional)
    limit := int32(56) // int32 | Maximum number of resources to be returned in result. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProblemApi.GetProblemCausingEvents(context.Background(), problemId).TopologyTime(topologyTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetProblemCausingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProblemCausingEvents`: GetCausingEventsResult
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetProblemCausingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **int64** | The problem id number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemCausingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried for current time. | 
 **limit** | **int32** | Maximum number of resources to be returned in result. | 

### Return type

[**GetCausingEventsResult**](GetCausingEventsResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

