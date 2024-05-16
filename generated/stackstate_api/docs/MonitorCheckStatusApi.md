# \MonitorCheckStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitorCheckStatus**](MonitorCheckStatusAPI.md#GetMonitorCheckStatus) | **Get** /monitor/checkStatus/{id} | Get a monitor check status
[**GetMonitorCheckStatusHealthHistory**](MonitorCheckStatusAPI.md#GetMonitorCheckStatusHealthHistory) | **Get** /monitor/checkStatus/{id}/healthHistory | Get a monitor check health history
[**GetMonitorCheckStatusRelatedFailures**](MonitorCheckStatusAPI.md#GetMonitorCheckStatusRelatedFailures) | **Get** /monitor/checkStatus/{id}/relatedFailures | Get a monitor check related failures



## GetMonitorCheckStatus

> MonitorCheckStatus GetMonitorCheckStatus(ctx, id).TopologyTime(topologyTime).Execute()

Get a monitor check status



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
    id := int64(789) // int64 | The id of a monitor check status
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusAPI.GetMonitorCheckStatus(context.Background(), id).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusAPI.GetMonitorCheckStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatus`: MonitorCheckStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusAPI.GetMonitorCheckStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MonitorCheckStatus**](MonitorCheckStatus.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorCheckStatusHealthHistory

> MonitorCheckStatusHealthHistory GetMonitorCheckStatusHealthHistory(ctx, id).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()

Get a monitor check health history



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
    id := int64(789) // int64 | The id of a monitor check status
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusAPI.GetMonitorCheckStatusHealthHistory(context.Background(), id).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusAPI.GetMonitorCheckStatusHealthHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatusHealthHistory`: MonitorCheckStatusHealthHistory
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusAPI.GetMonitorCheckStatusHealthHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusHealthHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MonitorCheckStatusHealthHistory**](MonitorCheckStatusHealthHistory.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorCheckStatusRelatedFailures

> MonitorCheckStatusRelatedFailures GetMonitorCheckStatusRelatedFailures(ctx, id).TopologyTime(topologyTime).Execute()

Get a monitor check related failures



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
    id := int64(789) // int64 | The id of a monitor check status
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusAPI.GetMonitorCheckStatusRelatedFailures(context.Background(), id).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusAPI.GetMonitorCheckStatusRelatedFailures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatusRelatedFailures`: MonitorCheckStatusRelatedFailures
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusAPI.GetMonitorCheckStatusRelatedFailures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusRelatedFailuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MonitorCheckStatusRelatedFailures**](MonitorCheckStatusRelatedFailures.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

