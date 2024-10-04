# AgentRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**Lease** | [**AgentLease**](AgentLease.md) |  | 
**LeaseUntilEpochMs** | **int64** |  | 
**RegisteredEpochMs** | **int64** |  | 
**AgentData** | Pointer to [**AgentData**](AgentData.md) |  | [optional] 
**NodeBudgetCount** | **int32** | The number of standard (4CPU, 16Gb) nodes this agent counts for | 

## Methods

### NewAgentRegistration

`func NewAgentRegistration(agentId string, lease AgentLease, leaseUntilEpochMs int64, registeredEpochMs int64, nodeBudgetCount int32, ) *AgentRegistration`

NewAgentRegistration instantiates a new AgentRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRegistrationWithDefaults

`func NewAgentRegistrationWithDefaults() *AgentRegistration`

NewAgentRegistrationWithDefaults instantiates a new AgentRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentRegistration) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentRegistration) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentRegistration) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetLease

`func (o *AgentRegistration) GetLease() AgentLease`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *AgentRegistration) GetLeaseOk() (*AgentLease, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *AgentRegistration) SetLease(v AgentLease)`

SetLease sets Lease field to given value.


### GetLeaseUntilEpochMs

`func (o *AgentRegistration) GetLeaseUntilEpochMs() int64`

GetLeaseUntilEpochMs returns the LeaseUntilEpochMs field if non-nil, zero value otherwise.

### GetLeaseUntilEpochMsOk

`func (o *AgentRegistration) GetLeaseUntilEpochMsOk() (*int64, bool)`

GetLeaseUntilEpochMsOk returns a tuple with the LeaseUntilEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseUntilEpochMs

`func (o *AgentRegistration) SetLeaseUntilEpochMs(v int64)`

SetLeaseUntilEpochMs sets LeaseUntilEpochMs field to given value.


### GetRegisteredEpochMs

`func (o *AgentRegistration) GetRegisteredEpochMs() int64`

GetRegisteredEpochMs returns the RegisteredEpochMs field if non-nil, zero value otherwise.

### GetRegisteredEpochMsOk

`func (o *AgentRegistration) GetRegisteredEpochMsOk() (*int64, bool)`

GetRegisteredEpochMsOk returns a tuple with the RegisteredEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEpochMs

`func (o *AgentRegistration) SetRegisteredEpochMs(v int64)`

SetRegisteredEpochMs sets RegisteredEpochMs field to given value.


### GetAgentData

`func (o *AgentRegistration) GetAgentData() AgentData`

GetAgentData returns the AgentData field if non-nil, zero value otherwise.

### GetAgentDataOk

`func (o *AgentRegistration) GetAgentDataOk() (*AgentData, bool)`

GetAgentDataOk returns a tuple with the AgentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentData

`func (o *AgentRegistration) SetAgentData(v AgentData)`

SetAgentData sets AgentData field to given value.

### HasAgentData

`func (o *AgentRegistration) HasAgentData() bool`

HasAgentData returns a boolean if a field has been set.

### GetNodeBudgetCount

`func (o *AgentRegistration) GetNodeBudgetCount() int32`

GetNodeBudgetCount returns the NodeBudgetCount field if non-nil, zero value otherwise.

### GetNodeBudgetCountOk

`func (o *AgentRegistration) GetNodeBudgetCountOk() (*int32, bool)`

GetNodeBudgetCountOk returns a tuple with the NodeBudgetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeBudgetCount

`func (o *AgentRegistration) SetNodeBudgetCount(v int32)`

SetNodeBudgetCount sets NodeBudgetCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


