# MonitorStreamStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]MonitorStreamError**](MonitorStreamError.md) |  | [optional] 
**Metrics** | [**MonitorStreamMetrics**](MonitorStreamMetrics.md) |  | 
**MonitorHealthStateStateCount** | **int32** |  | 

## Methods

### NewMonitorStreamStatus

`func NewMonitorStreamStatus(metrics MonitorStreamMetrics, monitorHealthStateStateCount int32, ) *MonitorStreamStatus`

NewMonitorStreamStatus instantiates a new MonitorStreamStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStreamStatusWithDefaults

`func NewMonitorStreamStatusWithDefaults() *MonitorStreamStatus`

NewMonitorStreamStatusWithDefaults instantiates a new MonitorStreamStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *MonitorStreamStatus) GetErrors() []MonitorStreamError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MonitorStreamStatus) GetErrorsOk() (*[]MonitorStreamError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MonitorStreamStatus) SetErrors(v []MonitorStreamError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MonitorStreamStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetrics

`func (o *MonitorStreamStatus) GetMetrics() MonitorStreamMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MonitorStreamStatus) GetMetricsOk() (*MonitorStreamMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MonitorStreamStatus) SetMetrics(v MonitorStreamMetrics)`

SetMetrics sets Metrics field to given value.


### GetMonitorHealthStateStateCount

`func (o *MonitorStreamStatus) GetMonitorHealthStateStateCount() int32`

GetMonitorHealthStateStateCount returns the MonitorHealthStateStateCount field if non-nil, zero value otherwise.

### GetMonitorHealthStateStateCountOk

`func (o *MonitorStreamStatus) GetMonitorHealthStateStateCountOk() (*int32, bool)`

GetMonitorHealthStateStateCountOk returns a tuple with the MonitorHealthStateStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorHealthStateStateCount

`func (o *MonitorStreamStatus) SetMonitorHealthStateStateCount(v int32)`

SetMonitorHealthStateStateCount sets MonitorHealthStateStateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


