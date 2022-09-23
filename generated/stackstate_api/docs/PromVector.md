# PromVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **string** |  | 
**Result** | [**[]PromVectorResult**](PromVectorResult.md) |  | 

## Methods

### NewPromVector

`func NewPromVector(resultType string, result []PromVectorResult, ) *PromVector`

NewPromVector instantiates a new PromVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromVectorWithDefaults

`func NewPromVectorWithDefaults() *PromVector`

NewPromVectorWithDefaults instantiates a new PromVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *PromVector) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *PromVector) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *PromVector) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetResult

`func (o *PromVector) GetResult() []PromVectorResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PromVector) GetResultOk() (*[]PromVectorResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PromVector) SetResult(v []PromVectorResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


