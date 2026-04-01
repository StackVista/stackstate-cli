# RatioField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Numerator** | **float32** |  | 
**Denominator** | **float32** |  | 
**Status** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 

## Methods

### NewRatioField

`func NewRatioField(fieldId string, title string, type_ string, numerator float32, denominator float32, ) *RatioField`

NewRatioField instantiates a new RatioField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatioFieldWithDefaults

`func NewRatioFieldWithDefaults() *RatioField`

NewRatioFieldWithDefaults instantiates a new RatioField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *RatioField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *RatioField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *RatioField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *RatioField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RatioField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RatioField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RatioField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RatioField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RatioField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RatioField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RatioField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatioField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatioField) SetType(v string)`

SetType sets Type field to given value.


### GetNumerator

`func (o *RatioField) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *RatioField) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *RatioField) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.


### GetDenominator

`func (o *RatioField) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *RatioField) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *RatioField) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.


### GetStatus

`func (o *RatioField) GetStatus() HealthStateValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RatioField) GetStatusOk() (*HealthStateValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RatioField) SetStatus(v HealthStateValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RatioField) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


