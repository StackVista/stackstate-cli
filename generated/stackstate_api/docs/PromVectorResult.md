# PromVectorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **map[string]string** |  | 
**Value** | [**[]PromSampleInner**](PromSampleInner.md) | This is always a tuple represented as an array with in first position the unix timestamp as  a float with precision 3 in seconds) and in second position the sample value as a string.  | 

## Methods

### NewPromVectorResult

`func NewPromVectorResult(metric map[string]string, value []PromSampleInner, ) *PromVectorResult`

NewPromVectorResult instantiates a new PromVectorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromVectorResultWithDefaults

`func NewPromVectorResultWithDefaults() *PromVectorResult`

NewPromVectorResultWithDefaults instantiates a new PromVectorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *PromVectorResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PromVectorResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PromVectorResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.


### GetValue

`func (o *PromVectorResult) GetValue() []PromSampleInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PromVectorResult) GetValueOk() (*[]PromSampleInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PromVectorResult) SetValue(v []PromSampleInner)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


