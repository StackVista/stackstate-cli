# SpanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | Pointer to **[]string** | Filter spans by service name | [optional] 
**SpanName** | Pointer to **[]string** | Filter spans by name | [optional] 
**Attributes** | Pointer to **map[string][]string** | Filter spans by 1 or more attributes | [optional] 
**SpanKind** | Pointer to [**[]SpanKind**](SpanKind.md) | Filter span by kind | [optional] 
**SpanParentType** | Pointer to [**[]SpanParentType**](SpanParentType.md) | Filter span by parent type | [optional] 
**DurationFromNanos** | Pointer to **int64** | Filter spans by duration &gt;&#x3D; value, in nanoseconds | [optional] 
**DurationToNanos** | Pointer to **int64** | Filter spans by duration &lt; value, in nanoseconds | [optional] 
**StatusCode** | Pointer to [**[]StatusCode**](StatusCode.md) | Filter spans by the StatusCode | [optional] 
**TraceId** | Pointer to **[]string** | Filter spans by trace id, use only this filter to get a complete trace | [optional] 
**SpanId** | Pointer to **[]string** | Filter spans by span id, use only this filter to get a single span | [optional] 
**ScopeName** | Pointer to **[]string** | Filter spans by scope | [optional] 
**ScopeVersion** | Pointer to **[]string** | Filter spans by scope version | [optional] 

## Methods

### NewSpanFilter

`func NewSpanFilter() *SpanFilter`

NewSpanFilter instantiates a new SpanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanFilterWithDefaults

`func NewSpanFilterWithDefaults() *SpanFilter`

NewSpanFilterWithDefaults instantiates a new SpanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *SpanFilter) GetServiceName() []string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SpanFilter) GetServiceNameOk() (*[]string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SpanFilter) SetServiceName(v []string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SpanFilter) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSpanName

`func (o *SpanFilter) GetSpanName() []string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *SpanFilter) GetSpanNameOk() (*[]string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *SpanFilter) SetSpanName(v []string)`

SetSpanName sets SpanName field to given value.

### HasSpanName

`func (o *SpanFilter) HasSpanName() bool`

HasSpanName returns a boolean if a field has been set.

### GetAttributes

`func (o *SpanFilter) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanFilter) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanFilter) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SpanFilter) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSpanKind

`func (o *SpanFilter) GetSpanKind() []SpanKind`

GetSpanKind returns the SpanKind field if non-nil, zero value otherwise.

### GetSpanKindOk

`func (o *SpanFilter) GetSpanKindOk() (*[]SpanKind, bool)`

GetSpanKindOk returns a tuple with the SpanKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanKind

`func (o *SpanFilter) SetSpanKind(v []SpanKind)`

SetSpanKind sets SpanKind field to given value.

### HasSpanKind

`func (o *SpanFilter) HasSpanKind() bool`

HasSpanKind returns a boolean if a field has been set.

### GetSpanParentType

`func (o *SpanFilter) GetSpanParentType() []SpanParentType`

GetSpanParentType returns the SpanParentType field if non-nil, zero value otherwise.

### GetSpanParentTypeOk

`func (o *SpanFilter) GetSpanParentTypeOk() (*[]SpanParentType, bool)`

GetSpanParentTypeOk returns a tuple with the SpanParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanParentType

`func (o *SpanFilter) SetSpanParentType(v []SpanParentType)`

SetSpanParentType sets SpanParentType field to given value.

### HasSpanParentType

`func (o *SpanFilter) HasSpanParentType() bool`

HasSpanParentType returns a boolean if a field has been set.

### GetDurationFromNanos

`func (o *SpanFilter) GetDurationFromNanos() int64`

GetDurationFromNanos returns the DurationFromNanos field if non-nil, zero value otherwise.

### GetDurationFromNanosOk

`func (o *SpanFilter) GetDurationFromNanosOk() (*int64, bool)`

GetDurationFromNanosOk returns a tuple with the DurationFromNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFromNanos

`func (o *SpanFilter) SetDurationFromNanos(v int64)`

SetDurationFromNanos sets DurationFromNanos field to given value.

### HasDurationFromNanos

`func (o *SpanFilter) HasDurationFromNanos() bool`

HasDurationFromNanos returns a boolean if a field has been set.

### GetDurationToNanos

`func (o *SpanFilter) GetDurationToNanos() int64`

GetDurationToNanos returns the DurationToNanos field if non-nil, zero value otherwise.

### GetDurationToNanosOk

`func (o *SpanFilter) GetDurationToNanosOk() (*int64, bool)`

GetDurationToNanosOk returns a tuple with the DurationToNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToNanos

`func (o *SpanFilter) SetDurationToNanos(v int64)`

SetDurationToNanos sets DurationToNanos field to given value.

### HasDurationToNanos

`func (o *SpanFilter) HasDurationToNanos() bool`

HasDurationToNanos returns a boolean if a field has been set.

### GetStatusCode

`func (o *SpanFilter) GetStatusCode() []StatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SpanFilter) GetStatusCodeOk() (*[]StatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SpanFilter) SetStatusCode(v []StatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SpanFilter) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTraceId

`func (o *SpanFilter) GetTraceId() []string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SpanFilter) GetTraceIdOk() (*[]string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SpanFilter) SetTraceId(v []string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SpanFilter) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetSpanId

`func (o *SpanFilter) GetSpanId() []string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanFilter) GetSpanIdOk() (*[]string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanFilter) SetSpanId(v []string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *SpanFilter) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetScopeName

`func (o *SpanFilter) GetScopeName() []string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *SpanFilter) GetScopeNameOk() (*[]string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *SpanFilter) SetScopeName(v []string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *SpanFilter) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeVersion

`func (o *SpanFilter) GetScopeVersion() []string`

GetScopeVersion returns the ScopeVersion field if non-nil, zero value otherwise.

### GetScopeVersionOk

`func (o *SpanFilter) GetScopeVersionOk() (*[]string, bool)`

GetScopeVersionOk returns a tuple with the ScopeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeVersion

`func (o *SpanFilter) SetScopeVersion(v []string)`

SetScopeVersion sets ScopeVersion field to given value.

### HasScopeVersion

`func (o *SpanFilter) HasScopeVersion() bool`

HasScopeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


