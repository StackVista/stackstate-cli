# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **string** |  | 
**Result** | [**ValueTuple**](ValueTuple.md) |  | 

## Methods

### NewData

`func NewData(resultType string, result ValueTuple, ) *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *Data) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *Data) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *Data) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetResult

`func (o *Data) GetResult() ValueTuple`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Data) GetResultOk() (*ValueTuple, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Data) SetResult(v ValueTuple)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


