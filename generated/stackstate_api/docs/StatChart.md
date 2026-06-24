# StatChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Sparkline** | Pointer to **bool** |  | [optional] 
**Calculation** | Pointer to [**ChartCalculation**](ChartCalculation.md) |  | [optional] 
**Thresholds** | Pointer to [**ChartStatThresholds**](ChartStatThresholds.md) |  | [optional] 

## Methods

### NewStatChart

`func NewStatChart(type_ string, ) *StatChart`

NewStatChart instantiates a new StatChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatChartWithDefaults

`func NewStatChartWithDefaults() *StatChart`

NewStatChartWithDefaults instantiates a new StatChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StatChart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatChart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatChart) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *StatChart) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *StatChart) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *StatChart) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *StatChart) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimals

`func (o *StatChart) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *StatChart) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *StatChart) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *StatChart) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetSparkline

`func (o *StatChart) GetSparkline() bool`

GetSparkline returns the Sparkline field if non-nil, zero value otherwise.

### GetSparklineOk

`func (o *StatChart) GetSparklineOk() (*bool, bool)`

GetSparklineOk returns a tuple with the Sparkline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkline

`func (o *StatChart) SetSparkline(v bool)`

SetSparkline sets Sparkline field to given value.

### HasSparkline

`func (o *StatChart) HasSparkline() bool`

HasSparkline returns a boolean if a field has been set.

### GetCalculation

`func (o *StatChart) GetCalculation() ChartCalculation`

GetCalculation returns the Calculation field if non-nil, zero value otherwise.

### GetCalculationOk

`func (o *StatChart) GetCalculationOk() (*ChartCalculation, bool)`

GetCalculationOk returns a tuple with the Calculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculation

`func (o *StatChart) SetCalculation(v ChartCalculation)`

SetCalculation sets Calculation field to given value.

### HasCalculation

`func (o *StatChart) HasCalculation() bool`

HasCalculation returns a boolean if a field has been set.

### GetThresholds

`func (o *StatChart) GetThresholds() ChartStatThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *StatChart) GetThresholdsOk() (*ChartStatThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *StatChart) SetThresholds(v ChartStatThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *StatChart) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


