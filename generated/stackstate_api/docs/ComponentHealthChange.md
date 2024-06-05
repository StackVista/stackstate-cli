# ComponentHealthChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int64** |  | 
**NewHealth** | [**HealthStateValue**](HealthStateValue.md) |  | 

## Methods

### NewComponentHealthChange

`func NewComponentHealthChange(timestamp int64, newHealth HealthStateValue, ) *ComponentHealthChange`

NewComponentHealthChange instantiates a new ComponentHealthChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentHealthChangeWithDefaults

`func NewComponentHealthChangeWithDefaults() *ComponentHealthChange`

NewComponentHealthChangeWithDefaults instantiates a new ComponentHealthChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ComponentHealthChange) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ComponentHealthChange) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ComponentHealthChange) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNewHealth

`func (o *ComponentHealthChange) GetNewHealth() HealthStateValue`

GetNewHealth returns the NewHealth field if non-nil, zero value otherwise.

### GetNewHealthOk

`func (o *ComponentHealthChange) GetNewHealthOk() (*HealthStateValue, bool)`

GetNewHealthOk returns a tuple with the NewHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewHealth

`func (o *ComponentHealthChange) SetNewHealth(v HealthStateValue)`

SetNewHealth sets NewHealth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


