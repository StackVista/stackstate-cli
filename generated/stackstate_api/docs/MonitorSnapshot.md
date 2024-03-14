# MonitorSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstBatch** | Pointer to **bool** |  | [optional] [default to true]
**LastBatch** | Pointer to **bool** |  | [optional] [default to true]
**HealthStates** | [**[]MonitorHealthState**](MonitorHealthState.md) |  | 

## Methods

### NewMonitorSnapshot

`func NewMonitorSnapshot(healthStates []MonitorHealthState, ) *MonitorSnapshot`

NewMonitorSnapshot instantiates a new MonitorSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSnapshotWithDefaults

`func NewMonitorSnapshotWithDefaults() *MonitorSnapshot`

NewMonitorSnapshotWithDefaults instantiates a new MonitorSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstBatch

`func (o *MonitorSnapshot) GetFirstBatch() bool`

GetFirstBatch returns the FirstBatch field if non-nil, zero value otherwise.

### GetFirstBatchOk

`func (o *MonitorSnapshot) GetFirstBatchOk() (*bool, bool)`

GetFirstBatchOk returns a tuple with the FirstBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBatch

`func (o *MonitorSnapshot) SetFirstBatch(v bool)`

SetFirstBatch sets FirstBatch field to given value.

### HasFirstBatch

`func (o *MonitorSnapshot) HasFirstBatch() bool`

HasFirstBatch returns a boolean if a field has been set.

### GetLastBatch

`func (o *MonitorSnapshot) GetLastBatch() bool`

GetLastBatch returns the LastBatch field if non-nil, zero value otherwise.

### GetLastBatchOk

`func (o *MonitorSnapshot) GetLastBatchOk() (*bool, bool)`

GetLastBatchOk returns a tuple with the LastBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatch

`func (o *MonitorSnapshot) SetLastBatch(v bool)`

SetLastBatch sets LastBatch field to given value.

### HasLastBatch

`func (o *MonitorSnapshot) HasLastBatch() bool`

HasLastBatch returns a boolean if a field has been set.

### GetHealthStates

`func (o *MonitorSnapshot) GetHealthStates() []MonitorHealthState`

GetHealthStates returns the HealthStates field if non-nil, zero value otherwise.

### GetHealthStatesOk

`func (o *MonitorSnapshot) GetHealthStatesOk() (*[]MonitorHealthState, bool)`

GetHealthStatesOk returns a tuple with the HealthStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStates

`func (o *MonitorSnapshot) SetHealthStates(v []MonitorHealthState)`

SetHealthStates sets HealthStates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


