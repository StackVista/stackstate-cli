# OtelInputDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A Cel expression that must return a boolean | [optional] 
**Action** | Pointer to [**OtelInputConditionAction**](OtelInputConditionAction.md) |  | [optional] 

## Methods

### NewOtelInputDatapoint

`func NewOtelInputDatapoint() *OtelInputDatapoint`

NewOtelInputDatapoint instantiates a new OtelInputDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelInputDatapointWithDefaults

`func NewOtelInputDatapointWithDefaults() *OtelInputDatapoint`

NewOtelInputDatapointWithDefaults instantiates a new OtelInputDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *OtelInputDatapoint) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OtelInputDatapoint) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OtelInputDatapoint) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *OtelInputDatapoint) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAction

`func (o *OtelInputDatapoint) GetAction() OtelInputConditionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OtelInputDatapoint) GetActionOk() (*OtelInputConditionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OtelInputDatapoint) SetAction(v OtelInputConditionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *OtelInputDatapoint) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


