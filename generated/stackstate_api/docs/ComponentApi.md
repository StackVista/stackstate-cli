# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentCheckStates**](ComponentApi.md#GetComponentCheckStates) | **Get** /components/{componentIdOrUrn}/checkStates | Get a component checkstates
[**GetComponentHealthHistory**](ComponentApi.md#GetComponentHealthHistory) | **Get** /components/{componentIdOrUrn}/healthHistory | Get a component health history
[**GetComponentMetricBinding**](ComponentApi.md#GetComponentMetricBinding) | **Get** /components/{componentIdOrUrn}/bindmetric | Get a bound metric binding to a component
[**GetComponentMetricsWithData**](ComponentApi.md#GetComponentMetricsWithData) | **Get** /components/{componentIdOrUrn}/boundMetricsWithData | Bound metric bindings that have data for a component
[**GetFullComponent**](ComponentApi.md#GetFullComponent) | **Get** /components/{componentIdOrUrn} | Get full component



## GetComponentCheckStates

> ComponentCheckStates GetComponentCheckStates(ctx, componentIdOrUrn).StartTime(startTime).EndTime(endTime).Execute()

Get a component checkstates



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentCheckStates(context.Background(), componentIdOrUrn).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentCheckStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentCheckStates`: ComponentCheckStates
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentCheckStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrUrn** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentCheckStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 

### Return type

[**ComponentCheckStates**](ComponentCheckStates.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentHealthHistory

> ComponentHealthHistory GetComponentHealthHistory(ctx, componentIdOrUrn).StartTime(startTime).EndTime(endTime).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentHealthHistory(context.Background(), componentIdOrUrn).StartTime(startTime).EndTime(endTime).Execute()
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


## GetComponentMetricBinding

> BoundMetric GetComponentMetricBinding(ctx, componentIdOrUrn).MetricBindingIdentifier(metricBindingIdentifier).TopologyTime(topologyTime).Execute()

Get a bound metric binding to a component



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
    metricBindingIdentifier := "metricBindingIdentifier_example" // string | The identifier (urn) of a metric binding
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentMetricBinding(context.Background(), componentIdOrUrn).MetricBindingIdentifier(metricBindingIdentifier).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentMetricBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentMetricBinding`: BoundMetric
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentMetricBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrUrn** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentMetricBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricBindingIdentifier** | **string** | The identifier (urn) of a metric binding | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**BoundMetric**](BoundMetric.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentMetricsWithData

> BoundMetrics GetComponentMetricsWithData(ctx, componentIdOrUrn).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()

Bound metric bindings that have data for a component



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
    endTime := int32(56) // int32 | The end time of a time range to query resources.
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentMetricsWithData(context.Background(), componentIdOrUrn).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentMetricsWithData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentMetricsWithData`: BoundMetrics
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentMetricsWithData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrUrn** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentMetricsWithDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**BoundMetrics**](BoundMetrics.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullComponent

> FullComponent GetFullComponent(ctx, componentIdOrUrn).TopologyTime(topologyTime).Execute()

Get full component



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
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetFullComponent(context.Background(), componentIdOrUrn).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetFullComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFullComponent`: FullComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetFullComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrUrn** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**FullComponent**](FullComponent.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

