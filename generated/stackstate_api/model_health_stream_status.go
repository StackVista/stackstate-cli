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

// HealthStreamStatus struct for HealthStreamStatus
type HealthStreamStatus struct {
	Partition int32 `json:"partition"`
	ConsistencyModel string `json:"consistencyModel"`
	RecoverMessage *string `json:"recoverMessage,omitempty"`
	GlobalErrors *[]HealthStreamError `json:"globalErrors,omitempty"`
	AggregateMetrics HealthStreamMetrics `json:"aggregateMetrics"`
	MainStreamStatus *HealthSubStreamStatus `json:"mainStreamStatus,omitempty"`
}

// NewHealthStreamStatus instantiates a new HealthStreamStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthStreamStatus(partition int32, consistencyModel string, aggregateMetrics HealthStreamMetrics) *HealthStreamStatus {
	this := HealthStreamStatus{}
	this.Partition = partition
	this.ConsistencyModel = consistencyModel
	this.AggregateMetrics = aggregateMetrics
	return &this
}

// NewHealthStreamStatusWithDefaults instantiates a new HealthStreamStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthStreamStatusWithDefaults() *HealthStreamStatus {
	this := HealthStreamStatus{}
	return &this
}

// GetPartition returns the Partition field value
func (o *HealthStreamStatus) GetPartition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetPartitionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value
func (o *HealthStreamStatus) SetPartition(v int32) {
	o.Partition = v
}

// GetConsistencyModel returns the ConsistencyModel field value
func (o *HealthStreamStatus) GetConsistencyModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsistencyModel
}

// GetConsistencyModelOk returns a tuple with the ConsistencyModel field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetConsistencyModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsistencyModel, true
}

// SetConsistencyModel sets field value
func (o *HealthStreamStatus) SetConsistencyModel(v string) {
	o.ConsistencyModel = v
}

// GetRecoverMessage returns the RecoverMessage field value if set, zero value otherwise.
func (o *HealthStreamStatus) GetRecoverMessage() string {
	if o == nil || o.RecoverMessage == nil {
		var ret string
		return ret
	}
	return *o.RecoverMessage
}

// GetRecoverMessageOk returns a tuple with the RecoverMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetRecoverMessageOk() (*string, bool) {
	if o == nil || o.RecoverMessage == nil {
		return nil, false
	}
	return o.RecoverMessage, true
}

// HasRecoverMessage returns a boolean if a field has been set.
func (o *HealthStreamStatus) HasRecoverMessage() bool {
	if o != nil && o.RecoverMessage != nil {
		return true
	}

	return false
}

// SetRecoverMessage gets a reference to the given string and assigns it to the RecoverMessage field.
func (o *HealthStreamStatus) SetRecoverMessage(v string) {
	o.RecoverMessage = &v
}

// GetGlobalErrors returns the GlobalErrors field value if set, zero value otherwise.
func (o *HealthStreamStatus) GetGlobalErrors() []HealthStreamError {
	if o == nil || o.GlobalErrors == nil {
		var ret []HealthStreamError
		return ret
	}
	return *o.GlobalErrors
}

// GetGlobalErrorsOk returns a tuple with the GlobalErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetGlobalErrorsOk() (*[]HealthStreamError, bool) {
	if o == nil || o.GlobalErrors == nil {
		return nil, false
	}
	return o.GlobalErrors, true
}

// HasGlobalErrors returns a boolean if a field has been set.
func (o *HealthStreamStatus) HasGlobalErrors() bool {
	if o != nil && o.GlobalErrors != nil {
		return true
	}

	return false
}

// SetGlobalErrors gets a reference to the given []HealthStreamError and assigns it to the GlobalErrors field.
func (o *HealthStreamStatus) SetGlobalErrors(v []HealthStreamError) {
	o.GlobalErrors = &v
}

// GetAggregateMetrics returns the AggregateMetrics field value
func (o *HealthStreamStatus) GetAggregateMetrics() HealthStreamMetrics {
	if o == nil {
		var ret HealthStreamMetrics
		return ret
	}

	return o.AggregateMetrics
}

// GetAggregateMetricsOk returns a tuple with the AggregateMetrics field value
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetAggregateMetricsOk() (*HealthStreamMetrics, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AggregateMetrics, true
}

// SetAggregateMetrics sets field value
func (o *HealthStreamStatus) SetAggregateMetrics(v HealthStreamMetrics) {
	o.AggregateMetrics = v
}

// GetMainStreamStatus returns the MainStreamStatus field value if set, zero value otherwise.
func (o *HealthStreamStatus) GetMainStreamStatus() HealthSubStreamStatus {
	if o == nil || o.MainStreamStatus == nil {
		var ret HealthSubStreamStatus
		return ret
	}
	return *o.MainStreamStatus
}

// GetMainStreamStatusOk returns a tuple with the MainStreamStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamStatus) GetMainStreamStatusOk() (*HealthSubStreamStatus, bool) {
	if o == nil || o.MainStreamStatus == nil {
		return nil, false
	}
	return o.MainStreamStatus, true
}

// HasMainStreamStatus returns a boolean if a field has been set.
func (o *HealthStreamStatus) HasMainStreamStatus() bool {
	if o != nil && o.MainStreamStatus != nil {
		return true
	}

	return false
}

// SetMainStreamStatus gets a reference to the given HealthSubStreamStatus and assigns it to the MainStreamStatus field.
func (o *HealthStreamStatus) SetMainStreamStatus(v HealthSubStreamStatus) {
	o.MainStreamStatus = &v
}

func (o HealthStreamStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["partition"] = o.Partition
	}
	if true {
		toSerialize["consistencyModel"] = o.ConsistencyModel
	}
	if o.RecoverMessage != nil {
		toSerialize["recoverMessage"] = o.RecoverMessage
	}
	if o.GlobalErrors != nil {
		toSerialize["globalErrors"] = o.GlobalErrors
	}
	if true {
		toSerialize["aggregateMetrics"] = o.AggregateMetrics
	}
	if o.MainStreamStatus != nil {
		toSerialize["mainStreamStatus"] = o.MainStreamStatus
	}
	return json.Marshal(toSerialize)
}

type NullableHealthStreamStatus struct {
	value *HealthStreamStatus
	isSet bool
}

func (v NullableHealthStreamStatus) Get() *HealthStreamStatus {
	return v.value
}

func (v *NullableHealthStreamStatus) Set(val *HealthStreamStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStreamStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStreamStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStreamStatus(val *HealthStreamStatus) *NullableHealthStreamStatus {
	return &NullableHealthStreamStatus{value: val, isSet: true}
}

func (v NullableHealthStreamStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStreamStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


