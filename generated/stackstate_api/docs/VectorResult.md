# VectorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **map[string]string** |  | 
**Value** | [**ValueTuple**](ValueTuple.md) |  | 

## Methods

### NewVectorResult

`func NewVectorResult(metric map[string]string, value ValueTuple, ) *VectorResult`

NewVectorResult instantiates a new VectorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorResultWithDefaults

`func NewVectorResultWithDefaults() *VectorResult`

NewVectorResultWithDefaults instantiates a new VectorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *VectorResult) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *VectorResult) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *VectorResult) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.


### GetValue

`func (o *VectorResult) GetValue() ValueTuple`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VectorResult) GetValueOk() (*ValueTuple, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VectorResult) SetValue(v ValueTuple)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


