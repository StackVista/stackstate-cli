# TimelineSummaryHealthChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **int64** |  | 
**NewHealth** | [**HealthStateValue**](HealthStateValue.md) |  | 

## Methods

### NewTimelineSummaryHealthChange

`func NewTimelineSummaryHealthChange(type_ string, timestamp int64, newHealth HealthStateValue, ) *TimelineSummaryHealthChange`

NewTimelineSummaryHealthChange instantiates a new TimelineSummaryHealthChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSummaryHealthChangeWithDefaults

`func NewTimelineSummaryHealthChangeWithDefaults() *TimelineSummaryHealthChange`

NewTimelineSummaryHealthChangeWithDefaults instantiates a new TimelineSummaryHealthChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimelineSummaryHealthChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineSummaryHealthChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineSummaryHealthChange) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *TimelineSummaryHealthChange) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimelineSummaryHealthChange) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimelineSummaryHealthChange) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNewHealth

`func (o *TimelineSummaryHealthChange) GetNewHealth() HealthStateValue`

GetNewHealth returns the NewHealth field if non-nil, zero value otherwise.

### GetNewHealthOk

`func (o *TimelineSummaryHealthChange) GetNewHealthOk() (*HealthStateValue, bool)`

GetNewHealthOk returns a tuple with the NewHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewHealth

`func (o *TimelineSummaryHealthChange) SetNewHealth(v HealthStateValue)`

SetNewHealth sets NewHealth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


