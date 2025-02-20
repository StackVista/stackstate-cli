# \TimelineApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimelineSummary**](TimelineApi.md#GetTimelineSummary) | **Post** /timeline/summary | Timeline summary



## GetTimelineSummary

> TimelineSummary GetTimelineSummary(ctx).TimelineSummaryRequest(timelineSummaryRequest).Execute()

Timeline summary



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
    timelineSummaryRequest := *openapiclient.NewTimelineSummaryRequest(openapiclient.TimelineSummaryRequest_arguments{ComponentViewArguments: openapiclient.NewComponentViewArguments("Type_example", "ComponentIdentifier_example")}, int32(123), int32(123)) // TimelineSummaryRequest | Request for event summary and aggregated health over time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimelineApi.GetTimelineSummary(context.Background()).TimelineSummaryRequest(timelineSummaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimelineApi.GetTimelineSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimelineSummary`: TimelineSummary
    fmt.Fprintf(os.Stdout, "Response from `TimelineApi.GetTimelineSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimelineSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timelineSummaryRequest** | [**TimelineSummaryRequest**](TimelineSummaryRequest.md) | Request for event summary and aggregated health over time | 

### Return type

[**TimelineSummary**](TimelineSummary.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

