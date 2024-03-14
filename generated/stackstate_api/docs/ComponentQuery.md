# ComponentQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceAttributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 
**Instant** | Pointer to **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | [optional] 

## Methods

### NewComponentQuery

`func NewComponentQuery(resourceAttributes map[string]string, ) *ComponentQuery`

NewComponentQuery instantiates a new ComponentQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentQueryWithDefaults

`func NewComponentQueryWithDefaults() *ComponentQuery`

NewComponentQueryWithDefaults instantiates a new ComponentQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceAttributes

`func (o *ComponentQuery) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *ComponentQuery) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *ComponentQuery) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetInstant

`func (o *ComponentQuery) GetInstant() int32`

GetInstant returns the Instant field if non-nil, zero value otherwise.

### GetInstantOk

`func (o *ComponentQuery) GetInstantOk() (*int32, bool)`

GetInstantOk returns a tuple with the Instant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstant

`func (o *ComponentQuery) SetInstant(v int32)`

SetInstant sets Instant field to given value.

### HasInstant

`func (o *ComponentQuery) HasInstant() bool`

HasInstant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


