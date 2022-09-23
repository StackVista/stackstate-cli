# PromScalar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **string** |  | 
**Result** | [**[]PromSampleInner**](PromSampleInner.md) | This is always a tuple represented as an array with in first position the unix timestamp as  a float with precision 3 in seconds) and in second position the sample value as a string.  | 

## Methods

### NewPromScalar

`func NewPromScalar(resultType string, result []PromSampleInner, ) *PromScalar`

NewPromScalar instantiates a new PromScalar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromScalarWithDefaults

`func NewPromScalarWithDefaults() *PromScalar`

NewPromScalarWithDefaults instantiates a new PromScalar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *PromScalar) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *PromScalar) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *PromScalar) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetResult

`func (o *PromScalar) GetResult() []PromSampleInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PromScalar) GetResultOk() (*[]PromSampleInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PromScalar) SetResult(v []PromSampleInner)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


