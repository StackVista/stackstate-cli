# \SystemNotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllSystemNotifications**](SystemNotificationsApi.md#AllSystemNotifications) | **Get** /system/notifications | Overview of system notifications



## AllSystemNotifications

> SystemNotifications AllSystemNotifications(ctx).Execute()

Overview of system notifications



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
    resp, r, err := apiClient.SystemNotificationsApi.AllSystemNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemNotificationsApi.AllSystemNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllSystemNotifications`: SystemNotifications
    fmt.Fprintf(os.Stdout, "Response from `SystemNotificationsApi.AllSystemNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAllSystemNotificationsRequest struct via the builder pattern


### Return type

[**SystemNotifications**](SystemNotifications.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

