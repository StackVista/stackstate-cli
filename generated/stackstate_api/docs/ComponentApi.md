# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentHealthHistory**](ComponentApi.md#GetComponentHealthHistory) | **Get** /components/{componentIdOrUrn}/healthHistory | Get a component health history



## GetComponentHealthHistory

> ComponentHealthHistory GetComponentHealthHistory(ctx, componentIdOrUrn).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()

Get a component health history



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
    componentIdOrUrn := "componentIdOrUrn_example" // string | The id or identifier (urn) of a component
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentHealthHistory(context.Background(), componentIdOrUrn).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentHealthHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentHealthHistory`: ComponentHealthHistory
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentHealthHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrUrn** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentHealthHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**ComponentHealthHistory**](ComponentHealthHistory.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

