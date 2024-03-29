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

// Subscription struct for Subscription
type Subscription struct {
	Tenant            string `json:"tenant"`
	ExpiryTimestampMs *int64 `json:"expiryTimestampMs,omitempty"`
	Plan              string `json:"plan"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(tenant string, plan string) *Subscription {
	this := Subscription{}
	this.Tenant = tenant
	this.Plan = plan
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetTenant returns the Tenant field value
func (o *Subscription) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *Subscription) SetTenant(v string) {
	o.Tenant = v
}

// GetExpiryTimestampMs returns the ExpiryTimestampMs field value if set, zero value otherwise.
func (o *Subscription) GetExpiryTimestampMs() int64 {
	if o == nil || o.ExpiryTimestampMs == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTimestampMs
}

// GetExpiryTimestampMsOk returns a tuple with the ExpiryTimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetExpiryTimestampMsOk() (*int64, bool) {
	if o == nil || o.ExpiryTimestampMs == nil {
		return nil, false
	}
	return o.ExpiryTimestampMs, true
}

// HasExpiryTimestampMs returns a boolean if a field has been set.
func (o *Subscription) HasExpiryTimestampMs() bool {
	if o != nil && o.ExpiryTimestampMs != nil {
		return true
	}

	return false
}

// SetExpiryTimestampMs gets a reference to the given int64 and assigns it to the ExpiryTimestampMs field.
func (o *Subscription) SetExpiryTimestampMs(v int64) {
	o.ExpiryTimestampMs = &v
}

// GetPlan returns the Plan field value
func (o *Subscription) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *Subscription) SetPlan(v string) {
	o.Plan = v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if o.ExpiryTimestampMs != nil {
		toSerialize["expiryTimestampMs"] = o.ExpiryTimestampMs
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
