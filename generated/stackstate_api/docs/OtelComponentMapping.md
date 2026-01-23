# OtelComponentMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Input** | [**OtelInput**](OtelInput.md) |  | 
**Output** | [**OtelComponentMappingOutput**](OtelComponentMappingOutput.md) |  | 
**Vars** | Pointer to [**[]OtelVariableMapping**](OtelVariableMapping.md) |  | [optional] 
**ExpireAfter** | **int64** |  | 

## Methods

### NewOtelComponentMapping

`func NewOtelComponentMapping(identifier string, name string, input OtelInput, output OtelComponentMappingOutput, expireAfter int64, ) *OtelComponentMapping`

NewOtelComponentMapping instantiates a new OtelComponentMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelComponentMappingWithDefaults

`func NewOtelComponentMappingWithDefaults() *OtelComponentMapping`

NewOtelComponentMappingWithDefaults instantiates a new OtelComponentMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *OtelComponentMapping) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OtelComponentMapping) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OtelComponentMapping) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *OtelComponentMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelComponentMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelComponentMapping) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OtelComponentMapping) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OtelComponentMapping) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OtelComponentMapping) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OtelComponentMapping) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInput

`func (o *OtelComponentMapping) GetInput() OtelInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *OtelComponentMapping) GetInputOk() (*OtelInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *OtelComponentMapping) SetInput(v OtelInput)`

SetInput sets Input field to given value.


### GetOutput

`func (o *OtelComponentMapping) GetOutput() OtelComponentMappingOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OtelComponentMapping) GetOutputOk() (*OtelComponentMappingOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OtelComponentMapping) SetOutput(v OtelComponentMappingOutput)`

SetOutput sets Output field to given value.


### GetVars

`func (o *OtelComponentMapping) GetVars() []OtelVariableMapping`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *OtelComponentMapping) GetVarsOk() (*[]OtelVariableMapping, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *OtelComponentMapping) SetVars(v []OtelVariableMapping)`

SetVars sets Vars field to given value.

### HasVars

`func (o *OtelComponentMapping) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetExpireAfter

`func (o *OtelComponentMapping) GetExpireAfter() int64`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *OtelComponentMapping) GetExpireAfterOk() (*int64, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *OtelComponentMapping) SetExpireAfter(v int64)`

SetExpireAfter sets ExpireAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


