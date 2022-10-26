# MonitorValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalSeconds** | **int32** |  | 
**Type** | **string** |  | 
**MonitorIdOrUrn** | **string** |  | 

## Methods

### NewMonitorValidationError

`func NewMonitorValidationError(intervalSeconds int32, type_ string, monitorIdOrUrn string, ) *MonitorValidationError`

NewMonitorValidationError instantiates a new MonitorValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorValidationErrorWithDefaults

`func NewMonitorValidationErrorWithDefaults() *MonitorValidationError`

NewMonitorValidationErrorWithDefaults instantiates a new MonitorValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalSeconds

`func (o *MonitorValidationError) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *MonitorValidationError) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *MonitorValidationError) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.


### GetType

`func (o *MonitorValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorValidationError) SetType(v string)`

SetType sets Type field to given value.


### GetMonitorIdOrUrn

`func (o *MonitorValidationError) GetMonitorIdOrUrn() string`

GetMonitorIdOrUrn returns the MonitorIdOrUrn field if non-nil, zero value otherwise.

### GetMonitorIdOrUrnOk

`func (o *MonitorValidationError) GetMonitorIdOrUrnOk() (*string, bool)`

GetMonitorIdOrUrnOk returns a tuple with the MonitorIdOrUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIdOrUrn

`func (o *MonitorValidationError) SetMonitorIdOrUrn(v string)`

SetMonitorIdOrUrn sets MonitorIdOrUrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


