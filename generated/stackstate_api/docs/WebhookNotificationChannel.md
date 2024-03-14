# WebhookNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NotificationConfigurationId** | Pointer to **int64** |  | [optional] 
**Status** | [**NotificationChannelStatus**](NotificationChannelStatus.md) |  | 
**Url** | **string** |  | 
**Token** | **string** |  | 
**VerifySsl** | **bool** |  | 
**Metadata** | **map[string]string** |  | 
**Type** | **string** |  | 

## Methods

### NewWebhookNotificationChannel

`func NewWebhookNotificationChannel(id int64, status NotificationChannelStatus, url string, token string, verifySsl bool, metadata map[string]string, type_ string, ) *WebhookNotificationChannel`

NewWebhookNotificationChannel instantiates a new WebhookNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationChannelWithDefaults

`func NewWebhookNotificationChannelWithDefaults() *WebhookNotificationChannel`

NewWebhookNotificationChannelWithDefaults instantiates a new WebhookNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookNotificationChannel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookNotificationChannel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookNotificationChannel) SetId(v int64)`

SetId sets Id field to given value.


### GetNotificationConfigurationId

`func (o *WebhookNotificationChannel) GetNotificationConfigurationId() int64`

GetNotificationConfigurationId returns the NotificationConfigurationId field if non-nil, zero value otherwise.

### GetNotificationConfigurationIdOk

`func (o *WebhookNotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool)`

GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfigurationId

`func (o *WebhookNotificationChannel) SetNotificationConfigurationId(v int64)`

SetNotificationConfigurationId sets NotificationConfigurationId field to given value.

### HasNotificationConfigurationId

`func (o *WebhookNotificationChannel) HasNotificationConfigurationId() bool`

HasNotificationConfigurationId returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookNotificationChannel) GetStatus() NotificationChannelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookNotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookNotificationChannel) SetStatus(v NotificationChannelStatus)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *WebhookNotificationChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookNotificationChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookNotificationChannel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *WebhookNotificationChannel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebhookNotificationChannel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebhookNotificationChannel) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerifySsl

`func (o *WebhookNotificationChannel) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *WebhookNotificationChannel) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *WebhookNotificationChannel) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.


### GetMetadata

`func (o *WebhookNotificationChannel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookNotificationChannel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookNotificationChannel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetType

`func (o *WebhookNotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookNotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookNotificationChannel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


