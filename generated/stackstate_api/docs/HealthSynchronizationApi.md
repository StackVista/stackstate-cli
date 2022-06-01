# \HealthSynchronizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHealthSynchronizationStream**](HealthSynchronizationApi.md#DeleteHealthSynchronizationStream) | **Delete** /synchronization/health/stream/{healthStreamUrn} | Delete health sync stream
[**GetHealthSynchronizationStreamStatus**](HealthSynchronizationApi.md#GetHealthSynchronizationStreamStatus) | **Get** /synchronization/health/stream/{healthStreamUrn}/status | Get health sync stream status
[**GetHealthSynchronizationStreamTopologyMatches**](HealthSynchronizationApi.md#GetHealthSynchronizationStreamTopologyMatches) | **Get** /synchronization/health/stream/{healthStreamUrn}/topologyMatches | List health sync stream check-states
[**GetHealthSynchronizationStreamsOverview**](HealthSynchronizationApi.md#GetHealthSynchronizationStreamsOverview) | **Get** /synchronization/health/streams | List health sync streams
[**GetHealthSynchronizationSubStreamOverview**](HealthSynchronizationApi.md#GetHealthSynchronizationSubStreamOverview) | **Get** /synchronization/health/stream/{healthStreamUrn}/substreams | List health sync sub-streams
[**GetHealthSynchronizationSubStreamStatus**](HealthSynchronizationApi.md#GetHealthSynchronizationSubStreamStatus) | **Get** /synchronization/health/stream/{healthStreamUrn}/substream/{healthSyncSubStreamId}/status | Get health sync sub-stream status
[**GetHealthSynchronizationSubStreamTopologyMatches**](HealthSynchronizationApi.md#GetHealthSynchronizationSubStreamTopologyMatches) | **Get** /synchronization/health/stream/{healthStreamUrn}/substream/{healthSyncSubStreamId}/topologyMatches | List health sync sub-stream check-states
[**PostHealthSynchronizationStreamClearErrors**](HealthSynchronizationApi.md#PostHealthSynchronizationStreamClearErrors) | **Post** /synchronization/health/stream/{healthStreamUrn}/clearErrors | Clear health sync stream errors



## DeleteHealthSynchronizationStream

> DeleteHealthSynchronizationStream(ctx, healthStreamUrn).Execute()

Delete health sync stream



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.DeleteHealthSynchronizationStream(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.DeleteHealthSynchronizationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHealthSynchronizationStreamRequest struct via the builder pattern


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


## GetHealthSynchronizationStreamStatus

> HealthStreamStatus GetHealthSynchronizationStreamStatus(ctx, healthStreamUrn).Execute()

Get health sync stream status



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationStreamStatus(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationStreamStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamStatus`: HealthStreamStatus
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationStreamStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationStreamStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HealthStreamStatus**](HealthStreamStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthSynchronizationStreamTopologyMatches

> TopologyMatchResult GetHealthSynchronizationStreamTopologyMatches(ctx, healthStreamUrn).Execute()

List health sync stream check-states



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatches(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamTopologyMatches`: TopologyMatchResult
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationStreamTopologyMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopologyMatchResult**](TopologyMatchResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthSynchronizationStreamsOverview

> StreamList GetHealthSynchronizationStreamsOverview(ctx).Execute()

List health sync streams



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
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationStreamsOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationStreamsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamsOverview`: StreamList
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationStreamsOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationStreamsOverviewRequest struct via the builder pattern


### Return type

[**StreamList**](StreamList.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthSynchronizationSubStreamOverview

> SubStreamList GetHealthSynchronizationSubStreamOverview(ctx, healthStreamUrn).Execute()

List health sync sub-streams



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverview(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamOverview`: SubStreamList
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationSubStreamOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubStreamList**](SubStreamList.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthSynchronizationSubStreamStatus

> HealthSubStreamStatus GetHealthSynchronizationSubStreamStatus(ctx, healthStreamUrn, healthSyncSubStreamId).Execute()

Get health sync sub-stream status



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.
    healthSyncSubStreamId := "healthSyncSubStreamId_example" // string | Health synchronization sub stream id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus(context.Background(), healthStreamUrn, healthSyncSubStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamStatus`: HealthSubStreamStatus
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 
**healthSyncSubStreamId** | **string** | Health synchronization sub stream id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationSubStreamStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HealthSubStreamStatus**](HealthSubStreamStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthSynchronizationSubStreamTopologyMatches

> TopologyMatchResult GetHealthSynchronizationSubStreamTopologyMatches(ctx, healthStreamUrn, healthSyncSubStreamId).Execute()

List health sync sub-stream check-states



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.
    healthSyncSubStreamId := "healthSyncSubStreamId_example" // string | Health synchronization sub stream id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.GetHealthSynchronizationSubStreamTopologyMatches(context.Background(), healthStreamUrn, healthSyncSubStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.GetHealthSynchronizationSubStreamTopologyMatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamTopologyMatches`: TopologyMatchResult
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationApi.GetHealthSynchronizationSubStreamTopologyMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 
**healthSyncSubStreamId** | **string** | Health synchronization sub stream id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationSubStreamTopologyMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopologyMatchResult**](TopologyMatchResult.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHealthSynchronizationStreamClearErrors

> PostHealthSynchronizationStreamClearErrors(ctx, healthStreamUrn).Execute()

Clear health sync stream errors



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
    healthStreamUrn := "healthStreamUrn_example" // string | Urn of the health stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthSynchronizationApi.PostHealthSynchronizationStreamClearErrors(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationApi.PostHealthSynchronizationStreamClearErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthStreamUrn** | **string** | Urn of the health stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostHealthSynchronizationStreamClearErrorsRequest struct via the builder pattern


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

