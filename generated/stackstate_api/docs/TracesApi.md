# \TracesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpan**](TracesApi.md#GetSpan) | **Get** /traces/{traceId}/spans/{spanId} | Get a span
[**GetTrace**](TracesApi.md#GetTrace) | **Get** /traces/{traceId} | Fetch a trace
[**QueryDuration**](TracesApi.md#QueryDuration) | **Post** /traces/duration/histogram | Query duration distribution
[**QueryTraces**](TracesApi.md#QueryTraces) | **Post** /traces/query | Query for traces
[**SpanComponents**](TracesApi.md#SpanComponents) | **Post** /traces/components | Fetch components based on resource attributes
[**SuggestionsAttributeName**](TracesApi.md#SuggestionsAttributeName) | **Get** /traces/spans/fields/attributes | Suggestions for attribute names
[**SuggestionsAttributeValue**](TracesApi.md#SuggestionsAttributeValue) | **Get** /traces/spans/fields/attributes/{attributeName}/values | Suggestions for attribute values
[**SuggestionsFieldValues**](TracesApi.md#SuggestionsFieldValues) | **Get** /traces/spans/fields/{field}/values | Suggestions for span fields



## GetSpan

> SpanResponse GetSpan(ctx, traceId, spanId).Execute()

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
    // response from `GetSpan`: SpanResponse
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

[**SpanResponse**](SpanResponse.md)

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

> DurationHistogram QueryDuration(ctx).Start(start).End(end).BucketsCount(bucketsCount).TraceFilter(traceFilter).Execute()

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
    traceFilter := *openapiclient.NewTraceFilter(*openapiclient.NewSpanFilter()) // TraceFilter | Filter for traces

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.QueryDuration(context.Background()).Start(start).End(end).BucketsCount(bucketsCount).TraceFilter(traceFilter).Execute()
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
 **traceFilter** | [**TraceFilter**](TraceFilter.md) | Filter for traces | 

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


## QueryTraces

> TraceQueryResponse QueryTraces(ctx).Start(start).End(end).TraceQuery(traceQuery).PageSize(pageSize).Page(page).Execute()

Query for traces



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
    traceQuery := *openapiclient.NewTraceQuery(*openapiclient.NewSpanFilter()) // TraceQuery | Query for traces
    pageSize := int32(30) // int32 | Number of spans in 1 page (optional) (default to 20)
    page := int32(4) // int32 | Get the specified page (with pageSize # of spans), defaults to page 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.QueryTraces(context.Background()).Start(start).End(end).TraceQuery(traceQuery).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.QueryTraces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTraces`: TraceQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.QueryTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Filter spans by start time &gt;&#x3D; value | 
 **end** | **int32** | Filter spans by start time &lt; value | 
 **traceQuery** | [**TraceQuery**](TraceQuery.md) | Query for traces | 
 **pageSize** | **int32** | Number of spans in 1 page | [default to 20]
 **page** | **int32** | Get the specified page (with pageSize # of spans), defaults to page 0 | [default to 0]

### Return type

[**TraceQueryResponse**](TraceQueryResponse.md)

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

