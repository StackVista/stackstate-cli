# \RetentionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetentionEpoch**](RetentionAPI.md#GetRetentionEpoch) | **Get** /retention/currentEpoch | Get retention epoch
[**GetRetentionWindow**](RetentionAPI.md#GetRetentionWindow) | **Get** /retention/window | Get retention window
[**RemoveExpiredData**](RetentionAPI.md#RemoveExpiredData) | **Post** /retention/removeExpiredData | Remove expired data from StackGraph
[**SetRetentionWindow**](RetentionAPI.md#SetRetentionWindow) | **Post** /retention/window | Set retention window



## GetRetentionEpoch

> EpochTx GetRetentionEpoch(ctx).Execute()

Get retention epoch



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
    resp, r, err := apiClient.RetentionAPI.GetRetentionEpoch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.GetRetentionEpoch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionEpoch`: EpochTx
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.GetRetentionEpoch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionEpochRequest struct via the builder pattern


### Return type

[**EpochTx**](EpochTx.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionWindow

> WindowMs GetRetentionWindow(ctx).Execute()

Get retention window



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
    resp, r, err := apiClient.RetentionAPI.GetRetentionWindow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.GetRetentionWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionWindow`: WindowMs
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.GetRetentionWindow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionWindowRequest struct via the builder pattern


### Return type

[**WindowMs**](WindowMs.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExpiredData

> RemovalProgress RemoveExpiredData(ctx).ExpireImmediatelyAndRestart(expireImmediatelyAndRestart).Execute()

Remove expired data from StackGraph



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
    expireImmediatelyAndRestart := true // bool | If set, makes StackState remove expired data immediately and restart afterwards. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.RemoveExpiredData(context.Background()).ExpireImmediatelyAndRestart(expireImmediatelyAndRestart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.RemoveExpiredData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExpiredData`: RemovalProgress
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.RemoveExpiredData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExpiredDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expireImmediatelyAndRestart** | **bool** | If set, makes StackState remove expired data immediately and restart afterwards. | 

### Return type

[**RemovalProgress**](RemovalProgress.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetentionWindow

> WindowMs SetRetentionWindow(ctx).WindowMs(windowMs).ScheduleRemoval(scheduleRemoval).Execute()

Set retention window



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
    windowMs := *openapiclient.NewWindowMs(int64(123)) // WindowMs | The new retention window value.
    scheduleRemoval := true // bool | If set, makes StackState schedule removal of expired data according to the set retention window. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.SetRetentionWindow(context.Background()).WindowMs(windowMs).ScheduleRemoval(scheduleRemoval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.SetRetentionWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRetentionWindow`: WindowMs
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.SetRetentionWindow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRetentionWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **windowMs** | [**WindowMs**](WindowMs.md) | The new retention window value. | 
 **scheduleRemoval** | **bool** | If set, makes StackState schedule removal of expired data according to the set retention window. | 

### Return type

[**WindowMs**](WindowMs.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

