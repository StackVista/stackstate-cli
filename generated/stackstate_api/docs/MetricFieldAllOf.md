# MetricFieldAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | **string** | The frontend can use the query to refresh the metrics (at interval). Allows to keep the current behaviour of making individual calls for each row. | 
**Unit** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetricFieldAllOf

`func NewMetricFieldAllOf(type_ string, query string, ) *MetricFieldAllOf`

NewMetricFieldAllOf instantiates a new MetricFieldAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricFieldAllOfWithDefaults

`func NewMetricFieldAllOfWithDefaults() *MetricFieldAllOf`

NewMetricFieldAllOfWithDefaults instantiates a new MetricFieldAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricFieldAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricFieldAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricFieldAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *MetricFieldAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricFieldAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricFieldAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUnit

`func (o *MetricFieldAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricFieldAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricFieldAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricFieldAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *MetricFieldAllOf) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *MetricFieldAllOf) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *MetricFieldAllOf) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *MetricFieldAllOf) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


