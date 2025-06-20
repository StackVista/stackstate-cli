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

// PersesListVariable struct for PersesListVariable
type PersesListVariable struct {
	Kind string                  `json:"kind"`
	Spec *PersesListVariableSpec `json:"spec,omitempty"`
}

// NewPersesListVariable instantiates a new PersesListVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersesListVariable(kind string) *PersesListVariable {
	this := PersesListVariable{}
	this.Kind = kind
	return &this
}

// NewPersesListVariableWithDefaults instantiates a new PersesListVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersesListVariableWithDefaults() *PersesListVariable {
	this := PersesListVariable{}
	return &this
}

// GetKind returns the Kind field value
func (o *PersesListVariable) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PersesListVariable) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PersesListVariable) SetKind(v string) {
	o.Kind = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *PersesListVariable) GetSpec() PersesListVariableSpec {
	if o == nil || o.Spec == nil {
		var ret PersesListVariableSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesListVariable) GetSpecOk() (*PersesListVariableSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *PersesListVariable) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given PersesListVariableSpec and assigns it to the Spec field.
func (o *PersesListVariable) SetSpec(v PersesListVariableSpec) {
	o.Spec = &v
}

func (o PersesListVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullablePersesListVariable struct {
	value *PersesListVariable
	isSet bool
}

func (v NullablePersesListVariable) Get() *PersesListVariable {
	return v.value
}

func (v *NullablePersesListVariable) Set(val *PersesListVariable) {
	v.value = val
	v.isSet = true
}

func (v NullablePersesListVariable) IsSet() bool {
	return v.isSet
}

func (v *NullablePersesListVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersesListVariable(val *PersesListVariable) *NullablePersesListVariable {
	return &NullablePersesListVariable{value: val, isSet: true}
}

func (v NullablePersesListVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersesListVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
