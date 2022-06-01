# \MonitorUrnApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitorByURN**](MonitorUrnApi.md#DeleteMonitorByURN) | **Delete** /monitorUrn/{monitorUrnId} | Delete a monitor
[**DryRunMonitorByURN**](MonitorUrnApi.md#DryRunMonitorByURN) | **Post** /monitorUrn/{monitorUrnId}/dryRun | Dry run a monitor and show a result
[**GetMonitorByURN**](MonitorUrnApi.md#GetMonitorByURN) | **Get** /monitorUrn/{monitorUrnId} | Get a monitor
[**GetMonitorWithStatusByURN**](MonitorUrnApi.md#GetMonitorWithStatusByURN) | **Get** /monitorUrn/{monitorUrnId}/status | Get a monitor with stream information
[**RunMonitorByURN**](MonitorUrnApi.md#RunMonitorByURN) | **Post** /monitorUrn/{monitorUrnId}/run | Run a monitor



## DeleteMonitorByURN

> DeleteMonitorByURN(ctx, monitorUrnId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUrnApi.DeleteMonitorByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.DeleteMonitorByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUrnApi.DryRunMonitorByURN(context.Background(), monitorUrnId).Execute()
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

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUrnApi.GetMonitorByURN(context.Background(), monitorUrnId).Execute()
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

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorWithStatusByURN

> MonitorStatus GetMonitorWithStatusByURN(ctx, monitorUrnId).Execute()

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
    monitorUrnId := "monitorUrnId_example" // string | The identifier of a monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUrnApi.GetMonitorWithStatusByURN(context.Background(), monitorUrnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUrnApi.GetMonitorWithStatusByURN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorWithStatusByURN`: MonitorStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorUrnApi.GetMonitorWithStatusByURN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorUrnId** | **string** | The identifier of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorWithStatusByURNRequest struct via the builder pattern


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUrnApi.RunMonitorByURN(context.Background(), monitorUrnId).Execute()
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

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

