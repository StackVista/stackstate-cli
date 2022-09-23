# PromData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **string** |  | 
**Result** | [**[]PromSampleInner**](PromSampleInner.md) | This is always a tuple represented as an array with in first position the unix timestamp as  a float with precision 3 in seconds) and in second position the sample value as a string.  | 

## Methods

### NewPromData

`func NewPromData(resultType string, result []PromSampleInner, ) *PromData`

NewPromData instantiates a new PromData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromDataWithDefaults

`func NewPromDataWithDefaults() *PromData`

NewPromDataWithDefaults instantiates a new PromData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *PromData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *PromData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *PromData) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetResult

`func (o *PromData) GetResult() []PromSampleInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PromData) GetResultOk() (*[]PromSampleInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PromData) SetResult(v []PromSampleInner)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


