# SubjectConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** |  | 
**Source** | [**SubjectSource**](SubjectSource.md) |  | 

## Methods

### NewSubjectConfig

`func NewSubjectConfig(handle string, source SubjectSource, ) *SubjectConfig`

NewSubjectConfig instantiates a new SubjectConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectConfigWithDefaults

`func NewSubjectConfigWithDefaults() *SubjectConfig`

NewSubjectConfigWithDefaults instantiates a new SubjectConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *SubjectConfig) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SubjectConfig) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SubjectConfig) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetSource

`func (o *SubjectConfig) GetSource() SubjectSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SubjectConfig) GetSourceOk() (*SubjectSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SubjectConfig) SetSource(v SubjectSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


