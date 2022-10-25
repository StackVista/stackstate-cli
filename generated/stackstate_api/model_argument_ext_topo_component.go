/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// ArgumentExtTopoComponent struct for ArgumentExtTopoComponent
type ArgumentExtTopoComponent struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	ExtTopoElement      int64  `json:"extTopoElement"`
}

// NewArgumentExtTopoComponent instantiates a new ArgumentExtTopoComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentExtTopoComponent(type_ string, parameter int64, extTopoElement int64) *ArgumentExtTopoComponent {
	this := ArgumentExtTopoComponent{}
	this.Type = type_
	this.Parameter = parameter
	this.ExtTopoElement = extTopoElement
	return &this
}

// NewArgumentExtTopoComponentWithDefaults instantiates a new ArgumentExtTopoComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentExtTopoComponentWithDefaults() *ArgumentExtTopoComponent {
	this := ArgumentExtTopoComponent{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentExtTopoComponent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentExtTopoComponent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentExtTopoComponent) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentExtTopoComponent) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentExtTopoComponent) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentExtTopoComponent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentExtTopoComponent) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentExtTopoComponent) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentExtTopoComponent) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentExtTopoComponent) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentExtTopoComponent) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentExtTopoComponent) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentExtTopoComponent) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentExtTopoComponent) SetParameter(v int64) {
	o.Parameter = v
}

// GetExtTopoElement returns the ExtTopoElement field value
func (o *ArgumentExtTopoComponent) GetExtTopoElement() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExtTopoElement
}

// GetExtTopoElementOk returns a tuple with the ExtTopoElement field value
// and a boolean to check if the value has been set.
func (o *ArgumentExtTopoComponent) GetExtTopoElementOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtTopoElement, true
}

// SetExtTopoElement sets field value
func (o *ArgumentExtTopoComponent) SetExtTopoElement(v int64) {
	o.ExtTopoElement = v
}

func (o ArgumentExtTopoComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	if true {
		toSerialize["extTopoElement"] = o.ExtTopoElement
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentExtTopoComponent struct {
	value *ArgumentExtTopoComponent
	isSet bool
}

func (v NullableArgumentExtTopoComponent) Get() *ArgumentExtTopoComponent {
	return v.value
}

func (v *NullableArgumentExtTopoComponent) Set(val *ArgumentExtTopoComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentExtTopoComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentExtTopoComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentExtTopoComponent(val *ArgumentExtTopoComponent) *NullableArgumentExtTopoComponent {
	return &NullableArgumentExtTopoComponent{value: val, isSet: true}
}

func (v NullableArgumentExtTopoComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentExtTopoComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
