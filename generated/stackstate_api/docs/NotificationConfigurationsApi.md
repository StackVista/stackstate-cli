# \NotificationConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationConfiguration**](NotificationConfigurationsAPI.md#CreateNotificationConfiguration) | **Post** /notifications/configurations | Create a new notification configuration
[**DeleteNotificationConfiguration**](NotificationConfigurationsAPI.md#DeleteNotificationConfiguration) | **Delete** /notifications/configurations/{notificationConfigurationIdOrUrn} | Delete the notification configuration
[**GetNotificationConfiguration**](NotificationConfigurationsAPI.md#GetNotificationConfiguration) | **Get** /notifications/configurations/{notificationConfigurationIdOrUrn} | Get the notification configuration
[**GetNotificationConfigurationChannels**](NotificationConfigurationsAPI.md#GetNotificationConfigurationChannels) | **Get** /notifications/configurations/{notificationConfigurationIdOrUrn}/channels | Get the channels for the notification configuration
[**GetNotificationConfigurations**](NotificationConfigurationsAPI.md#GetNotificationConfigurations) | **Get** /notifications/configurations | Get all notification configurations
[**UpdateNotificationConfiguration**](NotificationConfigurationsAPI.md#UpdateNotificationConfiguration) | **Put** /notifications/configurations/{notificationConfigurationIdOrUrn} | Update the notification configuration



## CreateNotificationConfiguration

> NotificationConfigurationReadSchema CreateNotificationConfiguration(ctx).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()

Create a new notification configuration



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
    notificationConfigurationWriteSchema := *openapiclient.NewNotificationConfigurationWriteSchema("Name_example", openapiclient.NotifyOnOptions("CRITICAL"), []openapiclient.MonitorReferenceId{openapiclient.MonitorReferenceId{ExternalMonitorDefId: openapiclient.NewExternalMonitorDefId("Type_example", int64(123))}}, []string{"MonitorTags_example"}, []int64{int64(123)}, []string{"ComponentTags_example"}, openapiclient.NotificationConfigurationStatusValue("ENABLED"), []openapiclient.ChannelReferenceId{openapiclient.ChannelReferenceId{OpsgenieChannelRefId: openapiclient.NewOpsgenieChannelRefId("Type_example", int64(123))}}) // NotificationConfigurationWriteSchema | Create or update a notification configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsAPI.CreateNotificationConfiguration(context.Background()).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.CreateNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsAPI.CreateNotificationConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfigurationWriteSchema** | [**NotificationConfigurationWriteSchema**](NotificationConfigurationWriteSchema.md) | Create or update a notification configuration | 

### Return type

[**NotificationConfigurationReadSchema**](NotificationConfigurationReadSchema.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationConfiguration

> DeleteNotificationConfiguration(ctx, notificationConfigurationIdOrUrn).Execute()

Delete the notification configuration



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
    notificationConfigurationIdOrUrn := "notificationConfigurationIdOrUrn_example" // string | Notification identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsAPI.DeleteNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.DeleteNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationConfigurationIdOrUrn** | **string** | Notification identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationConfiguration

> NotificationConfigurationReadSchema GetNotificationConfiguration(ctx, notificationConfigurationIdOrUrn).Execute()

Get the notification configuration



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
    notificationConfigurationIdOrUrn := "notificationConfigurationIdOrUrn_example" // string | Notification identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.GetNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsAPI.GetNotificationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationConfigurationIdOrUrn** | **string** | Notification identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationConfigurationReadSchema**](NotificationConfigurationReadSchema.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationConfigurationChannels

> []NotificationChannel GetNotificationConfigurationChannels(ctx, notificationConfigurationIdOrUrn).Execute()

Get the channels for the notification configuration



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
    notificationConfigurationIdOrUrn := "notificationConfigurationIdOrUrn_example" // string | Notification identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfigurationChannels(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.GetNotificationConfigurationChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfigurationChannels`: []NotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsAPI.GetNotificationConfigurationChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationConfigurationIdOrUrn** | **string** | Notification identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigurationChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationChannel**](NotificationChannel.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationConfigurations

> []NotificationConfigurationReadSchema GetNotificationConfigurations(ctx).Execute()

Get all notification configurations



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
    resp, r, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.GetNotificationConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfigurations`: []NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsAPI.GetNotificationConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationConfigurationReadSchema**](NotificationConfigurationReadSchema.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationConfiguration

> NotificationConfigurationReadSchema UpdateNotificationConfiguration(ctx, notificationConfigurationIdOrUrn).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()

Update the notification configuration



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
    notificationConfigurationIdOrUrn := "notificationConfigurationIdOrUrn_example" // string | Notification identifier
    notificationConfigurationWriteSchema := *openapiclient.NewNotificationConfigurationWriteSchema("Name_example", openapiclient.NotifyOnOptions("CRITICAL"), []openapiclient.MonitorReferenceId{openapiclient.MonitorReferenceId{ExternalMonitorDefId: openapiclient.NewExternalMonitorDefId("Type_example", int64(123))}}, []string{"MonitorTags_example"}, []int64{int64(123)}, []string{"ComponentTags_example"}, openapiclient.NotificationConfigurationStatusValue("ENABLED"), []openapiclient.ChannelReferenceId{openapiclient.ChannelReferenceId{OpsgenieChannelRefId: openapiclient.NewOpsgenieChannelRefId("Type_example", int64(123))}}) // NotificationConfigurationWriteSchema | Create or update a notification configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsAPI.UpdateNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsAPI.UpdateNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsAPI.UpdateNotificationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationConfigurationIdOrUrn** | **string** | Notification identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfigurationWriteSchema** | [**NotificationConfigurationWriteSchema**](NotificationConfigurationWriteSchema.md) | Create or update a notification configuration | 

### Return type

[**NotificationConfigurationReadSchema**](NotificationConfigurationReadSchema.md)

### Authorization

[ServiceToken](../README.md#ServiceToken), [ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

