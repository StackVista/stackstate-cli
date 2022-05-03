# \MonitorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitor**](MonitorApi.md#DeleteMonitor) | **Delete** /monitor/{monitorId} | Delete a monitor
[**DryRunMonitor**](MonitorApi.md#DryRunMonitor) | **Post** /monitor/{monitorId}/dryRun | Dry run a monitor and show a result
[**GetAllMonitors**](MonitorApi.md#GetAllMonitors) | **Get** /monitors | List monitors
[**GetMonitor**](MonitorApi.md#GetMonitor) | **Get** /monitor/{monitorId} | Get a monitor
[**GetMonitorWithStatus**](MonitorApi.md#GetMonitorWithStatus) | **Get** /monitor/{monitorId}/status | Get a monitor with stream information
[**RunMonitor**](MonitorApi.md#RunMonitor) | **Post** /monitor/{monitorId}/run | Run a monitor



## DeleteMonitor

> DeleteMonitor(ctx, monitorId).Execute()

Delete a monitor



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
    monitorId := int64(789) // int64 | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.DeleteMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.DeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunMonitor

> MonitorRunResult DryRunMonitor(ctx, monitorId).Execute()

Dry run a monitor and show a result



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
    monitorId := int64(789) // int64 | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.DryRunMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.DryRunMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DryRunMonitor`: MonitorRunResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.DryRunMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDryRunMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorRunResult**](MonitorRunResult.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMonitors

> MonitorList GetAllMonitors(ctx).Execute()

List monitors



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.GetAllMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetAllMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMonitors`: MonitorList
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetAllMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMonitorsRequest struct via the builder pattern


### Return type

[**MonitorList**](MonitorList.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> Monitor GetMonitor(ctx, monitorId).Execute()

Get a monitor



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
    monitorId := int64(789) // int64 | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.GetMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Monitor**](Monitor.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorWithStatus

> MonitorStatus GetMonitorWithStatus(ctx, monitorId).Execute()

Get a monitor with stream information



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
    monitorId := int64(789) // int64 | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.GetMonitorWithStatus(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetMonitorWithStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorWithStatus`: MonitorStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetMonitorWithStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorWithStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorStatus**](MonitorStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunMonitor

> MonitorRunResult RunMonitor(ctx, monitorId).Execute()

Run a monitor



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
    monitorId := int64(789) // int64 | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.RunMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.RunMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMonitor`: MonitorRunResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.RunMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorRunResult**](MonitorRunResult.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

