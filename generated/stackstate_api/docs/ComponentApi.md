# \ComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentBoundMetric**](ComponentApi.md#GetComponentBoundMetric) | **Post** /components/{componentIdOrIdentifier}/bindmetric | Get a bound metric for a component
[**GetComponentCheckStates**](ComponentApi.md#GetComponentCheckStates) | **Get** /components/{componentIdOrIdentifier}/checkStates | Get a component checkstates
[**GetComponentHealthHistory**](ComponentApi.md#GetComponentHealthHistory) | **Get** /components/{componentIdOrIdentifier}/healthHistory | Get a component health history
[**GetFullComponent**](ComponentApi.md#GetFullComponent) | **Get** /components/{componentIdOrIdentifier} | Get full component
[**GetMetricPerspectiveData**](ComponentApi.md#GetMetricPerspectiveData) | **Get** /components/{componentIdOrIdentifier}/metricPerspectiveData | Bound metric bindings that have data for a component



## GetComponentBoundMetric

> BoundMetric GetComponentBoundMetric(ctx, componentIdOrIdentifier).BoundMetricId(boundMetricId).TopologyTime(topologyTime).Execute()

Get a bound metric for a component



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
    componentIdOrIdentifier := "componentIdOrIdentifier_example" // string | The id or identifier (urn) of a component
    boundMetricId := openapiclient.BoundMetricId{BoundMetricBindingId: openapiclient.NewBoundMetricBindingId("Type_example", "Identifier_example")} // BoundMetricId | 
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentBoundMetric(context.Background(), componentIdOrIdentifier).BoundMetricId(boundMetricId).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetComponentBoundMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentBoundMetric`: BoundMetric
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetComponentBoundMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrIdentifier** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentBoundMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boundMetricId** | [**BoundMetricId**](BoundMetricId.md) |  | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**BoundMetric**](BoundMetric.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentCheckStates

> ComponentCheckStates GetComponentCheckStates(ctx, componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).Execute()

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
    componentIdOrIdentifier := "componentIdOrIdentifier_example" // string | The id or identifier (urn) of a component
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentCheckStates(context.Background(), componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).Execute()
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
**componentIdOrIdentifier** | **string** | The id or identifier (urn) of a component | 

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

> ComponentHealthHistory GetComponentHealthHistory(ctx, componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).Execute()

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
    componentIdOrIdentifier := "componentIdOrIdentifier_example" // string | The id or identifier (urn) of a component
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetComponentHealthHistory(context.Background(), componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).Execute()
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
**componentIdOrIdentifier** | **string** | The id or identifier (urn) of a component | 

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


## GetFullComponent

> FullComponent GetFullComponent(ctx, componentIdOrIdentifier).TopologyTime(topologyTime).Execute()

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
    componentIdOrIdentifier := "componentIdOrIdentifier_example" // string | The id or identifier (urn) of a component
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetFullComponent(context.Background(), componentIdOrIdentifier).TopologyTime(topologyTime).Execute()
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
**componentIdOrIdentifier** | **string** | The id or identifier (urn) of a component | 

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


## GetMetricPerspectiveData

> MetricPerspectiveData GetMetricPerspectiveData(ctx, componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()

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
    componentIdOrIdentifier := "componentIdOrIdentifier_example" // string | The id or identifier (urn) of a component
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    endTime := int32(56) // int32 | The end time of a time range to query resources.
    topologyTime := int32(56) // int32 | A timestamp at which resources will be queried. If not given the resources are queried at current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.GetMetricPerspectiveData(context.Background(), componentIdOrIdentifier).StartTime(startTime).EndTime(endTime).TopologyTime(topologyTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.GetMetricPerspectiveData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricPerspectiveData`: MetricPerspectiveData
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.GetMetricPerspectiveData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentIdOrIdentifier** | **string** | The id or identifier (urn) of a component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricPerspectiveDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | The start time of a time range to query resources. | 
 **endTime** | **int32** | The end time of a time range to query resources. | 
 **topologyTime** | **int32** | A timestamp at which resources will be queried. If not given the resources are queried at current time. | 

### Return type

[**MetricPerspectiveData**](MetricPerspectiveData.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

