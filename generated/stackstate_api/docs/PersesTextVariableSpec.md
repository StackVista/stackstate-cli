# PersesTextVariableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Display** | Pointer to [**PersesVariableDisplaySpec**](PersesVariableDisplaySpec.md) |  | [optional] 
**Value** | **string** |  | 
**Constant** | Pointer to **bool** |  | [optional] 

## Methods

### NewPersesTextVariableSpec

`func NewPersesTextVariableSpec(name string, value string, ) *PersesTextVariableSpec`

NewPersesTextVariableSpec instantiates a new PersesTextVariableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesTextVariableSpecWithDefaults

`func NewPersesTextVariableSpecWithDefaults() *PersesTextVariableSpec`

NewPersesTextVariableSpecWithDefaults instantiates a new PersesTextVariableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PersesTextVariableSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersesTextVariableSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersesTextVariableSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *PersesTextVariableSpec) GetDisplay() PersesVariableDisplaySpec`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesTextVariableSpec) GetDisplayOk() (*PersesVariableDisplaySpec, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesTextVariableSpec) SetDisplay(v PersesVariableDisplaySpec)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesTextVariableSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetValue

`func (o *PersesTextVariableSpec) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PersesTextVariableSpec) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PersesTextVariableSpec) SetValue(v string)`

SetValue sets Value field to given value.


### GetConstant

`func (o *PersesTextVariableSpec) GetConstant() bool`

GetConstant returns the Constant field if non-nil, zero value otherwise.

### GetConstantOk

`func (o *PersesTextVariableSpec) GetConstantOk() (*bool, bool)`

GetConstantOk returns a tuple with the Constant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstant

`func (o *PersesTextVariableSpec) SetConstant(v bool)`

SetConstant sets Constant field to given value.

### HasConstant

`func (o *PersesTextVariableSpec) HasConstant() bool`

HasConstant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


