# TextField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Value** | **string** |  | 
**AsTag** | **bool** |  | [default to false]

## Methods

### NewTextField

`func NewTextField(fieldId string, title string, type_ string, value string, asTag bool, ) *TextField`

NewTextField instantiates a new TextField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextFieldWithDefaults

`func NewTextFieldWithDefaults() *TextField`

NewTextFieldWithDefaults instantiates a new TextField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *TextField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *TextField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *TextField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *TextField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TextField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TextField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *TextField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TextField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TextField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TextField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TextField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextField) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *TextField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TextField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TextField) SetValue(v string)`

SetValue sets Value field to given value.


### GetAsTag

`func (o *TextField) GetAsTag() bool`

GetAsTag returns the AsTag field if non-nil, zero value otherwise.

### GetAsTagOk

`func (o *TextField) GetAsTagOk() (*bool, bool)`

GetAsTagOk returns a tuple with the AsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTag

`func (o *TextField) SetAsTag(v bool)`

SetAsTag sets AsTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


