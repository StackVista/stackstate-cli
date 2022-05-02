# StackpackConfigurationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LastUpdateTimestamp** | **int64** |  | 
**StackPackVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewStackpackConfigurationsInner

`func NewStackpackConfigurationsInner(lastUpdateTimestamp int64, ) *StackpackConfigurationsInner`

NewStackpackConfigurationsInner instantiates a new StackpackConfigurationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackpackConfigurationsInnerWithDefaults

`func NewStackpackConfigurationsInnerWithDefaults() *StackpackConfigurationsInner`

NewStackpackConfigurationsInnerWithDefaults instantiates a new StackpackConfigurationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackpackConfigurationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackpackConfigurationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackpackConfigurationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StackpackConfigurationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *StackpackConfigurationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackpackConfigurationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackpackConfigurationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StackpackConfigurationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *StackpackConfigurationsInner) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *StackpackConfigurationsInner) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *StackpackConfigurationsInner) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.


### GetStackPackVersion

`func (o *StackpackConfigurationsInner) GetStackPackVersion() string`

GetStackPackVersion returns the StackPackVersion field if non-nil, zero value otherwise.

### GetStackPackVersionOk

`func (o *StackpackConfigurationsInner) GetStackPackVersionOk() (*string, bool)`

GetStackPackVersionOk returns a tuple with the StackPackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPackVersion

`func (o *StackpackConfigurationsInner) SetStackPackVersion(v string)`

SetStackPackVersion sets StackPackVersion field to given value.

### HasStackPackVersion

`func (o *StackpackConfigurationsInner) HasStackPackVersion() bool`

HasStackPackVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


