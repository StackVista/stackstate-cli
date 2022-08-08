# DataStream

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
**Query** | [**MetricTelemetryQuery**](MetricTelemetryQuery.md) |  | 
**SyncCreated** | **bool** |  | 
**StaleAfter** | Pointer to **int64** |  | [optional] 

## Methods

### NewDataStream

`func NewDataStream(type_ string, dataSource int64, dataType DataType, name string, query MetricTelemetryQuery, syncCreated bool, ) *DataStream`

NewDataStream instantiates a new DataStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamWithDefaults

`func NewDataStreamWithDefaults() *DataStream`

NewDataStreamWithDefaults instantiates a new DataStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStream) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStream) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStream) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *DataStream) GetDataSource() int64`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataStream) GetDataSourceOk() (*int64, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataStream) SetDataSource(v int64)`

SetDataSource sets DataSource field to given value.


### GetDataType

`func (o *DataStream) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DataStream) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DataStream) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetDescription

`func (o *DataStream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataStream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataStream) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataStream) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DataStream) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStream) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStream) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DataStream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *DataStream) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *DataStream) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *DataStream) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *DataStream) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *DataStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStream) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *DataStream) GetPriority() StreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DataStream) GetPriorityOk() (*StreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DataStream) SetPriority(v StreamPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DataStream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *DataStream) GetQuery() MetricTelemetryQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DataStream) GetQueryOk() (*MetricTelemetryQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DataStream) SetQuery(v MetricTelemetryQuery)`

SetQuery sets Query field to given value.


### GetSyncCreated

`func (o *DataStream) GetSyncCreated() bool`

GetSyncCreated returns the SyncCreated field if non-nil, zero value otherwise.

### GetSyncCreatedOk

`func (o *DataStream) GetSyncCreatedOk() (*bool, bool)`

GetSyncCreatedOk returns a tuple with the SyncCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCreated

`func (o *DataStream) SetSyncCreated(v bool)`

SetSyncCreated sets SyncCreated field to given value.


### GetStaleAfter

`func (o *DataStream) GetStaleAfter() int64`

GetStaleAfter returns the StaleAfter field if non-nil, zero value otherwise.

### GetStaleAfterOk

`func (o *DataStream) GetStaleAfterOk() (*int64, bool)`

GetStaleAfterOk returns a tuple with the StaleAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAfter

`func (o *DataStream) SetStaleAfter(v int64)`

SetStaleAfter sets StaleAfter field to given value.

### HasStaleAfter

`func (o *DataStream) HasStaleAfter() bool`

HasStaleAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


