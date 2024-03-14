# SpanLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**TraceState** | Pointer to **string** |  | [optional] 
**Attributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 

## Methods

### NewSpanLink

`func NewSpanLink(traceId string, spanId string, attributes map[string]string, ) *SpanLink`

NewSpanLink instantiates a new SpanLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanLinkWithDefaults

`func NewSpanLinkWithDefaults() *SpanLink`

NewSpanLinkWithDefaults instantiates a new SpanLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *SpanLink) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SpanLink) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SpanLink) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *SpanLink) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanLink) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanLink) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetTraceState

`func (o *SpanLink) GetTraceState() string`

GetTraceState returns the TraceState field if non-nil, zero value otherwise.

### GetTraceStateOk

`func (o *SpanLink) GetTraceStateOk() (*string, bool)`

GetTraceStateOk returns a tuple with the TraceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceState

`func (o *SpanLink) SetTraceState(v string)`

SetTraceState sets TraceState field to given value.

### HasTraceState

`func (o *SpanLink) HasTraceState() bool`

HasTraceState returns a boolean if a field has been set.

### GetAttributes

`func (o *SpanLink) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanLink) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanLink) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


