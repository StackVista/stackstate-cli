# GaugeChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**Calculation** | Pointer to [**ChartCalculation**](ChartCalculation.md) |  | [optional] 
**Thresholds** | Pointer to [**ChartThresholds**](ChartThresholds.md) |  | [optional] 

## Methods

### NewGaugeChart

`func NewGaugeChart(type_ string, ) *GaugeChart`

NewGaugeChart instantiates a new GaugeChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaugeChartWithDefaults

`func NewGaugeChartWithDefaults() *GaugeChart`

NewGaugeChartWithDefaults instantiates a new GaugeChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GaugeChart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GaugeChart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GaugeChart) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *GaugeChart) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GaugeChart) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GaugeChart) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GaugeChart) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimals

`func (o *GaugeChart) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *GaugeChart) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *GaugeChart) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *GaugeChart) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetMax

`func (o *GaugeChart) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GaugeChart) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GaugeChart) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *GaugeChart) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetCalculation

`func (o *GaugeChart) GetCalculation() ChartCalculation`

GetCalculation returns the Calculation field if non-nil, zero value otherwise.

### GetCalculationOk

`func (o *GaugeChart) GetCalculationOk() (*ChartCalculation, bool)`

GetCalculationOk returns a tuple with the Calculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculation

`func (o *GaugeChart) SetCalculation(v ChartCalculation)`

SetCalculation sets Calculation field to given value.

### HasCalculation

`func (o *GaugeChart) HasCalculation() bool`

HasCalculation returns a boolean if a field has been set.

### GetThresholds

`func (o *GaugeChart) GetThresholds() ChartThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *GaugeChart) GetThresholdsOk() (*ChartThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *GaugeChart) SetThresholds(v ChartThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *GaugeChart) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


