# MonitorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Count** | **int32** |  | 
**Level** | [**MessageLevel**](MessageLevel.md) |  | 

## Methods

### NewMonitorError

`func NewMonitorError(error_ string, count int32, level MessageLevel, ) *MonitorError`

NewMonitorError instantiates a new MonitorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorErrorWithDefaults

`func NewMonitorErrorWithDefaults() *MonitorError`

NewMonitorErrorWithDefaults instantiates a new MonitorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MonitorError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MonitorError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MonitorError) SetError(v string)`

SetError sets Error field to given value.


### GetCount

`func (o *MonitorError) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MonitorError) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MonitorError) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLevel

`func (o *MonitorError) GetLevel() MessageLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MonitorError) GetLevelOk() (*MessageLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MonitorError) SetLevel(v MessageLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


