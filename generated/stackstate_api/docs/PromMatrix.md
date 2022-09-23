# PromMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | **string** |  | 
**Result** | [**[]PromDataResult**](PromDataResult.md) |  | 

## Methods

### NewPromMatrix

`func NewPromMatrix(resultType string, result []PromDataResult, ) *PromMatrix`

NewPromMatrix instantiates a new PromMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromMatrixWithDefaults

`func NewPromMatrixWithDefaults() *PromMatrix`

NewPromMatrixWithDefaults instantiates a new PromMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *PromMatrix) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *PromMatrix) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *PromMatrix) SetResultType(v string)`

SetResultType sets ResultType field to given value.


### GetResult

`func (o *PromMatrix) GetResult() []PromDataResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PromMatrix) GetResultOk() (*[]PromDataResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PromMatrix) SetResult(v []PromDataResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


