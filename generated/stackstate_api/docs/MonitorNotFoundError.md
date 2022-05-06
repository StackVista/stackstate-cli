# MonitorNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**MonitorId** | **int64** |  | 

## Methods

### NewMonitorNotFoundError

`func NewMonitorNotFoundError(type_ string, monitorId int64, ) *MonitorNotFoundError`

NewMonitorNotFoundError instantiates a new MonitorNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorNotFoundErrorWithDefaults

`func NewMonitorNotFoundErrorWithDefaults() *MonitorNotFoundError`

NewMonitorNotFoundErrorWithDefaults instantiates a new MonitorNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorNotFoundError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorNotFoundError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorNotFoundError) SetType(v string)`

SetType sets Type field to given value.


### GetMonitorId

`func (o *MonitorNotFoundError) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorNotFoundError) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorNotFoundError) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


