/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
)

// StackPackFaqs struct for StackPackFaqs
type StackPackFaqs struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
}

// NewStackPackFaqs instantiates a new StackPackFaqs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackPackFaqs(question string, answer string) *StackPackFaqs {
	this := StackPackFaqs{}
	this.Question = question
	this.Answer = answer
	return &this
}

// NewStackPackFaqsWithDefaults instantiates a new StackPackFaqs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackPackFaqsWithDefaults() *StackPackFaqs {
	this := StackPackFaqs{}
	return &this
}

// GetQuestion returns the Question field value
func (o *StackPackFaqs) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *StackPackFaqs) GetQuestionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *StackPackFaqs) SetQuestion(v string) {
	o.Question = v
}

// GetAnswer returns the Answer field value
func (o *StackPackFaqs) GetAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value
// and a boolean to check if the value has been set.
func (o *StackPackFaqs) GetAnswerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Answer, true
}

// SetAnswer sets field value
func (o *StackPackFaqs) SetAnswer(v string) {
	o.Answer = v
}

func (o StackPackFaqs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["question"] = o.Question
	}
	if true {
		toSerialize["answer"] = o.Answer
	}
	return json.Marshal(toSerialize)
}

type NullableStackPackFaqs struct {
	value *StackPackFaqs
	isSet bool
}

func (v NullableStackPackFaqs) Get() *StackPackFaqs {
	return v.value
}

func (v *NullableStackPackFaqs) Set(val *StackPackFaqs) {
	v.value = val
	v.isSet = true
}

func (v NullableStackPackFaqs) IsSet() bool {
	return v.isSet
}

func (v *NullableStackPackFaqs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackPackFaqs(val *StackPackFaqs) *NullableStackPackFaqs {
	return &NullableStackPackFaqs{value: val, isSet: true}
}

func (v NullableStackPackFaqs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackPackFaqs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


