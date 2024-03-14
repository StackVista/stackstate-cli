# BaseNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NotificationConfigurationId** | Pointer to **int64** |  | [optional] 
**Status** | [**NotificationChannelStatus**](NotificationChannelStatus.md) |  | 

## Methods

### NewBaseNotificationChannel

`func NewBaseNotificationChannel(id int64, status NotificationChannelStatus, ) *BaseNotificationChannel`

NewBaseNotificationChannel instantiates a new BaseNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotificationChannelWithDefaults

`func NewBaseNotificationChannelWithDefaults() *BaseNotificationChannel`

NewBaseNotificationChannelWithDefaults instantiates a new BaseNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseNotificationChannel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNotificationChannel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseNotificationChannel) SetId(v int64)`

SetId sets Id field to given value.


### GetNotificationConfigurationId

`func (o *BaseNotificationChannel) GetNotificationConfigurationId() int64`

GetNotificationConfigurationId returns the NotificationConfigurationId field if non-nil, zero value otherwise.

### GetNotificationConfigurationIdOk

`func (o *BaseNotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool)`

GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfigurationId

`func (o *BaseNotificationChannel) SetNotificationConfigurationId(v int64)`

SetNotificationConfigurationId sets NotificationConfigurationId field to given value.

### HasNotificationConfigurationId

`func (o *BaseNotificationChannel) HasNotificationConfigurationId() bool`

HasNotificationConfigurationId returns a boolean if a field has been set.

### GetStatus

`func (o *BaseNotificationChannel) GetStatus() NotificationChannelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseNotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseNotificationChannel) SetStatus(v NotificationChannelStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


