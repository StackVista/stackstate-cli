# PersesJSONRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** | The JSON reference pointing to the actual object. The property name &#39;$ref&#39; contains a reserved char, which gets stripped away when generated (it&#39;s possible to handle this but makes the generation code more complex). For now, consumers of this spec will need to omit the &#39;$&#39; in the property name. | 

## Methods

### NewPersesJSONRef

`func NewPersesJSONRef(ref string, ) *PersesJSONRef`

NewPersesJSONRef instantiates a new PersesJSONRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesJSONRefWithDefaults

`func NewPersesJSONRefWithDefaults() *PersesJSONRef`

NewPersesJSONRefWithDefaults instantiates a new PersesJSONRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *PersesJSONRef) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PersesJSONRef) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PersesJSONRef) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


