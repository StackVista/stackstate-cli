# OtelInputMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A Cel expression that must return a boolean | [optional] 
**Action** | Pointer to [**OtelInputConditionAction**](OtelInputConditionAction.md) |  | [optional] 
**Datapoint** | Pointer to [**OtelInputDatapoint**](OtelInputDatapoint.md) |  | [optional] 

## Methods

### NewOtelInputMetric

`func NewOtelInputMetric() *OtelInputMetric`

NewOtelInputMetric instantiates a new OtelInputMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelInputMetricWithDefaults

`func NewOtelInputMetricWithDefaults() *OtelInputMetric`

NewOtelInputMetricWithDefaults instantiates a new OtelInputMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *OtelInputMetric) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OtelInputMetric) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OtelInputMetric) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *OtelInputMetric) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAction

`func (o *OtelInputMetric) GetAction() OtelInputConditionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OtelInputMetric) GetActionOk() (*OtelInputConditionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OtelInputMetric) SetAction(v OtelInputConditionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *OtelInputMetric) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDatapoint

`func (o *OtelInputMetric) GetDatapoint() OtelInputDatapoint`

GetDatapoint returns the Datapoint field if non-nil, zero value otherwise.

### GetDatapointOk

`func (o *OtelInputMetric) GetDatapointOk() (*OtelInputDatapoint, bool)`

GetDatapointOk returns a tuple with the Datapoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoint

`func (o *OtelInputMetric) SetDatapoint(v OtelInputDatapoint)`

SetDatapoint sets Datapoint field to given value.

### HasDatapoint

`func (o *OtelInputMetric) HasDatapoint() bool`

HasDatapoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


