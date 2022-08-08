# MetricStream

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
**StaleAfter** | Pointer to **int64** |  | [optional] 
**SyncCreated** | **bool** |  | 

## Methods

### NewMetricStream

`func NewMetricStream(type_ string, dataSource int64, dataType DataType, name string, query MetricTelemetryQuery, syncCreated bool, ) *MetricStream`

NewMetricStream instantiates a new MetricStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStreamWithDefaults

`func NewMetricStreamWithDefaults() *MetricStream`

NewMetricStreamWithDefaults instantiates a new MetricStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricStream) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricStream) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricStream) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *MetricStream) GetDataSource() int64`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MetricStream) GetDataSourceOk() (*int64, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MetricStream) SetDataSource(v int64)`

SetDataSource sets DataSource field to given value.


### GetDataType

`func (o *MetricStream) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MetricStream) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MetricStream) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetDescription

`func (o *MetricStream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricStream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricStream) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricStream) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MetricStream) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricStream) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricStream) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetricStream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *MetricStream) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *MetricStream) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *MetricStream) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *MetricStream) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *MetricStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricStream) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *MetricStream) GetPriority() StreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MetricStream) GetPriorityOk() (*StreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MetricStream) SetPriority(v StreamPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MetricStream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *MetricStream) GetQuery() MetricTelemetryQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricStream) GetQueryOk() (*MetricTelemetryQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricStream) SetQuery(v MetricTelemetryQuery)`

SetQuery sets Query field to given value.


### GetStaleAfter

`func (o *MetricStream) GetStaleAfter() int64`

GetStaleAfter returns the StaleAfter field if non-nil, zero value otherwise.

### GetStaleAfterOk

`func (o *MetricStream) GetStaleAfterOk() (*int64, bool)`

GetStaleAfterOk returns a tuple with the StaleAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAfter

`func (o *MetricStream) SetStaleAfter(v int64)`

SetStaleAfter sets StaleAfter field to given value.

### HasStaleAfter

`func (o *MetricStream) HasStaleAfter() bool`

HasStaleAfter returns a boolean if a field has been set.

### GetSyncCreated

`func (o *MetricStream) GetSyncCreated() bool`

GetSyncCreated returns the SyncCreated field if non-nil, zero value otherwise.

### GetSyncCreatedOk

`func (o *MetricStream) GetSyncCreatedOk() (*bool, bool)`

GetSyncCreatedOk returns a tuple with the SyncCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCreated

`func (o *MetricStream) SetSyncCreated(v bool)`

SetSyncCreated sets SyncCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


