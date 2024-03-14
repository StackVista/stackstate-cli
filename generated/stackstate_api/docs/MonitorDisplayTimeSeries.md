# MonitorDisplayTimeSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Queries** | [**[]MonitorDisplayQuery**](MonitorDisplayQuery.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitorDisplayTimeSeries

`func NewMonitorDisplayTimeSeries(name string, queries []MonitorDisplayQuery, ) *MonitorDisplayTimeSeries`

NewMonitorDisplayTimeSeries instantiates a new MonitorDisplayTimeSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorDisplayTimeSeriesWithDefaults

`func NewMonitorDisplayTimeSeriesWithDefaults() *MonitorDisplayTimeSeries`

NewMonitorDisplayTimeSeriesWithDefaults instantiates a new MonitorDisplayTimeSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorDisplayTimeSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorDisplayTimeSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorDisplayTimeSeries) SetName(v string)`

SetName sets Name field to given value.


### GetQueries

`func (o *MonitorDisplayTimeSeries) GetQueries() []MonitorDisplayQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorDisplayTimeSeries) GetQueriesOk() (*[]MonitorDisplayQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorDisplayTimeSeries) SetQueries(v []MonitorDisplayQuery)`

SetQueries sets Queries field to given value.


### GetDescription

`func (o *MonitorDisplayTimeSeries) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorDisplayTimeSeries) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorDisplayTimeSeries) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorDisplayTimeSeries) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *MonitorDisplayTimeSeries) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MonitorDisplayTimeSeries) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MonitorDisplayTimeSeries) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MonitorDisplayTimeSeries) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


