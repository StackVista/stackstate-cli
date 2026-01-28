# \NotificationChannelsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailNotificationChannel**](NotificationChannelsApi.md#CreateEmailNotificationChannel) | **Post** /notifications/channels/email | Create a Email Notification channel
[**CreateOpsgenieNotificationChannel**](NotificationChannelsApi.md#CreateOpsgenieNotificationChannel) | **Post** /notifications/channels/opsgenie | Create a Opsgenie Notification channel
[**CreateTeamsNotificationChannel**](NotificationChannelsApi.md#CreateTeamsNotificationChannel) | **Post** /notifications/channels/teams | Create a Teams Notification channel
[**CreateWebhookNotificationChannel**](NotificationChannelsApi.md#CreateWebhookNotificationChannel) | **Post** /notifications/channels/webhook | Create a Webhook Notification channel
[**DeleteEmailNotificationChannel**](NotificationChannelsApi.md#DeleteEmailNotificationChannel) | **Delete** /notifications/channels/email/{channelId} | Delete the Email Notification channel by id
[**DeleteOpsgenieNotificationChannel**](NotificationChannelsApi.md#DeleteOpsgenieNotificationChannel) | **Delete** /notifications/channels/opsgenie/{channelId} | Delete the Opsgenie Notification channel by id
[**DeleteSlackNotificationChannel**](NotificationChannelsApi.md#DeleteSlackNotificationChannel) | **Delete** /notifications/channels/slack/{channelId} | Delete the Slack Notification channel by id
[**DeleteTeamsNotificationChannel**](NotificationChannelsApi.md#DeleteTeamsNotificationChannel) | **Delete** /notifications/channels/teams/{channelId} | Delete the Teams Notification channel by id
[**DeleteWebhookNotificationChannel**](NotificationChannelsApi.md#DeleteWebhookNotificationChannel) | **Delete** /notifications/channels/webhook/{channelId} | Delete the Webhook Notification channel by id
[**GetEmailNotificationChannel**](NotificationChannelsApi.md#GetEmailNotificationChannel) | **Get** /notifications/channels/email/{channelId} | Get the Email Notification channel by id
[**GetEmailNotificationStatus**](NotificationChannelsApi.md#GetEmailNotificationStatus) | **Get** /notifications/channels/email/status | Email Notification channel status
[**GetOpsgenieNotificationChannel**](NotificationChannelsApi.md#GetOpsgenieNotificationChannel) | **Get** /notifications/channels/opsgenie/{channelId} | Get the Opsgenie Notification channel by id
[**GetSlackNotificationChannel**](NotificationChannelsApi.md#GetSlackNotificationChannel) | **Get** /notifications/channels/slack/{channelId} | Get the Slack Notification channel by id
[**GetTeamsNotificationChannel**](NotificationChannelsApi.md#GetTeamsNotificationChannel) | **Get** /notifications/channels/teams/{channelId} | Get the Teams Notification channel by id
[**GetWebhookNotificationChannel**](NotificationChannelsApi.md#GetWebhookNotificationChannel) | **Get** /notifications/channels/webhook/{channelId} | Get the Webhook Notification channel by id
[**JoinSlackChannel**](NotificationChannelsApi.md#JoinSlackChannel) | **Post** /notifications/channels/slack/{channelId}/joinSlackChannel | Join the specified Slack channel to send notifications
[**ListOpsgenieResponders**](NotificationChannelsApi.md#ListOpsgenieResponders) | **Get** /notifications/channels/opsgenie/responders | List Opsgenie responders
[**ListSlackChannels**](NotificationChannelsApi.md#ListSlackChannels) | **Get** /notifications/channels/slack/{channelId}/listSlackChannels | List all public Slack channels
[**SlackChannelDetails**](NotificationChannelsApi.md#SlackChannelDetails) | **Get** /notifications/channels/slack/{channelId}/slackChannelDetails/{slackChannelId} | Get Slack channel details
[**SlackOAuthCallback**](NotificationChannelsApi.md#SlackOAuthCallback) | **Get** /notifications/channels/slack/oauth-callback | The OAuth callback for Slack
[**SlackOauthRedirect**](NotificationChannelsApi.md#SlackOauthRedirect) | **Get** /notifications/channels/slack/oauth-redirect | Starts Slack OAuth2 flow
[**TestEmailChannel**](NotificationChannelsApi.md#TestEmailChannel) | **Post** /notifications/channels/email/{channelId}/test | Test the Email notification channel
[**TestOpsgenieChannel**](NotificationChannelsApi.md#TestOpsgenieChannel) | **Post** /notifications/channels/opsgenie/{channelId}/test | Test the Opsgenie notification channel
[**TestSlackChannel**](NotificationChannelsApi.md#TestSlackChannel) | **Post** /notifications/channels/slack/{channelId}/test | Test the Notification channel
[**TestTeamsChannel**](NotificationChannelsApi.md#TestTeamsChannel) | **Post** /notifications/channels/teams/{channelId}/test | Test the Teams notification channel
[**TestWebhookChannel**](NotificationChannelsApi.md#TestWebhookChannel) | **Post** /notifications/channels/webhook/{channelId}/test | Test the Webhook notification channel
[**UpdateEmailNotificationChannel**](NotificationChannelsApi.md#UpdateEmailNotificationChannel) | **Put** /notifications/channels/email/{channelId} | Update the Email Notification channel by id
[**UpdateOpsgenieNotificationChannel**](NotificationChannelsApi.md#UpdateOpsgenieNotificationChannel) | **Put** /notifications/channels/opsgenie/{channelId} | Update the Opsgenie Notification channel by id
[**UpdateTeamsNotificationChannel**](NotificationChannelsApi.md#UpdateTeamsNotificationChannel) | **Put** /notifications/channels/teams/{channelId} | Update the Teams Notification channel by id
[**UpdateWebhookNotificationChannel**](NotificationChannelsApi.md#UpdateWebhookNotificationChannel) | **Put** /notifications/channels/webhook/{channelId} | Update the Webhook Notification channel by id



## CreateEmailNotificationChannel

> EmailNotificationChannel CreateEmailNotificationChannel(ctx).EmailChannelWriteSchema(emailChannelWriteSchema).Execute()

Create a Email Notification channel



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
    emailChannelWriteSchema := *openapiclient.NewEmailChannelWriteSchema([]string{"To_example"}, []string{"Cc_example"}) // EmailChannelWriteSchema | Create or update a email channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.CreateEmailNotificationChannel(context.Background()).EmailChannelWriteSchema(emailChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.CreateEmailNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailNotificationChannel`: EmailNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.CreateEmailNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailChannelWriteSchema** | [**EmailChannelWriteSchema**](EmailChannelWriteSchema.md) | Create or update a email channel | 

### Return type

[**EmailNotificationChannel**](EmailNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOpsgenieNotificationChannel

> OpsgenieNotificationChannel CreateOpsgenieNotificationChannel(ctx).OpsgenieChannelWriteSchema(opsgenieChannelWriteSchema).Execute()

Create a Opsgenie Notification channel



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
    opsgenieChannelWriteSchema := *openapiclient.NewOpsgenieChannelWriteSchema(openapiclient.OpsgenieRegion("EU"), "GenieKey_example", []openapiclient.OpsgenieResponder{*openapiclient.NewOpsgenieResponder(openapiclient.OpsgenieResponderType("TEAM"), "Responder_example")}, openapiclient.OpsgeniePriority("P1")) // OpsgenieChannelWriteSchema | Create or update a opsgenie channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.CreateOpsgenieNotificationChannel(context.Background()).OpsgenieChannelWriteSchema(opsgenieChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.CreateOpsgenieNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOpsgenieNotificationChannel`: OpsgenieNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.CreateOpsgenieNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpsgenieNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsgenieChannelWriteSchema** | [**OpsgenieChannelWriteSchema**](OpsgenieChannelWriteSchema.md) | Create or update a opsgenie channel | 

### Return type

[**OpsgenieNotificationChannel**](OpsgenieNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamsNotificationChannel

> TeamsNotificationChannel CreateTeamsNotificationChannel(ctx).TeamsChannelWriteSchema(teamsChannelWriteSchema).Execute()

Create a Teams Notification channel



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
    teamsChannelWriteSchema := *openapiclient.NewTeamsChannelWriteSchema("Url_example") // TeamsChannelWriteSchema | Create or update a teams channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.CreateTeamsNotificationChannel(context.Background()).TeamsChannelWriteSchema(teamsChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.CreateTeamsNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamsNotificationChannel`: TeamsNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.CreateTeamsNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamsNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsChannelWriteSchema** | [**TeamsChannelWriteSchema**](TeamsChannelWriteSchema.md) | Create or update a teams channel | 

### Return type

[**TeamsNotificationChannel**](TeamsNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhookNotificationChannel

> WebhookNotificationChannel CreateWebhookNotificationChannel(ctx).WebhookChannelWriteSchema(webhookChannelWriteSchema).Execute()

Create a Webhook Notification channel



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
    webhookChannelWriteSchema := *openapiclient.NewWebhookChannelWriteSchema("Url_example", "Token_example", false, map[string]string{"key": "Inner_example"}) // WebhookChannelWriteSchema | Create or update a webhook channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.CreateWebhookNotificationChannel(context.Background()).WebhookChannelWriteSchema(webhookChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.CreateWebhookNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhookNotificationChannel`: WebhookNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.CreateWebhookNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookChannelWriteSchema** | [**WebhookChannelWriteSchema**](WebhookChannelWriteSchema.md) | Create or update a webhook channel | 

### Return type

[**WebhookNotificationChannel**](WebhookNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailNotificationChannel

> DeleteEmailNotificationChannel(ctx, channelId).Execute()

Delete the Email Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.DeleteEmailNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.DeleteEmailNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailNotificationChannelRequest struct via the builder pattern


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


## DeleteOpsgenieNotificationChannel

> DeleteOpsgenieNotificationChannel(ctx, channelId).Execute()

Delete the Opsgenie Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.DeleteOpsgenieNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.DeleteOpsgenieNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpsgenieNotificationChannelRequest struct via the builder pattern


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


## DeleteSlackNotificationChannel

> DeleteSlackNotificationChannel(ctx, channelId).Execute()

Delete the Slack Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.DeleteSlackNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.DeleteSlackNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackNotificationChannelRequest struct via the builder pattern


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


## DeleteTeamsNotificationChannel

> DeleteTeamsNotificationChannel(ctx, channelId).Execute()

Delete the Teams Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.DeleteTeamsNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.DeleteTeamsNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamsNotificationChannelRequest struct via the builder pattern


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


## DeleteWebhookNotificationChannel

> DeleteWebhookNotificationChannel(ctx, channelId).Execute()

Delete the Webhook Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.DeleteWebhookNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.DeleteWebhookNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookNotificationChannelRequest struct via the builder pattern


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


## GetEmailNotificationChannel

> EmailNotificationChannel GetEmailNotificationChannel(ctx, channelId).Execute()

Get the Email Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.GetEmailNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetEmailNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotificationChannel`: EmailNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetEmailNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailNotificationChannel**](EmailNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailNotificationStatus

> EmailNotificationStatus GetEmailNotificationStatus(ctx).Execute()

Email Notification channel status



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
    resp, r, err := apiClient.NotificationChannelsApi.GetEmailNotificationStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetEmailNotificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotificationStatus`: EmailNotificationStatus
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetEmailNotificationStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationStatusRequest struct via the builder pattern


### Return type

[**EmailNotificationStatus**](EmailNotificationStatus.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpsgenieNotificationChannel

> OpsgenieNotificationChannel GetOpsgenieNotificationChannel(ctx, channelId).Execute()

Get the Opsgenie Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.GetOpsgenieNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetOpsgenieNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpsgenieNotificationChannel`: OpsgenieNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetOpsgenieNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpsgenieNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsgenieNotificationChannel**](OpsgenieNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackNotificationChannel

> SlackNotificationChannel GetSlackNotificationChannel(ctx, channelId).Execute()

Get the Slack Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.GetSlackNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetSlackNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackNotificationChannel`: SlackNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetSlackNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlackNotificationChannel**](SlackNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsNotificationChannel

> TeamsNotificationChannel GetTeamsNotificationChannel(ctx, channelId).Execute()

Get the Teams Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.GetTeamsNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetTeamsNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamsNotificationChannel`: TeamsNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetTeamsNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamsNotificationChannel**](TeamsNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookNotificationChannel

> WebhookNotificationChannel GetWebhookNotificationChannel(ctx, channelId).Execute()

Get the Webhook Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.GetWebhookNotificationChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.GetWebhookNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookNotificationChannel`: WebhookNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.GetWebhookNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookNotificationChannel**](WebhookNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinSlackChannel

> SlackNotificationChannel JoinSlackChannel(ctx, channelId).SlackChannelId(slackChannelId).Execute()

Join the specified Slack channel to send notifications



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
    channelId := int64(789) // int64 | Channel identifier
    slackChannelId := *openapiclient.NewSlackChannelId("Id_example") // SlackChannelId | Provide a Slack channel id to join the specified Slack channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.JoinSlackChannel(context.Background(), channelId).SlackChannelId(slackChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.JoinSlackChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinSlackChannel`: SlackNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.JoinSlackChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinSlackChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slackChannelId** | [**SlackChannelId**](SlackChannelId.md) | Provide a Slack channel id to join the specified Slack channel | 

### Return type

[**SlackNotificationChannel**](SlackNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOpsgenieResponders

> []OpsgenieResponder ListOpsgenieResponders(ctx).GenieKey(genieKey).Region(region).Execute()

List Opsgenie responders



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
    genieKey := "genieKey_example" // string | OpsGenie API key
    region := openapiclient.OpsgenieRegion("EU") // OpsgenieRegion | OpsGenie region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.ListOpsgenieResponders(context.Background()).GenieKey(genieKey).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.ListOpsgenieResponders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOpsgenieResponders`: []OpsgenieResponder
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.ListOpsgenieResponders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOpsgenieRespondersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genieKey** | **string** | OpsGenie API key | 
 **region** | [**OpsgenieRegion**](OpsgenieRegion.md) | OpsGenie region | 

### Return type

[**[]OpsgenieResponder**](OpsgenieResponder.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackChannels

> SlackChannelsChunk ListSlackChannels(ctx, channelId).Cursor(cursor).Execute()

List all public Slack channels



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
    channelId := int64(789) // int64 | Channel identifier
    cursor := "cursor_example" // string | Slack list cursor (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.ListSlackChannels(context.Background(), channelId).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.ListSlackChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSlackChannels`: SlackChannelsChunk
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.ListSlackChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Slack list cursor | 

### Return type

[**SlackChannelsChunk**](SlackChannelsChunk.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackChannelDetails

> SlackChannel SlackChannelDetails(ctx, channelId, slackChannelId).Execute()

Get Slack channel details



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
    channelId := int64(789) // int64 | Channel identifier
    slackChannelId := "slackChannelId_example" // string | Slack Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.SlackChannelDetails(context.Background(), channelId, slackChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.SlackChannelDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlackChannelDetails`: SlackChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.SlackChannelDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 
**slackChannelId** | **string** | Slack Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlackChannelDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SlackChannel**](SlackChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackOAuthCallback

> SlackOAuthCallback(ctx).State(state).Code(code).Error_(error_).Execute()

The OAuth callback for Slack



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
    state := "state_example" // string | State parameter that was passed to Slack, should have the same value as the one passed to Slack.
    code := "code_example" // string | OAuth code from Slack. Either the code is present for the success case or the error parameter is present for the error case. (optional)
    error_ := "error__example" // string | Error parameter. Either the code is present for the success case or the error parameter is present for the error case. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.SlackOAuthCallback(context.Background()).State(state).Code(code).Error_(error_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.SlackOAuthCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackOAuthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | State parameter that was passed to Slack, should have the same value as the one passed to Slack. | 
 **code** | **string** | OAuth code from Slack. Either the code is present for the success case or the error parameter is present for the error case. | 
 **error_** | **string** | Error parameter. Either the code is present for the success case or the error parameter is present for the error case. | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackOauthRedirect

> SlackOauthRedirect(ctx).RedirectPath(redirectPath).Execute()

Starts Slack OAuth2 flow



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
    redirectPath := "redirectPath_example" // string | After completing the oauth flow the user will be redirected back to this path, in the UI, on StackState, to continue further setup of the Slack notification channel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.SlackOauthRedirect(context.Background()).RedirectPath(redirectPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.SlackOauthRedirect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackOauthRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectPath** | **string** | After completing the oauth flow the user will be redirected back to this path, in the UI, on StackState, to continue further setup of the Slack notification channel. | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestEmailChannel

> TestEmailChannel(ctx, channelId).Execute()

Test the Email notification channel



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.TestEmailChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.TestEmailChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestEmailChannelRequest struct via the builder pattern


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


## TestOpsgenieChannel

> TestOpsgenieChannel(ctx, channelId).Execute()

Test the Opsgenie notification channel



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.TestOpsgenieChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.TestOpsgenieChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestOpsgenieChannelRequest struct via the builder pattern


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


## TestSlackChannel

> TestSlackChannel(ctx, channelId).Execute()

Test the Notification channel



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.TestSlackChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.TestSlackChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSlackChannelRequest struct via the builder pattern


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


## TestTeamsChannel

> TestTeamsChannel(ctx, channelId).Execute()

Test the Teams notification channel



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.TestTeamsChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.TestTeamsChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestTeamsChannelRequest struct via the builder pattern


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


## TestWebhookChannel

> TestWebhookChannel(ctx, channelId).Execute()

Test the Webhook notification channel



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
    channelId := int64(789) // int64 | Channel identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.TestWebhookChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.TestWebhookChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookChannelRequest struct via the builder pattern


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


## UpdateEmailNotificationChannel

> EmailNotificationChannel UpdateEmailNotificationChannel(ctx, channelId).EmailChannelWriteSchema(emailChannelWriteSchema).Execute()

Update the Email Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier
    emailChannelWriteSchema := *openapiclient.NewEmailChannelWriteSchema([]string{"To_example"}, []string{"Cc_example"}) // EmailChannelWriteSchema | Create or update a email channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.UpdateEmailNotificationChannel(context.Background(), channelId).EmailChannelWriteSchema(emailChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.UpdateEmailNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailNotificationChannel`: EmailNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.UpdateEmailNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailChannelWriteSchema** | [**EmailChannelWriteSchema**](EmailChannelWriteSchema.md) | Create or update a email channel | 

### Return type

[**EmailNotificationChannel**](EmailNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOpsgenieNotificationChannel

> OpsgenieNotificationChannel UpdateOpsgenieNotificationChannel(ctx, channelId).OpsgenieChannelWriteSchema(opsgenieChannelWriteSchema).Execute()

Update the Opsgenie Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier
    opsgenieChannelWriteSchema := *openapiclient.NewOpsgenieChannelWriteSchema(openapiclient.OpsgenieRegion("EU"), "GenieKey_example", []openapiclient.OpsgenieResponder{*openapiclient.NewOpsgenieResponder(openapiclient.OpsgenieResponderType("TEAM"), "Responder_example")}, openapiclient.OpsgeniePriority("P1")) // OpsgenieChannelWriteSchema | Create or update a opsgenie channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.UpdateOpsgenieNotificationChannel(context.Background(), channelId).OpsgenieChannelWriteSchema(opsgenieChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.UpdateOpsgenieNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOpsgenieNotificationChannel`: OpsgenieNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.UpdateOpsgenieNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpsgenieNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsgenieChannelWriteSchema** | [**OpsgenieChannelWriteSchema**](OpsgenieChannelWriteSchema.md) | Create or update a opsgenie channel | 

### Return type

[**OpsgenieNotificationChannel**](OpsgenieNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamsNotificationChannel

> TeamsNotificationChannel UpdateTeamsNotificationChannel(ctx, channelId).TeamsChannelWriteSchema(teamsChannelWriteSchema).Execute()

Update the Teams Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier
    teamsChannelWriteSchema := *openapiclient.NewTeamsChannelWriteSchema("Url_example") // TeamsChannelWriteSchema | Create or update a teams channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.UpdateTeamsNotificationChannel(context.Background(), channelId).TeamsChannelWriteSchema(teamsChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.UpdateTeamsNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamsNotificationChannel`: TeamsNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.UpdateTeamsNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamsNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsChannelWriteSchema** | [**TeamsChannelWriteSchema**](TeamsChannelWriteSchema.md) | Create or update a teams channel | 

### Return type

[**TeamsNotificationChannel**](TeamsNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookNotificationChannel

> WebhookNotificationChannel UpdateWebhookNotificationChannel(ctx, channelId).WebhookChannelWriteSchema(webhookChannelWriteSchema).Execute()

Update the Webhook Notification channel by id



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
    channelId := int64(789) // int64 | Channel identifier
    webhookChannelWriteSchema := *openapiclient.NewWebhookChannelWriteSchema("Url_example", "Token_example", false, map[string]string{"key": "Inner_example"}) // WebhookChannelWriteSchema | Create or update a webhook channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationChannelsApi.UpdateWebhookNotificationChannel(context.Background(), channelId).WebhookChannelWriteSchema(webhookChannelWriteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsApi.UpdateWebhookNotificationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookNotificationChannel`: WebhookNotificationChannel
    fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsApi.UpdateWebhookNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **int64** | Channel identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookChannelWriteSchema** | [**WebhookChannelWriteSchema**](WebhookChannelWriteSchema.md) | Create or update a webhook channel | 

### Return type

[**WebhookNotificationChannel**](WebhookNotificationChannel.md)

### Authorization

[ApiToken](../README.md#ApiToken), [ServiceBearer](../README.md#ServiceBearer), [ServiceToken](../README.md#ServiceToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

