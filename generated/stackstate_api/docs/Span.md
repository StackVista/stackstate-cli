# Span

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | [**InstantNanoPrecision**](InstantNanoPrecision.md) |  | 
**EndTime** | [**InstantNanoPrecision**](InstantNanoPrecision.md) |  | 
**DurationNanos** | **int64** | Duration of the span in nanoseconds | 
**TraceId** | **string** | The unique identifier for the trace, all spans of the same trace share the same trace_id | 
**SpanId** | **string** | The unique identifier of the span within the trace | 
**ParentSpanId** | Pointer to **string** | The id of the parent span of this span. Empty if this is the root span | [optional] 
**TraceState** | Pointer to **string** | Can contain vendor specific trace identification information https://www.w3.org/TR/trace-context/#tracestate-header | [optional] 
**SpanName** | **string** | A description of the span&#39;s operation. For example, the name can be a qualified method name or a file name and a line number where the operation is called | 
**ServiceName** | **string** | Logical name of the service for the span | 
**SpanKind** | [**SpanKind**](SpanKind.md) |  | 
**SpanParentType** | [**SpanParentType**](SpanParentType.md) |  | 
**ResourceAttributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 
**SpanAttributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 
**StatusCode** | [**StatusCode**](StatusCode.md) |  | 
**StatusMessage** | Pointer to **string** | Human readable message for the status | [optional] 
**ScopeName** | Pointer to **string** | The name of the instrumentation scope for the span https://opentelemetry.io/docs/specs/otel/glossary/#instrumentation-scope | [optional] 
**ScopeVersion** | Pointer to **string** | The version for the instrumentation scope for the span https://opentelemetry.io/docs/specs/otel/glossary/#instrumentation-scope | [optional] 
**Events** | [**[]SpanEvent**](SpanEvent.md) | Time-stamped annotations on the span providing extra application context | 
**Links** | [**[]SpanLink**](SpanLink.md) | Links to related spans in the same or in other traces | 

## Methods

### NewSpan

`func NewSpan(startTime InstantNanoPrecision, endTime InstantNanoPrecision, durationNanos int64, traceId string, spanId string, spanName string, serviceName string, spanKind SpanKind, spanParentType SpanParentType, resourceAttributes map[string]string, spanAttributes map[string]string, statusCode StatusCode, events []SpanEvent, links []SpanLink, ) *Span`

NewSpan instantiates a new Span object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanWithDefaults

`func NewSpanWithDefaults() *Span`

NewSpanWithDefaults instantiates a new Span object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *Span) GetStartTime() InstantNanoPrecision`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Span) GetStartTimeOk() (*InstantNanoPrecision, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Span) SetStartTime(v InstantNanoPrecision)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *Span) GetEndTime() InstantNanoPrecision`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Span) GetEndTimeOk() (*InstantNanoPrecision, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Span) SetEndTime(v InstantNanoPrecision)`

SetEndTime sets EndTime field to given value.


### GetDurationNanos

`func (o *Span) GetDurationNanos() int64`

GetDurationNanos returns the DurationNanos field if non-nil, zero value otherwise.

### GetDurationNanosOk

`func (o *Span) GetDurationNanosOk() (*int64, bool)`

GetDurationNanosOk returns a tuple with the DurationNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNanos

`func (o *Span) SetDurationNanos(v int64)`

SetDurationNanos sets DurationNanos field to given value.


### GetTraceId

`func (o *Span) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Span) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Span) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *Span) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *Span) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *Span) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetParentSpanId

`func (o *Span) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *Span) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *Span) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *Span) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetTraceState

`func (o *Span) GetTraceState() string`

GetTraceState returns the TraceState field if non-nil, zero value otherwise.

### GetTraceStateOk

`func (o *Span) GetTraceStateOk() (*string, bool)`

GetTraceStateOk returns a tuple with the TraceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceState

`func (o *Span) SetTraceState(v string)`

SetTraceState sets TraceState field to given value.

### HasTraceState

`func (o *Span) HasTraceState() bool`

HasTraceState returns a boolean if a field has been set.

### GetSpanName

`func (o *Span) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *Span) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *Span) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetServiceName

`func (o *Span) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Span) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Span) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSpanKind

`func (o *Span) GetSpanKind() SpanKind`

GetSpanKind returns the SpanKind field if non-nil, zero value otherwise.

### GetSpanKindOk

`func (o *Span) GetSpanKindOk() (*SpanKind, bool)`

GetSpanKindOk returns a tuple with the SpanKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanKind

`func (o *Span) SetSpanKind(v SpanKind)`

SetSpanKind sets SpanKind field to given value.


### GetSpanParentType

`func (o *Span) GetSpanParentType() SpanParentType`

GetSpanParentType returns the SpanParentType field if non-nil, zero value otherwise.

### GetSpanParentTypeOk

`func (o *Span) GetSpanParentTypeOk() (*SpanParentType, bool)`

GetSpanParentTypeOk returns a tuple with the SpanParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanParentType

`func (o *Span) SetSpanParentType(v SpanParentType)`

SetSpanParentType sets SpanParentType field to given value.


### GetResourceAttributes

`func (o *Span) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *Span) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *Span) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetSpanAttributes

`func (o *Span) GetSpanAttributes() map[string]string`

GetSpanAttributes returns the SpanAttributes field if non-nil, zero value otherwise.

### GetSpanAttributesOk

`func (o *Span) GetSpanAttributesOk() (*map[string]string, bool)`

GetSpanAttributesOk returns a tuple with the SpanAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributes

`func (o *Span) SetSpanAttributes(v map[string]string)`

SetSpanAttributes sets SpanAttributes field to given value.


### GetStatusCode

`func (o *Span) GetStatusCode() StatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Span) GetStatusCodeOk() (*StatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Span) SetStatusCode(v StatusCode)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *Span) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Span) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Span) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Span) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetScopeName

`func (o *Span) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *Span) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *Span) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *Span) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeVersion

`func (o *Span) GetScopeVersion() string`

GetScopeVersion returns the ScopeVersion field if non-nil, zero value otherwise.

### GetScopeVersionOk

`func (o *Span) GetScopeVersionOk() (*string, bool)`

GetScopeVersionOk returns a tuple with the ScopeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeVersion

`func (o *Span) SetScopeVersion(v string)`

SetScopeVersion sets ScopeVersion field to given value.

### HasScopeVersion

`func (o *Span) HasScopeVersion() bool`

HasScopeVersion returns a boolean if a field has been set.

### GetEvents

`func (o *Span) GetEvents() []SpanEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Span) GetEventsOk() (*[]SpanEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Span) SetEvents(v []SpanEvent)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *Span) GetLinks() []SpanLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Span) GetLinksOk() (*[]SpanLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Span) SetLinks(v []SpanLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


