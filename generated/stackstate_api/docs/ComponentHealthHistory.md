# ComponentHealthHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**StartTime** | **int32** |  | 
**EndTime** | **int32** |  | 
**HealthStateChanges** | [**[]ComponentHealthChange**](ComponentHealthChange.md) | List of health state changes ordered from most recent to oldest. | 

## Methods

### NewComponentHealthHistory

`func NewComponentHealthHistory(id int64, startTime int32, endTime int32, healthStateChanges []ComponentHealthChange, ) *ComponentHealthHistory`

NewComponentHealthHistory instantiates a new ComponentHealthHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentHealthHistoryWithDefaults

`func NewComponentHealthHistoryWithDefaults() *ComponentHealthHistory`

NewComponentHealthHistoryWithDefaults instantiates a new ComponentHealthHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentHealthHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentHealthHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentHealthHistory) SetId(v int64)`

SetId sets Id field to given value.


### GetStartTime

`func (o *ComponentHealthHistory) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ComponentHealthHistory) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ComponentHealthHistory) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ComponentHealthHistory) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ComponentHealthHistory) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ComponentHealthHistory) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetHealthStateChanges

`func (o *ComponentHealthHistory) GetHealthStateChanges() []ComponentHealthChange`

GetHealthStateChanges returns the HealthStateChanges field if non-nil, zero value otherwise.

### GetHealthStateChangesOk

`func (o *ComponentHealthHistory) GetHealthStateChangesOk() (*[]ComponentHealthChange, bool)`

GetHealthStateChangesOk returns a tuple with the HealthStateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStateChanges

`func (o *ComponentHealthHistory) SetHealthStateChanges(v []ComponentHealthChange)`

SetHealthStateChanges sets HealthStateChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


