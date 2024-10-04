# TraceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanFilter** | [**SpanFilter**](SpanFilter.md) |  | 
**TraceAttributes** | **map[string][]string** | Filter traces by 1 or more attributes | 

## Methods

### NewTraceFilter

`func NewTraceFilter(spanFilter SpanFilter, traceAttributes map[string][]string, ) *TraceFilter`

NewTraceFilter instantiates a new TraceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceFilterWithDefaults

`func NewTraceFilterWithDefaults() *TraceFilter`

NewTraceFilterWithDefaults instantiates a new TraceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanFilter

`func (o *TraceFilter) GetSpanFilter() SpanFilter`

GetSpanFilter returns the SpanFilter field if non-nil, zero value otherwise.

### GetSpanFilterOk

`func (o *TraceFilter) GetSpanFilterOk() (*SpanFilter, bool)`

GetSpanFilterOk returns a tuple with the SpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanFilter

`func (o *TraceFilter) SetSpanFilter(v SpanFilter)`

SetSpanFilter sets SpanFilter field to given value.


### GetTraceAttributes

`func (o *TraceFilter) GetTraceAttributes() map[string][]string`

GetTraceAttributes returns the TraceAttributes field if non-nil, zero value otherwise.

### GetTraceAttributesOk

`func (o *TraceFilter) GetTraceAttributesOk() (*map[string][]string, bool)`

GetTraceAttributesOk returns a tuple with the TraceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceAttributes

`func (o *TraceFilter) SetTraceAttributes(v map[string][]string)`

SetTraceAttributes sets TraceAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


