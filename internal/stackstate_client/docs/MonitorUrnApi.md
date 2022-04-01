# \MonitorUrnApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitorByURN**](MonitorUrnApi.md#DeleteMonitorByURN) | **Delete** /monitorUrn/{monitorUrnId} | Delete a monitor
[**DryRunMonitorByURN**](MonitorUrnApi.md#DryRunMonitorByURN) | **Post** /monitorUrn/{monitorUrnId}/dryRun | Dry run a monitor and show a result
[**GetMonitorByURN**](MonitorUrnApi.md#GetMonitorByURN) | **Get** /monitorUrn/{monitorUrnId} | Get a monitor
[**RunMonitorByURN**](MonitorUrnApi.md#RunMonitorByURN) | **Post** /monitorUrn/{monitorUrnId}/run | Run a monitor



## DeleteMonitorByURN

> string DeleteMonitorByURN(ctx, monitorUrnId).Execute()

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
    monitorUrnId := "monitorUrnId_example" // string | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorUrnApi.DeleteMonitorByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.DeleteMonitorByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitorByURN`: string
    fmt.Fprintf(os.Stdout, "Response from `MonitorUrnApi.DeleteMonitorByURN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorUrnId** | **string** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorByURNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunMonitorByURN

> MonitorRunResult DryRunMonitorByURN(ctx, monitorUrnId).Execute()

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
    monitorUrnId := "monitorUrnId_example" // string | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorUrnApi.DryRunMonitorByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.DryRunMonitorByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DryRunMonitorByURN`: MonitorRunResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorUrnApi.DryRunMonitorByURN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorUrnId** | **string** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDryRunMonitorByURNRequest struct via the builder pattern


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


## GetMonitorByURN

> Monitor GetMonitorByURN(ctx, monitorUrnId).Execute()

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
    monitorUrnId := "monitorUrnId_example" // string | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorUrnApi.GetMonitorByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.GetMonitorByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorByURN`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUrnApi.GetMonitorByURN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorUrnId** | **string** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorByURNRequest struct via the builder pattern


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


## RunMonitorByURN

> MonitorRunResult RunMonitorByURN(ctx, monitorUrnId).Execute()

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
    monitorUrnId := "monitorUrnId_example" // string | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorUrnApi.RunMonitorByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.RunMonitorByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMonitorByURN`: MonitorRunResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorUrnApi.RunMonitorByURN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorUrnId** | **string** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunMonitorByURNRequest struct via the builder pattern


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

