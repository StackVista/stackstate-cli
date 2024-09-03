# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FunctionId** | **int64** |  | 
**Arguments** | **[]map[string]interface{}** |  | 
**RemediationHint** | Pointer to **string** |  | [optional] 
**IntervalSeconds** | **int32** |  | 
**Tags** | **[]string** |  | 
**Source** | **string** |  | 
**SourceDetails** | Pointer to **string** |  | [optional] 
**CanEdit** | **bool** |  | 
**CanClone** | **bool** |  | 
**Status** | [**MonitorStatusValue**](MonitorStatusValue.md) |  | 
**RuntimeStatus** | [**MonitorRuntimeStatusValue**](MonitorRuntimeStatusValue.md) |  | 
**Dummy** | Pointer to **bool** |  | [optional] 
**LastUpdateTimestamp** | **int64** |  | 

## Methods

### NewMonitor

`func NewMonitor(id int64, name string, functionId int64, arguments []map[string]interface{}, intervalSeconds int32, tags []string, source string, canEdit bool, canClone bool, status MonitorStatusValue, runtimeStatus MonitorRuntimeStatusValue, lastUpdateTimestamp int64, ) *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitor) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *Monitor) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Monitor) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Monitor) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Monitor) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *Monitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Monitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Monitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Monitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionId

`func (o *Monitor) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Monitor) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Monitor) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetArguments

`func (o *Monitor) GetArguments() []map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Monitor) GetArgumentsOk() (*[]map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Monitor) SetArguments(v []map[string]interface{})`

SetArguments sets Arguments field to given value.


### GetRemediationHint

`func (o *Monitor) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *Monitor) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *Monitor) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *Monitor) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetIntervalSeconds

`func (o *Monitor) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *Monitor) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *Monitor) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.


### GetTags

`func (o *Monitor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Monitor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Monitor) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetSource

`func (o *Monitor) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Monitor) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Monitor) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceDetails

`func (o *Monitor) GetSourceDetails() string`

GetSourceDetails returns the SourceDetails field if non-nil, zero value otherwise.

### GetSourceDetailsOk

`func (o *Monitor) GetSourceDetailsOk() (*string, bool)`

GetSourceDetailsOk returns a tuple with the SourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDetails

`func (o *Monitor) SetSourceDetails(v string)`

SetSourceDetails sets SourceDetails field to given value.

### HasSourceDetails

`func (o *Monitor) HasSourceDetails() bool`

HasSourceDetails returns a boolean if a field has been set.

### GetCanEdit

`func (o *Monitor) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *Monitor) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *Monitor) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.


### GetCanClone

`func (o *Monitor) GetCanClone() bool`

GetCanClone returns the CanClone field if non-nil, zero value otherwise.

### GetCanCloneOk

`func (o *Monitor) GetCanCloneOk() (*bool, bool)`

GetCanCloneOk returns a tuple with the CanClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanClone

`func (o *Monitor) SetCanClone(v bool)`

SetCanClone sets CanClone field to given value.


### GetStatus

`func (o *Monitor) GetStatus() MonitorStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Monitor) GetStatusOk() (*MonitorStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Monitor) SetStatus(v MonitorStatusValue)`

SetStatus sets Status field to given value.


### GetRuntimeStatus

`func (o *Monitor) GetRuntimeStatus() MonitorRuntimeStatusValue`

GetRuntimeStatus returns the RuntimeStatus field if non-nil, zero value otherwise.

### GetRuntimeStatusOk

`func (o *Monitor) GetRuntimeStatusOk() (*MonitorRuntimeStatusValue, bool)`

GetRuntimeStatusOk returns a tuple with the RuntimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStatus

`func (o *Monitor) SetRuntimeStatus(v MonitorRuntimeStatusValue)`

SetRuntimeStatus sets RuntimeStatus field to given value.


### GetDummy

`func (o *Monitor) GetDummy() bool`

GetDummy returns the Dummy field if non-nil, zero value otherwise.

### GetDummyOk

`func (o *Monitor) GetDummyOk() (*bool, bool)`

GetDummyOk returns a tuple with the Dummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummy

`func (o *Monitor) SetDummy(v bool)`

SetDummy sets Dummy field to given value.

### HasDummy

`func (o *Monitor) HasDummy() bool`

HasDummy returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Monitor) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Monitor) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Monitor) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


