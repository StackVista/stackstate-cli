# ComponentTypeField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Label** | [**ComponentTypeLabel**](ComponentTypeLabel.md) |  | 
**ValueExtractor** | [**ComponentTypeValueExtractor**](ComponentTypeValueExtractor.md) |  | 
**Display** | [**ComponentTypeFieldDisplay**](ComponentTypeFieldDisplay.md) |  | 

## Methods

### NewComponentTypeField

`func NewComponentTypeField(fieldId string, label ComponentTypeLabel, valueExtractor ComponentTypeValueExtractor, display ComponentTypeFieldDisplay, ) *ComponentTypeField`

NewComponentTypeField instantiates a new ComponentTypeField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTypeFieldWithDefaults

`func NewComponentTypeFieldWithDefaults() *ComponentTypeField`

NewComponentTypeFieldWithDefaults instantiates a new ComponentTypeField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *ComponentTypeField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ComponentTypeField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ComponentTypeField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetLabel

`func (o *ComponentTypeField) GetLabel() ComponentTypeLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComponentTypeField) GetLabelOk() (*ComponentTypeLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComponentTypeField) SetLabel(v ComponentTypeLabel)`

SetLabel sets Label field to given value.


### GetValueExtractor

`func (o *ComponentTypeField) GetValueExtractor() ComponentTypeValueExtractor`

GetValueExtractor returns the ValueExtractor field if non-nil, zero value otherwise.

### GetValueExtractorOk

`func (o *ComponentTypeField) GetValueExtractorOk() (*ComponentTypeValueExtractor, bool)`

GetValueExtractorOk returns a tuple with the ValueExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueExtractor

`func (o *ComponentTypeField) SetValueExtractor(v ComponentTypeValueExtractor)`

SetValueExtractor sets ValueExtractor field to given value.


### GetDisplay

`func (o *ComponentTypeField) GetDisplay() ComponentTypeFieldDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ComponentTypeField) GetDisplayOk() (*ComponentTypeFieldDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ComponentTypeField) SetDisplay(v ComponentTypeFieldDisplay)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


