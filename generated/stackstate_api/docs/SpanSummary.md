# SpanSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | [**InstantNanoPrecision**](InstantNanoPrecision.md) |  | 
**EndTime** | [**InstantNanoPrecision**](InstantNanoPrecision.md) |  | 
**DurationNanos** | **int64** | Duration of the span in nanoseconds | 
**TraceId** | **string** | The unique identifier for the trace, all spans of the same trace share the same trace_id | 
**SpanId** | **string** | The unique identifier of the span within the trace | 
**ParentSpanId** | Pointer to **string** | The id of the parent span of this span. Empty if this is the root span | [optional] 
**SpanName** | **string** | A description of the span&#39;s operation. For example, the name can be a qualified method name or a file name and a line number where the operation is called | 
**ServiceName** | **string** | Logical name of the service for the span | 
**StatusCode** | [**StatusCode**](StatusCode.md) |  | 

## Methods

### NewSpanSummary

`func NewSpanSummary(startTime InstantNanoPrecision, endTime InstantNanoPrecision, durationNanos int64, traceId string, spanId string, spanName string, serviceName string, statusCode StatusCode, ) *SpanSummary`

NewSpanSummary instantiates a new SpanSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanSummaryWithDefaults

`func NewSpanSummaryWithDefaults() *SpanSummary`

NewSpanSummaryWithDefaults instantiates a new SpanSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SpanSummary) GetStartTime() InstantNanoPrecision`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SpanSummary) GetStartTimeOk() (*InstantNanoPrecision, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SpanSummary) SetStartTime(v InstantNanoPrecision)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SpanSummary) GetEndTime() InstantNanoPrecision`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SpanSummary) GetEndTimeOk() (*InstantNanoPrecision, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SpanSummary) SetEndTime(v InstantNanoPrecision)`

SetEndTime sets EndTime field to given value.


### GetDurationNanos

`func (o *SpanSummary) GetDurationNanos() int64`

GetDurationNanos returns the DurationNanos field if non-nil, zero value otherwise.

### GetDurationNanosOk

`func (o *SpanSummary) GetDurationNanosOk() (*int64, bool)`

GetDurationNanosOk returns a tuple with the DurationNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNanos

`func (o *SpanSummary) SetDurationNanos(v int64)`

SetDurationNanos sets DurationNanos field to given value.


### GetTraceId

`func (o *SpanSummary) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SpanSummary) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SpanSummary) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *SpanSummary) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanSummary) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanSummary) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetParentSpanId

`func (o *SpanSummary) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *SpanSummary) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *SpanSummary) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *SpanSummary) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetSpanName

`func (o *SpanSummary) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *SpanSummary) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *SpanSummary) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetServiceName

`func (o *SpanSummary) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SpanSummary) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SpanSummary) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatusCode

`func (o *SpanSummary) GetStatusCode() StatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SpanSummary) GetStatusCodeOk() (*StatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SpanSummary) SetStatusCode(v StatusCode)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


