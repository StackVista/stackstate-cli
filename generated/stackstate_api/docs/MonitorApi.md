# \MonitorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitor**](MonitorApi.md#DeleteMonitor) | **Delete** /monitors/{monitorIdOrUrn} | Delete a monitor
[**GetAllMonitors**](MonitorApi.md#GetAllMonitors) | **Get** /monitors | List monitors
[**GetMonitor**](MonitorApi.md#GetMonitor) | **Get** /monitors/{monitorIdOrUrn} | Get a monitor
[**GetMonitorWithStatus**](MonitorApi.md#GetMonitorWithStatus) | **Get** /monitors/{monitorIdOrUrn}/status | Get a monitor with stream information
[**RunMonitor**](MonitorApi.md#RunMonitor) | **Post** /monitors/{monitorIdOrUrn}/run | Run a monitor



## DeleteMonitor

> DeleteMonitor(ctx, monitorIdOrUrn).Execute()

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
    monitorIdOrUrn := "monitorIdOrUrn_example" // string | The id or identifier (urn) of a monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.DeleteMonitor(context.Background(), monitorIdOrUrn).Execute()
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
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetAllMonitors(context.Background()).Execute()
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

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> Monitor GetMonitor(ctx, monitorIdOrUrn).Execute()

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
    monitorIdOrUrn := "monitorIdOrUrn_example" // string | The id or identifier (urn) of a monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetMonitor(context.Background(), monitorIdOrUrn).Execute()
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
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Monitor**](Monitor.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorWithStatus

> MonitorStatus GetMonitorWithStatus(ctx, monitorIdOrUrn).Execute()

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
    monitorIdOrUrn := "monitorIdOrUrn_example" // string | The id or identifier (urn) of a monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetMonitorWithStatus(context.Background(), monitorIdOrUrn).Execute()
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
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorWithStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorStatus**](MonitorStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunMonitor

> MonitorRunResult RunMonitor(ctx, monitorIdOrUrn).DryRun(dryRun).Execute()

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
    monitorIdOrUrn := "monitorIdOrUrn_example" // string | The id or identifier (urn) of a monitor
    dryRun := true // bool | If set, the topology state will not be modified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.RunMonitor(context.Background(), monitorIdOrUrn).DryRun(dryRun).Execute()
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
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor |

### Other Parameters

Other parameters are passed through a pointer to a apiRunMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | If set, the topology state will not be modified |

### Return type

[**MonitorRunResult**](MonitorRunResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

