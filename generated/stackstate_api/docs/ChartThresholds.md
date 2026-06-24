# ChartThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**DefaultColor** | Pointer to **string** |  | [optional] 
**Steps** | [**[]ChartThresholdStep**](ChartThresholdStep.md) |  | 

## Methods

### NewChartThresholds

`func NewChartThresholds(steps []ChartThresholdStep, ) *ChartThresholds`

NewChartThresholds instantiates a new ChartThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartThresholdsWithDefaults

`func NewChartThresholdsWithDefaults() *ChartThresholds`

NewChartThresholdsWithDefaults instantiates a new ChartThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ChartThresholds) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ChartThresholds) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ChartThresholds) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ChartThresholds) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDefaultColor

`func (o *ChartThresholds) GetDefaultColor() string`

GetDefaultColor returns the DefaultColor field if non-nil, zero value otherwise.

### GetDefaultColorOk

`func (o *ChartThresholds) GetDefaultColorOk() (*string, bool)`

GetDefaultColorOk returns a tuple with the DefaultColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColor

`func (o *ChartThresholds) SetDefaultColor(v string)`

SetDefaultColor sets DefaultColor field to given value.

### HasDefaultColor

`func (o *ChartThresholds) HasDefaultColor() bool`

HasDefaultColor returns a boolean if a field has been set.

### GetSteps

`func (o *ChartThresholds) GetSteps() []ChartThresholdStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ChartThresholds) GetStepsOk() (*[]ChartThresholdStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ChartThresholds) SetSteps(v []ChartThresholdStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


