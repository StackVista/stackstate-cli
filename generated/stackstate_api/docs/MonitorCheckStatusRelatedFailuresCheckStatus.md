# MonitorCheckStatusRelatedFailuresCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Health** | [**HealthStateValue**](HealthStateValue.md) |  | 
**Name** | **string** |  | 
**Component** | [**MonitorCheckStatusComponent**](MonitorCheckStatusComponent.md) |  | 
**Topology** | [**MonitorCheckStatusRelatedFailuresTopology**](MonitorCheckStatusRelatedFailuresTopology.md) |  | 
**TriggeredTimestamp** | **int32** |  | 

## Methods

### NewMonitorCheckStatusRelatedFailuresCheckStatus

`func NewMonitorCheckStatusRelatedFailuresCheckStatus(identifier string, health HealthStateValue, name string, component MonitorCheckStatusComponent, topology MonitorCheckStatusRelatedFailuresTopology, triggeredTimestamp int32, ) *MonitorCheckStatusRelatedFailuresCheckStatus`

NewMonitorCheckStatusRelatedFailuresCheckStatus instantiates a new MonitorCheckStatusRelatedFailuresCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorCheckStatusRelatedFailuresCheckStatusWithDefaults

`func NewMonitorCheckStatusRelatedFailuresCheckStatusWithDefaults() *MonitorCheckStatusRelatedFailuresCheckStatus`

NewMonitorCheckStatusRelatedFailuresCheckStatusWithDefaults instantiates a new MonitorCheckStatusRelatedFailuresCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetHealth

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetHealth() HealthStateValue`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetHealthOk() (*HealthStateValue, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetHealth(v HealthStateValue)`

SetHealth sets Health field to given value.


### GetName

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetComponent() MonitorCheckStatusComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetComponentOk() (*MonitorCheckStatusComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetComponent(v MonitorCheckStatusComponent)`

SetComponent sets Component field to given value.


### GetTopology

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetTopology() MonitorCheckStatusRelatedFailuresTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetTopologyOk() (*MonitorCheckStatusRelatedFailuresTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetTopology(v MonitorCheckStatusRelatedFailuresTopology)`

SetTopology sets Topology field to given value.


### GetTriggeredTimestamp

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetTriggeredTimestamp() int32`

GetTriggeredTimestamp returns the TriggeredTimestamp field if non-nil, zero value otherwise.

### GetTriggeredTimestampOk

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) GetTriggeredTimestampOk() (*int32, bool)`

GetTriggeredTimestampOk returns a tuple with the TriggeredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTimestamp

`func (o *MonitorCheckStatusRelatedFailuresCheckStatus) SetTriggeredTimestamp(v int32)`

SetTriggeredTimestamp sets TriggeredTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


