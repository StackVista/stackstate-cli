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
	"fmt"
)

// DataStream - struct for DataStream
type DataStream struct {
	EventStream *EventStream
	MetricStream *MetricStream
}

// EventStreamAsDataStream is a convenience function that returns EventStream wrapped in DataStream
func EventStreamAsDataStream(v *EventStream) DataStream {
	return DataStream{
		EventStream: v,
	}
}

// MetricStreamAsDataStream is a convenience function that returns MetricStream wrapped in DataStream
func MetricStreamAsDataStream(v *MetricStream) DataStream {
	return DataStream{
		MetricStream: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataStream) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'EventStream'
	if jsonDict["_type"] == "EventStream" {
		// try to unmarshal JSON data into EventStream
		err = json.Unmarshal(data, &dst.EventStream)
		if err == nil {
			return nil // data stored in dst.EventStream, return on the first match
		} else {
			dst.EventStream = nil
			return fmt.Errorf("Failed to unmarshal DataStream as EventStream: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MetricStream'
	if jsonDict["_type"] == "MetricStream" {
		// try to unmarshal JSON data into MetricStream
		err = json.Unmarshal(data, &dst.MetricStream)
		if err == nil {
			return nil // data stored in dst.MetricStream, return on the first match
		} else {
			dst.MetricStream = nil
			return fmt.Errorf("Failed to unmarshal DataStream as MetricStream: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataStream) MarshalJSON() ([]byte, error) {
	if src.EventStream != nil {
		return json.Marshal(&src.EventStream)
	}

	if src.MetricStream != nil {
		return json.Marshal(&src.MetricStream)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataStream) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EventStream != nil {
		return obj.EventStream
	}

	if obj.MetricStream != nil {
		return obj.MetricStream
	}

	// all schemas are nil
	return nil
}

type NullableDataStream struct {
	value *DataStream
	isSet bool
}

func (v NullableDataStream) Get() *DataStream {
	return v.value
}

func (v *NullableDataStream) Set(val *DataStream) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStream) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStream(val *DataStream) *NullableDataStream {
	return &NullableDataStream{value: val, isSet: true}
}

func (v NullableDataStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

