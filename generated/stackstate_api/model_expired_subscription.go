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

// ExpiredSubscription struct for ExpiredSubscription
type ExpiredSubscription struct {
	Type string `json:"_type"`
	Subscription Subscription `json:"subscription"`
}

// NewExpiredSubscription instantiates a new ExpiredSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiredSubscription(type_ string, subscription Subscription) *ExpiredSubscription {
	this := ExpiredSubscription{}
	this.Type = type_
	this.Subscription = subscription
	return &this
}

// NewExpiredSubscriptionWithDefaults instantiates a new ExpiredSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiredSubscriptionWithDefaults() *ExpiredSubscription {
	this := ExpiredSubscription{}
	return &this
}

// GetType returns the Type field value
func (o *ExpiredSubscription) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExpiredSubscription) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExpiredSubscription) SetType(v string) {
	o.Type = v
}

// GetSubscription returns the Subscription field value
func (o *ExpiredSubscription) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *ExpiredSubscription) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *ExpiredSubscription) SetSubscription(v Subscription) {
	o.Subscription = v
}

func (o ExpiredSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableExpiredSubscription struct {
	value *ExpiredSubscription
	isSet bool
}

func (v NullableExpiredSubscription) Get() *ExpiredSubscription {
	return v.value
}

func (v *NullableExpiredSubscription) Set(val *ExpiredSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiredSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiredSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiredSubscription(val *ExpiredSubscription) *NullableExpiredSubscription {
	return &NullableExpiredSubscription{value: val, isSet: true}
}

func (v NullableExpiredSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiredSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

