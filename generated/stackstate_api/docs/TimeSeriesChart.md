# TimeSeriesChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**ConnectNulls** | Pointer to **bool** |  | [optional] 
**Thresholds** | Pointer to [**ChartThresholds**](ChartThresholds.md) |  | [optional] 

## Methods

### NewTimeSeriesChart

`func NewTimeSeriesChart(type_ string, ) *TimeSeriesChart`

NewTimeSeriesChart instantiates a new TimeSeriesChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesChartWithDefaults

`func NewTimeSeriesChartWithDefaults() *TimeSeriesChart`

NewTimeSeriesChartWithDefaults instantiates a new TimeSeriesChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeSeriesChart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSeriesChart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSeriesChart) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *TimeSeriesChart) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeSeriesChart) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeSeriesChart) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeSeriesChart) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimals

`func (o *TimeSeriesChart) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TimeSeriesChart) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TimeSeriesChart) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TimeSeriesChart) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetMin

`func (o *TimeSeriesChart) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *TimeSeriesChart) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *TimeSeriesChart) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *TimeSeriesChart) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *TimeSeriesChart) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *TimeSeriesChart) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *TimeSeriesChart) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *TimeSeriesChart) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetConnectNulls

`func (o *TimeSeriesChart) GetConnectNulls() bool`

GetConnectNulls returns the ConnectNulls field if non-nil, zero value otherwise.

### GetConnectNullsOk

`func (o *TimeSeriesChart) GetConnectNullsOk() (*bool, bool)`

GetConnectNullsOk returns a tuple with the ConnectNulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNulls

`func (o *TimeSeriesChart) SetConnectNulls(v bool)`

SetConnectNulls sets ConnectNulls field to given value.

### HasConnectNulls

`func (o *TimeSeriesChart) HasConnectNulls() bool`

HasConnectNulls returns a boolean if a field has been set.

### GetThresholds

`func (o *TimeSeriesChart) GetThresholds() ChartThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *TimeSeriesChart) GetThresholdsOk() (*ChartThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *TimeSeriesChart) SetThresholds(v ChartThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *TimeSeriesChart) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


