# \NotificationConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationConfiguration**](NotificationConfigurationsApi.md#CreateNotificationConfiguration) | **Post** /notifications/configurations | Create a new notification configuration
[**DeleteNotificationConfiguration**](NotificationConfigurationsApi.md#DeleteNotificationConfiguration) | **Delete** /notifications/configurations/{notificationConfigurationIdOrUrn} | Delete the notification configuration
[**GetNotificationConfiguration**](NotificationConfigurationsApi.md#GetNotificationConfiguration) | **Get** /notifications/configurations/{notificationConfigurationIdOrUrn} | Get the notification configuration
[**GetNotificationConfigurationChannels**](NotificationConfigurationsApi.md#GetNotificationConfigurationChannels) | **Get** /notifications/configurations/{notificationConfigurationIdOrUrn}/channels | Get the channels for the notification configuration
[**GetNotificationConfigurations**](NotificationConfigurationsApi.md#GetNotificationConfigurations) | **Get** /notifications/configurations | Get all notification configurations
[**UpdateNotificationConfiguration**](NotificationConfigurationsApi.md#UpdateNotificationConfiguration) | **Put** /notifications/configurations/{notificationConfigurationIdOrUrn} | Update the notification configuration



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
    notificationConfigurationWriteSchema := *openapiclient.NewNotificationConfigurationWriteSchema("Name_example", openapiclient.NotifyOnOptions("CRITICAL"), []openapiclient.MonitorReferenceId{openapiclient.MonitorReferenceId{ExternalMonitorDefId: openapiclient.NewExternalMonitorDefId("Type_example", int64(123))}}, []string{"MonitorTags_example"}, []string{"ComponentTypeNames_example"}, []string{"ComponentTags_example"}, openapiclient.NotificationConfigurationStatusValue("ENABLED"), []openapiclient.ChannelReferenceId{openapiclient.ChannelReferenceId{EmailChannelRefId: openapiclient.NewEmailChannelRefId("Type_example", int64(123))}}) // NotificationConfigurationWriteSchema | Create or update a notification configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsApi.CreateNotificationConfiguration(context.Background()).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.CreateNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsApi.CreateNotificationConfiguration`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.NotificationConfigurationsApi.DeleteNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.DeleteNotificationConfiguration``: %v\n", err)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.NotificationConfigurationsApi.GetNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.GetNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsApi.GetNotificationConfiguration`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.NotificationConfigurationsApi.GetNotificationConfigurationChannels(context.Background(), notificationConfigurationIdOrUrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.GetNotificationConfigurationChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfigurationChannels`: []NotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsApi.GetNotificationConfigurationChannels`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    resp, r, err := apiClient.NotificationConfigurationsApi.GetNotificationConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.GetNotificationConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfigurations`: []NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsApi.GetNotificationConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigurationsRequest struct via the builder pattern


### Return type

[**[]NotificationConfigurationReadSchema**](NotificationConfigurationReadSchema.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

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
    notificationConfigurationWriteSchema := *openapiclient.NewNotificationConfigurationWriteSchema("Name_example", openapiclient.NotifyOnOptions("CRITICAL"), []openapiclient.MonitorReferenceId{openapiclient.MonitorReferenceId{ExternalMonitorDefId: openapiclient.NewExternalMonitorDefId("Type_example", int64(123))}}, []string{"MonitorTags_example"}, []string{"ComponentTypeNames_example"}, []string{"ComponentTags_example"}, openapiclient.NotificationConfigurationStatusValue("ENABLED"), []openapiclient.ChannelReferenceId{openapiclient.ChannelReferenceId{EmailChannelRefId: openapiclient.NewEmailChannelRefId("Type_example", int64(123))}}) // NotificationConfigurationWriteSchema | Create or update a notification configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationConfigurationsApi.UpdateNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).NotificationConfigurationWriteSchema(notificationConfigurationWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationConfigurationsApi.UpdateNotificationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationConfiguration`: NotificationConfigurationReadSchema
    fmt.Fprintf(os.Stdout, "Response from `NotificationConfigurationsApi.UpdateNotificationConfiguration`: %v\n", resp)
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

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

