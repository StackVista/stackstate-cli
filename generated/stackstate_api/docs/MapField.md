# MapField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Values** | **map[string]string** |  | 
**AsTag** | **bool** |  | [default to false]

## Methods

### NewMapField

`func NewMapField(fieldId string, title string, type_ string, values map[string]string, asTag bool, ) *MapField`

NewMapField instantiates a new MapField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapFieldWithDefaults

`func NewMapFieldWithDefaults() *MapField`

NewMapFieldWithDefaults instantiates a new MapField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *MapField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *MapField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *MapField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *MapField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MapField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MapField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *MapField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MapField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MapField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MapField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *MapField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapField) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *MapField) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MapField) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MapField) SetValues(v map[string]string)`

SetValues sets Values field to given value.


### GetAsTag

`func (o *MapField) GetAsTag() bool`

GetAsTag returns the AsTag field if non-nil, zero value otherwise.

### GetAsTagOk

`func (o *MapField) GetAsTagOk() (*bool, bool)`

GetAsTagOk returns a tuple with the AsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTag

`func (o *MapField) SetAsTag(v bool)`

SetAsTag sets AsTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


