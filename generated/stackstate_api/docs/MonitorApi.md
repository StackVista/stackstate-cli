# \MonitorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitor**](MonitorApi.md#DeleteMonitor) | **Delete** /monitors/{monitorIdOrUrn} | Delete a monitor
[**GetAllMonitors**](MonitorApi.md#GetAllMonitors) | **Get** /monitors | List monitors
[**GetMonitor**](MonitorApi.md#GetMonitor) | **Get** /monitors/{monitorIdOrUrn} | Get a monitor
[**GetMonitorCheckStates**](MonitorApi.md#GetMonitorCheckStates) | **Get** /monitors/{monitorIdOrUrn}/checkStates | Get the check states for a monitor
[**GetMonitorWithStatus**](MonitorApi.md#GetMonitorWithStatus) | **Get** /monitors/{monitorIdOrUrn}/status | Get a monitor with stream information
[**GetMonitorsOverview**](MonitorApi.md#GetMonitorsOverview) | **Get** /monitors/overview | List monitors overview
[**LookupIdentifier**](MonitorApi.md#LookupIdentifier) | **Post** /monitors/identifierLookup | Multiple component identifier lookup
[**PatchMonitor**](MonitorApi.md#PatchMonitor) | **Patch** /monitors/{monitorIdOrUrn} | Update some monitor properties
[**PreviewMonitor**](MonitorApi.md#PreviewMonitor) | **Post** /monitors/{monitorIdOrUrn}/preview | Preview a monitor
[**PreviewMonitorCheckStates**](MonitorApi.md#PreviewMonitorCheckStates) | **Post** /monitors/{monitorIdOrUrn}/preview/checkStates | Preview a monitor
[**PublishHealthStates**](MonitorApi.md#PublishHealthStates) | **Post** /monitors/{monitorIdOrUrn}/publish | Post monitor health states
[**RunMonitor**](MonitorApi.md#RunMonitor) | **Post** /monitors/{monitorIdOrUrn}/run | Run a monitor
[**TestMonitorFunction**](MonitorApi.md#TestMonitorFunction) | **Post** /monitors/{monitorFunctionIdOrUrn}/test | Test a monitor
[**TestMonitorFunctionCheckStates**](MonitorApi.md#TestMonitorFunctionCheckStates) | **Post** /monitors/{monitorFunctionIdOrUrn}/test/checkStates | Test a monitor



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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorCheckStates

> MonitorCheckStates GetMonitorCheckStates(ctx, monitorIdOrUrn).HealthState(healthState).Limit(limit).Timestamp(timestamp).Execute()

Get the check states for a monitor



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
    healthState := openapiclient.HealthStateValue("UNINITIALIZED") // HealthStateValue | Health state of check states (optional)
    limit := int32(56) // int32 |  (optional)
    timestamp := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetMonitorCheckStates(context.Background(), monitorIdOrUrn).HealthState(healthState).Limit(limit).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetMonitorCheckStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorCheckStates`: MonitorCheckStates
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetMonitorCheckStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorCheckStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **healthState** | [**HealthStateValue**](HealthStateValue.md) | Health state of check states | 
 **limit** | **int32** |  | 
 **timestamp** | **int64** |  | 

### Return type

[**MonitorCheckStates**](MonitorCheckStates.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorWithStatus

> MonitorStatus GetMonitorWithStatus(ctx, monitorIdOrUrn).Timestamp(timestamp).Execute()

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
    timestamp := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetMonitorWithStatus(context.Background(), monitorIdOrUrn).Timestamp(timestamp).Execute()
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

 **timestamp** | **int64** |  | 

### Return type

[**MonitorStatus**](MonitorStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorsOverview

> MonitorOverviewList GetMonitorsOverview(ctx).Execute()

List monitors overview



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
    resp, r, err := apiClient.MonitorApi.GetMonitorsOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetMonitorsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorsOverview`: MonitorOverviewList
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetMonitorsOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsOverviewRequest struct via the builder pattern


### Return type

[**MonitorOverviewList**](MonitorOverviewList.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupIdentifier

> MonitorIdentifierSuggestions LookupIdentifier(ctx).MonitorIdentifierLookup(monitorIdentifierLookup).Execute()

Multiple component identifier lookup



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
    monitorIdentifierLookup := *openapiclient.NewMonitorIdentifierLookup("MetricQuery_example", "ComponentTypeName_example") // MonitorIdentifierLookup | Component type and metric query for identifier lookup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.LookupIdentifier(context.Background()).MonitorIdentifierLookup(monitorIdentifierLookup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.LookupIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupIdentifier`: MonitorIdentifierSuggestions
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.LookupIdentifier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitorIdentifierLookup** | [**MonitorIdentifierLookup**](MonitorIdentifierLookup.md) | Component type and metric query for identifier lookup | 

### Return type

[**MonitorIdentifierSuggestions**](MonitorIdentifierSuggestions.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMonitor

> Monitor PatchMonitor(ctx, monitorIdOrUrn).MonitorPatch(monitorPatch).Execute()

Update some monitor properties



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
    monitorPatch := *openapiclient.NewMonitorPatch() // MonitorPatch | Monitor base properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.PatchMonitor(context.Background(), monitorIdOrUrn).MonitorPatch(monitorPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.PatchMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.PatchMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorPatch** | [**MonitorPatch**](MonitorPatch.md) | Monitor base properties | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewMonitor

> MonitorPreviewResult PreviewMonitor(ctx, monitorIdOrUrn).MonitorPreview(monitorPreview).Execute()

Preview a monitor



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
    monitorPreview := *openapiclient.NewMonitorPreview() // MonitorPreview | Monitor overrides in order to run a preview

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.PreviewMonitor(context.Background(), monitorIdOrUrn).MonitorPreview(monitorPreview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.PreviewMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewMonitor`: MonitorPreviewResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.PreviewMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorPreview** | [**MonitorPreview**](MonitorPreview.md) | Monitor overrides in order to run a preview | 

### Return type

[**MonitorPreviewResult**](MonitorPreviewResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewMonitorCheckStates

> MonitorCheckStates PreviewMonitorCheckStates(ctx, monitorIdOrUrn).MonitorPreview(monitorPreview).HealthState(healthState).Limit(limit).Execute()

Preview a monitor



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
    monitorPreview := *openapiclient.NewMonitorPreview() // MonitorPreview | Monitor overrides in order to run a preview
    healthState := openapiclient.HealthStateValue("UNINITIALIZED") // HealthStateValue | Health state of check states (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.PreviewMonitorCheckStates(context.Background(), monitorIdOrUrn).MonitorPreview(monitorPreview).HealthState(healthState).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.PreviewMonitorCheckStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewMonitorCheckStates`: MonitorCheckStates
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.PreviewMonitorCheckStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorIdOrUrn** | **string** | The id or identifier (urn) of a monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMonitorCheckStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorPreview** | [**MonitorPreview**](MonitorPreview.md) | Monitor overrides in order to run a preview | 
 **healthState** | [**HealthStateValue**](HealthStateValue.md) | Health state of check states | 
 **limit** | **int32** |  | 

### Return type

[**MonitorCheckStates**](MonitorCheckStates.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishHealthStates

> PublishHealthStates(ctx, monitorIdOrUrn).MonitorSnapshot(monitorSnapshot).Execute()

Post monitor health states



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
    monitorSnapshot := *openapiclient.NewMonitorSnapshot([]openapiclient.MonitorHealthState{*openapiclient.NewMonitorHealthState("Id_example", openapiclient.HealthStateValue("UNINITIALIZED"), "TopologyIdentifier_example")}) // MonitorSnapshot | Monitor snapshot of health states

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.PublishHealthStates(context.Background(), monitorIdOrUrn).MonitorSnapshot(monitorSnapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.PublishHealthStates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPublishHealthStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorSnapshot** | [**MonitorSnapshot**](MonitorSnapshot.md) | Monitor snapshot of health states | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestMonitorFunction

> MonitorPreviewResult TestMonitorFunction(ctx, monitorFunctionIdOrUrn).MonitorFunctionTest(monitorFunctionTest).Execute()

Test a monitor



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
    monitorFunctionIdOrUrn := "monitorFunctionIdOrUrn_example" // string | The id or identifier (urn) of a monitor function
    monitorFunctionTest := *openapiclient.NewMonitorFunctionTest([]openapiclient.Argument{openapiclient.Argument{ArgumentBooleanVal: openapiclient.NewArgumentBooleanVal("Type_example", int64(123), false)}}) // MonitorFunctionTest | Monitor function arguments to test

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.TestMonitorFunction(context.Background(), monitorFunctionIdOrUrn).MonitorFunctionTest(monitorFunctionTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.TestMonitorFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMonitorFunction`: MonitorPreviewResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.TestMonitorFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorFunctionIdOrUrn** | **string** | The id or identifier (urn) of a monitor function | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestMonitorFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorFunctionTest** | [**MonitorFunctionTest**](MonitorFunctionTest.md) | Monitor function arguments to test | 

### Return type

[**MonitorPreviewResult**](MonitorPreviewResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestMonitorFunctionCheckStates

> MonitorCheckStates TestMonitorFunctionCheckStates(ctx, monitorFunctionIdOrUrn).MonitorFunctionTest(monitorFunctionTest).HealthState(healthState).Limit(limit).Execute()

Test a monitor



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
    monitorFunctionIdOrUrn := "monitorFunctionIdOrUrn_example" // string | The id or identifier (urn) of a monitor function
    monitorFunctionTest := *openapiclient.NewMonitorFunctionTest([]openapiclient.Argument{openapiclient.Argument{ArgumentBooleanVal: openapiclient.NewArgumentBooleanVal("Type_example", int64(123), false)}}) // MonitorFunctionTest | Monitor function arguments to test
    healthState := openapiclient.HealthStateValue("UNINITIALIZED") // HealthStateValue | Health state of check states (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.TestMonitorFunctionCheckStates(context.Background(), monitorFunctionIdOrUrn).MonitorFunctionTest(monitorFunctionTest).HealthState(healthState).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.TestMonitorFunctionCheckStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMonitorFunctionCheckStates`: MonitorCheckStates
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.TestMonitorFunctionCheckStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorFunctionIdOrUrn** | **string** | The id or identifier (urn) of a monitor function | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestMonitorFunctionCheckStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorFunctionTest** | [**MonitorFunctionTest**](MonitorFunctionTest.md) | Monitor function arguments to test | 
 **healthState** | [**HealthStateValue**](HealthStateValue.md) | Health state of check states | 
 **limit** | **int32** |  | 

### Return type

[**MonitorCheckStates**](MonitorCheckStates.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

