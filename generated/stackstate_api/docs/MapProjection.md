# MapProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | Cel expression that returns a map&lt;string,dyn&gt; | 
**AsTags** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewMapProjection

`func NewMapProjection(type_ string, value string, ) *MapProjection`

NewMapProjection instantiates a new MapProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapProjectionWithDefaults

`func NewMapProjectionWithDefaults() *MapProjection`

NewMapProjectionWithDefaults instantiates a new MapProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MapProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapProjection) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *MapProjection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MapProjection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MapProjection) SetValue(v string)`

SetValue sets Value field to given value.


### GetAsTags

`func (o *MapProjection) GetAsTags() bool`

GetAsTags returns the AsTags field if non-nil, zero value otherwise.

### GetAsTagsOk

`func (o *MapProjection) GetAsTagsOk() (*bool, bool)`

GetAsTagsOk returns a tuple with the AsTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTags

`func (o *MapProjection) SetAsTags(v bool)`

SetAsTags sets AsTags field to given value.

### HasAsTags

`func (o *MapProjection) HasAsTags() bool`

HasAsTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


