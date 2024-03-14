# InstantNanoPrecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Date/time representation in milliseconds since epoch (1970-01-01 00:00:00) | 
**OffsetNanos** | **int32** | Offset in nanoseconds (relative to the timestamp). Especially useful when comparing start and/or end times of spans, for example when rendering a trace chart. | 

## Methods

### NewInstantNanoPrecision

`func NewInstantNanoPrecision(timestamp int32, offsetNanos int32, ) *InstantNanoPrecision`

NewInstantNanoPrecision instantiates a new InstantNanoPrecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantNanoPrecisionWithDefaults

`func NewInstantNanoPrecisionWithDefaults() *InstantNanoPrecision`

NewInstantNanoPrecisionWithDefaults instantiates a new InstantNanoPrecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InstantNanoPrecision) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InstantNanoPrecision) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InstantNanoPrecision) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetOffsetNanos

`func (o *InstantNanoPrecision) GetOffsetNanos() int32`

GetOffsetNanos returns the OffsetNanos field if non-nil, zero value otherwise.

### GetOffsetNanosOk

`func (o *InstantNanoPrecision) GetOffsetNanosOk() (*int32, bool)`

GetOffsetNanosOk returns a tuple with the OffsetNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetNanos

`func (o *InstantNanoPrecision) SetOffsetNanos(v int32)`

SetOffsetNanos sets OffsetNanos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


