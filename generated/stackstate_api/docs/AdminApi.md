# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetentionEpoch**](AdminApi.md#GetRetentionEpoch) | **Get** /admin/retention/currentEpoch | Get retention epoch
[**GetRetentionWindow**](AdminApi.md#GetRetentionWindow) | **Get** /admin/retention/window | Get retention window
[**RemoveExpiredData**](AdminApi.md#RemoveExpiredData) | **Post** /admin/retention/removeExpiredData | Remove expired data from StackGraph
[**SetRetentionWindow**](AdminApi.md#SetRetentionWindow) | **Post** /admin/retention/window | Set retention window



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
    resp, r, err := apiClient.AdminApi.GetRetentionEpoch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.GetRetentionEpoch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionEpoch`: EpochTx
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.GetRetentionEpoch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionEpochRequest struct via the builder pattern


### Return type

[**EpochTx**](EpochTx.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.AdminApi.GetRetentionWindow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.GetRetentionWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetentionWindow`: WindowMs
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.GetRetentionWindow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionWindowRequest struct via the builder pattern


### Return type

[**WindowMs**](WindowMs.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.AdminApi.RemoveExpiredData(context.Background()).ExpireImmediatelyAndRestart(expireImmediatelyAndRestart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.RemoveExpiredData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExpiredData`: RemovalProgress
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.RemoveExpiredData`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    windowMs := *openapiclient.NewWindowMs() // WindowMs | The new retention window value.
    scheduleRemoval := true // bool | If set, makes StackState schedule removal of expired data according to the set retention window. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.SetRetentionWindow(context.Background()).WindowMs(windowMs).ScheduleRemoval(scheduleRemoval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.SetRetentionWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRetentionWindow`: WindowMs
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.SetRetentionWindow`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

