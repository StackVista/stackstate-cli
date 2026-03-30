# PresentationHighlightField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **float64** |  | 
**Projection** | Pointer to [**ComponentHighlightProjection**](ComponentHighlightProjection.md) |  | [optional] 

## Methods

### NewPresentationHighlightField

`func NewPresentationHighlightField(fieldId string, title string, order float64, ) *PresentationHighlightField`

NewPresentationHighlightField instantiates a new PresentationHighlightField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationHighlightFieldWithDefaults

`func NewPresentationHighlightFieldWithDefaults() *PresentationHighlightField`

NewPresentationHighlightFieldWithDefaults instantiates a new PresentationHighlightField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *PresentationHighlightField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *PresentationHighlightField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *PresentationHighlightField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetTitle

`func (o *PresentationHighlightField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationHighlightField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationHighlightField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *PresentationHighlightField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PresentationHighlightField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PresentationHighlightField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PresentationHighlightField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *PresentationHighlightField) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationHighlightField) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationHighlightField) SetOrder(v float64)`

SetOrder sets Order field to given value.


### GetProjection

`func (o *PresentationHighlightField) GetProjection() ComponentHighlightProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *PresentationHighlightField) GetProjectionOk() (*ComponentHighlightProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *PresentationHighlightField) SetProjection(v ComponentHighlightProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *PresentationHighlightField) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


