# DeleteVersionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | **[]string** |  | 
**SkippedInUse** | **[]string** |  | 

## Methods

### NewDeleteVersionsResult

`func NewDeleteVersionsResult(deleted []string, skippedInUse []string, ) *DeleteVersionsResult`

NewDeleteVersionsResult instantiates a new DeleteVersionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVersionsResultWithDefaults

`func NewDeleteVersionsResultWithDefaults() *DeleteVersionsResult`

NewDeleteVersionsResultWithDefaults instantiates a new DeleteVersionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteVersionsResult) GetDeleted() []string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteVersionsResult) GetDeletedOk() (*[]string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteVersionsResult) SetDeleted(v []string)`

SetDeleted sets Deleted field to given value.


### GetSkippedInUse

`func (o *DeleteVersionsResult) GetSkippedInUse() []string`

GetSkippedInUse returns the SkippedInUse field if non-nil, zero value otherwise.

### GetSkippedInUseOk

`func (o *DeleteVersionsResult) GetSkippedInUseOk() (*[]string, bool)`

GetSkippedInUseOk returns a tuple with the SkippedInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedInUse

`func (o *DeleteVersionsResult) SetSkippedInUse(v []string)`

SetSkippedInUse sets SkippedInUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


