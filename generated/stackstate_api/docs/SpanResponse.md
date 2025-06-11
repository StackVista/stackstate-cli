# SpanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Span** | [**Span**](Span.md) |  | 
**ResourceAttributes** | **map[string]string** | Set of key/value pairs providing extra contextual information. Keys are unique. | 

## Methods

### NewSpanResponse

`func NewSpanResponse(span Span, resourceAttributes map[string]string, ) *SpanResponse`

NewSpanResponse instantiates a new SpanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanResponseWithDefaults

`func NewSpanResponseWithDefaults() *SpanResponse`

NewSpanResponseWithDefaults instantiates a new SpanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpan

`func (o *SpanResponse) GetSpan() Span`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *SpanResponse) GetSpanOk() (*Span, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *SpanResponse) SetSpan(v Span)`

SetSpan sets Span field to given value.


### GetResourceAttributes

`func (o *SpanResponse) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *SpanResponse) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *SpanResponse) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


