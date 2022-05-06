# Reference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StreamNodeId** | **int64** |  | 
**ElementIdentifiers** | **[]string** |  | 

## Methods

### NewReference

`func NewReference(type_ string, streamNodeId int64, elementIdentifiers []string, ) *Reference`

NewReference instantiates a new Reference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceWithDefaults

`func NewReferenceWithDefaults() *Reference`

NewReferenceWithDefaults instantiates a new Reference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Reference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reference) SetType(v string)`

SetType sets Type field to given value.


### GetStreamNodeId

`func (o *Reference) GetStreamNodeId() int64`

GetStreamNodeId returns the StreamNodeId field if non-nil, zero value otherwise.

### GetStreamNodeIdOk

`func (o *Reference) GetStreamNodeIdOk() (*int64, bool)`

GetStreamNodeIdOk returns a tuple with the StreamNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamNodeId

`func (o *Reference) SetStreamNodeId(v int64)`

SetStreamNodeId sets StreamNodeId field to given value.


### GetElementIdentifiers

`func (o *Reference) GetElementIdentifiers() []string`

GetElementIdentifiers returns the ElementIdentifiers field if non-nil, zero value otherwise.

### GetElementIdentifiersOk

`func (o *Reference) GetElementIdentifiersOk() (*[]string, bool)`

GetElementIdentifiersOk returns a tuple with the ElementIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIdentifiers

`func (o *Reference) SetElementIdentifiers(v []string)`

SetElementIdentifiers sets ElementIdentifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


