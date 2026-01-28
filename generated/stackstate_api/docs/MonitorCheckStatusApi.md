# \MonitorCheckStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitorCheckStatus**](MonitorCheckStatusApi.md#GetMonitorCheckStatus) | **Get** /monitor/checkStatus/{identifier} | Get a monitor check status
[**GetMonitorCheckStatusHealthHistory**](MonitorCheckStatusApi.md#GetMonitorCheckStatusHealthHistory) | **Get** /monitor/checkStatus/{identifier}/healthHistory | Get a monitor check health history
[**GetMonitorCheckStatusRelatedFailures**](MonitorCheckStatusApi.md#GetMonitorCheckStatusRelatedFailures) | **Get** /monitor/checkStatus/{identifier}/relatedFailures | Get a monitor check related failures



## GetMonitorCheckStatus

> MonitorCheckStatus GetMonitorCheckStatus(ctx, identifier).TopologyTime(topologyTime).Execute()

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
    identifier := "identifier_example" // string | The id of a monitor check status
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusApi.GetMonitorCheckStatus(context.Background(), identifier).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusApi.GetMonitorCheckStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatus`: MonitorCheckStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusApi.GetMonitorCheckStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MonitorCheckStatus**](MonitorCheckStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorCheckStatusHealthHistory

> MonitorCheckStatusHealthHistory GetMonitorCheckStatusHealthHistory(ctx, identifier).StartTime(startTime).EndTime(endTime).Execute()

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
    identifier := "identifier_example" // string | The id of a monitor check status
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusApi.GetMonitorCheckStatusHealthHistory(context.Background(), identifier).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusApi.GetMonitorCheckStatusHealthHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatusHealthHistory`: MonitorCheckStatusHealthHistory
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusApi.GetMonitorCheckStatusHealthHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusHealthHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 

### Return type

[**MonitorCheckStatusHealthHistory**](MonitorCheckStatusHealthHistory.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorCheckStatusRelatedFailures

> MonitorCheckStatusRelatedFailures GetMonitorCheckStatusRelatedFailures(ctx, identifier).TopologyTime(topologyTime).Execute()

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
    identifier := "identifier_example" // string | The id of a monitor check status
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorCheckStatusApi.GetMonitorCheckStatusRelatedFailures(context.Background(), identifier).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorCheckStatusApi.GetMonitorCheckStatusRelatedFailures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStatusRelatedFailures`: MonitorCheckStatusRelatedFailures
    fmt.Fprintf(os.Stdout, "Response from `MonitorCheckStatusApi.GetMonitorCheckStatusRelatedFailures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The id of a monitor check status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatusRelatedFailuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MonitorCheckStatusRelatedFailures**](MonitorCheckStatusRelatedFailures.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

