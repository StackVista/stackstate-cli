# EventTelemetryStreamDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**BindQuery** | **string** |  | 
**DataSourceId** | **int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Labels** | **[]string** |  | 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Priority** | [**TelemetryStreamPriority**](TelemetryStreamPriority.md) |  | 
**TelemetryQuery** | [**EventTelemetryQuery**](EventTelemetryQuery.md) |  | 
**TopologyMapping** | [**[]TopologyMapping**](TopologyMapping.md) |  | 

## Methods

### NewEventTelemetryStreamDefinition

`func NewEventTelemetryStreamDefinition(type_ string, bindQuery string, dataSourceId int64, labels []string, name string, priority TelemetryStreamPriority, telemetryQuery EventTelemetryQuery, topologyMapping []TopologyMapping, ) *EventTelemetryStreamDefinition`

NewEventTelemetryStreamDefinition instantiates a new EventTelemetryStreamDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTelemetryStreamDefinitionWithDefaults

`func NewEventTelemetryStreamDefinitionWithDefaults() *EventTelemetryStreamDefinition`

NewEventTelemetryStreamDefinitionWithDefaults instantiates a new EventTelemetryStreamDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventTelemetryStreamDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventTelemetryStreamDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventTelemetryStreamDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetBindQuery

`func (o *EventTelemetryStreamDefinition) GetBindQuery() string`

GetBindQuery returns the BindQuery field if non-nil, zero value otherwise.

### GetBindQueryOk

`func (o *EventTelemetryStreamDefinition) GetBindQueryOk() (*string, bool)`

GetBindQueryOk returns a tuple with the BindQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindQuery

`func (o *EventTelemetryStreamDefinition) SetBindQuery(v string)`

SetBindQuery sets BindQuery field to given value.


### GetDataSourceId

`func (o *EventTelemetryStreamDefinition) GetDataSourceId() int64`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *EventTelemetryStreamDefinition) GetDataSourceIdOk() (*int64, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *EventTelemetryStreamDefinition) SetDataSourceId(v int64)`

SetDataSourceId sets DataSourceId field to given value.


### GetDescription

`func (o *EventTelemetryStreamDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventTelemetryStreamDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventTelemetryStreamDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventTelemetryStreamDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EventTelemetryStreamDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTelemetryStreamDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTelemetryStreamDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventTelemetryStreamDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *EventTelemetryStreamDefinition) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EventTelemetryStreamDefinition) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EventTelemetryStreamDefinition) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EventTelemetryStreamDefinition) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLabels

`func (o *EventTelemetryStreamDefinition) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EventTelemetryStreamDefinition) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EventTelemetryStreamDefinition) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLastUpdateTimestamp

`func (o *EventTelemetryStreamDefinition) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *EventTelemetryStreamDefinition) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *EventTelemetryStreamDefinition) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *EventTelemetryStreamDefinition) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *EventTelemetryStreamDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventTelemetryStreamDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventTelemetryStreamDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOwnedBy

`func (o *EventTelemetryStreamDefinition) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *EventTelemetryStreamDefinition) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *EventTelemetryStreamDefinition) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *EventTelemetryStreamDefinition) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetPriority

`func (o *EventTelemetryStreamDefinition) GetPriority() TelemetryStreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EventTelemetryStreamDefinition) GetPriorityOk() (*TelemetryStreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EventTelemetryStreamDefinition) SetPriority(v TelemetryStreamPriority)`

SetPriority sets Priority field to given value.


### GetTelemetryQuery

`func (o *EventTelemetryStreamDefinition) GetTelemetryQuery() EventTelemetryQuery`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *EventTelemetryStreamDefinition) GetTelemetryQueryOk() (*EventTelemetryQuery, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *EventTelemetryStreamDefinition) SetTelemetryQuery(v EventTelemetryQuery)`

SetTelemetryQuery sets TelemetryQuery field to given value.


### GetTopologyMapping

`func (o *EventTelemetryStreamDefinition) GetTopologyMapping() []TopologyMapping`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *EventTelemetryStreamDefinition) GetTopologyMappingOk() (*[]TopologyMapping, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *EventTelemetryStreamDefinition) SetTopologyMapping(v []TopologyMapping)`

SetTopologyMapping sets TopologyMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


