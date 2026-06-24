# MetricProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | **string** | Individual metric query that returns a timeseries for a specific cell. | 
**Unit** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Sparkline** | Pointer to **bool** |  | [optional] 
**MetricId** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricProjection

`func NewMetricProjection(type_ string, query string, ) *MetricProjection`

NewMetricProjection instantiates a new MetricProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricProjectionWithDefaults

`func NewMetricProjectionWithDefaults() *MetricProjection`

NewMetricProjectionWithDefaults instantiates a new MetricProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricProjection) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *MetricProjection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricProjection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricProjection) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUnit

`func (o *MetricProjection) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricProjection) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricProjection) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricProjection) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *MetricProjection) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *MetricProjection) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *MetricProjection) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *MetricProjection) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetSparkline

`func (o *MetricProjection) GetSparkline() bool`

GetSparkline returns the Sparkline field if non-nil, zero value otherwise.

### GetSparklineOk

`func (o *MetricProjection) GetSparklineOk() (*bool, bool)`

GetSparklineOk returns a tuple with the Sparkline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkline

`func (o *MetricProjection) SetSparkline(v bool)`

SetSparkline sets Sparkline field to given value.

### HasSparkline

`func (o *MetricProjection) HasSparkline() bool`

HasSparkline returns a boolean if a field has been set.

### GetMetricId

`func (o *MetricProjection) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricProjection) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricProjection) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *MetricProjection) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


