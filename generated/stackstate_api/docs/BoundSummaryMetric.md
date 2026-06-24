# BoundSummaryMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundMetricId** | Pointer to [**BoundMetricId**](BoundMetricId.md) |  | [optional] 
**Name** | **string** |  | 
**Query** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Dummy** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoundSummaryMetric

`func NewBoundSummaryMetric(name string, query string, ) *BoundSummaryMetric`

NewBoundSummaryMetric instantiates a new BoundSummaryMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundSummaryMetricWithDefaults

`func NewBoundSummaryMetricWithDefaults() *BoundSummaryMetric`

NewBoundSummaryMetricWithDefaults instantiates a new BoundSummaryMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundMetricId

`func (o *BoundSummaryMetric) GetBoundMetricId() BoundMetricId`

GetBoundMetricId returns the BoundMetricId field if non-nil, zero value otherwise.

### GetBoundMetricIdOk

`func (o *BoundSummaryMetric) GetBoundMetricIdOk() (*BoundMetricId, bool)`

GetBoundMetricIdOk returns a tuple with the BoundMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundMetricId

`func (o *BoundSummaryMetric) SetBoundMetricId(v BoundMetricId)`

SetBoundMetricId sets BoundMetricId field to given value.

### HasBoundMetricId

`func (o *BoundSummaryMetric) HasBoundMetricId() bool`

HasBoundMetricId returns a boolean if a field has been set.

### GetName

`func (o *BoundSummaryMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoundSummaryMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoundSummaryMetric) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *BoundSummaryMetric) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BoundSummaryMetric) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BoundSummaryMetric) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUnit

`func (o *BoundSummaryMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BoundSummaryMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BoundSummaryMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BoundSummaryMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *BoundSummaryMetric) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *BoundSummaryMetric) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *BoundSummaryMetric) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *BoundSummaryMetric) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetDummy

`func (o *BoundSummaryMetric) GetDummy() bool`

GetDummy returns the Dummy field if non-nil, zero value otherwise.

### GetDummyOk

`func (o *BoundSummaryMetric) GetDummyOk() (*bool, bool)`

GetDummyOk returns a tuple with the Dummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummy

`func (o *BoundSummaryMetric) SetDummy(v bool)`

SetDummy sets Dummy field to given value.

### HasDummy

`func (o *BoundSummaryMetric) HasDummy() bool`

HasDummy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


