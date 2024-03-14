# MonitorPreviewResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatesCount** | **int32** | Representing the total count of the monitor results. | 
**UnknownCount** | **int32** | Representing the count of the monitor results that are UNKNOWN in the result. | 
**ClearCount** | **int32** | Representing the count of the monitor results that are CLEAR in the result. | 
**DeviatingCount** | **int32** | Representing the count of the monitor results that are DEVIATING in the result. | 
**CriticalCount** | **int32** | Representing the count of the monitor results that are CRITICAL in the result. | 
**ComponentCounts** | Pointer to [**MonitorPreviewComponentCount**](MonitorPreviewComponentCount.md) |  | [optional] 
**Errors** | **[]string** |  | 

## Methods

### NewMonitorPreviewResult

`func NewMonitorPreviewResult(healthStatesCount int32, unknownCount int32, clearCount int32, deviatingCount int32, criticalCount int32, errors []string, ) *MonitorPreviewResult`

NewMonitorPreviewResult instantiates a new MonitorPreviewResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorPreviewResultWithDefaults

`func NewMonitorPreviewResultWithDefaults() *MonitorPreviewResult`

NewMonitorPreviewResultWithDefaults instantiates a new MonitorPreviewResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatesCount

`func (o *MonitorPreviewResult) GetHealthStatesCount() int32`

GetHealthStatesCount returns the HealthStatesCount field if non-nil, zero value otherwise.

### GetHealthStatesCountOk

`func (o *MonitorPreviewResult) GetHealthStatesCountOk() (*int32, bool)`

GetHealthStatesCountOk returns a tuple with the HealthStatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatesCount

`func (o *MonitorPreviewResult) SetHealthStatesCount(v int32)`

SetHealthStatesCount sets HealthStatesCount field to given value.


### GetUnknownCount

`func (o *MonitorPreviewResult) GetUnknownCount() int32`

GetUnknownCount returns the UnknownCount field if non-nil, zero value otherwise.

### GetUnknownCountOk

`func (o *MonitorPreviewResult) GetUnknownCountOk() (*int32, bool)`

GetUnknownCountOk returns a tuple with the UnknownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownCount

`func (o *MonitorPreviewResult) SetUnknownCount(v int32)`

SetUnknownCount sets UnknownCount field to given value.


### GetClearCount

`func (o *MonitorPreviewResult) GetClearCount() int32`

GetClearCount returns the ClearCount field if non-nil, zero value otherwise.

### GetClearCountOk

`func (o *MonitorPreviewResult) GetClearCountOk() (*int32, bool)`

GetClearCountOk returns a tuple with the ClearCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCount

`func (o *MonitorPreviewResult) SetClearCount(v int32)`

SetClearCount sets ClearCount field to given value.


### GetDeviatingCount

`func (o *MonitorPreviewResult) GetDeviatingCount() int32`

GetDeviatingCount returns the DeviatingCount field if non-nil, zero value otherwise.

### GetDeviatingCountOk

`func (o *MonitorPreviewResult) GetDeviatingCountOk() (*int32, bool)`

GetDeviatingCountOk returns a tuple with the DeviatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviatingCount

`func (o *MonitorPreviewResult) SetDeviatingCount(v int32)`

SetDeviatingCount sets DeviatingCount field to given value.


### GetCriticalCount

`func (o *MonitorPreviewResult) GetCriticalCount() int32`

GetCriticalCount returns the CriticalCount field if non-nil, zero value otherwise.

### GetCriticalCountOk

`func (o *MonitorPreviewResult) GetCriticalCountOk() (*int32, bool)`

GetCriticalCountOk returns a tuple with the CriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCount

`func (o *MonitorPreviewResult) SetCriticalCount(v int32)`

SetCriticalCount sets CriticalCount field to given value.


### GetComponentCounts

`func (o *MonitorPreviewResult) GetComponentCounts() MonitorPreviewComponentCount`

GetComponentCounts returns the ComponentCounts field if non-nil, zero value otherwise.

### GetComponentCountsOk

`func (o *MonitorPreviewResult) GetComponentCountsOk() (*MonitorPreviewComponentCount, bool)`

GetComponentCountsOk returns a tuple with the ComponentCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCounts

`func (o *MonitorPreviewResult) SetComponentCounts(v MonitorPreviewComponentCount)`

SetComponentCounts sets ComponentCounts field to given value.

### HasComponentCounts

`func (o *MonitorPreviewResult) HasComponentCounts() bool`

HasComponentCounts returns a boolean if a field has been set.

### GetErrors

`func (o *MonitorPreviewResult) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MonitorPreviewResult) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MonitorPreviewResult) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


