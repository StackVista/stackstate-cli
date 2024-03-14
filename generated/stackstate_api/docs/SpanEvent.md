# SpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**InstantNanoPrecision**](InstantNanoPrecision.md) |  | 
**Name** | **string** |  | 
**Attributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 

## Methods

### NewSpanEvent

`func NewSpanEvent(timestamp InstantNanoPrecision, name string, attributes map[string]string, ) *SpanEvent`

NewSpanEvent instantiates a new SpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanEventWithDefaults

`func NewSpanEventWithDefaults() *SpanEvent`

NewSpanEventWithDefaults instantiates a new SpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SpanEvent) GetTimestamp() InstantNanoPrecision`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SpanEvent) GetTimestampOk() (*InstantNanoPrecision, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SpanEvent) SetTimestamp(v InstantNanoPrecision)`

SetTimestamp sets Timestamp field to given value.


### GetName

`func (o *SpanEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanEvent) SetName(v string)`

SetName sets Name field to given value.


### GetAttributes

`func (o *SpanEvent) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanEvent) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanEvent) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


