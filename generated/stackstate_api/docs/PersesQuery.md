# PersesQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**PersesQuerySpec**](PersesQuerySpec.md) |  | [optional] 

## Methods

### NewPersesQuery

`func NewPersesQuery() *PersesQuery`

NewPersesQuery instantiates a new PersesQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesQueryWithDefaults

`func NewPersesQueryWithDefaults() *PersesQuery`

NewPersesQueryWithDefaults instantiates a new PersesQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PersesQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PersesQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PersesQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PersesQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSpec

`func (o *PersesQuery) GetSpec() PersesQuerySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PersesQuery) GetSpecOk() (*PersesQuerySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PersesQuery) SetSpec(v PersesQuerySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PersesQuery) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


