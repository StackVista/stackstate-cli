# MonitorCheckStatusHealthHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatusId** | **int64** |  | 
**StartTime** | **int32** |  | 
**EndTime** | **int32** |  | 
**HealthStateChanges** | [**[]MonitorCheckStatusHealthChange**](MonitorCheckStatusHealthChange.md) | List of health state changes ordered from most recent to oldest. | 

## Methods

### NewMonitorCheckStatusHealthHistory

`func NewMonitorCheckStatusHealthHistory(checkStatusId int64, startTime int32, endTime int32, healthStateChanges []MonitorCheckStatusHealthChange, ) *MonitorCheckStatusHealthHistory`

NewMonitorCheckStatusHealthHistory instantiates a new MonitorCheckStatusHealthHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusHealthHistoryWithDefaults

`func NewMonitorCheckStatusHealthHistoryWithDefaults() *MonitorCheckStatusHealthHistory`

NewMonitorCheckStatusHealthHistoryWithDefaults instantiates a new MonitorCheckStatusHealthHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatusId

`func (o *MonitorCheckStatusHealthHistory) GetCheckStatusId() int64`

GetCheckStatusId returns the CheckStatusId field if non-nil, zero value otherwise.

### GetCheckStatusIdOk

`func (o *MonitorCheckStatusHealthHistory) GetCheckStatusIdOk() (*int64, bool)`

GetCheckStatusIdOk returns a tuple with the CheckStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatusId

`func (o *MonitorCheckStatusHealthHistory) SetCheckStatusId(v int64)`

SetCheckStatusId sets CheckStatusId field to given value.


### GetStartTime

`func (o *MonitorCheckStatusHealthHistory) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MonitorCheckStatusHealthHistory) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MonitorCheckStatusHealthHistory) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *MonitorCheckStatusHealthHistory) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MonitorCheckStatusHealthHistory) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MonitorCheckStatusHealthHistory) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetHealthStateChanges

`func (o *MonitorCheckStatusHealthHistory) GetHealthStateChanges() []MonitorCheckStatusHealthChange`

GetHealthStateChanges returns the HealthStateChanges field if non-nil, zero value otherwise.

### GetHealthStateChangesOk

`func (o *MonitorCheckStatusHealthHistory) GetHealthStateChangesOk() (*[]MonitorCheckStatusHealthChange, bool)`

GetHealthStateChangesOk returns a tuple with the HealthStateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStateChanges

`func (o *MonitorCheckStatusHealthHistory) SetHealthStateChanges(v []MonitorCheckStatusHealthChange)`

SetHealthStateChanges sets HealthStateChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


