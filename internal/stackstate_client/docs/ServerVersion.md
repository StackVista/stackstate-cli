# ServerVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Major** | **int32** |  | 
**Minor** | **int32** |  | 
**Patch** | **int32** |  | 
**Diff** | **string** |  | 
**Commit** | **string** |  | 
**IsDev** | **bool** |  | 

## Methods

### NewServerVersion

`func NewServerVersion(major int32, minor int32, patch int32, diff string, commit string, isDev bool, ) *ServerVersion`

NewServerVersion instantiates a new ServerVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVersionWithDefaults

`func NewServerVersionWithDefaults() *ServerVersion`

NewServerVersionWithDefaults instantiates a new ServerVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajor

`func (o *ServerVersion) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *ServerVersion) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *ServerVersion) SetMajor(v int32)`

SetMajor sets Major field to given value.


### GetMinor

`func (o *ServerVersion) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *ServerVersion) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *ServerVersion) SetMinor(v int32)`

SetMinor sets Minor field to given value.


### GetPatch

`func (o *ServerVersion) GetPatch() int32`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ServerVersion) GetPatchOk() (*int32, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ServerVersion) SetPatch(v int32)`

SetPatch sets Patch field to given value.


### GetDiff

`func (o *ServerVersion) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *ServerVersion) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *ServerVersion) SetDiff(v string)`

SetDiff sets Diff field to given value.


### GetCommit

`func (o *ServerVersion) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ServerVersion) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ServerVersion) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetIsDev

`func (o *ServerVersion) GetIsDev() bool`

GetIsDev returns the IsDev field if non-nil, zero value otherwise.

### GetIsDevOk

`func (o *ServerVersion) GetIsDevOk() (*bool, bool)`

GetIsDevOk returns a tuple with the IsDev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDev

`func (o *ServerVersion) SetIsDev(v bool)`

SetIsDev sets IsDev field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


