# \QueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExemplarsQuery**](QueryApi.md#GetExemplarsQuery) | **Get** /v1alpha1/query_exemplars | Experimental: Exemplars for a specific time range
[**GetInstantQuery**](QueryApi.md#GetInstantQuery) | **Get** /v1alpha1/query | Instant query at a single point in time
[**GetRangeQuery**](QueryApi.md#GetRangeQuery) | **Get** /v1alpha1/query_range | Query over a range of time



## GetExemplarsQuery

> Envelope GetExemplarsQuery(ctx).Query(query).Start(start).End(end).Execute()

Experimental: Exemplars for a specific time range



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
    query := "query_example" // string | Prometheus expression query string
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.GetExemplarsQuery(context.Background()).Query(query).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetExemplarsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExemplarsQuery`: Envelope
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetExemplarsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExemplarsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Prometheus expression query string | 
 **start** | **string** | Start timestamp in rfc3339 format or unix format | 
 **end** | **string** | End timestamp in rfc3339 format or unix format | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstantQuery

> Envelope GetInstantQuery(ctx).Query(query).Time(time).Timeout(timeout).Execute()

Instant query at a single point in time



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
    query := "query_example" // string | Prometheus expression query string
    time := "2015-07-01T20:10:51.781Z or 1660817432" // string | Evaluation timestamp in rfc3339 format or unix format (optional)
    timeout := "timeout_example" // string | Evaluation timeout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.GetInstantQuery(context.Background()).Query(query).Time(time).Timeout(timeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetInstantQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstantQuery`: Envelope
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetInstantQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstantQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Prometheus expression query string | 
 **time** | **string** | Evaluation timestamp in rfc3339 format or unix format | 
 **timeout** | **string** | Evaluation timeout | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRangeQuery

> Envelope GetRangeQuery(ctx).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()

Query over a range of time



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
    query := "query_example" // string | Prometheus expression query string
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format
    step := "5m or 300" // string | Query resolution step width in duration format or float number of seconds.
    timeout := "timeout_example" // string | Evaluation timeout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.GetRangeQuery(context.Background()).Query(query).Start(start).End(end).Step(step).Timeout(timeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetRangeQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRangeQuery`: Envelope
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.GetRangeQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRangeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Prometheus expression query string | 
 **start** | **string** | Start timestamp in rfc3339 format or unix format | 
 **end** | **string** | End timestamp in rfc3339 format or unix format | 
 **step** | **string** | Query resolution step width in duration format or float number of seconds. | 
 **timeout** | **string** | Evaluation timeout | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

