# EventStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**DataSource** | **int64** |  | 
**DataType** | [**DataType**](DataType.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Priority** | Pointer to [**StreamPriority**](StreamPriority.md) |  | [optional] 
**Query** | [**EventTelemetryQuery**](EventTelemetryQuery.md) |  | 
**SyncCreated** | **bool** |  | 

## Methods

### NewEventStream

`func NewEventStream(type_ string, dataSource int64, dataType DataType, name string, query EventTelemetryQuery, syncCreated bool, ) *EventStream`

NewEventStream instantiates a new EventStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamWithDefaults

`func NewEventStreamWithDefaults() *EventStream`

NewEventStreamWithDefaults instantiates a new EventStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventStream) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventStream) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventStream) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *EventStream) GetDataSource() int64`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *EventStream) GetDataSourceOk() (*int64, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *EventStream) SetDataSource(v int64)`

SetDataSource sets DataSource field to given value.


### GetDataType

`func (o *EventStream) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EventStream) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EventStream) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetDescription

`func (o *EventStream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventStream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventStream) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventStream) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EventStream) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventStream) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventStream) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventStream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *EventStream) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *EventStream) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *EventStream) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *EventStream) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *EventStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventStream) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *EventStream) GetPriority() StreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EventStream) GetPriorityOk() (*StreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EventStream) SetPriority(v StreamPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EventStream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *EventStream) GetQuery() EventTelemetryQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventStream) GetQueryOk() (*EventTelemetryQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventStream) SetQuery(v EventTelemetryQuery)`

SetQuery sets Query field to given value.


### GetSyncCreated

`func (o *EventStream) GetSyncCreated() bool`

GetSyncCreated returns the SyncCreated field if non-nil, zero value otherwise.

### GetSyncCreatedOk

`func (o *EventStream) GetSyncCreatedOk() (*bool, bool)`

GetSyncCreatedOk returns a tuple with the SyncCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCreated

`func (o *EventStream) SetSyncCreated(v bool)`

SetSyncCreated sets SyncCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


