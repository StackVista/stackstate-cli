# MonitorStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | [**Monitor1**](Monitor1.md) |  | 
**Status** | [**MonitorStreamStatus1**](MonitorStreamStatus1.md) |  | 
**TopologyMatchResult** | [**TopologyMatchResult**](TopologyMatchResult.md) |  | 

## Methods

### NewMonitorStatus1

`func NewMonitorStatus1(monitor Monitor1, status MonitorStreamStatus1, topologyMatchResult TopologyMatchResult, ) *MonitorStatus1`

NewMonitorStatus1 instantiates a new MonitorStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatus1WithDefaults

`func NewMonitorStatus1WithDefaults() *MonitorStatus1`

NewMonitorStatus1WithDefaults instantiates a new MonitorStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *MonitorStatus1) GetMonitor() Monitor1`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MonitorStatus1) GetMonitorOk() (*Monitor1, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MonitorStatus1) SetMonitor(v Monitor1)`

SetMonitor sets Monitor field to given value.


### GetStatus

`func (o *MonitorStatus1) GetStatus() MonitorStreamStatus1`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorStatus1) GetStatusOk() (*MonitorStreamStatus1, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorStatus1) SetStatus(v MonitorStreamStatus1)`

SetStatus sets Status field to given value.


### GetTopologyMatchResult

`func (o *MonitorStatus1) GetTopologyMatchResult() TopologyMatchResult`

GetTopologyMatchResult returns the TopologyMatchResult field if non-nil, zero value otherwise.

### GetTopologyMatchResultOk

`func (o *MonitorStatus1) GetTopologyMatchResultOk() (*TopologyMatchResult, bool)`

GetTopologyMatchResultOk returns a tuple with the TopologyMatchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMatchResult

`func (o *MonitorStatus1) SetTopologyMatchResult(v TopologyMatchResult)`

SetTopologyMatchResult sets TopologyMatchResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


