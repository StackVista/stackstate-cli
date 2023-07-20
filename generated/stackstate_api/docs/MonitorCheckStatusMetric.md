# MonitorCheckStatusMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**Queries** | [**[]MonitorCheckStatusQuery**](MonitorCheckStatusQuery.md) |  | 

## Methods

### NewMonitorCheckStatusMetric

`func NewMonitorCheckStatusMetric(type_ string, name string, queries []MonitorCheckStatusQuery, ) *MonitorCheckStatusMetric`

NewMonitorCheckStatusMetric instantiates a new MonitorCheckStatusMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusMetricWithDefaults

`func NewMonitorCheckStatusMetricWithDefaults() *MonitorCheckStatusMetric`

NewMonitorCheckStatusMetricWithDefaults instantiates a new MonitorCheckStatusMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorCheckStatusMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorCheckStatusMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorCheckStatusMetric) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *MonitorCheckStatusMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorCheckStatusMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorCheckStatusMetric) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MonitorCheckStatusMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorCheckStatusMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorCheckStatusMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorCheckStatusMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *MonitorCheckStatusMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MonitorCheckStatusMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MonitorCheckStatusMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MonitorCheckStatusMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStep

`func (o *MonitorCheckStatusMetric) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *MonitorCheckStatusMetric) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *MonitorCheckStatusMetric) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *MonitorCheckStatusMetric) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorCheckStatusMetric) GetQueries() []MonitorCheckStatusQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorCheckStatusMetric) GetQueriesOk() (*[]MonitorCheckStatusQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorCheckStatusMetric) SetQueries(v []MonitorCheckStatusQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


