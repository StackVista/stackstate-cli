# OtelInputScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A Cel expression that must return a boolean | [optional] 
**Action** | Pointer to [**OtelInputConditionAction**](OtelInputConditionAction.md) |  | [optional] 
**Metric** | Pointer to [**OtelInputMetric**](OtelInputMetric.md) |  | [optional] 
**Span** | Pointer to [**OtelInputSpan**](OtelInputSpan.md) |  | [optional] 

## Methods

### NewOtelInputScope

`func NewOtelInputScope() *OtelInputScope`

NewOtelInputScope instantiates a new OtelInputScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelInputScopeWithDefaults

`func NewOtelInputScopeWithDefaults() *OtelInputScope`

NewOtelInputScopeWithDefaults instantiates a new OtelInputScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *OtelInputScope) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OtelInputScope) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OtelInputScope) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *OtelInputScope) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAction

`func (o *OtelInputScope) GetAction() OtelInputConditionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OtelInputScope) GetActionOk() (*OtelInputConditionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OtelInputScope) SetAction(v OtelInputConditionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *OtelInputScope) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMetric

`func (o *OtelInputScope) GetMetric() OtelInputMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *OtelInputScope) GetMetricOk() (*OtelInputMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *OtelInputScope) SetMetric(v OtelInputMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *OtelInputScope) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetSpan

`func (o *OtelInputScope) GetSpan() OtelInputSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *OtelInputScope) GetSpanOk() (*OtelInputSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *OtelInputScope) SetSpan(v OtelInputSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *OtelInputScope) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


