# ErrorField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewErrorField

`func NewErrorField(fieldId string, title string, type_ string, message string, ) *ErrorField`

NewErrorField instantiates a new ErrorField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorFieldWithDefaults

`func NewErrorFieldWithDefaults() *ErrorField`

NewErrorFieldWithDefaults instantiates a new ErrorField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *ErrorField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ErrorField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ErrorField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *ErrorField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ErrorField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ErrorField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorField) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ErrorField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorField) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


