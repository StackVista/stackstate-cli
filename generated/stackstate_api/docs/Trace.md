# Trace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | The id of the trace | 
**Spans** | [**[]Span**](Span.md) | All spans for the trace | 
**Resources** | **map[string]map[string]string** | Map of resource identifiers to their corresponding resource attributes (key/value pairs). Keys are UUIDs used in the spans’ &#x60;resourceId&#x60; fields.  | 

## Methods

### NewTrace

`func NewTrace(traceId string, spans []Span, resources map[string]map[string]string, ) *Trace`

NewTrace instantiates a new Trace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceWithDefaults

`func NewTraceWithDefaults() *Trace`

NewTraceWithDefaults instantiates a new Trace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *Trace) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Trace) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Trace) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpans

`func (o *Trace) GetSpans() []Span`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *Trace) GetSpansOk() (*[]Span, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *Trace) SetSpans(v []Span)`

SetSpans sets Spans field to given value.


### GetResources

`func (o *Trace) GetResources() map[string]map[string]string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Trace) GetResourcesOk() (*map[string]map[string]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Trace) SetResources(v map[string]map[string]string)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


