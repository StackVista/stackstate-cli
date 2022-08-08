# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Arguments** | [**[]Argument**](Argument.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Function** | **int64** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**RemediationHint** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**CheckState**](CheckState.md) |  | [optional] 
**SyncCreated** | **bool** |  | 

## Methods

### NewCheck

`func NewCheck(type_ string, arguments []Argument, function int64, name string, syncCreated bool, ) *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Check) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Check) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Check) SetType(v string)`

SetType sets Type field to given value.


### GetArguments

`func (o *Check) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Check) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Check) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.


### GetDescription

`func (o *Check) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Check) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Check) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Check) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunction

`func (o *Check) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Check) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Check) SetFunction(v int64)`

SetFunction sets Function field to given value.


### GetId

`func (o *Check) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Check) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Check) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Check) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Check) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Check) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *Check) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Check) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Check) SetName(v string)`

SetName sets Name field to given value.


### GetRemediationHint

`func (o *Check) GetRemediationHint() string`

GetRemediationHint returns the RemediationHint field if non-nil, zero value otherwise.

### GetRemediationHintOk

`func (o *Check) GetRemediationHintOk() (*string, bool)`

GetRemediationHintOk returns a tuple with the RemediationHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationHint

`func (o *Check) SetRemediationHint(v string)`

SetRemediationHint sets RemediationHint field to given value.

### HasRemediationHint

`func (o *Check) HasRemediationHint() bool`

HasRemediationHint returns a boolean if a field has been set.

### GetState

`func (o *Check) GetState() CheckState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Check) GetStateOk() (*CheckState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Check) SetState(v CheckState)`

SetState sets State field to given value.

### HasState

`func (o *Check) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSyncCreated

`func (o *Check) GetSyncCreated() bool`

GetSyncCreated returns the SyncCreated field if non-nil, zero value otherwise.

### GetSyncCreatedOk

`func (o *Check) GetSyncCreatedOk() (*bool, bool)`

GetSyncCreatedOk returns a tuple with the SyncCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCreated

`func (o *Check) SetSyncCreated(v bool)`

SetSyncCreated sets SyncCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


