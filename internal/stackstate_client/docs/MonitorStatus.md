# MonitorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | [**Monitor**](Monitor.md) |  | 
**Status** | [**HealthStreamStatus**](HealthStreamStatus.md) |  | 
**TopologyMatchResult** | [**TopologyMatchResult**](TopologyMatchResult.md) |  | 

## Methods

### NewMonitorStatus

`func NewMonitorStatus(monitor Monitor, status HealthStreamStatus, topologyMatchResult TopologyMatchResult, ) *MonitorStatus`

NewMonitorStatus instantiates a new MonitorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusWithDefaults

`func NewMonitorStatusWithDefaults() *MonitorStatus`

NewMonitorStatusWithDefaults instantiates a new MonitorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *MonitorStatus) GetMonitor() Monitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MonitorStatus) GetMonitorOk() (*Monitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MonitorStatus) SetMonitor(v Monitor)`

SetMonitor sets Monitor field to given value.


### GetStatus

`func (o *MonitorStatus) GetStatus() HealthStreamStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorStatus) GetStatusOk() (*HealthStreamStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorStatus) SetStatus(v HealthStreamStatus)`

SetStatus sets Status field to given value.


### GetTopologyMatchResult

`func (o *MonitorStatus) GetTopologyMatchResult() TopologyMatchResult`

GetTopologyMatchResult returns the TopologyMatchResult field if non-nil, zero value otherwise.

### GetTopologyMatchResultOk

`func (o *MonitorStatus) GetTopologyMatchResultOk() (*TopologyMatchResult, bool)`

GetTopologyMatchResultOk returns a tuple with the TopologyMatchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMatchResult

`func (o *MonitorStatus) SetTopologyMatchResult(v TopologyMatchResult)`

SetTopologyMatchResult sets TopologyMatchResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


