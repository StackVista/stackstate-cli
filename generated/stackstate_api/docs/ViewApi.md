# \ViewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetView**](ViewApi.md#GetView) | **Get** /views/{viewIdOrIdentifier} | Get a single view
[**GetViews**](ViewApi.md#GetViews) | **Get** /views | Get a list of views



## GetView

> View GetView(ctx, viewIdOrIdentifier).Execute()

Get a single view



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
    viewIdOrIdentifier := "viewIdOrIdentifier_example" // string | The system ID or identifier of a QueryView.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.GetView(context.Background(), viewIdOrIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.GetView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.GetView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewIdOrIdentifier** | **string** | The system ID or identifier of a QueryView. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**View**](View.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViews

> ViewList GetViews(ctx).Execute()

Get a list of views



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
    resp, r, err := apiClient.ViewApi.GetViews(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.GetViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViews`: ViewList
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.GetViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


### Return type

[**ViewList**](ViewList.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

