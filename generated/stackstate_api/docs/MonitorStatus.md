# MonitorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | [**Monitor**](Monitor.md) |  | 
**Errors** | Pointer to [**[]MonitorError**](MonitorError.md) |  | [optional] 
**Metrics** | [**MonitorMetrics**](MonitorMetrics.md) |  | 
**MonitorHealthStateStateCount** | **int32** |  | 
**TopologyMatchResult** | [**TopologyMatchResult**](TopologyMatchResult.md) |  | 

## Methods

### NewMonitorStatus

`func NewMonitorStatus(monitor Monitor, metrics MonitorMetrics, monitorHealthStateStateCount int32, topologyMatchResult TopologyMatchResult, ) *MonitorStatus`

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


### GetErrors

`func (o *MonitorStatus) GetErrors() []MonitorError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MonitorStatus) GetErrorsOk() (*[]MonitorError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MonitorStatus) SetErrors(v []MonitorError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MonitorStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetrics

`func (o *MonitorStatus) GetMetrics() MonitorMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MonitorStatus) GetMetricsOk() (*MonitorMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MonitorStatus) SetMetrics(v MonitorMetrics)`

SetMetrics sets Metrics field to given value.


### GetMonitorHealthStateStateCount

`func (o *MonitorStatus) GetMonitorHealthStateStateCount() int32`

GetMonitorHealthStateStateCount returns the MonitorHealthStateStateCount field if non-nil, zero value otherwise.

### GetMonitorHealthStateStateCountOk

`func (o *MonitorStatus) GetMonitorHealthStateStateCountOk() (*int32, bool)`

GetMonitorHealthStateStateCountOk returns a tuple with the MonitorHealthStateStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorHealthStateStateCount

`func (o *MonitorStatus) SetMonitorHealthStateStateCount(v int32)`

SetMonitorHealthStateStateCount sets MonitorHealthStateStateCount field to given value.


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


