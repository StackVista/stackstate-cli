# Chart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**ConnectNulls** | Pointer to **bool** |  | [optional] 
**Thresholds** | Pointer to [**ChartStatThresholds**](ChartStatThresholds.md) |  | [optional] 
**Calculation** | Pointer to [**ChartCalculation**](ChartCalculation.md) |  | [optional] 
**Sparkline** | Pointer to **bool** |  | [optional] 

## Methods

### NewChart

`func NewChart(type_ string, ) *Chart`

NewChart instantiates a new Chart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartWithDefaults

`func NewChartWithDefaults() *Chart`

NewChartWithDefaults instantiates a new Chart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Chart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Chart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Chart) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *Chart) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Chart) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Chart) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Chart) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimals

`func (o *Chart) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Chart) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Chart) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *Chart) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetMin

`func (o *Chart) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Chart) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Chart) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *Chart) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *Chart) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Chart) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Chart) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *Chart) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetConnectNulls

`func (o *Chart) GetConnectNulls() bool`

GetConnectNulls returns the ConnectNulls field if non-nil, zero value otherwise.

### GetConnectNullsOk

`func (o *Chart) GetConnectNullsOk() (*bool, bool)`

GetConnectNullsOk returns a tuple with the ConnectNulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNulls

`func (o *Chart) SetConnectNulls(v bool)`

SetConnectNulls sets ConnectNulls field to given value.

### HasConnectNulls

`func (o *Chart) HasConnectNulls() bool`

HasConnectNulls returns a boolean if a field has been set.

### GetThresholds

`func (o *Chart) GetThresholds() ChartStatThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Chart) GetThresholdsOk() (*ChartStatThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Chart) SetThresholds(v ChartStatThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Chart) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetCalculation

`func (o *Chart) GetCalculation() ChartCalculation`

GetCalculation returns the Calculation field if non-nil, zero value otherwise.

### GetCalculationOk

`func (o *Chart) GetCalculationOk() (*ChartCalculation, bool)`

GetCalculationOk returns a tuple with the Calculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculation

`func (o *Chart) SetCalculation(v ChartCalculation)`

SetCalculation sets Calculation field to given value.

### HasCalculation

`func (o *Chart) HasCalculation() bool`

HasCalculation returns a boolean if a field has been set.

### GetSparkline

`func (o *Chart) GetSparkline() bool`

GetSparkline returns the Sparkline field if non-nil, zero value otherwise.

### GetSparklineOk

`func (o *Chart) GetSparklineOk() (*bool, bool)`

GetSparklineOk returns a tuple with the Sparkline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkline

`func (o *Chart) SetSparkline(v bool)`

SetSparkline sets Sparkline field to given value.

### HasSparkline

`func (o *Chart) HasSparkline() bool`

HasSparkline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


