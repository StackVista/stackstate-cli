# MetricChartProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ShowChart** | Pointer to **bool** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Query** | **string** | Individual metric query that returns a timeseries for a specific cell. | 

## Methods

### NewMetricChartProjection

`func NewMetricChartProjection(type_ string, query string, ) *MetricChartProjection`

NewMetricChartProjection instantiates a new MetricChartProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricChartProjectionWithDefaults

`func NewMetricChartProjectionWithDefaults() *MetricChartProjection`

NewMetricChartProjectionWithDefaults instantiates a new MetricChartProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricChartProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricChartProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricChartProjection) SetType(v string)`

SetType sets Type field to given value.


### GetShowChart

`func (o *MetricChartProjection) GetShowChart() bool`

GetShowChart returns the ShowChart field if non-nil, zero value otherwise.

### GetShowChartOk

`func (o *MetricChartProjection) GetShowChartOk() (*bool, bool)`

GetShowChartOk returns a tuple with the ShowChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowChart

`func (o *MetricChartProjection) SetShowChart(v bool)`

SetShowChart sets ShowChart field to given value.

### HasShowChart

`func (o *MetricChartProjection) HasShowChart() bool`

HasShowChart returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *MetricChartProjection) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *MetricChartProjection) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *MetricChartProjection) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *MetricChartProjection) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetUnit

`func (o *MetricChartProjection) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricChartProjection) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricChartProjection) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricChartProjection) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetQuery

`func (o *MetricChartProjection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricChartProjection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricChartProjection) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


