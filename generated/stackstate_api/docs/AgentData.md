# AgentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**CoreCount** | **int32** |  | 
**MemoryBytes** | **int64** |  | 
**KernelVersion** | **string** |  | 

## Methods

### NewAgentData

`func NewAgentData(platform string, coreCount int32, memoryBytes int64, kernelVersion string, ) *AgentData`

NewAgentData instantiates a new AgentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDataWithDefaults

`func NewAgentDataWithDefaults() *AgentData`

NewAgentDataWithDefaults instantiates a new AgentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AgentData) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AgentData) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AgentData) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCoreCount

`func (o *AgentData) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *AgentData) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *AgentData) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.


### GetMemoryBytes

`func (o *AgentData) GetMemoryBytes() int64`

GetMemoryBytes returns the MemoryBytes field if non-nil, zero value otherwise.

### GetMemoryBytesOk

`func (o *AgentData) GetMemoryBytesOk() (*int64, bool)`

GetMemoryBytesOk returns a tuple with the MemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBytes

`func (o *AgentData) SetMemoryBytes(v int64)`

SetMemoryBytes sets MemoryBytes field to given value.


### GetKernelVersion

`func (o *AgentData) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *AgentData) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *AgentData) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


