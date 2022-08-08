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
    startTime := int64(789) // int64 | Beginning of timerange of to be exported anomalies.  Timestamp in unix millis.
    feedback := "feedback_example" // string | Type of filtering to do on feedback.  Filtering on feedback is currently mandatory, with only the 'present' value being supporeted (feedback is available).
    endTime := int64(789) // int64 | End of timerange of to be exported anomalies.  Timestamp in unix millis. (optional)
    history := int64(789) // int64 | Amount of historic data, leading up to the anomaly, to be exported.  Duration in unix millis. (optional)

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
 **startTime** | **int64** | Beginning of timerange of to be exported anomalies.  Timestamp in unix millis. | 
 **feedback** | **string** | Type of filtering to do on feedback.  Filtering on feedback is currently mandatory, with only the &#39;present&#39; value being supporeted (feedback is available). | 
 **endTime** | **int64** | End of timerange of to be exported anomalies.  Timestamp in unix millis. | 
 **history** | **int64** | Amount of historic data, leading up to the anomaly, to be exported.  Duration in unix millis. | 

### Return type

[**[]AnomalyWithContext**](AnomalyWithContext.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

