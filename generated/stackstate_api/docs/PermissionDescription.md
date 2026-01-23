# PermissionDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectHandle** | **string** |  | 
**FromSources** | Pointer to [**[]SubjectSource**](SubjectSource.md) |  | [optional] 
**Permissions** | [**map[string]map[string][]string**](map.md) |  | 

## Methods

### NewPermissionDescription

`func NewPermissionDescription(subjectHandle string, permissions map[string]map[string][]string, ) *PermissionDescription`

NewPermissionDescription instantiates a new PermissionDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDescriptionWithDefaults

`func NewPermissionDescriptionWithDefaults() *PermissionDescription`

NewPermissionDescriptionWithDefaults instantiates a new PermissionDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectHandle

`func (o *PermissionDescription) GetSubjectHandle() string`

GetSubjectHandle returns the SubjectHandle field if non-nil, zero value otherwise.

### GetSubjectHandleOk

`func (o *PermissionDescription) GetSubjectHandleOk() (*string, bool)`

GetSubjectHandleOk returns a tuple with the SubjectHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectHandle

`func (o *PermissionDescription) SetSubjectHandle(v string)`

SetSubjectHandle sets SubjectHandle field to given value.


### GetFromSources

`func (o *PermissionDescription) GetFromSources() []SubjectSource`

GetFromSources returns the FromSources field if non-nil, zero value otherwise.

### GetFromSourcesOk

`func (o *PermissionDescription) GetFromSourcesOk() (*[]SubjectSource, bool)`

GetFromSourcesOk returns a tuple with the FromSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSources

`func (o *PermissionDescription) SetFromSources(v []SubjectSource)`

SetFromSources sets FromSources field to given value.

### HasFromSources

`func (o *PermissionDescription) HasFromSources() bool`

HasFromSources returns a boolean if a field has been set.

### GetPermissions

`func (o *PermissionDescription) GetPermissions() map[string]map[string][]string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionDescription) GetPermissionsOk() (*map[string]map[string][]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionDescription) SetPermissions(v map[string]map[string][]string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


