# DataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | [**[]ValueTuple**](ValueTuple.md) |  | 
**Values** | [**[]ValueTuple**](ValueTuple.md) |  | 

## Methods

### NewDataResult

`func NewDataResult(metric []ValueTuple, values []ValueTuple, ) *DataResult`

NewDataResult instantiates a new DataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataResultWithDefaults

`func NewDataResultWithDefaults() *DataResult`

NewDataResultWithDefaults instantiates a new DataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *DataResult) GetMetric() []ValueTuple`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *DataResult) GetMetricOk() (*[]ValueTuple, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *DataResult) SetMetric(v []ValueTuple)`

SetMetric sets Metric field to given value.


### GetValues

`func (o *DataResult) GetValues() []ValueTuple`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DataResult) GetValuesOk() (*[]ValueTuple, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DataResult) SetValues(v []ValueTuple)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


