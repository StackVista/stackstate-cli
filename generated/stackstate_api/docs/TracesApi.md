# \TracesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpan**](TracesApi.md#GetSpan) | **Get** /traces/{traceId}/spans/{spanId} | Get a span
[**GetTrace**](TracesApi.md#GetTrace) | **Get** /traces/{traceId} | Fetch a trace
[**QueryDuration**](TracesApi.md#QueryDuration) | **Post** /traces/spans/duration/histogram | Query duration distribution
[**QuerySpans**](TracesApi.md#QuerySpans) | **Post** /traces/spans | Query for spans
[**SpanComponents**](TracesApi.md#SpanComponents) | **Post** /traces/components | Fetch components based on resource attributes
[**SuggestionsAttributeName**](TracesApi.md#SuggestionsAttributeName) | **Get** /traces/spans/fields/attributes | Suggestions for attribute names
[**SuggestionsAttributeValue**](TracesApi.md#SuggestionsAttributeValue) | **Get** /traces/spans/fields/attributes/{attributeName}/values | Suggestions for attribute values
[**SuggestionsFieldValues**](TracesApi.md#SuggestionsFieldValues) | **Get** /traces/spans/fields/{field}/values | Suggestions for span fields



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
    resp, r, err := apiClient.TracesApi.GetSpan(context.Background(), traceId, spanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.GetSpan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpan`: Span
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.GetSpan`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.TracesApi.GetTrace(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.GetTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrace`: Trace
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.GetTrace`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryDuration

> DurationHistogram QueryDuration(ctx).Start(start).End(end).BucketsCount(bucketsCount).SpanFilter(spanFilter).Execute()

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
    start := int32(1707117737) // int32 | Filter spans by start time >= value
    end := int32(1707121359) // int32 | Filter spans by start time < value
    bucketsCount := int32(56) // int32 | The number of histogram buckets.
    spanFilter := *openapiclient.NewSpanFilter() // SpanFilter | Filter for spans

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.QueryDuration(context.Background()).Start(start).End(end).BucketsCount(bucketsCount).SpanFilter(spanFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.QueryDuration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryDuration`: DurationHistogram
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.QueryDuration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Filter spans by start time &gt;&#x3D; value | 
 **end** | **int32** | Filter spans by start time &lt; value | 
 **bucketsCount** | **int32** | The number of histogram buckets. | 
 **spanFilter** | [**SpanFilter**](SpanFilter.md) | Filter for spans | 

### Return type

[**DurationHistogram**](DurationHistogram.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySpans

> Spans QuerySpans(ctx).Start(start).End(end).SpanQuery(spanQuery).PageSize(pageSize).Page(page).Execute()

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
    start := int32(1707117737) // int32 | Filter spans by start time >= value
    end := int32(1707121359) // int32 | Filter spans by start time < value
    spanQuery := *openapiclient.NewSpanQuery() // SpanQuery | Query for spans
    pageSize := int32(30) // int32 | Number of spans in 1 page (optional) (default to 20)
    page := int32(4) // int32 | Get the specified page (with pageSize # of spans), defaults to page 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.QuerySpans(context.Background()).Start(start).End(end).SpanQuery(spanQuery).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.QuerySpans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySpans`: Spans
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.QuerySpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuerySpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Filter spans by start time &gt;&#x3D; value | 
 **end** | **int32** | Filter spans by start time &lt; value | 
 **spanQuery** | [**SpanQuery**](SpanQuery.md) | Query for spans | 
 **pageSize** | **int32** | Number of spans in 1 page | [default to 20]
 **page** | **int32** | Get the specified page (with pageSize # of spans), defaults to page 0 | [default to 0]

### Return type

[**Spans**](Spans.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.TracesApi.SpanComponents(context.Background()).ComponentQuery(componentQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.SpanComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpanComponents`: SpanComponents
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.SpanComponents`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.TracesApi.SuggestionsAttributeName(context.Background()).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.SuggestionsAttributeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsAttributeName`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.SuggestionsAttributeName`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.TracesApi.SuggestionsAttributeValue(context.Background(), attributeName).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.SuggestionsAttributeValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsAttributeValue`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.SuggestionsAttributeValue`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.TracesApi.SuggestionsFieldValues(context.Background(), field).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.SuggestionsFieldValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestionsFieldValues`: Suggestions
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.SuggestionsFieldValues`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

