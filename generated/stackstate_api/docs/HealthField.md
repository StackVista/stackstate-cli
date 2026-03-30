# HealthField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 

## Methods

### NewHealthField

`func NewHealthField(fieldId string, title string, type_ string, state HealthStateValue, ) *HealthField`

NewHealthField instantiates a new HealthField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthFieldWithDefaults

`func NewHealthFieldWithDefaults() *HealthField`

NewHealthFieldWithDefaults instantiates a new HealthField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *HealthField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *HealthField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *HealthField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *HealthField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HealthField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HealthField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *HealthField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *HealthField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthField) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *HealthField) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HealthField) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HealthField) SetState(v HealthStateValue)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


