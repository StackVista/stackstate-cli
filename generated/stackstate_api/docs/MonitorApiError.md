# MonitorApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | Pointer to **string** |  | [optional] 
**MonitorIdType** | Pointer to **string** |  | [optional] 
**StatusCode** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewMonitorApiError

`func NewMonitorApiError(statusCode string, message string, ) *MonitorApiError`

NewMonitorApiError instantiates a new MonitorApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorApiErrorWithDefaults

`func NewMonitorApiErrorWithDefaults() *MonitorApiError`

NewMonitorApiErrorWithDefaults instantiates a new MonitorApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *MonitorApiError) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorApiError) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorApiError) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *MonitorApiError) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetMonitorIdType

`func (o *MonitorApiError) GetMonitorIdType() string`

GetMonitorIdType returns the MonitorIdType field if non-nil, zero value otherwise.

### GetMonitorIdTypeOk

`func (o *MonitorApiError) GetMonitorIdTypeOk() (*string, bool)`

GetMonitorIdTypeOk returns a tuple with the MonitorIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIdType

`func (o *MonitorApiError) SetMonitorIdType(v string)`

SetMonitorIdType sets MonitorIdType field to given value.

### HasMonitorIdType

`func (o *MonitorApiError) HasMonitorIdType() bool`

HasMonitorIdType returns a boolean if a field has been set.

### GetStatusCode

`func (o *MonitorApiError) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *MonitorApiError) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *MonitorApiError) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *MonitorApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitorApiError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


