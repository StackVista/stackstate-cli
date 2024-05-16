# \TracesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpan**](TracesAPI.md#GetSpan) | **Get** /traces/{traceId}/spans/{spanId} | Get a span
[**GetTrace**](TracesAPI.md#GetTrace) | **Get** /traces/{traceId} | Fetch a trace
[**QueryDuration**](TracesAPI.md#QueryDuration) | **Post** /traces/spans/duration/histogram | Query duration distribution
[**QuerySpans**](TracesAPI.md#QuerySpans) | **Post** /traces/spans | Query for spans
[**SpanComponents**](TracesAPI.md#SpanComponents) | **Post** /traces/components | Fetch components based on resource attributes
[**SuggestionsAttributeName**](TracesAPI.md#SuggestionsAttributeName) | **Get** /traces/spans/fields/attributes | Suggestions for attribute names
[**SuggestionsAttributeValue**](TracesAPI.md#SuggestionsAttributeValue) | **Get** /traces/spans/fields/attributes/{attributeName}/values | Suggestions for attribute values
[**SuggestionsFieldValues**](TracesAPI.md#SuggestionsFieldValues) | **Get** /traces/spans/fields/{field}/values | Suggestions for span fields



## GetSpan

> Span GetSpan(ctx, traceId, spanId).Execute()

Get a span



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
    traceId := "traceId_example" // string | The id of the trace
    spanId := "spanId_example" // string | The id of the span

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetSpan(context.Background(), traceId, spanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetSpan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpan`: Span
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetSpan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | The id of the trace | 
**spanId** | **string** | The id of the span | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Span**](Span.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrace

> Trace GetTrace(ctx, traceId).Execute()

Fetch a trace



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
    traceId := "traceId_example" // string | The id of the trace

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetTrace(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrace`: Trace
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | The id of the trace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trace**](Trace.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryDuration

> DurationHistogram QueryDuration(ctx).StartTime(startTime).BucketsCount(bucketsCount).SpanFilter(spanFilter).EndTime(endTime).Execute()

Query duration distribution



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
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    bucketsCount := int32(56) // int32 | The number of histogram buckets.
    spanFilter := *openapiclient.NewSpanFilter() // SpanFilter | Filter for spans
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.QueryDuration(context.Background()).StartTime(startTime).BucketsCount(bucketsCount).SpanFilter(spanFilter).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.QueryDuration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryDuration`: DurationHistogram
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.QueryDuration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | The start time of a time range to query resources. | 
 **bucketsCount** | **int32** | The number of histogram buckets. | 
 **spanFilter** | [**SpanFilter**](SpanFilter.md) | Filter for spans | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 

### Return type

[**DurationHistogram**](DurationHistogram.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySpans

> Spans QuerySpans(ctx).StartTime(startTime).SpanQuery(spanQuery).EndTime(endTime).PageSize(pageSize).Page(page).Execute()

Query for spans



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
    startTime := int32(56) // int32 | The start time of a time range to query resources.
    spanQuery := *openapiclient.NewSpanQuery() // SpanQuery | Query for spans
    endTime := int32(56) // int32 | The end time of a time range to query resources. If not given the endTime is set to current time. (optional)
    pageSize := int32(56) // int32 | Maximum number of the log lines in the result. (optional) (default to 25)
    page := int32(56) // int32 | The page for which the log lines of pageSize must be returned. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.QuerySpans(context.Background()).StartTime(startTime).SpanQuery(spanQuery).EndTime(endTime).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.QuerySpans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySpans`: Spans
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.QuerySpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuerySpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | The start time of a time range to query resources. | 
 **spanQuery** | [**SpanQuery**](SpanQuery.md) | Query for spans | 
 **endTime** | **int32** | The end time of a time range to query resources. If not given the endTime is set to current time. | 
 **pageSize** | **int32** | Maximum number of the log lines in the result. | [default to 25]
 **page** | **int32** | The page for which the log lines of pageSize must be returned. | [default to 1]

### Return type

[**Spans**](Spans.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpanComponents

> SpanComponents SpanComponents(ctx).ComponentQuery(componentQuery).Execute()

Fetch components based on resource attributes



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
    componentQuery := *openapiclient.NewComponentQuery(map[string]string{"key": "Inner_example"}) // ComponentQuery | Span properties to find matching components

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.SpanComponents(context.Background()).ComponentQuery(componentQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SpanComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanComponents`: SpanComponents
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SpanComponents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpanComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentQuery** | [**ComponentQuery**](ComponentQuery.md) | Span properties to find matching components | 

### Return type

[**SpanComponents**](SpanComponents.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestionsAttributeName

> Suggestions SuggestionsAttributeName(ctx).Contains(contains).Execute()

Suggestions for attribute names



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
    contains := "contains_example" // string | Get suggestions based of this partial name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.SuggestionsAttributeName(context.Background()).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SuggestionsAttributeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsAttributeName`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SuggestionsAttributeName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestionsAttributeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contains** | **string** | Get suggestions based of this partial name | 

### Return type

[**Suggestions**](Suggestions.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestionsAttributeValue

> Suggestions SuggestionsAttributeValue(ctx, attributeName).Contains(contains).Execute()

Suggestions for attribute values



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
    attributeName := "attributeName_example" // string | Get suggestions for this attribute
    contains := "contains_example" // string | Get suggestions based on this partial value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.SuggestionsAttributeValue(context.Background(), attributeName).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SuggestionsAttributeValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsAttributeValue`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SuggestionsAttributeValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeName** | **string** | Get suggestions for this attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuggestionsAttributeValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contains** | **string** | Get suggestions based on this partial value | 

### Return type

[**Suggestions**](Suggestions.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestionsFieldValues

> Suggestions SuggestionsFieldValues(ctx, field).Contains(contains).Execute()

Suggestions for span fields



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
    field := openapiclient.SpanSuggestionField("ServiceName") // SpanSuggestionField | Get suggestions for this field
    contains := "contains_example" // string | Get suggestions for this field based on this partial value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.SuggestionsFieldValues(context.Background(), field).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SuggestionsFieldValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsFieldValues`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SuggestionsFieldValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**field** | [**SpanSuggestionField**](.md) | Get suggestions for this field | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuggestionsFieldValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contains** | **string** | Get suggestions for this field based on this partial value | 

### Return type

[**Suggestions**](Suggestions.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

