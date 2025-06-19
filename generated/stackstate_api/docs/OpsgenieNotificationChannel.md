# OpsgenieNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NotificationConfigurationId** | Pointer to **int64** |  | [optional] 
**Status** | [**NotificationChannelStatus**](NotificationChannelStatus.md) |  | 
**Region** | [**OpsgenieRegion**](OpsgenieRegion.md) |  | 
**GenieKey** | **string** |  | 
**Responders** | [**[]OpsgenieResponder**](OpsgenieResponder.md) |  | 
**Priority** | [**OpsgeniePriority**](OpsgeniePriority.md) |  | 
**Type** | **string** |  | 

## Methods

### NewOpsgenieNotificationChannel

`func NewOpsgenieNotificationChannel(id int64, status NotificationChannelStatus, region OpsgenieRegion, genieKey string, responders []OpsgenieResponder, priority OpsgeniePriority, type_ string, ) *OpsgenieNotificationChannel`

NewOpsgenieNotificationChannel instantiates a new OpsgenieNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsgenieNotificationChannelWithDefaults

`func NewOpsgenieNotificationChannelWithDefaults() *OpsgenieNotificationChannel`

NewOpsgenieNotificationChannelWithDefaults instantiates a new OpsgenieNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpsgenieNotificationChannel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsgenieNotificationChannel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsgenieNotificationChannel) SetId(v int64)`

SetId sets Id field to given value.


### GetNotificationConfigurationId

`func (o *OpsgenieNotificationChannel) GetNotificationConfigurationId() int64`

GetNotificationConfigurationId returns the NotificationConfigurationId field if non-nil, zero value otherwise.

### GetNotificationConfigurationIdOk

`func (o *OpsgenieNotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool)`

GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfigurationId

`func (o *OpsgenieNotificationChannel) SetNotificationConfigurationId(v int64)`

SetNotificationConfigurationId sets NotificationConfigurationId field to given value.

### HasNotificationConfigurationId

`func (o *OpsgenieNotificationChannel) HasNotificationConfigurationId() bool`

HasNotificationConfigurationId returns a boolean if a field has been set.

### GetStatus

`func (o *OpsgenieNotificationChannel) GetStatus() NotificationChannelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsgenieNotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsgenieNotificationChannel) SetStatus(v NotificationChannelStatus)`

SetStatus sets Status field to given value.


### GetRegion

`func (o *OpsgenieNotificationChannel) GetRegion() OpsgenieRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OpsgenieNotificationChannel) GetRegionOk() (*OpsgenieRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OpsgenieNotificationChannel) SetRegion(v OpsgenieRegion)`

SetRegion sets Region field to given value.


### GetGenieKey

`func (o *OpsgenieNotificationChannel) GetGenieKey() string`

GetGenieKey returns the GenieKey field if non-nil, zero value otherwise.

### GetGenieKeyOk

`func (o *OpsgenieNotificationChannel) GetGenieKeyOk() (*string, bool)`

GetGenieKeyOk returns a tuple with the GenieKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenieKey

`func (o *OpsgenieNotificationChannel) SetGenieKey(v string)`

SetGenieKey sets GenieKey field to given value.


### GetResponders

`func (o *OpsgenieNotificationChannel) GetResponders() []OpsgenieResponder`

GetResponders returns the Responders field if non-nil, zero value otherwise.

### GetRespondersOk

`func (o *OpsgenieNotificationChannel) GetRespondersOk() (*[]OpsgenieResponder, bool)`

GetRespondersOk returns a tuple with the Responders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponders

`func (o *OpsgenieNotificationChannel) SetResponders(v []OpsgenieResponder)`

SetResponders sets Responders field to given value.


### GetPriority

`func (o *OpsgenieNotificationChannel) GetPriority() OpsgeniePriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OpsgenieNotificationChannel) GetPriorityOk() (*OpsgeniePriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OpsgenieNotificationChannel) SetPriority(v OpsgeniePriority)`

SetPriority sets Priority field to given value.


### GetType

`func (o *OpsgenieNotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpsgenieNotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpsgenieNotificationChannel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


