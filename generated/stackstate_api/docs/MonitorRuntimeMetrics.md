# MonitorRuntimeMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatesCount** | Pointer to **int32** | Representing the total count of the monitors results. | [optional] 
**UnmappedHealthStatesCount** | Pointer to **int32** | Representing the total count of the monitors results that are not mapped to topology/ | [optional] 
**UnknownCount** | Pointer to **int32** | Representing the count of the monitors results that are UNKNOWN and are mapped to topology. | [optional] 
**ClearCount** | Pointer to **int32** | Representing the count of the monitors results that are CLEAR and are mapped to topology. | [optional] 
**DeviatingCount** | Pointer to **int32** | Representing the count of the monitors results that are DEVIATING and are mapped to topology. | [optional] 
**CriticalCount** | Pointer to **int32** | Representing the count of the monitors results that are CRITICAL and are mapped to topology. | [optional] 
**LastRunTimestamp** | Pointer to **int64** | Representing the epoch millis of the last monitor run. | [optional] 
**LastSuccessfulRunTimestamp** | Pointer to **int64** | Representing the epoch millis of the last monitor successful run. | [optional] 
**LastFailedRunTimestamp** | Pointer to **int64** | Representing the epoch millis of the last monitor failed run. | [optional] 

## Methods

### NewMonitorRuntimeMetrics

`func NewMonitorRuntimeMetrics() *MonitorRuntimeMetrics`

NewMonitorRuntimeMetrics instantiates a new MonitorRuntimeMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorRuntimeMetricsWithDefaults

`func NewMonitorRuntimeMetricsWithDefaults() *MonitorRuntimeMetrics`

NewMonitorRuntimeMetricsWithDefaults instantiates a new MonitorRuntimeMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatesCount

`func (o *MonitorRuntimeMetrics) GetHealthStatesCount() int32`

GetHealthStatesCount returns the HealthStatesCount field if non-nil, zero value otherwise.

### GetHealthStatesCountOk

`func (o *MonitorRuntimeMetrics) GetHealthStatesCountOk() (*int32, bool)`

GetHealthStatesCountOk returns a tuple with the HealthStatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatesCount

`func (o *MonitorRuntimeMetrics) SetHealthStatesCount(v int32)`

SetHealthStatesCount sets HealthStatesCount field to given value.

### HasHealthStatesCount

`func (o *MonitorRuntimeMetrics) HasHealthStatesCount() bool`

HasHealthStatesCount returns a boolean if a field has been set.

### GetUnmappedHealthStatesCount

`func (o *MonitorRuntimeMetrics) GetUnmappedHealthStatesCount() int32`

GetUnmappedHealthStatesCount returns the UnmappedHealthStatesCount field if non-nil, zero value otherwise.

### GetUnmappedHealthStatesCountOk

`func (o *MonitorRuntimeMetrics) GetUnmappedHealthStatesCountOk() (*int32, bool)`

GetUnmappedHealthStatesCountOk returns a tuple with the UnmappedHealthStatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmappedHealthStatesCount

`func (o *MonitorRuntimeMetrics) SetUnmappedHealthStatesCount(v int32)`

SetUnmappedHealthStatesCount sets UnmappedHealthStatesCount field to given value.

### HasUnmappedHealthStatesCount

`func (o *MonitorRuntimeMetrics) HasUnmappedHealthStatesCount() bool`

HasUnmappedHealthStatesCount returns a boolean if a field has been set.

### GetUnknownCount

`func (o *MonitorRuntimeMetrics) GetUnknownCount() int32`

GetUnknownCount returns the UnknownCount field if non-nil, zero value otherwise.

### GetUnknownCountOk

`func (o *MonitorRuntimeMetrics) GetUnknownCountOk() (*int32, bool)`

GetUnknownCountOk returns a tuple with the UnknownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownCount

`func (o *MonitorRuntimeMetrics) SetUnknownCount(v int32)`

SetUnknownCount sets UnknownCount field to given value.

### HasUnknownCount

`func (o *MonitorRuntimeMetrics) HasUnknownCount() bool`

HasUnknownCount returns a boolean if a field has been set.

### GetClearCount

`func (o *MonitorRuntimeMetrics) GetClearCount() int32`

GetClearCount returns the ClearCount field if non-nil, zero value otherwise.

### GetClearCountOk

`func (o *MonitorRuntimeMetrics) GetClearCountOk() (*int32, bool)`

GetClearCountOk returns a tuple with the ClearCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCount

`func (o *MonitorRuntimeMetrics) SetClearCount(v int32)`

SetClearCount sets ClearCount field to given value.

### HasClearCount

`func (o *MonitorRuntimeMetrics) HasClearCount() bool`

HasClearCount returns a boolean if a field has been set.

### GetDeviatingCount

`func (o *MonitorRuntimeMetrics) GetDeviatingCount() int32`

GetDeviatingCount returns the DeviatingCount field if non-nil, zero value otherwise.

### GetDeviatingCountOk

`func (o *MonitorRuntimeMetrics) GetDeviatingCountOk() (*int32, bool)`

GetDeviatingCountOk returns a tuple with the DeviatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviatingCount

`func (o *MonitorRuntimeMetrics) SetDeviatingCount(v int32)`

SetDeviatingCount sets DeviatingCount field to given value.

### HasDeviatingCount

`func (o *MonitorRuntimeMetrics) HasDeviatingCount() bool`

HasDeviatingCount returns a boolean if a field has been set.

### GetCriticalCount

`func (o *MonitorRuntimeMetrics) GetCriticalCount() int32`

GetCriticalCount returns the CriticalCount field if non-nil, zero value otherwise.

### GetCriticalCountOk

`func (o *MonitorRuntimeMetrics) GetCriticalCountOk() (*int32, bool)`

GetCriticalCountOk returns a tuple with the CriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCount

`func (o *MonitorRuntimeMetrics) SetCriticalCount(v int32)`

SetCriticalCount sets CriticalCount field to given value.

### HasCriticalCount

`func (o *MonitorRuntimeMetrics) HasCriticalCount() bool`

HasCriticalCount returns a boolean if a field has been set.

### GetLastRunTimestamp

`func (o *MonitorRuntimeMetrics) GetLastRunTimestamp() int64`

GetLastRunTimestamp returns the LastRunTimestamp field if non-nil, zero value otherwise.

### GetLastRunTimestampOk

`func (o *MonitorRuntimeMetrics) GetLastRunTimestampOk() (*int64, bool)`

GetLastRunTimestampOk returns a tuple with the LastRunTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunTimestamp

`func (o *MonitorRuntimeMetrics) SetLastRunTimestamp(v int64)`

SetLastRunTimestamp sets LastRunTimestamp field to given value.

### HasLastRunTimestamp

`func (o *MonitorRuntimeMetrics) HasLastRunTimestamp() bool`

HasLastRunTimestamp returns a boolean if a field has been set.

### GetLastSuccessfulRunTimestamp

`func (o *MonitorRuntimeMetrics) GetLastSuccessfulRunTimestamp() int64`

GetLastSuccessfulRunTimestamp returns the LastSuccessfulRunTimestamp field if non-nil, zero value otherwise.

### GetLastSuccessfulRunTimestampOk

`func (o *MonitorRuntimeMetrics) GetLastSuccessfulRunTimestampOk() (*int64, bool)`

GetLastSuccessfulRunTimestampOk returns a tuple with the LastSuccessfulRunTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulRunTimestamp

`func (o *MonitorRuntimeMetrics) SetLastSuccessfulRunTimestamp(v int64)`

SetLastSuccessfulRunTimestamp sets LastSuccessfulRunTimestamp field to given value.

### HasLastSuccessfulRunTimestamp

`func (o *MonitorRuntimeMetrics) HasLastSuccessfulRunTimestamp() bool`

HasLastSuccessfulRunTimestamp returns a boolean if a field has been set.

### GetLastFailedRunTimestamp

`func (o *MonitorRuntimeMetrics) GetLastFailedRunTimestamp() int64`

GetLastFailedRunTimestamp returns the LastFailedRunTimestamp field if non-nil, zero value otherwise.

### GetLastFailedRunTimestampOk

`func (o *MonitorRuntimeMetrics) GetLastFailedRunTimestampOk() (*int64, bool)`

GetLastFailedRunTimestampOk returns a tuple with the LastFailedRunTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedRunTimestamp

`func (o *MonitorRuntimeMetrics) SetLastFailedRunTimestamp(v int64)`

SetLastFailedRunTimestamp sets LastFailedRunTimestamp field to given value.

### HasLastFailedRunTimestamp

`func (o *MonitorRuntimeMetrics) HasLastFailedRunTimestamp() bool`

HasLastFailedRunTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


