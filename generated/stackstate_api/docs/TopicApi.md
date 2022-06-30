# \TopicApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopic**](TopicApi.md#GetTopic) | **Get** /topic/{topicName} | topic details API
[**ListTopics**](TopicApi.md#ListTopics) | **Get** /topic | Topic list API



## GetTopic

> InlineResponse2001 GetTopic(ctx, topicName).Execute()

topic details API



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
    topicName := "topicName_example" // string | Name of the topic

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicApi.GetTopic(context.Background(), topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TopicApi.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Name of the topic | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopics

> []Topic ListTopics(ctx).Execute()

Topic list API



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
    resp, r, err := apiClient.TopicApi.ListTopics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicApi.ListTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTopics`: []Topic
    fmt.Fprintf(os.Stdout, "Response from `TopicApi.ListTopics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTopicsRequest struct via the builder pattern


### Return type

[**[]Topic**](Topic.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

