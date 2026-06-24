# ChartStatThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultColor** | Pointer to **string** |  | [optional] 
**Steps** | [**[]ChartThresholdStep**](ChartThresholdStep.md) |  | 

## Methods

### NewChartStatThresholds

`func NewChartStatThresholds(steps []ChartThresholdStep, ) *ChartStatThresholds`

NewChartStatThresholds instantiates a new ChartStatThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartStatThresholdsWithDefaults

`func NewChartStatThresholdsWithDefaults() *ChartStatThresholds`

NewChartStatThresholdsWithDefaults instantiates a new ChartStatThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultColor

`func (o *ChartStatThresholds) GetDefaultColor() string`

GetDefaultColor returns the DefaultColor field if non-nil, zero value otherwise.

### GetDefaultColorOk

`func (o *ChartStatThresholds) GetDefaultColorOk() (*string, bool)`

GetDefaultColorOk returns a tuple with the DefaultColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColor

`func (o *ChartStatThresholds) SetDefaultColor(v string)`

SetDefaultColor sets DefaultColor field to given value.

### HasDefaultColor

`func (o *ChartStatThresholds) HasDefaultColor() bool`

HasDefaultColor returns a boolean if a field has been set.

### GetSteps

`func (o *ChartStatThresholds) GetSteps() []ChartThresholdStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ChartStatThresholds) GetStepsOk() (*[]ChartThresholdStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ChartStatThresholds) SetSteps(v []ChartThresholdStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


