# \SubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscription**](SubscriptionAPI.md#GetSubscription) | **Get** /subscription | Get subscription info
[**PostSubscription**](SubscriptionAPI.md#PostSubscription) | **Post** /subscription | Submit a new license to update the subscription



## GetSubscription

> SubscriptionState GetSubscription(ctx).Execute()

Get subscription info



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
    resp, r, err := apiClient.SubscriptionAPI.GetSubscription(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscription`: SubscriptionState
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


### Return type

[**SubscriptionState**](SubscriptionState.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSubscription

> SubscriptionState PostSubscription(ctx).NewLicense(newLicense).Execute()

Submit a new license to update the subscription



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
    newLicense := *openapiclient.NewNewLicense("Key_example") // NewLicense | The new license

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.PostSubscription(context.Background()).NewLicense(newLicense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.PostSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSubscription`: SubscriptionState
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.PostSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newLicense** | [**NewLicense**](NewLicense.md) | The new license | 

### Return type

[**SubscriptionState**](SubscriptionState.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

