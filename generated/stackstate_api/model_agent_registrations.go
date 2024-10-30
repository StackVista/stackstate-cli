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

// AgentRegistrations struct for AgentRegistrations
type AgentRegistrations struct {
	Agents []AgentRegistration `json:"agents"`
}

// NewAgentRegistrations instantiates a new AgentRegistrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRegistrations(agents []AgentRegistration) *AgentRegistrations {
	this := AgentRegistrations{}
	this.Agents = agents
	return &this
}

// NewAgentRegistrationsWithDefaults instantiates a new AgentRegistrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRegistrationsWithDefaults() *AgentRegistrations {
	this := AgentRegistrations{}
	return &this
}

// GetAgents returns the Agents field value
func (o *AgentRegistrations) GetAgents() []AgentRegistration {
	if o == nil {
		var ret []AgentRegistration
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *AgentRegistrations) GetAgentsOk() ([]AgentRegistration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agents, true
}

// SetAgents sets field value
func (o *AgentRegistrations) SetAgents(v []AgentRegistration) {
	o.Agents = v
}

func (o AgentRegistrations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["agents"] = o.Agents
	}
	return json.Marshal(toSerialize)
}

type NullableAgentRegistrations struct {
	value *AgentRegistrations
	isSet bool
}

func (v NullableAgentRegistrations) Get() *AgentRegistrations {
	return v.value
}

func (v *NullableAgentRegistrations) Set(val *AgentRegistrations) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRegistrations) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRegistrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRegistrations(val *AgentRegistrations) *NullableAgentRegistrations {
	return &NullableAgentRegistrations{value: val, isSet: true}
}

func (v NullableAgentRegistrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRegistrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}