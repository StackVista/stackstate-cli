# MonitorSaveError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalSeconds** | **int32** |  | 
**Type** | **string** |  | 
**MonitorIdOrUrn** | **string** |  | 

## Methods

### NewMonitorSaveError

`func NewMonitorSaveError(intervalSeconds int32, type_ string, monitorIdOrUrn string, ) *MonitorSaveError`

NewMonitorSaveError instantiates a new MonitorSaveError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSaveErrorWithDefaults

`func NewMonitorSaveErrorWithDefaults() *MonitorSaveError`

NewMonitorSaveErrorWithDefaults instantiates a new MonitorSaveError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalSeconds

`func (o *MonitorSaveError) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *MonitorSaveError) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *MonitorSaveError) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.


### GetType

`func (o *MonitorSaveError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorSaveError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorSaveError) SetType(v string)`

SetType sets Type field to given value.


### GetMonitorIdOrUrn

`func (o *MonitorSaveError) GetMonitorIdOrUrn() string`

GetMonitorIdOrUrn returns the MonitorIdOrUrn field if non-nil, zero value otherwise.

### GetMonitorIdOrUrnOk

`func (o *MonitorSaveError) GetMonitorIdOrUrnOk() (*string, bool)`

GetMonitorIdOrUrnOk returns a tuple with the MonitorIdOrUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIdOrUrn

`func (o *MonitorSaveError) SetMonitorIdOrUrn(v string)`

SetMonitorIdOrUrn sets MonitorIdOrUrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


