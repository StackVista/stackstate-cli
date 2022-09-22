# MonitorPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RemediationHint** | Pointer to **string** |  | [optional] 
**IntervalSeconds** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**MonitorStatusValue**](MonitorStatusValue.md) |  | [optional] 

## Methods

### NewMonitorPatch

`func NewMonitorPatch() *MonitorPatch`

NewMonitorPatch instantiates a new MonitorPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorPatchWithDefaults

`func NewMonitorPatchWithDefaults() *MonitorPatch`

NewMonitorPatchWithDefaults instantiates a new MonitorPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *MonitorPatch) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MonitorPatch) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MonitorPatch) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MonitorPatch) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *MonitorPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRemediationHint

`func (o *MonitorPatch) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *MonitorPatch) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *MonitorPatch) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *MonitorPatch) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetIntervalSeconds

`func (o *MonitorPatch) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *MonitorPatch) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *MonitorPatch) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.

### HasIntervalSeconds

`func (o *MonitorPatch) HasIntervalSeconds() bool`

HasIntervalSeconds returns a boolean if a field has been set.

### GetTags

`func (o *MonitorPatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitorPatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitorPatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitorPatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *MonitorPatch) GetStatus() MonitorStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorPatch) GetStatusOk() (*MonitorStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorPatch) SetStatus(v MonitorStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


