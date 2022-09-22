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

// MonitorPatch struct for MonitorPatch
type MonitorPatch struct {
	Name *string `json:"name,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	RemediationHint *string `json:"remediationHint,omitempty"`
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Status *MonitorStatusValue `json:"status,omitempty"`
}

// NewMonitorPatch instantiates a new MonitorPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorPatch() *MonitorPatch {
	this := MonitorPatch{}
	return &this
}

// NewMonitorPatchWithDefaults instantiates a new MonitorPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorPatchWithDefaults() *MonitorPatch {
	this := MonitorPatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorPatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorPatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorPatch) SetName(v string) {
	o.Name = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MonitorPatch) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MonitorPatch) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *MonitorPatch) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MonitorPatch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorPatch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MonitorPatch) SetDescription(v string) {
	o.Description = &v
}

// GetRemediationHint returns the RemediationHint field value if set, zero value otherwise.
func (o *MonitorPatch) GetRemediationHint() string {
	if o == nil || o.RemediationHint == nil {
		var ret string
		return ret
	}
	return *o.RemediationHint
}

// GetRemediationHintOk returns a tuple with the RemediationHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetRemediationHintOk() (*string, bool) {
	if o == nil || o.RemediationHint == nil {
		return nil, false
	}
	return o.RemediationHint, true
}

// HasRemediationHint returns a boolean if a field has been set.
func (o *MonitorPatch) HasRemediationHint() bool {
	if o != nil && o.RemediationHint != nil {
		return true
	}

	return false
}

// SetRemediationHint gets a reference to the given string and assigns it to the RemediationHint field.
func (o *MonitorPatch) SetRemediationHint(v string) {
	o.RemediationHint = &v
}

// GetIntervalSeconds returns the IntervalSeconds field value if set, zero value otherwise.
func (o *MonitorPatch) GetIntervalSeconds() int32 {
	if o == nil || o.IntervalSeconds == nil {
		var ret int32
		return ret
	}
	return *o.IntervalSeconds
}

// GetIntervalSecondsOk returns a tuple with the IntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetIntervalSecondsOk() (*int32, bool) {
	if o == nil || o.IntervalSeconds == nil {
		return nil, false
	}
	return o.IntervalSeconds, true
}

// HasIntervalSeconds returns a boolean if a field has been set.
func (o *MonitorPatch) HasIntervalSeconds() bool {
	if o != nil && o.IntervalSeconds != nil {
		return true
	}

	return false
}

// SetIntervalSeconds gets a reference to the given int32 and assigns it to the IntervalSeconds field.
func (o *MonitorPatch) SetIntervalSeconds(v int32) {
	o.IntervalSeconds = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonitorPatch) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonitorPatch) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MonitorPatch) SetTags(v []string) {
	o.Tags = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorPatch) GetStatus() MonitorStatusValue {
	if o == nil || o.Status == nil {
		var ret MonitorStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorPatch) GetStatusOk() (*MonitorStatusValue, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorPatch) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MonitorStatusValue and assigns it to the Status field.
func (o *MonitorPatch) SetStatus(v MonitorStatusValue) {
	o.Status = &v
}

func (o MonitorPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RemediationHint != nil {
		toSerialize["remediationHint"] = o.RemediationHint
	}
	if o.IntervalSeconds != nil {
		toSerialize["intervalSeconds"] = o.IntervalSeconds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorPatch struct {
	value *MonitorPatch
	isSet bool
}

func (v NullableMonitorPatch) Get() *MonitorPatch {
	return v.value
}

func (v *NullableMonitorPatch) Set(val *MonitorPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorPatch(val *MonitorPatch) *NullableMonitorPatch {
	return &NullableMonitorPatch{value: val, isSet: true}
}

func (v NullableMonitorPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


