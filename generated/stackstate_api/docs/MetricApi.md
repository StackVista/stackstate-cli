# \MetricAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExemplarsQuery**](MetricAPI.md#GetExemplarsQuery) | **Get** /metrics/query_exemplars | Experimental: Exemplars for a specific time range
[**GetInstantQuery**](MetricAPI.md#GetInstantQuery) | **Get** /metrics/query | Instant query at a single point in time
[**GetLabelValues**](MetricAPI.md#GetLabelValues) | **Get** /metrics/label/{label}/values | List of label values for a provided label name
[**GetLabels**](MetricAPI.md#GetLabels) | **Get** /metrics/labels | List of label names
[**GetMetadata**](MetricAPI.md#GetMetadata) | **Get** /metrics/metadata | Metadata about metrics currently scraped from targets
[**GetRangeQuery**](MetricAPI.md#GetRangeQuery) | **Get** /metrics/query_range | Query over a range of time
[**GetSeries**](MetricAPI.md#GetSeries) | **Get** /metrics/series | List of time series that match a certain label set
[**PostExemplarsQuery**](MetricAPI.md#PostExemplarsQuery) | **Post** /metrics/query_exemplars | Experimental: Exemplars for a specific time range
[**PostInstantQuery**](MetricAPI.md#PostInstantQuery) | **Post** /metrics/query | Instant query at a single point in time
[**PostLabelValues**](MetricAPI.md#PostLabelValues) | **Post** /metrics/label/{label}/values | List of label values for a provided label name
[**PostLabels**](MetricAPI.md#PostLabels) | **Post** /metrics/labels | List of label names
[**PostMetadata**](MetricAPI.md#PostMetadata) | **Post** /metrics/metadata | Metadata about metrics currently scraped from targets
[**PostRangeQuery**](MetricAPI.md#PostRangeQuery) | **Post** /metrics/query_range | Query over a range of time
[**PostSeries**](MetricAPI.md#PostSeries) | **Post** /metrics/series | List of time series that match a certain label set



## GetExemplarsQuery

> PromExemplarEnvelope GetExemplarsQuery(ctx).Query(query).Start(start).End(end).Execute()

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
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format (optional)
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetExemplarsQuery(context.Background()).Query(query).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetExemplarsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExemplarsQuery`: PromExemplarEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetExemplarsQuery`: %v\n", resp)
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

[**PromExemplarEnvelope**](PromExemplarEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstantQuery

> PromEnvelope GetInstantQuery(ctx).Query(query).Time(time).Timeout(timeout).PostFilter(postFilter).Execute()

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
    postFilter := []string{"Inner_example"} // []string | Enforce additional label filters for queries (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetInstantQuery(context.Background()).Query(query).Time(time).Timeout(timeout).PostFilter(postFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetInstantQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstantQuery`: PromEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetInstantQuery`: %v\n", resp)
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
 **postFilter** | **[]string** | Enforce additional label filters for queries | 

### Return type

[**PromEnvelope**](PromEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabelValues

> PromLabelsEnvelope GetLabelValues(ctx, label).Start(start).End(end).Match(match).Execute()

List of label values for a provided label name



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
    label := "label_example" // string | Prometheus query label
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format (optional)
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format (optional)
    match := []string{"Inner_example"} // []string | Repeated series selector argument that selects the series from which to read the label names. Optional. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetLabelValues(context.Background(), label).Start(start).End(end).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetLabelValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabelValues`: PromLabelsEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetLabelValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** | Prometheus query label | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** | Start timestamp in rfc3339 format or unix format | 
 **end** | **string** | End timestamp in rfc3339 format or unix format | 
 **match** | **[]string** | Repeated series selector argument that selects the series from which to read the label names. Optional. | 

### Return type

[**PromLabelsEnvelope**](PromLabelsEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> PromLabelsEnvelope GetLabels(ctx).Start(start).End(end).Match(match).Execute()

List of label names



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
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format (optional)
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format (optional)
    match := []string{"Inner_example"} // []string | Repeated series selector argument that selects the series from which to read the label names. Optional. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetLabels(context.Background()).Start(start).End(end).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabels`: PromLabelsEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start timestamp in rfc3339 format or unix format | 
 **end** | **string** | End timestamp in rfc3339 format or unix format | 
 **match** | **[]string** | Repeated series selector argument that selects the series from which to read the label names. Optional. | 

### Return type

[**PromLabelsEnvelope**](PromLabelsEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadata

> PromMetadataEnvelope GetMetadata(ctx).Limit(limit).Metric(metric).Execute()

Metadata about metrics currently scraped from targets



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
    limit := int64(2) // int64 | Maximum number of metrics to return. (optional)
    metric := "http_requests_total" // string | A metric name to filter metadata for. All metric metadata is retrieved if left empty. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetMetadata(context.Background()).Limit(limit).Metric(metric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadata`: PromMetadataEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Maximum number of metrics to return. | 
 **metric** | **string** | A metric name to filter metadata for. All metric metadata is retrieved if left empty. | 

### Return type

[**PromMetadataEnvelope**](PromMetadataEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRangeQuery

> PromEnvelope GetRangeQuery(ctx).Query(query).Start(start).End(end).Step(step).Timeout(timeout).MaxNumberOfDataPoints(maxNumberOfDataPoints).PostFilter(postFilter).Execute()

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
    maxNumberOfDataPoints := int64(2) // int64 | Maximum number of data points to return. (optional)
    postFilter := []string{"Inner_example"} // []string | Enforce additional label filters for queries (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetRangeQuery(context.Background()).Query(query).Start(start).End(end).Step(step).Timeout(timeout).MaxNumberOfDataPoints(maxNumberOfDataPoints).PostFilter(postFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetRangeQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRangeQuery`: PromEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetRangeQuery`: %v\n", resp)
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
 **maxNumberOfDataPoints** | **int64** | Maximum number of data points to return. | 
 **postFilter** | **[]string** | Enforce additional label filters for queries | 

### Return type

[**PromEnvelope**](PromEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeries

> PromSeriesEnvelope GetSeries(ctx).Match(match).Start(start).End(end).Execute()

List of time series that match a certain label set



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
    match := []string{"Inner_example"} // []string | Repeated series selector argument that selects the series to return. At least one match[] argument must be provided.
    start := "2015-07-01T20:10:51.781Z or 1660817432" // string | Start timestamp in rfc3339 format or unix format (optional)
    end := "2015-07-01T20:10:51.781Z or 1660817432" // string | End timestamp in rfc3339 format or unix format (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.GetSeries(context.Background()).Match(match).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.GetSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeries`: PromSeriesEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.GetSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **match** | **[]string** | Repeated series selector argument that selects the series to return. At least one match[] argument must be provided. | 
 **start** | **string** | Start timestamp in rfc3339 format or unix format | 
 **end** | **string** | End timestamp in rfc3339 format or unix format | 

### Return type

[**PromSeriesEnvelope**](PromSeriesEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExemplarsQuery

> PromExemplarEnvelope PostExemplarsQuery(ctx).Query(query).Start(start).End(end).Execute()

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
    query := "query_example" // string | 
    start := "start_example" // string |  (optional)
    end := "end_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostExemplarsQuery(context.Background()).Query(query).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostExemplarsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostExemplarsQuery`: PromExemplarEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostExemplarsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExemplarsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 

### Return type

[**PromExemplarEnvelope**](PromExemplarEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstantQuery

> PromEnvelope PostInstantQuery(ctx).Query(query).Time(time).Timeout(timeout).PostFilter(postFilter).Execute()

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
    query := "query_example" // string | 
    time := "time_example" // string |  (optional)
    timeout := "timeout_example" // string |  (optional)
    postFilter := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostInstantQuery(context.Background()).Query(query).Time(time).Timeout(timeout).PostFilter(postFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostInstantQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostInstantQuery`: PromEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostInstantQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInstantQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **time** | **string** |  | 
 **timeout** | **string** |  | 
 **postFilter** | **[]string** |  | 

### Return type

[**PromEnvelope**](PromEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLabelValues

> PromLabelsEnvelope PostLabelValues(ctx, label).Start(start).End(end).Match(match).Execute()

List of label values for a provided label name



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
    label := "label_example" // string | Prometheus query label
    start := "start_example" // string |  (optional)
    end := "end_example" // string |  (optional)
    match := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostLabelValues(context.Background(), label).Start(start).End(end).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostLabelValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLabelValues`: PromLabelsEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostLabelValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** | Prometheus query label | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLabelValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** |  | 
 **end** | **string** |  | 
 **match** | **[]string** |  | 

### Return type

[**PromLabelsEnvelope**](PromLabelsEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLabels

> PromLabelsEnvelope PostLabels(ctx).Start(start).End(end).Match(match).Execute()

List of label names



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
    start := "start_example" // string |  (optional)
    end := "end_example" // string |  (optional)
    match := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostLabels(context.Background()).Start(start).End(end).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLabels`: PromLabelsEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** |  | 
 **end** | **string** |  | 
 **match** | **[]string** |  | 

### Return type

[**PromLabelsEnvelope**](PromLabelsEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMetadata

> PromMetadataEnvelope PostMetadata(ctx).Limit(limit).Metric(metric).Execute()

Metadata about metrics currently scraped from targets



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
    limit := int64(789) // int64 |  (optional)
    metric := "metric_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostMetadata(context.Background()).Limit(limit).Metric(metric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMetadata`: PromMetadataEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** |  | 
 **metric** | **string** |  | 

### Return type

[**PromMetadataEnvelope**](PromMetadataEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRangeQuery

> PromEnvelope PostRangeQuery(ctx).Query(query).Start(start).End(end).Step(step).Timeout(timeout).MaxNumberOfDataPoints(maxNumberOfDataPoints).PostFilter(postFilter).Execute()

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
    query := "query_example" // string | 
    start := "start_example" // string | 
    end := "end_example" // string | 
    step := "step_example" // string | 
    timeout := "timeout_example" // string |  (optional)
    maxNumberOfDataPoints := int64(789) // int64 |  (optional)
    postFilter := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostRangeQuery(context.Background()).Query(query).Start(start).End(end).Step(step).Timeout(timeout).MaxNumberOfDataPoints(maxNumberOfDataPoints).PostFilter(postFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostRangeQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRangeQuery`: PromEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostRangeQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRangeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 
 **step** | **string** |  | 
 **timeout** | **string** |  | 
 **maxNumberOfDataPoints** | **int64** |  | 
 **postFilter** | **[]string** |  | 

### Return type

[**PromEnvelope**](PromEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSeries

> PromSeriesEnvelope PostSeries(ctx).Match(match).Start(start).End(end).Execute()

List of time series that match a certain label set



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
    match := []string{"Inner_example"} // []string | 
    start := "start_example" // string |  (optional)
    end := "end_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricAPI.PostSeries(context.Background()).Match(match).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricAPI.PostSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSeries`: PromSeriesEnvelope
    fmt.Fprintf(os.Stdout, "Response from `MetricAPI.PostSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **match** | **[]string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 

### Return type

[**PromSeriesEnvelope**](PromSeriesEnvelope.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

