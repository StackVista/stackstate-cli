# \EventApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventApi.md#GetEvent) | **Get** /events/{eventId} | Get single event
[**GetEventSources**](EventApi.md#GetEventSources) | **Get** /eventSources | Get event sources
[**GetEventTags**](EventApi.md#GetEventTags) | **Get** /eventTags | Get event tags
[**GetEventTypes**](EventApi.md#GetEventTypes) | **Get** /eventTypes | Get event types
[**GetEvents**](EventApi.md#GetEvents) | **Post** /events | Get events



## GetEvent

> TopologyEvent GetEvent(ctx, eventId).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).PlayHeadTimestampMs(playHeadTimestampMs).Execute()

Get single event



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
    eventId := "eventId_example" // string | The Identifier of an event.
    startTimestampMs := int32(56) // int32 | 
    endTimestampMs := int32(56) // int32 | 
    playHeadTimestampMs := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventApi.GetEvent(context.Background(), eventId).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).PlayHeadTimestampMs(playHeadTimestampMs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: TopologyEvent
    fmt.Fprintf(os.Stdout, "Response from `EventApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The Identifier of an event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTimestampMs** | **int32** |  | 
 **endTimestampMs** | **int32** |  | 
 **playHeadTimestampMs** | **int32** |  | 

### Return type

[**TopologyEvent**](TopologyEvent.md)

### Authorization

[app_id](../README.md#app_id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSources

> StringItemsWithTotal GetEventSources(ctx).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()

Get event sources



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
    startTimestampMs := int32(56) // int32 | 
    endTimestampMs := int32(56) // int32 | 
    topologyQuery := "topologyQuery_example" // string | 
    limit := int32(56) // int32 | 
    rootCauseMode := openapiclient.RootCauseMode("no-cause") // RootCauseMode |  (optional)
    playHeadTimestampMs := int32(56) // int32 |  (optional)
    eventTypes := []string{"Inner_example"} // []string |  (optional)
    eventCategories := []openapiclient.EventCategory{openapiclient.EventCategory("Changes")} // []EventCategory | The category labels of an event. (optional)
    eventSources := []string{"Inner_example"} // []string |  (optional)
    eventTags := []string{"Inner_example"} // []string |  (optional)
    match := "match_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventApi.GetEventSources(context.Background()).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.GetEventSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventSources`: StringItemsWithTotal
    fmt.Fprintf(os.Stdout, "Response from `EventApi.GetEventSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestampMs** | **int32** |  | 
 **endTimestampMs** | **int32** |  | 
 **topologyQuery** | **string** |  | 
 **limit** | **int32** |  | 
 **rootCauseMode** | [**RootCauseMode**](RootCauseMode.md) |  | 
 **playHeadTimestampMs** | **int32** |  | 
 **eventTypes** | **[]string** |  | 
 **eventCategories** | [**[]EventCategory**](EventCategory.md) | The category labels of an event. | 
 **eventSources** | **[]string** |  | 
 **eventTags** | **[]string** |  | 
 **match** | **string** |  | [default to &quot;&quot;]

### Return type

[**StringItemsWithTotal**](StringItemsWithTotal.md)

### Authorization

[app_id](../README.md#app_id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTags

> StringItemsWithTotal GetEventTags(ctx).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()

Get event tags



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
    startTimestampMs := int32(56) // int32 | 
    endTimestampMs := int32(56) // int32 | 
    topologyQuery := "topologyQuery_example" // string | 
    limit := int32(56) // int32 | 
    rootCauseMode := openapiclient.RootCauseMode("no-cause") // RootCauseMode |  (optional)
    playHeadTimestampMs := int32(56) // int32 |  (optional)
    eventTypes := []string{"Inner_example"} // []string |  (optional)
    eventCategories := []openapiclient.EventCategory{openapiclient.EventCategory("Changes")} // []EventCategory | The category labels of an event. (optional)
    eventSources := []string{"Inner_example"} // []string |  (optional)
    eventTags := []string{"Inner_example"} // []string |  (optional)
    match := "match_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventApi.GetEventTags(context.Background()).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.GetEventTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventTags`: StringItemsWithTotal
    fmt.Fprintf(os.Stdout, "Response from `EventApi.GetEventTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestampMs** | **int32** |  | 
 **endTimestampMs** | **int32** |  | 
 **topologyQuery** | **string** |  | 
 **limit** | **int32** |  | 
 **rootCauseMode** | [**RootCauseMode**](RootCauseMode.md) |  | 
 **playHeadTimestampMs** | **int32** |  | 
 **eventTypes** | **[]string** |  | 
 **eventCategories** | [**[]EventCategory**](EventCategory.md) | The category labels of an event. | 
 **eventSources** | **[]string** |  | 
 **eventTags** | **[]string** |  | 
 **match** | **string** |  | [default to &quot;&quot;]

### Return type

[**StringItemsWithTotal**](StringItemsWithTotal.md)

### Authorization

[app_id](../README.md#app_id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> StringItemsWithTotal GetEventTypes(ctx).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()

Get event types



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
    startTimestampMs := int32(56) // int32 | 
    endTimestampMs := int32(56) // int32 | 
    topologyQuery := "topologyQuery_example" // string | 
    limit := int32(56) // int32 | 
    rootCauseMode := openapiclient.RootCauseMode("no-cause") // RootCauseMode |  (optional)
    playHeadTimestampMs := int32(56) // int32 |  (optional)
    eventTypes := []string{"Inner_example"} // []string |  (optional)
    eventCategories := []openapiclient.EventCategory{openapiclient.EventCategory("Changes")} // []EventCategory | The category labels of an event. (optional)
    eventSources := []string{"Inner_example"} // []string |  (optional)
    eventTags := []string{"Inner_example"} // []string |  (optional)
    match := "match_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventApi.GetEventTypes(context.Background()).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).TopologyQuery(topologyQuery).Limit(limit).RootCauseMode(rootCauseMode).PlayHeadTimestampMs(playHeadTimestampMs).EventTypes(eventTypes).EventCategories(eventCategories).EventSources(eventSources).EventTags(eventTags).Match(match).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.GetEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventTypes`: StringItemsWithTotal
    fmt.Fprintf(os.Stdout, "Response from `EventApi.GetEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestampMs** | **int32** |  | 
 **endTimestampMs** | **int32** |  | 
 **topologyQuery** | **string** |  | 
 **limit** | **int32** |  | 
 **rootCauseMode** | [**RootCauseMode**](RootCauseMode.md) |  | 
 **playHeadTimestampMs** | **int32** |  | 
 **eventTypes** | **[]string** |  | 
 **eventCategories** | [**[]EventCategory**](EventCategory.md) | The category labels of an event. | 
 **eventSources** | **[]string** |  | 
 **eventTags** | **[]string** |  | 
 **match** | **string** |  | [default to &quot;&quot;]

### Return type

[**StringItemsWithTotal**](StringItemsWithTotal.md)

### Authorization

[app_id](../README.md#app_id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> EventItemsWithTotal GetEvents(ctx).EventListRequest(eventListRequest).Execute()

Get events



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
    eventListRequest := *openapiclient.NewEventListRequest(int32(123), int32(123), "TopologyQuery_example", int32(123)) // EventListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventApi.GetEvents(context.Background()).EventListRequest(eventListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: EventItemsWithTotal
    fmt.Fprintf(os.Stdout, "Response from `EventApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventListRequest** | [**EventListRequest**](EventListRequest.md) |  | 

### Return type

[**EventItemsWithTotal**](EventItemsWithTotal.md)

### Authorization

[app_id](../README.md#app_id)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

