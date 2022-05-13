# MonitorStreamStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]MonitorStreamError**](MonitorStreamError.md) |  | [optional] 
**Metrics** | [**MonitorStreamMetrics**](MonitorStreamMetrics.md) |  | 
**MonitorHealthStateStateCount** | **int32** |  | 

## Methods

### NewMonitorStreamStatus1

`func NewMonitorStreamStatus1(metrics MonitorStreamMetrics, monitorHealthStateStateCount int32, ) *MonitorStreamStatus1`

NewMonitorStreamStatus1 instantiates a new MonitorStreamStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStreamStatus1WithDefaults

`func NewMonitorStreamStatus1WithDefaults() *MonitorStreamStatus1`

NewMonitorStreamStatus1WithDefaults instantiates a new MonitorStreamStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *MonitorStreamStatus1) GetErrors() []MonitorStreamError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MonitorStreamStatus1) GetErrorsOk() (*[]MonitorStreamError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MonitorStreamStatus1) SetErrors(v []MonitorStreamError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MonitorStreamStatus1) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetrics

`func (o *MonitorStreamStatus1) GetMetrics() MonitorStreamMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MonitorStreamStatus1) GetMetricsOk() (*MonitorStreamMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MonitorStreamStatus1) SetMetrics(v MonitorStreamMetrics)`

SetMetrics sets Metrics field to given value.


### GetMonitorHealthStateStateCount

`func (o *MonitorStreamStatus1) GetMonitorHealthStateStateCount() int32`

GetMonitorHealthStateStateCount returns the MonitorHealthStateStateCount field if non-nil, zero value otherwise.

### GetMonitorHealthStateStateCountOk

`func (o *MonitorStreamStatus1) GetMonitorHealthStateStateCountOk() (*int32, bool)`

GetMonitorHealthStateStateCountOk returns a tuple with the MonitorHealthStateStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorHealthStateStateCount

`func (o *MonitorStreamStatus1) SetMonitorHealthStateStateCount(v int32)`

SetMonitorHealthStateStateCount sets MonitorHealthStateStateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


