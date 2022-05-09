# \AnomalyFeedbackApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectAnomalyFeedback**](AnomalyFeedbackApi.md#CollectAnomalyFeedback) | **Get** /anomaly-feedback | Collect feedback on anomalies



## CollectAnomalyFeedback

> []FeedbackWithContext CollectAnomalyFeedback(ctx).StartTime(startTime).EndTime(endTime).History(history).Execute()

Collect feedback on anomalies



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
    startTime := int64(789) // int64 | 
    endTime := int64(789) // int64 |  (optional)
    history := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyFeedbackApi.CollectAnomalyFeedback(context.Background()).StartTime(startTime).EndTime(endTime).History(history).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyFeedbackApi.CollectAnomalyFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectAnomalyFeedback`: []FeedbackWithContext
    fmt.Fprintf(os.Stdout, "Response from `AnomalyFeedbackApi.CollectAnomalyFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectAnomalyFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **history** | **int64** |  | 

### Return type

[**[]FeedbackWithContext**](FeedbackWithContext.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
