# HealthSubStreamConsistencyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExpiryIntervalMs** | **int32** |  | 
**RepeatIntervalMs** | **int32** |  | 
**Offset** | **int64** |  | 
**BatchIndex** | Pointer to **int64** |  | [optional] 

## Methods

### NewHealthSubStreamConsistencyState

`func NewHealthSubStreamConsistencyState(type_ string, expiryIntervalMs int32, repeatIntervalMs int32, offset int64, ) *HealthSubStreamConsistencyState`

NewHealthSubStreamConsistencyState instantiates a new HealthSubStreamConsistencyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthSubStreamConsistencyStateWithDefaults

`func NewHealthSubStreamConsistencyStateWithDefaults() *HealthSubStreamConsistencyState`

NewHealthSubStreamConsistencyStateWithDefaults instantiates a new HealthSubStreamConsistencyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthSubStreamConsistencyState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthSubStreamConsistencyState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthSubStreamConsistencyState) SetType(v string)`

SetType sets Type field to given value.


### GetExpiryIntervalMs

`func (o *HealthSubStreamConsistencyState) GetExpiryIntervalMs() int32`

GetExpiryIntervalMs returns the ExpiryIntervalMs field if non-nil, zero value otherwise.

### GetExpiryIntervalMsOk

`func (o *HealthSubStreamConsistencyState) GetExpiryIntervalMsOk() (*int32, bool)`

GetExpiryIntervalMsOk returns a tuple with the ExpiryIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryIntervalMs

`func (o *HealthSubStreamConsistencyState) SetExpiryIntervalMs(v int32)`

SetExpiryIntervalMs sets ExpiryIntervalMs field to given value.


### GetRepeatIntervalMs

`func (o *HealthSubStreamConsistencyState) GetRepeatIntervalMs() int32`

GetRepeatIntervalMs returns the RepeatIntervalMs field if non-nil, zero value otherwise.

### GetRepeatIntervalMsOk

`func (o *HealthSubStreamConsistencyState) GetRepeatIntervalMsOk() (*int32, bool)`

GetRepeatIntervalMsOk returns a tuple with the RepeatIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIntervalMs

`func (o *HealthSubStreamConsistencyState) SetRepeatIntervalMs(v int32)`

SetRepeatIntervalMs sets RepeatIntervalMs field to given value.


### GetOffset

`func (o *HealthSubStreamConsistencyState) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *HealthSubStreamConsistencyState) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *HealthSubStreamConsistencyState) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetBatchIndex

`func (o *HealthSubStreamConsistencyState) GetBatchIndex() int64`

GetBatchIndex returns the BatchIndex field if non-nil, zero value otherwise.

### GetBatchIndexOk

`func (o *HealthSubStreamConsistencyState) GetBatchIndexOk() (*int64, bool)`

GetBatchIndexOk returns a tuple with the BatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIndex

`func (o *HealthSubStreamConsistencyState) SetBatchIndex(v int64)`

SetBatchIndex sets BatchIndex field to given value.

### HasBatchIndex

`func (o *HealthSubStreamConsistencyState) HasBatchIndex() bool`

HasBatchIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


