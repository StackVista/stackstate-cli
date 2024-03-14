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

// KubernetesLogSeverityHistogramBucket struct for KubernetesLogSeverityHistogramBucket
type KubernetesLogSeverityHistogramBucket struct {
	// Total logs record count for a particular log severity in the bucket.
	Count    int64       `json:"count"`
	Severity LogSeverity `json:"severity"`
}

// NewKubernetesLogSeverityHistogramBucket instantiates a new KubernetesLogSeverityHistogramBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesLogSeverityHistogramBucket(count int64, severity LogSeverity) *KubernetesLogSeverityHistogramBucket {
	this := KubernetesLogSeverityHistogramBucket{}
	this.Count = count
	this.Severity = severity
	return &this
}

// NewKubernetesLogSeverityHistogramBucketWithDefaults instantiates a new KubernetesLogSeverityHistogramBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesLogSeverityHistogramBucketWithDefaults() *KubernetesLogSeverityHistogramBucket {
	this := KubernetesLogSeverityHistogramBucket{}
	var severity LogSeverity = LOGSEVERITY_OTHER
	this.Severity = severity
	return &this
}

// GetCount returns the Count field value
func (o *KubernetesLogSeverityHistogramBucket) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *KubernetesLogSeverityHistogramBucket) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *KubernetesLogSeverityHistogramBucket) SetCount(v int64) {
	o.Count = v
}

// GetSeverity returns the Severity field value
func (o *KubernetesLogSeverityHistogramBucket) GetSeverity() LogSeverity {
	if o == nil {
		var ret LogSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *KubernetesLogSeverityHistogramBucket) GetSeverityOk() (*LogSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *KubernetesLogSeverityHistogramBucket) SetSeverity(v LogSeverity) {
	o.Severity = v
}

func (o KubernetesLogSeverityHistogramBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesLogSeverityHistogramBucket struct {
	value *KubernetesLogSeverityHistogramBucket
	isSet bool
}

func (v NullableKubernetesLogSeverityHistogramBucket) Get() *KubernetesLogSeverityHistogramBucket {
	return v.value
}

func (v *NullableKubernetesLogSeverityHistogramBucket) Set(val *KubernetesLogSeverityHistogramBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesLogSeverityHistogramBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesLogSeverityHistogramBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesLogSeverityHistogramBucket(val *KubernetesLogSeverityHistogramBucket) *NullableKubernetesLogSeverityHistogramBucket {
	return &NullableKubernetesLogSeverityHistogramBucket{value: val, isSet: true}
}

func (v NullableKubernetesLogSeverityHistogramBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesLogSeverityHistogramBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}