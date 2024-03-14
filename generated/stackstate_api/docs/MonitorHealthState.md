# MonitorHealthState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 
**TopologyIdentifier** | **string** |  | 
**DisplayTimeSeries** | Pointer to [**[]MonitorDisplayTimeSeries**](MonitorDisplayTimeSeries.md) |  | [optional] 
**RemediationHintTemplateData** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitorHealthState

`func NewMonitorHealthState(id string, state HealthStateValue, topologyIdentifier string, ) *MonitorHealthState`

NewMonitorHealthState instantiates a new MonitorHealthState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorHealthStateWithDefaults

`func NewMonitorHealthStateWithDefaults() *MonitorHealthState`

NewMonitorHealthStateWithDefaults instantiates a new MonitorHealthState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitorHealthState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorHealthState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorHealthState) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *MonitorHealthState) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MonitorHealthState) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MonitorHealthState) SetState(v HealthStateValue)`

SetState sets State field to given value.


### GetTopologyIdentifier

`func (o *MonitorHealthState) GetTopologyIdentifier() string`

GetTopologyIdentifier returns the TopologyIdentifier field if non-nil, zero value otherwise.

### GetTopologyIdentifierOk

`func (o *MonitorHealthState) GetTopologyIdentifierOk() (*string, bool)`

GetTopologyIdentifierOk returns a tuple with the TopologyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyIdentifier

`func (o *MonitorHealthState) SetTopologyIdentifier(v string)`

SetTopologyIdentifier sets TopologyIdentifier field to given value.


### GetDisplayTimeSeries

`func (o *MonitorHealthState) GetDisplayTimeSeries() []MonitorDisplayTimeSeries`

GetDisplayTimeSeries returns the DisplayTimeSeries field if non-nil, zero value otherwise.

### GetDisplayTimeSeriesOk

`func (o *MonitorHealthState) GetDisplayTimeSeriesOk() (*[]MonitorDisplayTimeSeries, bool)`

GetDisplayTimeSeriesOk returns a tuple with the DisplayTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimeSeries

`func (o *MonitorHealthState) SetDisplayTimeSeries(v []MonitorDisplayTimeSeries)`

SetDisplayTimeSeries sets DisplayTimeSeries field to given value.

### HasDisplayTimeSeries

`func (o *MonitorHealthState) HasDisplayTimeSeries() bool`

HasDisplayTimeSeries returns a boolean if a field has been set.

### GetRemediationHintTemplateData

`func (o *MonitorHealthState) GetRemediationHintTemplateData() map[string]interface{}`

GetRemediationHintTemplateData returns the RemediationHintTemplateData field if non-nil, zero value otherwise.

### GetRemediationHintTemplateDataOk

`func (o *MonitorHealthState) GetRemediationHintTemplateDataOk() (*map[string]interface{}, bool)`

GetRemediationHintTemplateDataOk returns a tuple with the RemediationHintTemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHintTemplateData

`func (o *MonitorHealthState) SetRemediationHintTemplateData(v map[string]interface{})`

SetRemediationHintTemplateData sets RemediationHintTemplateData field to given value.

### HasRemediationHintTemplateData

`func (o *MonitorHealthState) HasRemediationHintTemplateData() bool`

HasRemediationHintTemplateData returns a boolean if a field has been set.

### GetMessage

`func (o *MonitorHealthState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorHealthState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitorHealthState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MonitorHealthState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *MonitorHealthState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MonitorHealthState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MonitorHealthState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MonitorHealthState) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


