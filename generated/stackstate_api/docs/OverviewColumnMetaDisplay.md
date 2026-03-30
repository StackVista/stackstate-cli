# OverviewColumnMetaDisplay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **NullableString** |  | [optional] 
**DecimalPlaces** | Pointer to **NullableInt32** |  | [optional] 
**ShowChart** | Pointer to **NullableBool** |  | [optional] 
**Locked** | **bool** |  | 
**MetricId** | Pointer to [**MetricBindingId**](MetricBindingId.md) |  | [optional] 
**External** | **bool** |  | 

## Methods

### NewOverviewColumnMetaDisplay

`func NewOverviewColumnMetaDisplay(type_ string, locked bool, external bool, ) *OverviewColumnMetaDisplay`

NewOverviewColumnMetaDisplay instantiates a new OverviewColumnMetaDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewColumnMetaDisplayWithDefaults

`func NewOverviewColumnMetaDisplayWithDefaults() *OverviewColumnMetaDisplay`

NewOverviewColumnMetaDisplayWithDefaults instantiates a new OverviewColumnMetaDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewColumnMetaDisplay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewColumnMetaDisplay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewColumnMetaDisplay) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *OverviewColumnMetaDisplay) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OverviewColumnMetaDisplay) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OverviewColumnMetaDisplay) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OverviewColumnMetaDisplay) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *OverviewColumnMetaDisplay) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *OverviewColumnMetaDisplay) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDecimalPlaces

`func (o *OverviewColumnMetaDisplay) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *OverviewColumnMetaDisplay) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *OverviewColumnMetaDisplay) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *OverviewColumnMetaDisplay) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *OverviewColumnMetaDisplay) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *OverviewColumnMetaDisplay) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetShowChart

`func (o *OverviewColumnMetaDisplay) GetShowChart() bool`

GetShowChart returns the ShowChart field if non-nil, zero value otherwise.

### GetShowChartOk

`func (o *OverviewColumnMetaDisplay) GetShowChartOk() (*bool, bool)`

GetShowChartOk returns a tuple with the ShowChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowChart

`func (o *OverviewColumnMetaDisplay) SetShowChart(v bool)`

SetShowChart sets ShowChart field to given value.

### HasShowChart

`func (o *OverviewColumnMetaDisplay) HasShowChart() bool`

HasShowChart returns a boolean if a field has been set.

### SetShowChartNil

`func (o *OverviewColumnMetaDisplay) SetShowChartNil(b bool)`

 SetShowChartNil sets the value for ShowChart to be an explicit nil

### UnsetShowChart
`func (o *OverviewColumnMetaDisplay) UnsetShowChart()`

UnsetShowChart ensures that no value is present for ShowChart, not even an explicit nil
### GetLocked

`func (o *OverviewColumnMetaDisplay) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OverviewColumnMetaDisplay) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OverviewColumnMetaDisplay) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMetricId

`func (o *OverviewColumnMetaDisplay) GetMetricId() MetricBindingId`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *OverviewColumnMetaDisplay) GetMetricIdOk() (*MetricBindingId, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *OverviewColumnMetaDisplay) SetMetricId(v MetricBindingId)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *OverviewColumnMetaDisplay) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetExternal

`func (o *OverviewColumnMetaDisplay) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *OverviewColumnMetaDisplay) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *OverviewColumnMetaDisplay) SetExternal(v bool)`

SetExternal sets External field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


