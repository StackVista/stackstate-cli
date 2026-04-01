# TextProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | Cel expression that returns a string | 
**AsTag** | Pointer to **bool** | Should the value be rendered as a tag or as plain text | [optional] [default to false]

## Methods

### NewTextProjection

`func NewTextProjection(type_ string, value string, ) *TextProjection`

NewTextProjection instantiates a new TextProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextProjectionWithDefaults

`func NewTextProjectionWithDefaults() *TextProjection`

NewTextProjectionWithDefaults instantiates a new TextProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextProjection) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *TextProjection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TextProjection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TextProjection) SetValue(v string)`

SetValue sets Value field to given value.


### GetAsTag

`func (o *TextProjection) GetAsTag() bool`

GetAsTag returns the AsTag field if non-nil, zero value otherwise.

### GetAsTagOk

`func (o *TextProjection) GetAsTagOk() (*bool, bool)`

GetAsTagOk returns a tuple with the AsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTag

`func (o *TextProjection) SetAsTag(v bool)`

SetAsTag sets AsTag field to given value.

### HasAsTag

`func (o *TextProjection) HasAsTag() bool`

HasAsTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


