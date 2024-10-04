# TraceIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | The unique identifier for the trace, all spans of the same trace share the same trace_id | 
**SpanId** | **string** | The unique identifier of the span within the trace | 

## Methods

### NewTraceIdentifier

`func NewTraceIdentifier(traceId string, spanId string, ) *TraceIdentifier`

NewTraceIdentifier instantiates a new TraceIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceIdentifierWithDefaults

`func NewTraceIdentifierWithDefaults() *TraceIdentifier`

NewTraceIdentifierWithDefaults instantiates a new TraceIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *TraceIdentifier) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TraceIdentifier) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TraceIdentifier) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *TraceIdentifier) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *TraceIdentifier) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *TraceIdentifier) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


