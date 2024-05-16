# ComponentNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComponentId** | **int64** |  | 

## Methods

### NewComponentNotFoundError

`func NewComponentNotFoundError(type_ string, componentId int64, ) *ComponentNotFoundError`

NewComponentNotFoundError instantiates a new ComponentNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentNotFoundErrorWithDefaults

`func NewComponentNotFoundErrorWithDefaults() *ComponentNotFoundError`

NewComponentNotFoundErrorWithDefaults instantiates a new ComponentNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentNotFoundError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentNotFoundError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentNotFoundError) SetType(v string)`

SetType sets Type field to given value.


### GetComponentId

`func (o *ComponentNotFoundError) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentNotFoundError) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentNotFoundError) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


