# UpdateMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FunctionId** | Pointer to **int64** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**RemediationHint** | Pointer to **string** |  | [optional] 
**TopologyMapping** | Pointer to **string** |  | [optional] 
**IntervalSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateMonitor

`func NewUpdateMonitor() *UpdateMonitor`

NewUpdateMonitor instantiates a new UpdateMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorWithDefaults

`func NewUpdateMonitorWithDefaults() *UpdateMonitor`

NewUpdateMonitorWithDefaults instantiates a new UpdateMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *UpdateMonitor) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UpdateMonitor) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UpdateMonitor) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UpdateMonitor) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionId

`func (o *UpdateMonitor) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *UpdateMonitor) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *UpdateMonitor) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *UpdateMonitor) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateMonitor) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateMonitor) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateMonitor) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateMonitor) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRemediationHint

`func (o *UpdateMonitor) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *UpdateMonitor) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *UpdateMonitor) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *UpdateMonitor) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetTopologyMapping

`func (o *UpdateMonitor) GetTopologyMapping() string`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *UpdateMonitor) GetTopologyMappingOk() (*string, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *UpdateMonitor) SetTopologyMapping(v string)`

SetTopologyMapping sets TopologyMapping field to given value.

### HasTopologyMapping

`func (o *UpdateMonitor) HasTopologyMapping() bool`

HasTopologyMapping returns a boolean if a field has been set.

### GetIntervalSeconds

`func (o *UpdateMonitor) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *UpdateMonitor) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *UpdateMonitor) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.

### HasIntervalSeconds

`func (o *UpdateMonitor) HasIntervalSeconds() bool`

HasIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


