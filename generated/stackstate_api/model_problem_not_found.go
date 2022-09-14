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

// ProblemNotFound struct for ProblemNotFound
type ProblemNotFound struct {
	Type string `json:"_type"`
	Message string `json:"message"`
	ProblemId int64 `json:"problemId"`
	RequestTimeMs int32 `json:"requestTimeMs"`
}

// NewProblemNotFound instantiates a new ProblemNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemNotFound(type_ string, message string, problemId int64, requestTimeMs int32) *ProblemNotFound {
	this := ProblemNotFound{}
	this.Type = type_
	this.Message = message
	this.ProblemId = problemId
	this.RequestTimeMs = requestTimeMs
	return &this
}

// NewProblemNotFoundWithDefaults instantiates a new ProblemNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemNotFoundWithDefaults() *ProblemNotFound {
	this := ProblemNotFound{}
	return &this
}

// GetType returns the Type field value
func (o *ProblemNotFound) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProblemNotFound) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProblemNotFound) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *ProblemNotFound) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ProblemNotFound) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ProblemNotFound) SetMessage(v string) {
	o.Message = v
}

// GetProblemId returns the ProblemId field value
func (o *ProblemNotFound) GetProblemId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProblemId
}

// GetProblemIdOk returns a tuple with the ProblemId field value
// and a boolean to check if the value has been set.
func (o *ProblemNotFound) GetProblemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProblemId, true
}

// SetProblemId sets field value
func (o *ProblemNotFound) SetProblemId(v int64) {
	o.ProblemId = v
}

// GetRequestTimeMs returns the RequestTimeMs field value
func (o *ProblemNotFound) GetRequestTimeMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestTimeMs
}

// GetRequestTimeMsOk returns a tuple with the RequestTimeMs field value
// and a boolean to check if the value has been set.
func (o *ProblemNotFound) GetRequestTimeMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTimeMs, true
}

// SetRequestTimeMs sets field value
func (o *ProblemNotFound) SetRequestTimeMs(v int32) {
	o.RequestTimeMs = v
}

func (o ProblemNotFound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["problemId"] = o.ProblemId
	}
	if true {
		toSerialize["requestTimeMs"] = o.RequestTimeMs
	}
	return json.Marshal(toSerialize)
}

type NullableProblemNotFound struct {
	value *ProblemNotFound
	isSet bool
}

func (v NullableProblemNotFound) Get() *ProblemNotFound {
	return v.value
}

func (v *NullableProblemNotFound) Set(val *ProblemNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemNotFound(val *ProblemNotFound) *NullableProblemNotFound {
	return &NullableProblemNotFound{value: val, isSet: true}
}

func (v NullableProblemNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

