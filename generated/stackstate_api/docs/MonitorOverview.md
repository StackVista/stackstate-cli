# MonitorOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | [**Monitor**](Monitor.md) |  | 
**Function** | [**MonitorFunction**](MonitorFunction.md) |  | 
**Errors** | Pointer to [**[]MonitorError**](MonitorError.md) |  | [optional] 

## Methods

### NewMonitorOverview

`func NewMonitorOverview(monitor Monitor, function MonitorFunction, ) *MonitorOverview`

NewMonitorOverview instantiates a new MonitorOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorOverviewWithDefaults

`func NewMonitorOverviewWithDefaults() *MonitorOverview`

NewMonitorOverviewWithDefaults instantiates a new MonitorOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *MonitorOverview) GetMonitor() Monitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MonitorOverview) GetMonitorOk() (*Monitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MonitorOverview) SetMonitor(v Monitor)`

SetMonitor sets Monitor field to given value.


### GetFunction

`func (o *MonitorOverview) GetFunction() MonitorFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *MonitorOverview) GetFunctionOk() (*MonitorFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *MonitorOverview) SetFunction(v MonitorFunction)`

SetFunction sets Function field to given value.


### GetErrors

`func (o *MonitorOverview) GetErrors() []MonitorError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MonitorOverview) GetErrorsOk() (*[]MonitorError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MonitorOverview) SetErrors(v []MonitorError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MonitorOverview) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


