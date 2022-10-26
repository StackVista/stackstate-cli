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

// PromLabelsEnvelope struct for PromLabelsEnvelope
type PromLabelsEnvelope struct {
	Status    string   `json:"status"`
	Data      []string `json:"data,omitempty"`
	ErrorType *string  `json:"errorType,omitempty"`
	Error     *string  `json:"error,omitempty"`
	Warnings  []string `json:"warnings,omitempty"`
}

// NewPromLabelsEnvelope instantiates a new PromLabelsEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromLabelsEnvelope(status string) *PromLabelsEnvelope {
	this := PromLabelsEnvelope{}
	this.Status = status
	return &this
}

// NewPromLabelsEnvelopeWithDefaults instantiates a new PromLabelsEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromLabelsEnvelopeWithDefaults() *PromLabelsEnvelope {
	this := PromLabelsEnvelope{}
	return &this
}

// GetStatus returns the Status field value
func (o *PromLabelsEnvelope) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PromLabelsEnvelope) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PromLabelsEnvelope) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PromLabelsEnvelope) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromLabelsEnvelope) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PromLabelsEnvelope) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *PromLabelsEnvelope) SetData(v []string) {
	o.Data = v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *PromLabelsEnvelope) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromLabelsEnvelope) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *PromLabelsEnvelope) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *PromLabelsEnvelope) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PromLabelsEnvelope) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromLabelsEnvelope) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PromLabelsEnvelope) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *PromLabelsEnvelope) SetError(v string) {
	o.Error = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PromLabelsEnvelope) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromLabelsEnvelope) GetWarningsOk() ([]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PromLabelsEnvelope) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *PromLabelsEnvelope) SetWarnings(v []string) {
	o.Warnings = v
}

func (o PromLabelsEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ErrorType != nil {
		toSerialize["errorType"] = o.ErrorType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullablePromLabelsEnvelope struct {
	value *PromLabelsEnvelope
	isSet bool
}

func (v NullablePromLabelsEnvelope) Get() *PromLabelsEnvelope {
	return v.value
}

func (v *NullablePromLabelsEnvelope) Set(val *PromLabelsEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullablePromLabelsEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullablePromLabelsEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromLabelsEnvelope(val *PromLabelsEnvelope) *NullablePromLabelsEnvelope {
	return &NullablePromLabelsEnvelope{value: val, isSet: true}
}

func (v NullablePromLabelsEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromLabelsEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}