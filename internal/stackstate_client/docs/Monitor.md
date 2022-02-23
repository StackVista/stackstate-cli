# Monitor

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

### NewMonitor

`func NewMonitor(id int64, name string, functionId int64, parameters map[string]interface{}, topologyMapping string, intervalSeconds int32, ) *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitor) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *Monitor) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Monitor) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Monitor) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Monitor) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *Monitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Monitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Monitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Monitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionId

`func (o *Monitor) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Monitor) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Monitor) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetParameters

`func (o *Monitor) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Monitor) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Monitor) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetRemediationHint

`func (o *Monitor) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *Monitor) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *Monitor) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *Monitor) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetTopologyMapping

`func (o *Monitor) GetTopologyMapping() string`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *Monitor) GetTopologyMappingOk() (*string, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *Monitor) SetTopologyMapping(v string)`

SetTopologyMapping sets TopologyMapping field to given value.


### GetIntervalSeconds

`func (o *Monitor) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *Monitor) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *Monitor) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


