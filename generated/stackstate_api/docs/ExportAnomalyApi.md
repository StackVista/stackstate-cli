# \ExportAnomalyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAnomaly**](ExportAnomalyApi.md#ExportAnomaly) | **Get** /anomalies/export | Export anomalies with metric history and feedback



## ExportAnomaly

> []AnomalyWithContext ExportAnomaly(ctx).StartTime(startTime).Feedback(feedback).EndTime(endTime).History(history).Execute()

Export anomalies with metric history and feedback



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
    feedback := "feedback_example" // string | 
    endTime := int64(789) // int64 |  (optional)
    history := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAnomalyApi.ExportAnomaly(context.Background()).StartTime(startTime).Feedback(feedback).EndTime(endTime).History(history).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAnomalyApi.ExportAnomaly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAnomaly`: []AnomalyWithContext
    fmt.Fprintf(os.Stdout, "Response from `ExportAnomalyApi.ExportAnomaly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAnomalyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **feedback** | **string** |  | 
 **endTime** | **int64** |  | 
 **history** | **int64** |  | 

### Return type

[**[]AnomalyWithContext**](AnomalyWithContext.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

