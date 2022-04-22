# Monitor1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FunctionId** | **int64** |  | 
**Parameters** | **map[string]interface{}** |  | 
**RemediationHint** | Pointer to **string** |  | [optional] 
**TopologyMapping** | **string** |  | 
**IntervalSeconds** | **int32** |  | 

## Methods

### NewMonitor1

`func NewMonitor1(id int64, name string, functionId int64, parameters map[string]interface{}, topologyMapping string, intervalSeconds int32, ) *Monitor1`

NewMonitor1 instantiates a new Monitor1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitor1WithDefaults

`func NewMonitor1WithDefaults() *Monitor1`

NewMonitor1WithDefaults instantiates a new Monitor1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monitor1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitor1) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Monitor1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor1) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *Monitor1) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Monitor1) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Monitor1) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Monitor1) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *Monitor1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Monitor1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Monitor1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Monitor1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionId

`func (o *Monitor1) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Monitor1) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Monitor1) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetParameters

`func (o *Monitor1) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Monitor1) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Monitor1) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetRemediationHint

`func (o *Monitor1) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *Monitor1) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *Monitor1) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *Monitor1) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetTopologyMapping

`func (o *Monitor1) GetTopologyMapping() string`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *Monitor1) GetTopologyMappingOk() (*string, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *Monitor1) SetTopologyMapping(v string)`

SetTopologyMapping sets TopologyMapping field to given value.


### GetIntervalSeconds

`func (o *Monitor1) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *Monitor1) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *Monitor1) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


