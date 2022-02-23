# CreateMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Identifier** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FunctionId** | **int64** |  | 
**Parameters** | **map[string]interface{}** |  | 
**RemediationHint** | Pointer to **string** |  | [optional] 
**TopologyMapping** | **string** | Placeholder for the topology element identifier | 
**IntervalSeconds** | **int32** |  | 

## Methods

### NewCreateMonitor

`func NewCreateMonitor(name string, identifier string, functionId int64, parameters map[string]interface{}, topologyMapping string, intervalSeconds int32, ) *CreateMonitor`

NewCreateMonitor instantiates a new CreateMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitorWithDefaults

`func NewCreateMonitorWithDefaults() *CreateMonitor`

NewCreateMonitorWithDefaults instantiates a new CreateMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMonitor) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *CreateMonitor) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateMonitor) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateMonitor) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDescription

`func (o *CreateMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionId

`func (o *CreateMonitor) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CreateMonitor) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CreateMonitor) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetParameters

`func (o *CreateMonitor) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateMonitor) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateMonitor) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetRemediationHint

`func (o *CreateMonitor) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *CreateMonitor) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *CreateMonitor) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *CreateMonitor) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetTopologyMapping

`func (o *CreateMonitor) GetTopologyMapping() string`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *CreateMonitor) GetTopologyMappingOk() (*string, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *CreateMonitor) SetTopologyMapping(v string)`

SetTopologyMapping sets TopologyMapping field to given value.


### GetIntervalSeconds

`func (o *CreateMonitor) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *CreateMonitor) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *CreateMonitor) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


