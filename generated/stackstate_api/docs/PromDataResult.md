# PromDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **map[string]string** |  | 
**Values** | [**[][]PromSampleInner**]([]PromSampleInner.md) |  | 

## Methods

### NewPromDataResult

`func NewPromDataResult(metric map[string]string, values [][]PromSampleInner, ) *PromDataResult`

NewPromDataResult instantiates a new PromDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromDataResultWithDefaults

`func NewPromDataResultWithDefaults() *PromDataResult`

NewPromDataResultWithDefaults instantiates a new PromDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *PromDataResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PromDataResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PromDataResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.


### GetValues

`func (o *PromDataResult) GetValues() [][]PromSampleInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PromDataResult) GetValuesOk() (*[][]PromSampleInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PromDataResult) SetValues(v [][]PromSampleInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


