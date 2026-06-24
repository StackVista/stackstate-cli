# MetricMetaDisplay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **NullableString** |  | [optional] 
**DecimalPlaces** | Pointer to **NullableInt32** |  | [optional] 
**Sparkline** | Pointer to **NullableBool** |  | [optional] 
**Locked** | **bool** |  | 
**BoundMetricId** | Pointer to [**BoundMetricId**](BoundMetricId.md) |  | [optional] 

## Methods

### NewMetricMetaDisplay

`func NewMetricMetaDisplay(type_ string, locked bool, ) *MetricMetaDisplay`

NewMetricMetaDisplay instantiates a new MetricMetaDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMetaDisplayWithDefaults

`func NewMetricMetaDisplayWithDefaults() *MetricMetaDisplay`

NewMetricMetaDisplayWithDefaults instantiates a new MetricMetaDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricMetaDisplay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricMetaDisplay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricMetaDisplay) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *MetricMetaDisplay) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricMetaDisplay) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricMetaDisplay) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricMetaDisplay) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *MetricMetaDisplay) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *MetricMetaDisplay) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDecimalPlaces

`func (o *MetricMetaDisplay) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *MetricMetaDisplay) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *MetricMetaDisplay) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *MetricMetaDisplay) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *MetricMetaDisplay) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *MetricMetaDisplay) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetSparkline

`func (o *MetricMetaDisplay) GetSparkline() bool`

GetSparkline returns the Sparkline field if non-nil, zero value otherwise.

### GetSparklineOk

`func (o *MetricMetaDisplay) GetSparklineOk() (*bool, bool)`

GetSparklineOk returns a tuple with the Sparkline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkline

`func (o *MetricMetaDisplay) SetSparkline(v bool)`

SetSparkline sets Sparkline field to given value.

### HasSparkline

`func (o *MetricMetaDisplay) HasSparkline() bool`

HasSparkline returns a boolean if a field has been set.

### SetSparklineNil

`func (o *MetricMetaDisplay) SetSparklineNil(b bool)`

 SetSparklineNil sets the value for Sparkline to be an explicit nil

### UnsetSparkline
`func (o *MetricMetaDisplay) UnsetSparkline()`

UnsetSparkline ensures that no value is present for Sparkline, not even an explicit nil
### GetLocked

`func (o *MetricMetaDisplay) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MetricMetaDisplay) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MetricMetaDisplay) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetBoundMetricId

`func (o *MetricMetaDisplay) GetBoundMetricId() BoundMetricId`

GetBoundMetricId returns the BoundMetricId field if non-nil, zero value otherwise.

### GetBoundMetricIdOk

`func (o *MetricMetaDisplay) GetBoundMetricIdOk() (*BoundMetricId, bool)`

GetBoundMetricIdOk returns a tuple with the BoundMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundMetricId

`func (o *MetricMetaDisplay) SetBoundMetricId(v BoundMetricId)`

SetBoundMetricId sets BoundMetricId field to given value.

### HasBoundMetricId

`func (o *MetricMetaDisplay) HasBoundMetricId() bool`

HasBoundMetricId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


