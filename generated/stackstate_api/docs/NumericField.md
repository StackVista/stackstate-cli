# NumericField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Value** | **float32** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 

## Methods

### NewNumericField

`func NewNumericField(fieldId string, title string, type_ string, value float32, ) *NumericField`

NewNumericField instantiates a new NumericField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericFieldWithDefaults

`func NewNumericFieldWithDefaults() *NumericField`

NewNumericFieldWithDefaults instantiates a new NumericField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *NumericField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *NumericField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *NumericField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *NumericField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NumericField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NumericField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NumericField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumericField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumericField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumericField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *NumericField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumericField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumericField) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *NumericField) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumericField) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumericField) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnit

`func (o *NumericField) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *NumericField) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *NumericField) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *NumericField) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *NumericField) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *NumericField) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *NumericField) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *NumericField) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


