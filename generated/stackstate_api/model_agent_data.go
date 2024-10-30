/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// AgentData struct for AgentData
type AgentData struct {
	Platform      string `json:"platform"`
	CoreCount     int32  `json:"coreCount"`
	MemoryBytes   int64  `json:"memoryBytes"`
	KernelVersion string `json:"kernelVersion"`
}

// NewAgentData instantiates a new AgentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentData(platform string, coreCount int32, memoryBytes int64, kernelVersion string) *AgentData {
	this := AgentData{}
	this.Platform = platform
	this.CoreCount = coreCount
	this.MemoryBytes = memoryBytes
	this.KernelVersion = kernelVersion
	return &this
}

// NewAgentDataWithDefaults instantiates a new AgentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentDataWithDefaults() *AgentData {
	this := AgentData{}
	return &this
}

// GetPlatform returns the Platform field value
func (o *AgentData) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *AgentData) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *AgentData) SetPlatform(v string) {
	o.Platform = v
}

// GetCoreCount returns the CoreCount field value
func (o *AgentData) GetCoreCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value
// and a boolean to check if the value has been set.
func (o *AgentData) GetCoreCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoreCount, true
}

// SetCoreCount sets field value
func (o *AgentData) SetCoreCount(v int32) {
	o.CoreCount = v
}

// GetMemoryBytes returns the MemoryBytes field value
func (o *AgentData) GetMemoryBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MemoryBytes
}

// GetMemoryBytesOk returns a tuple with the MemoryBytes field value
// and a boolean to check if the value has been set.
func (o *AgentData) GetMemoryBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryBytes, true
}

// SetMemoryBytes sets field value
func (o *AgentData) SetMemoryBytes(v int64) {
	o.MemoryBytes = v
}

// GetKernelVersion returns the KernelVersion field value
func (o *AgentData) GetKernelVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value
// and a boolean to check if the value has been set.
func (o *AgentData) GetKernelVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KernelVersion, true
}

// SetKernelVersion sets field value
func (o *AgentData) SetKernelVersion(v string) {
	o.KernelVersion = v
}

func (o AgentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["coreCount"] = o.CoreCount
	}
	if true {
		toSerialize["memoryBytes"] = o.MemoryBytes
	}
	if true {
		toSerialize["kernelVersion"] = o.KernelVersion
	}
	return json.Marshal(toSerialize)
}

type NullableAgentData struct {
	value *AgentData
	isSet bool
}

func (v NullableAgentData) Get() *AgentData {
	return v.value
}

func (v *NullableAgentData) Set(val *AgentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentData(val *AgentData) *NullableAgentData {
	return &NullableAgentData{value: val, isSet: true}
}

func (v NullableAgentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}