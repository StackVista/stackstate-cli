# NotificationConfigurationReadSchema

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
**Id** | **int64** |  | 
**LastUpdateTimestamp** | **int64** |  | 
**RuntimeStatus** | [**NotificationConfigurationRuntimeStatusValue**](NotificationConfigurationRuntimeStatusValue.md) |  | 

## Methods

### NewNotificationConfigurationReadSchema

`func NewNotificationConfigurationReadSchema(name string, notifyHealthStates NotifyOnOptions, monitors []MonitorReferenceId, monitorTags []string, componentTypes []int64, componentTags []string, status NotificationConfigurationStatusValue, notificationChannels []ChannelReferenceId, id int64, lastUpdateTimestamp int64, runtimeStatus NotificationConfigurationRuntimeStatusValue, ) *NotificationConfigurationReadSchema`

NewNotificationConfigurationReadSchema instantiates a new NotificationConfigurationReadSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigurationReadSchemaWithDefaults

`func NewNotificationConfigurationReadSchemaWithDefaults() *NotificationConfigurationReadSchema`

NewNotificationConfigurationReadSchemaWithDefaults instantiates a new NotificationConfigurationReadSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationConfigurationReadSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationConfigurationReadSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationConfigurationReadSchema) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *NotificationConfigurationReadSchema) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NotificationConfigurationReadSchema) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NotificationConfigurationReadSchema) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NotificationConfigurationReadSchema) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationConfigurationReadSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationConfigurationReadSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationConfigurationReadSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationConfigurationReadSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyHealthStates

`func (o *NotificationConfigurationReadSchema) GetNotifyHealthStates() NotifyOnOptions`

GetNotifyHealthStates returns the NotifyHealthStates field if non-nil, zero value otherwise.

### GetNotifyHealthStatesOk

`func (o *NotificationConfigurationReadSchema) GetNotifyHealthStatesOk() (*NotifyOnOptions, bool)`

GetNotifyHealthStatesOk returns a tuple with the NotifyHealthStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyHealthStates

`func (o *NotificationConfigurationReadSchema) SetNotifyHealthStates(v NotifyOnOptions)`

SetNotifyHealthStates sets NotifyHealthStates field to given value.


### GetMonitors

`func (o *NotificationConfigurationReadSchema) GetMonitors() []MonitorReferenceId`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *NotificationConfigurationReadSchema) GetMonitorsOk() (*[]MonitorReferenceId, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *NotificationConfigurationReadSchema) SetMonitors(v []MonitorReferenceId)`

SetMonitors sets Monitors field to given value.


### GetMonitorTags

`func (o *NotificationConfigurationReadSchema) GetMonitorTags() []string`

GetMonitorTags returns the MonitorTags field if non-nil, zero value otherwise.

### GetMonitorTagsOk

`func (o *NotificationConfigurationReadSchema) GetMonitorTagsOk() (*[]string, bool)`

GetMonitorTagsOk returns a tuple with the MonitorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTags

`func (o *NotificationConfigurationReadSchema) SetMonitorTags(v []string)`

SetMonitorTags sets MonitorTags field to given value.


### GetComponentTypes

`func (o *NotificationConfigurationReadSchema) GetComponentTypes() []int64`

GetComponentTypes returns the ComponentTypes field if non-nil, zero value otherwise.

### GetComponentTypesOk

`func (o *NotificationConfigurationReadSchema) GetComponentTypesOk() (*[]int64, bool)`

GetComponentTypesOk returns a tuple with the ComponentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTypes

`func (o *NotificationConfigurationReadSchema) SetComponentTypes(v []int64)`

SetComponentTypes sets ComponentTypes field to given value.


### GetComponentTags

`func (o *NotificationConfigurationReadSchema) GetComponentTags() []string`

GetComponentTags returns the ComponentTags field if non-nil, zero value otherwise.

### GetComponentTagsOk

`func (o *NotificationConfigurationReadSchema) GetComponentTagsOk() (*[]string, bool)`

GetComponentTagsOk returns a tuple with the ComponentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTags

`func (o *NotificationConfigurationReadSchema) SetComponentTags(v []string)`

SetComponentTags sets ComponentTags field to given value.


### GetStatus

`func (o *NotificationConfigurationReadSchema) GetStatus() NotificationConfigurationStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationConfigurationReadSchema) GetStatusOk() (*NotificationConfigurationStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationConfigurationReadSchema) SetStatus(v NotificationConfigurationStatusValue)`

SetStatus sets Status field to given value.


### GetNotificationChannels

`func (o *NotificationConfigurationReadSchema) GetNotificationChannels() []ChannelReferenceId`

GetNotificationChannels returns the NotificationChannels field if non-nil, zero value otherwise.

### GetNotificationChannelsOk

`func (o *NotificationConfigurationReadSchema) GetNotificationChannelsOk() (*[]ChannelReferenceId, bool)`

GetNotificationChannelsOk returns a tuple with the NotificationChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannels

`func (o *NotificationConfigurationReadSchema) SetNotificationChannels(v []ChannelReferenceId)`

SetNotificationChannels sets NotificationChannels field to given value.


### GetId

`func (o *NotificationConfigurationReadSchema) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationConfigurationReadSchema) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationConfigurationReadSchema) SetId(v int64)`

SetId sets Id field to given value.


### GetLastUpdateTimestamp

`func (o *NotificationConfigurationReadSchema) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *NotificationConfigurationReadSchema) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *NotificationConfigurationReadSchema) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.


### GetRuntimeStatus

`func (o *NotificationConfigurationReadSchema) GetRuntimeStatus() NotificationConfigurationRuntimeStatusValue`

GetRuntimeStatus returns the RuntimeStatus field if non-nil, zero value otherwise.

### GetRuntimeStatusOk

`func (o *NotificationConfigurationReadSchema) GetRuntimeStatusOk() (*NotificationConfigurationRuntimeStatusValue, bool)`

GetRuntimeStatusOk returns a tuple with the RuntimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStatus

`func (o *NotificationConfigurationReadSchema) SetRuntimeStatus(v NotificationConfigurationRuntimeStatusValue)`

SetRuntimeStatus sets RuntimeStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


