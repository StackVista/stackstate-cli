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

// TelemetryLatestSnapshotsResponse struct for TelemetryLatestSnapshotsResponse
type TelemetryLatestSnapshotsResponse struct {
	Type            string                         `json:"_type"`
	StreamSnapshots []LatestTelemetryStreamMetrics `json:"streamSnapshots"`
}

// NewTelemetryLatestSnapshotsResponse instantiates a new TelemetryLatestSnapshotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryLatestSnapshotsResponse(type_ string, streamSnapshots []LatestTelemetryStreamMetrics) *TelemetryLatestSnapshotsResponse {
	this := TelemetryLatestSnapshotsResponse{}
	this.Type = type_
	this.StreamSnapshots = streamSnapshots
	return &this
}

// NewTelemetryLatestSnapshotsResponseWithDefaults instantiates a new TelemetryLatestSnapshotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryLatestSnapshotsResponseWithDefaults() *TelemetryLatestSnapshotsResponse {
	this := TelemetryLatestSnapshotsResponse{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryLatestSnapshotsResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryLatestSnapshotsResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryLatestSnapshotsResponse) SetType(v string) {
	o.Type = v
}

// GetStreamSnapshots returns the StreamSnapshots field value
func (o *TelemetryLatestSnapshotsResponse) GetStreamSnapshots() []LatestTelemetryStreamMetrics {
	if o == nil {
		var ret []LatestTelemetryStreamMetrics
		return ret
	}

	return o.StreamSnapshots
}

// GetStreamSnapshotsOk returns a tuple with the StreamSnapshots field value
// and a boolean to check if the value has been set.
func (o *TelemetryLatestSnapshotsResponse) GetStreamSnapshotsOk() ([]LatestTelemetryStreamMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamSnapshots, true
}

// SetStreamSnapshots sets field value
func (o *TelemetryLatestSnapshotsResponse) SetStreamSnapshots(v []LatestTelemetryStreamMetrics) {
	o.StreamSnapshots = v
}

func (o TelemetryLatestSnapshotsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["streamSnapshots"] = o.StreamSnapshots
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryLatestSnapshotsResponse struct {
	value *TelemetryLatestSnapshotsResponse
	isSet bool
}

func (v NullableTelemetryLatestSnapshotsResponse) Get() *TelemetryLatestSnapshotsResponse {
	return v.value
}

func (v *NullableTelemetryLatestSnapshotsResponse) Set(val *TelemetryLatestSnapshotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryLatestSnapshotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryLatestSnapshotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryLatestSnapshotsResponse(val *TelemetryLatestSnapshotsResponse) *NullableTelemetryLatestSnapshotsResponse {
	return &NullableTelemetryLatestSnapshotsResponse{value: val, isSet: true}
}

func (v NullableTelemetryLatestSnapshotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryLatestSnapshotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
