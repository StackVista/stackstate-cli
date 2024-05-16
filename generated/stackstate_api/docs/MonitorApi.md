# \MonitorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMonitor**](MonitorAPI.md#DeleteMonitor) | **Delete** /monitors/{monitorIdOrUrn} | Delete a monitor
[**GetAllMonitors**](MonitorAPI.md#GetAllMonitors) | **Get** /monitors | List monitors
[**GetMonitor**](MonitorAPI.md#GetMonitor) | **Get** /monitors/{monitorIdOrUrn} | Get a monitor
[**GetMonitorWithStatus**](MonitorAPI.md#GetMonitorWithStatus) | **Get** /monitors/{monitorIdOrUrn}/status | Get a monitor with stream information
[**GetMonitorsOverview**](MonitorAPI.md#GetMonitorsOverview) | **Get** /monitors/overview | List monitors overview
[**LookupIdentifier**](MonitorAPI.md#LookupIdentifier) | **Post** /monitors/identifierLookup | Multiple component identifier lookup
[**PatchMonitor**](MonitorAPI.md#PatchMonitor) | **Patch** /monitors/{monitorIdOrUrn} | Update some monitor properties
[**PreviewMonitor**](MonitorAPI.md#PreviewMonitor) | **Post** /monitors/{monitorIdOrUrn}/preview | Preview a monitor
[**PublishHealthStates**](MonitorAPI.md#PublishHealthStates) | **Post** /monitors/{monitorIdOrUrn}/publish | Post monitor health states
[**RunMonitor**](MonitorAPI.md#RunMonitor) | **Post** /monitors/{monitorIdOrUrn}/run | Run a monitor
[**TestMonitorFunction**](MonitorAPI.md#TestMonitorFunction) | **Post** /monitors/{monitorFunctionIdOrUrn}/test | Test a monitor



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
    resp, r, err := apiClient.MonitorAPI.DeleteMonitor(context.Background(), monitorIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.DeleteMonitor``: %v\n", err)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.GetAllMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.GetAllMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMonitors`: MonitorList
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.GetAllMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMonitorsRequest struct via the builder pattern


### Return type

[**MonitorList**](MonitorList.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.GetMonitor(context.Background(), monitorIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.GetMonitor`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.GetMonitorWithStatus(context.Background(), monitorIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.GetMonitorWithStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorWithStatus`: MonitorStatus
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.GetMonitorWithStatus`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.GetMonitorsOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.GetMonitorsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorsOverview`: MonitorOverviewList
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.GetMonitorsOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsOverviewRequest struct via the builder pattern


### Return type

[**MonitorOverviewList**](MonitorOverviewList.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    monitorIdentifierLookup := *openapiclient.NewMonitorIdentifierLookup("MetricQuery_example", int64(123)) // MonitorIdentifierLookup | Component type and metric query for identifier lookup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorAPI.LookupIdentifier(context.Background()).MonitorIdentifierLookup(monitorIdentifierLookup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.LookupIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupIdentifier`: MonitorIdentifierSuggestions
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.LookupIdentifier`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.PatchMonitor(context.Background(), monitorIdOrUrn).MonitorPatch(monitorPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.PatchMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.PatchMonitor`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.PreviewMonitor(context.Background(), monitorIdOrUrn).MonitorPreview(monitorPreview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.PreviewMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewMonitor`: MonitorPreviewResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.PreviewMonitor`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.PublishHealthStates(context.Background(), monitorIdOrUrn).MonitorSnapshot(monitorSnapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.PublishHealthStates``: %v\n", err)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.RunMonitor(context.Background(), monitorIdOrUrn).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.RunMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMonitor`: MonitorRunResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.RunMonitor`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.MonitorAPI.TestMonitorFunction(context.Background(), monitorFunctionIdOrUrn).MonitorFunctionTest(monitorFunctionTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.TestMonitorFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMonitorFunction`: MonitorPreviewResult
    fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.TestMonitorFunction`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

