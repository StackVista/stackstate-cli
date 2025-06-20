# ServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 
**Roles** | **[]string** |  | 
**DedicatedSubject** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceToken

`func NewServiceToken(name string, roles []string, ) *ServiceToken`

NewServiceToken instantiates a new ServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenWithDefaults

`func NewServiceTokenWithDefaults() *ServiceToken`

NewServiceTokenWithDefaults instantiates a new ServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ServiceToken) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ServiceToken) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ServiceToken) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ServiceToken) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ServiceToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceToken) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *ServiceToken) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ServiceToken) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ServiceToken) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ServiceToken) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRoles

`func (o *ServiceToken) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ServiceToken) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ServiceToken) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetDedicatedSubject

`func (o *ServiceToken) GetDedicatedSubject() string`

GetDedicatedSubject returns the DedicatedSubject field if non-nil, zero value otherwise.

### GetDedicatedSubjectOk

`func (o *ServiceToken) GetDedicatedSubjectOk() (*string, bool)`

GetDedicatedSubjectOk returns a tuple with the DedicatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedSubject

`func (o *ServiceToken) SetDedicatedSubject(v string)`

SetDedicatedSubject sets DedicatedSubject field to given value.

### HasDedicatedSubject

`func (o *ServiceToken) HasDedicatedSubject() bool`

HasDedicatedSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


