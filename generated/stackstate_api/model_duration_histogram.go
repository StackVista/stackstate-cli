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

// DurationHistogram struct for DurationHistogram
type DurationHistogram struct {
	// List of duration buckets
	Buckets   []DurationHistogramBucket `json:"buckets"`
	Quantiles DurationQuantiles         `json:"quantiles"`
}

// NewDurationHistogram instantiates a new DurationHistogram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDurationHistogram(buckets []DurationHistogramBucket, quantiles DurationQuantiles) *DurationHistogram {
	this := DurationHistogram{}
	this.Buckets = buckets
	this.Quantiles = quantiles
	return &this
}

// NewDurationHistogramWithDefaults instantiates a new DurationHistogram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDurationHistogramWithDefaults() *DurationHistogram {
	this := DurationHistogram{}
	return &this
}

// GetBuckets returns the Buckets field value
func (o *DurationHistogram) GetBuckets() []DurationHistogramBucket {
	if o == nil {
		var ret []DurationHistogramBucket
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *DurationHistogram) GetBucketsOk() ([]DurationHistogramBucket, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *DurationHistogram) SetBuckets(v []DurationHistogramBucket) {
	o.Buckets = v
}

// GetQuantiles returns the Quantiles field value
func (o *DurationHistogram) GetQuantiles() DurationQuantiles {
	if o == nil {
		var ret DurationQuantiles
		return ret
	}

	return o.Quantiles
}

// GetQuantilesOk returns a tuple with the Quantiles field value
// and a boolean to check if the value has been set.
func (o *DurationHistogram) GetQuantilesOk() (*DurationQuantiles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantiles, true
}

// SetQuantiles sets field value
func (o *DurationHistogram) SetQuantiles(v DurationQuantiles) {
	o.Quantiles = v
}

func (o DurationHistogram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["buckets"] = o.Buckets
	}
	if true {
		toSerialize["quantiles"] = o.Quantiles
	}
	return json.Marshal(toSerialize)
}

type NullableDurationHistogram struct {
	value *DurationHistogram
	isSet bool
}

func (v NullableDurationHistogram) Get() *DurationHistogram {
	return v.value
}

func (v *NullableDurationHistogram) Set(val *DurationHistogram) {
	v.value = val
	v.isSet = true
}

func (v NullableDurationHistogram) IsSet() bool {
	return v.isSet
}

func (v *NullableDurationHistogram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDurationHistogram(val *DurationHistogram) *NullableDurationHistogram {
	return &NullableDurationHistogram{value: val, isSet: true}
}

func (v NullableDurationHistogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDurationHistogram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
