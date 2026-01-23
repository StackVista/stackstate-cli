# OtelInputResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A Cel expression that must return a boolean | [optional] 
**Action** | Pointer to [**OtelInputConditionAction**](OtelInputConditionAction.md) |  | [optional] 
**Scope** | Pointer to [**OtelInputScope**](OtelInputScope.md) |  | [optional] 

## Methods

### NewOtelInputResource

`func NewOtelInputResource() *OtelInputResource`

NewOtelInputResource instantiates a new OtelInputResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelInputResourceWithDefaults

`func NewOtelInputResourceWithDefaults() *OtelInputResource`

NewOtelInputResourceWithDefaults instantiates a new OtelInputResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *OtelInputResource) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OtelInputResource) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OtelInputResource) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *OtelInputResource) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAction

`func (o *OtelInputResource) GetAction() OtelInputConditionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OtelInputResource) GetActionOk() (*OtelInputConditionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OtelInputResource) SetAction(v OtelInputConditionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *OtelInputResource) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetScope

`func (o *OtelInputResource) GetScope() OtelInputScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OtelInputResource) GetScopeOk() (*OtelInputScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OtelInputResource) SetScope(v OtelInputScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OtelInputResource) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


