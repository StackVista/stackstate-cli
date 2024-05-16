# \HealthSynchronizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHealthSynchronizationStream**](HealthSynchronizationAPI.md#DeleteHealthSynchronizationStream) | **Delete** /synchronization/health/streams/{healthStreamUrn} | Delete health sync stream
[**GetHealthSynchronizationStreamStatus**](HealthSynchronizationAPI.md#GetHealthSynchronizationStreamStatus) | **Get** /synchronization/health/streams/{healthStreamUrn}/status | Get health sync stream status
[**GetHealthSynchronizationStreamTopologyMatches**](HealthSynchronizationAPI.md#GetHealthSynchronizationStreamTopologyMatches) | **Get** /synchronization/health/streams/{healthStreamUrn}/topologyMatches | List health sync stream check-states
[**GetHealthSynchronizationStreamsOverview**](HealthSynchronizationAPI.md#GetHealthSynchronizationStreamsOverview) | **Get** /synchronization/health/streams | List health sync streams
[**GetHealthSynchronizationSubStreamOverview**](HealthSynchronizationAPI.md#GetHealthSynchronizationSubStreamOverview) | **Get** /synchronization/health/streams/{healthStreamUrn}/substreams | List health sync sub-streams
[**GetHealthSynchronizationSubStreamStatus**](HealthSynchronizationAPI.md#GetHealthSynchronizationSubStreamStatus) | **Get** /synchronization/health/streams/{healthStreamUrn}/substreams/{healthSyncSubStreamId}/status | Get health sync sub-stream status
[**GetHealthSynchronizationSubStreamTopologyMatches**](HealthSynchronizationAPI.md#GetHealthSynchronizationSubStreamTopologyMatches) | **Get** /synchronization/health/streams/{healthStreamUrn}/substreams/{healthSyncSubStreamId}/topologyMatches | List health sync sub-stream check-states
[**PostHealthSynchronizationStreamClearErrors**](HealthSynchronizationAPI.md#PostHealthSynchronizationStreamClearErrors) | **Post** /synchronization/health/streams/{healthStreamUrn}/clearErrors | Clear health sync stream errors



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
    resp, r, err := apiClient.HealthSynchronizationAPI.DeleteHealthSynchronizationStream(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.DeleteHealthSynchronizationStream``: %v\n", err)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationStreamStatus(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationStreamStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamStatus`: HealthStreamStatus
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationStreamStatus`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationStreamTopologyMatches(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationStreamTopologyMatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamTopologyMatches`: TopologyMatchResult
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationStreamTopologyMatches`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationStreamsOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationStreamsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationStreamsOverview`: StreamList
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationStreamsOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthSynchronizationStreamsOverviewRequest struct via the builder pattern


### Return type

[**StreamList**](StreamList.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamOverview(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamOverview`: SubStreamList
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamOverview`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamStatus(context.Background(), healthStreamUrn, healthSyncSubStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamStatus`: HealthSubStreamStatus
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamStatus`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.GetHealthSynchronizationSubStreamTopologyMatches(context.Background(), healthStreamUrn, healthSyncSubStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamTopologyMatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthSynchronizationSubStreamTopologyMatches`: TopologyMatchResult
    fmt.Fprintf(os.Stdout, "Response from `HealthSynchronizationAPI.GetHealthSynchronizationSubStreamTopologyMatches`: %v\n", resp)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

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
    resp, r, err := apiClient.HealthSynchronizationAPI.PostHealthSynchronizationStreamClearErrors(context.Background(), healthStreamUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthSynchronizationAPI.PostHealthSynchronizationStreamClearErrors``: %v\n", err)
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

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

