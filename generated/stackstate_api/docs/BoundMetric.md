# BoundMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundMetricId** | Pointer to [**BoundMetricId**](BoundMetricId.md) |  | [optional] 
**Name** | **string** |  | 
**BoundQueries** | [**[]BoundMetricQuery**](BoundMetricQuery.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Dummy** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoundMetric

`func NewBoundMetric(name string, boundQueries []BoundMetricQuery, ) *BoundMetric`

NewBoundMetric instantiates a new BoundMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundMetricWithDefaults

`func NewBoundMetricWithDefaults() *BoundMetric`

NewBoundMetricWithDefaults instantiates a new BoundMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundMetricId

`func (o *BoundMetric) GetBoundMetricId() BoundMetricId`

GetBoundMetricId returns the BoundMetricId field if non-nil, zero value otherwise.

### GetBoundMetricIdOk

`func (o *BoundMetric) GetBoundMetricIdOk() (*BoundMetricId, bool)`

GetBoundMetricIdOk returns a tuple with the BoundMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundMetricId

`func (o *BoundMetric) SetBoundMetricId(v BoundMetricId)`

SetBoundMetricId sets BoundMetricId field to given value.

### HasBoundMetricId

`func (o *BoundMetric) HasBoundMetricId() bool`

HasBoundMetricId returns a boolean if a field has been set.

### GetName

`func (o *BoundMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundMetric) SetName(v string)`

SetName sets Name field to given value.


### GetBoundQueries

`func (o *BoundMetric) GetBoundQueries() []BoundMetricQuery`

GetBoundQueries returns the BoundQueries field if non-nil, zero value otherwise.

### GetBoundQueriesOk

`func (o *BoundMetric) GetBoundQueriesOk() (*[]BoundMetricQuery, bool)`

GetBoundQueriesOk returns a tuple with the BoundQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundQueries

`func (o *BoundMetric) SetBoundQueries(v []BoundMetricQuery)`

SetBoundQueries sets BoundQueries field to given value.


### GetDescription

`func (o *BoundMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BoundMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BoundMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BoundMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChart

`func (o *BoundMetric) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *BoundMetric) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *BoundMetric) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *BoundMetric) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetDummy

`func (o *BoundMetric) GetDummy() bool`

GetDummy returns the Dummy field if non-nil, zero value otherwise.

### GetDummyOk

`func (o *BoundMetric) GetDummyOk() (*bool, bool)`

GetDummyOk returns a tuple with the Dummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummy

`func (o *BoundMetric) SetDummy(v bool)`

SetDummy sets Dummy field to given value.

### HasDummy

`func (o *BoundMetric) HasDummy() bool`

HasDummy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


