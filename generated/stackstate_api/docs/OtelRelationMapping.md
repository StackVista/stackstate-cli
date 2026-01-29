# OtelRelationMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Identifier** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Input** | [**OtelInput**](OtelInput.md) |  | 
**Output** | [**OtelRelationMappingOutput**](OtelRelationMappingOutput.md) |  | 
**Vars** | Pointer to [**[]OtelVariableMapping**](OtelVariableMapping.md) |  | [optional] 
**ExpireAfter** | **int64** |  | 

## Methods

### NewOtelRelationMapping

`func NewOtelRelationMapping(type_ string, identifier string, name string, input OtelInput, output OtelRelationMappingOutput, expireAfter int64, ) *OtelRelationMapping`

NewOtelRelationMapping instantiates a new OtelRelationMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelRelationMappingWithDefaults

`func NewOtelRelationMappingWithDefaults() *OtelRelationMapping`

NewOtelRelationMappingWithDefaults instantiates a new OtelRelationMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OtelRelationMapping) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OtelRelationMapping) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OtelRelationMapping) SetType(v string)`

SetType sets Type field to given value.


### GetIdentifier

`func (o *OtelRelationMapping) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OtelRelationMapping) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OtelRelationMapping) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *OtelRelationMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtelRelationMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtelRelationMapping) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OtelRelationMapping) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OtelRelationMapping) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OtelRelationMapping) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OtelRelationMapping) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInput

`func (o *OtelRelationMapping) GetInput() OtelInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *OtelRelationMapping) GetInputOk() (*OtelInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *OtelRelationMapping) SetInput(v OtelInput)`

SetInput sets Input field to given value.


### GetOutput

`func (o *OtelRelationMapping) GetOutput() OtelRelationMappingOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OtelRelationMapping) GetOutputOk() (*OtelRelationMappingOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OtelRelationMapping) SetOutput(v OtelRelationMappingOutput)`

SetOutput sets Output field to given value.


### GetVars

`func (o *OtelRelationMapping) GetVars() []OtelVariableMapping`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *OtelRelationMapping) GetVarsOk() (*[]OtelVariableMapping, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *OtelRelationMapping) SetVars(v []OtelVariableMapping)`

SetVars sets Vars field to given value.

### HasVars

`func (o *OtelRelationMapping) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetExpireAfter

`func (o *OtelRelationMapping) GetExpireAfter() int64`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *OtelRelationMapping) GetExpireAfterOk() (*int64, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *OtelRelationMapping) SetExpireAfter(v int64)`

SetExpireAfter sets ExpireAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


