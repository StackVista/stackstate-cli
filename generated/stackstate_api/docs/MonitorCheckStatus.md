# MonitorCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Identifier** | **string** |  | 
**Message** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Health** | [**HealthStateValue**](HealthStateValue.md) |  | 
**TriggeredTimestamp** | **int32** |  | 
**Metrics** | [**[]MonitorCheckStatusMetric**](MonitorCheckStatusMetric.md) |  | 
**Component** | [**MonitorCheckStatusComponent**](MonitorCheckStatusComponent.md) |  | 
**MonitorId** | [**MonitorReferenceId**](MonitorReferenceId.md) |  | 
**MonitorName** | **string** |  | 
**MonitorDescription** | Pointer to **string** |  | [optional] 
**TroubleshootingSteps** | Pointer to **string** |  | [optional] 
**TopologyTime** | **int32** |  | 
**Dummy** | **bool** |  | 

## Methods

### NewMonitorCheckStatus

`func NewMonitorCheckStatus(id int64, identifier string, message string, health HealthStateValue, triggeredTimestamp int32, metrics []MonitorCheckStatusMetric, component MonitorCheckStatusComponent, monitorId MonitorReferenceId, monitorName string, topologyTime int32, dummy bool, ) *MonitorCheckStatus`

NewMonitorCheckStatus instantiates a new MonitorCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusWithDefaults

`func NewMonitorCheckStatusWithDefaults() *MonitorCheckStatus`

NewMonitorCheckStatusWithDefaults instantiates a new MonitorCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitorCheckStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorCheckStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorCheckStatus) SetId(v int64)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *MonitorCheckStatus) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MonitorCheckStatus) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MonitorCheckStatus) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMessage

`func (o *MonitorCheckStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorCheckStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitorCheckStatus) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *MonitorCheckStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MonitorCheckStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MonitorCheckStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MonitorCheckStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetHealth

`func (o *MonitorCheckStatus) GetHealth() HealthStateValue`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *MonitorCheckStatus) GetHealthOk() (*HealthStateValue, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *MonitorCheckStatus) SetHealth(v HealthStateValue)`

SetHealth sets Health field to given value.


### GetTriggeredTimestamp

`func (o *MonitorCheckStatus) GetTriggeredTimestamp() int32`

GetTriggeredTimestamp returns the TriggeredTimestamp field if non-nil, zero value otherwise.

### GetTriggeredTimestampOk

`func (o *MonitorCheckStatus) GetTriggeredTimestampOk() (*int32, bool)`

GetTriggeredTimestampOk returns a tuple with the TriggeredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTimestamp

`func (o *MonitorCheckStatus) SetTriggeredTimestamp(v int32)`

SetTriggeredTimestamp sets TriggeredTimestamp field to given value.


### GetMetrics

`func (o *MonitorCheckStatus) GetMetrics() []MonitorCheckStatusMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MonitorCheckStatus) GetMetricsOk() (*[]MonitorCheckStatusMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MonitorCheckStatus) SetMetrics(v []MonitorCheckStatusMetric)`

SetMetrics sets Metrics field to given value.


### GetComponent

`func (o *MonitorCheckStatus) GetComponent() MonitorCheckStatusComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MonitorCheckStatus) GetComponentOk() (*MonitorCheckStatusComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MonitorCheckStatus) SetComponent(v MonitorCheckStatusComponent)`

SetComponent sets Component field to given value.


### GetMonitorId

`func (o *MonitorCheckStatus) GetMonitorId() MonitorReferenceId`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorCheckStatus) GetMonitorIdOk() (*MonitorReferenceId, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorCheckStatus) SetMonitorId(v MonitorReferenceId)`

SetMonitorId sets MonitorId field to given value.


### GetMonitorName

`func (o *MonitorCheckStatus) GetMonitorName() string`

GetMonitorName returns the MonitorName field if non-nil, zero value otherwise.

### GetMonitorNameOk

`func (o *MonitorCheckStatus) GetMonitorNameOk() (*string, bool)`

GetMonitorNameOk returns a tuple with the MonitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorName

`func (o *MonitorCheckStatus) SetMonitorName(v string)`

SetMonitorName sets MonitorName field to given value.


### GetMonitorDescription

`func (o *MonitorCheckStatus) GetMonitorDescription() string`

GetMonitorDescription returns the MonitorDescription field if non-nil, zero value otherwise.

### GetMonitorDescriptionOk

`func (o *MonitorCheckStatus) GetMonitorDescriptionOk() (*string, bool)`

GetMonitorDescriptionOk returns a tuple with the MonitorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDescription

`func (o *MonitorCheckStatus) SetMonitorDescription(v string)`

SetMonitorDescription sets MonitorDescription field to given value.

### HasMonitorDescription

`func (o *MonitorCheckStatus) HasMonitorDescription() bool`

HasMonitorDescription returns a boolean if a field has been set.

### GetTroubleshootingSteps

`func (o *MonitorCheckStatus) GetTroubleshootingSteps() string`

GetTroubleshootingSteps returns the TroubleshootingSteps field if non-nil, zero value otherwise.

### GetTroubleshootingStepsOk

`func (o *MonitorCheckStatus) GetTroubleshootingStepsOk() (*string, bool)`

GetTroubleshootingStepsOk returns a tuple with the TroubleshootingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingSteps

`func (o *MonitorCheckStatus) SetTroubleshootingSteps(v string)`

SetTroubleshootingSteps sets TroubleshootingSteps field to given value.

### HasTroubleshootingSteps

`func (o *MonitorCheckStatus) HasTroubleshootingSteps() bool`

HasTroubleshootingSteps returns a boolean if a field has been set.

### GetTopologyTime

`func (o *MonitorCheckStatus) GetTopologyTime() int32`

GetTopologyTime returns the TopologyTime field if non-nil, zero value otherwise.

### GetTopologyTimeOk

`func (o *MonitorCheckStatus) GetTopologyTimeOk() (*int32, bool)`

GetTopologyTimeOk returns a tuple with the TopologyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyTime

`func (o *MonitorCheckStatus) SetTopologyTime(v int32)`

SetTopologyTime sets TopologyTime field to given value.


### GetDummy

`func (o *MonitorCheckStatus) GetDummy() bool`

GetDummy returns the Dummy field if non-nil, zero value otherwise.

### GetDummyOk

`func (o *MonitorCheckStatus) GetDummyOk() (*bool, bool)`

GetDummyOk returns a tuple with the Dummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummy

`func (o *MonitorCheckStatus) SetDummy(v bool)`

SetDummy sets Dummy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


