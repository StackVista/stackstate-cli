# UpsertOtelComponentMappingsRequest

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

### NewUpsertOtelComponentMappingsRequest

`func NewUpsertOtelComponentMappingsRequest(identifier string, name string, input OtelInput, output OtelComponentMappingOutput, expireAfter int64, ) *UpsertOtelComponentMappingsRequest`

NewUpsertOtelComponentMappingsRequest instantiates a new UpsertOtelComponentMappingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertOtelComponentMappingsRequestWithDefaults

`func NewUpsertOtelComponentMappingsRequestWithDefaults() *UpsertOtelComponentMappingsRequest`

NewUpsertOtelComponentMappingsRequestWithDefaults instantiates a new UpsertOtelComponentMappingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *UpsertOtelComponentMappingsRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UpsertOtelComponentMappingsRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UpsertOtelComponentMappingsRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *UpsertOtelComponentMappingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertOtelComponentMappingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertOtelComponentMappingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpsertOtelComponentMappingsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertOtelComponentMappingsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertOtelComponentMappingsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertOtelComponentMappingsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInput

`func (o *UpsertOtelComponentMappingsRequest) GetInput() OtelInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *UpsertOtelComponentMappingsRequest) GetInputOk() (*OtelInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *UpsertOtelComponentMappingsRequest) SetInput(v OtelInput)`

SetInput sets Input field to given value.


### GetOutput

`func (o *UpsertOtelComponentMappingsRequest) GetOutput() OtelComponentMappingOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *UpsertOtelComponentMappingsRequest) GetOutputOk() (*OtelComponentMappingOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *UpsertOtelComponentMappingsRequest) SetOutput(v OtelComponentMappingOutput)`

SetOutput sets Output field to given value.


### GetVars

`func (o *UpsertOtelComponentMappingsRequest) GetVars() []OtelVariableMapping`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *UpsertOtelComponentMappingsRequest) GetVarsOk() (*[]OtelVariableMapping, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *UpsertOtelComponentMappingsRequest) SetVars(v []OtelVariableMapping)`

SetVars sets Vars field to given value.

### HasVars

`func (o *UpsertOtelComponentMappingsRequest) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetExpireAfter

`func (o *UpsertOtelComponentMappingsRequest) GetExpireAfter() int64`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *UpsertOtelComponentMappingsRequest) GetExpireAfterOk() (*int64, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *UpsertOtelComponentMappingsRequest) SetExpireAfter(v int64)`

SetExpireAfter sets ExpireAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


