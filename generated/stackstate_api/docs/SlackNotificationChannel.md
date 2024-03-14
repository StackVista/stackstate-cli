# SlackNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NotificationConfigurationId** | Pointer to **int64** |  | [optional] 
**Status** | [**NotificationChannelStatus**](NotificationChannelStatus.md) |  | 
**Type** | **string** |  | 
**SlackWorkspace** | **string** |  | 
**SlackChannel** | Pointer to **string** |  | [optional] 
**SlackChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewSlackNotificationChannel

`func NewSlackNotificationChannel(id int64, status NotificationChannelStatus, type_ string, slackWorkspace string, ) *SlackNotificationChannel`

NewSlackNotificationChannel instantiates a new SlackNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackNotificationChannelWithDefaults

`func NewSlackNotificationChannelWithDefaults() *SlackNotificationChannel`

NewSlackNotificationChannelWithDefaults instantiates a new SlackNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlackNotificationChannel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlackNotificationChannel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlackNotificationChannel) SetId(v int64)`

SetId sets Id field to given value.


### GetNotificationConfigurationId

`func (o *SlackNotificationChannel) GetNotificationConfigurationId() int64`

GetNotificationConfigurationId returns the NotificationConfigurationId field if non-nil, zero value otherwise.

### GetNotificationConfigurationIdOk

`func (o *SlackNotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool)`

GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfigurationId

`func (o *SlackNotificationChannel) SetNotificationConfigurationId(v int64)`

SetNotificationConfigurationId sets NotificationConfigurationId field to given value.

### HasNotificationConfigurationId

`func (o *SlackNotificationChannel) HasNotificationConfigurationId() bool`

HasNotificationConfigurationId returns a boolean if a field has been set.

### GetStatus

`func (o *SlackNotificationChannel) GetStatus() NotificationChannelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlackNotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlackNotificationChannel) SetStatus(v NotificationChannelStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *SlackNotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlackNotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlackNotificationChannel) SetType(v string)`

SetType sets Type field to given value.


### GetSlackWorkspace

`func (o *SlackNotificationChannel) GetSlackWorkspace() string`

GetSlackWorkspace returns the SlackWorkspace field if non-nil, zero value otherwise.

### GetSlackWorkspaceOk

`func (o *SlackNotificationChannel) GetSlackWorkspaceOk() (*string, bool)`

GetSlackWorkspaceOk returns a tuple with the SlackWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWorkspace

`func (o *SlackNotificationChannel) SetSlackWorkspace(v string)`

SetSlackWorkspace sets SlackWorkspace field to given value.


### GetSlackChannel

`func (o *SlackNotificationChannel) GetSlackChannel() string`

GetSlackChannel returns the SlackChannel field if non-nil, zero value otherwise.

### GetSlackChannelOk

`func (o *SlackNotificationChannel) GetSlackChannelOk() (*string, bool)`

GetSlackChannelOk returns a tuple with the SlackChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannel

`func (o *SlackNotificationChannel) SetSlackChannel(v string)`

SetSlackChannel sets SlackChannel field to given value.

### HasSlackChannel

`func (o *SlackNotificationChannel) HasSlackChannel() bool`

HasSlackChannel returns a boolean if a field has been set.

### GetSlackChannelId

`func (o *SlackNotificationChannel) GetSlackChannelId() string`

GetSlackChannelId returns the SlackChannelId field if non-nil, zero value otherwise.

### GetSlackChannelIdOk

`func (o *SlackNotificationChannel) GetSlackChannelIdOk() (*string, bool)`

GetSlackChannelIdOk returns a tuple with the SlackChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannelId

`func (o *SlackNotificationChannel) SetSlackChannelId(v string)`

SetSlackChannelId sets SlackChannelId field to given value.

### HasSlackChannelId

`func (o *SlackNotificationChannel) HasSlackChannelId() bool`

HasSlackChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


