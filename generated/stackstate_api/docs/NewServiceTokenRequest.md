# NewServiceTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**Roles** | **[]string** |  | 
**DedicatedSubject** | Pointer to **string** |  | [optional] 

## Methods

### NewNewServiceTokenRequest

`func NewNewServiceTokenRequest(name string, roles []string, ) *NewServiceTokenRequest`

NewNewServiceTokenRequest instantiates a new NewServiceTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewServiceTokenRequestWithDefaults

`func NewNewServiceTokenRequestWithDefaults() *NewServiceTokenRequest`

NewNewServiceTokenRequestWithDefaults instantiates a new NewServiceTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewServiceTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewServiceTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewServiceTokenRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDate

`func (o *NewServiceTokenRequest) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewServiceTokenRequest) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewServiceTokenRequest) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewServiceTokenRequest) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetRoles

`func (o *NewServiceTokenRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NewServiceTokenRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NewServiceTokenRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetDedicatedSubject

`func (o *NewServiceTokenRequest) GetDedicatedSubject() string`

GetDedicatedSubject returns the DedicatedSubject field if non-nil, zero value otherwise.

### GetDedicatedSubjectOk

`func (o *NewServiceTokenRequest) GetDedicatedSubjectOk() (*string, bool)`

GetDedicatedSubjectOk returns a tuple with the DedicatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedSubject

`func (o *NewServiceTokenRequest) SetDedicatedSubject(v string)`

SetDedicatedSubject sets DedicatedSubject field to given value.

### HasDedicatedSubject

`func (o *NewServiceTokenRequest) HasDedicatedSubject() bool`

HasDedicatedSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


