# NotificationConfigurationWriteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotifyHealthStates** | [**NotifyOnOptions**](NotifyOnOptions.md) |  | 
**Monitors** | [**[]MonitorReferenceId**](MonitorReferenceId.md) |  | 
**MonitorTags** | **[]string** |  | 
**ComponentTypes** | **[]int64** |  | 
**ComponentTags** | **[]string** |  | 
**Status** | [**NotificationConfigurationStatusValue**](NotificationConfigurationStatusValue.md) |  | 
**NotificationChannels** | [**[]ChannelReferenceId**](ChannelReferenceId.md) |  | 

## Methods

### NewNotificationConfigurationWriteSchema

`func NewNotificationConfigurationWriteSchema(name string, notifyHealthStates NotifyOnOptions, monitors []MonitorReferenceId, monitorTags []string, componentTypes []int64, componentTags []string, status NotificationConfigurationStatusValue, notificationChannels []ChannelReferenceId, ) *NotificationConfigurationWriteSchema`

NewNotificationConfigurationWriteSchema instantiates a new NotificationConfigurationWriteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigurationWriteSchemaWithDefaults

`func NewNotificationConfigurationWriteSchemaWithDefaults() *NotificationConfigurationWriteSchema`

NewNotificationConfigurationWriteSchemaWithDefaults instantiates a new NotificationConfigurationWriteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationConfigurationWriteSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationConfigurationWriteSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationConfigurationWriteSchema) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *NotificationConfigurationWriteSchema) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NotificationConfigurationWriteSchema) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NotificationConfigurationWriteSchema) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NotificationConfigurationWriteSchema) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationConfigurationWriteSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationConfigurationWriteSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationConfigurationWriteSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationConfigurationWriteSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyHealthStates

`func (o *NotificationConfigurationWriteSchema) GetNotifyHealthStates() NotifyOnOptions`

GetNotifyHealthStates returns the NotifyHealthStates field if non-nil, zero value otherwise.

### GetNotifyHealthStatesOk

`func (o *NotificationConfigurationWriteSchema) GetNotifyHealthStatesOk() (*NotifyOnOptions, bool)`

GetNotifyHealthStatesOk returns a tuple with the NotifyHealthStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyHealthStates

`func (o *NotificationConfigurationWriteSchema) SetNotifyHealthStates(v NotifyOnOptions)`

SetNotifyHealthStates sets NotifyHealthStates field to given value.


### GetMonitors

`func (o *NotificationConfigurationWriteSchema) GetMonitors() []MonitorReferenceId`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *NotificationConfigurationWriteSchema) GetMonitorsOk() (*[]MonitorReferenceId, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *NotificationConfigurationWriteSchema) SetMonitors(v []MonitorReferenceId)`

SetMonitors sets Monitors field to given value.


### GetMonitorTags

`func (o *NotificationConfigurationWriteSchema) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *NotificationConfigurationWriteSchema) GetMonitorTagsOk() (*[]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTags

`func (o *NotificationConfigurationWriteSchema) SetMonitorTags(v []string)`

SetMonitorTags sets MonitorTags field to given value.


### GetComponentTypes

`func (o *NotificationConfigurationWriteSchema) GetComponentTypes() []int64`

GetComponentTypes returns the ComponentTypes field if non-nil, zero value otherwise.

### GetComponentTypesOk

`func (o *NotificationConfigurationWriteSchema) GetComponentTypesOk() (*[]int64, bool)`

GetComponentTypesOk returns a tuple with the ComponentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTypes

`func (o *NotificationConfigurationWriteSchema) SetComponentTypes(v []int64)`

SetComponentTypes sets ComponentTypes field to given value.


### GetComponentTags

`func (o *NotificationConfigurationWriteSchema) GetComponentTags() []string`

GetComponentTags returns the ComponentTags field if non-nil, zero value otherwise.

### GetComponentTagsOk

`func (o *NotificationConfigurationWriteSchema) GetComponentTagsOk() (*[]string, bool)`

GetComponentTagsOk returns a tuple with the ComponentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTags

`func (o *NotificationConfigurationWriteSchema) SetComponentTags(v []string)`

SetComponentTags sets ComponentTags field to given value.


### GetStatus

`func (o *NotificationConfigurationWriteSchema) GetStatus() NotificationConfigurationStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationConfigurationWriteSchema) GetStatusOk() (*NotificationConfigurationStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationConfigurationWriteSchema) SetStatus(v NotificationConfigurationStatusValue)`

SetStatus sets Status field to given value.


### GetNotificationChannels

`func (o *NotificationConfigurationWriteSchema) GetNotificationChannels() []ChannelReferenceId`

GetNotificationChannels returns the NotificationChannels field if non-nil, zero value otherwise.

### GetNotificationChannelsOk

`func (o *NotificationConfigurationWriteSchema) GetNotificationChannelsOk() (*[]ChannelReferenceId, bool)`

GetNotificationChannelsOk returns a tuple with the NotificationChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannels

`func (o *NotificationConfigurationWriteSchema) SetNotificationChannels(v []ChannelReferenceId)`

SetNotificationChannels sets NotificationChannels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


