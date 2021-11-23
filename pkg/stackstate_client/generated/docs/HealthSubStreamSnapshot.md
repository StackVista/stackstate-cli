# HealthSubStreamSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExpiryIntervalMs** | Pointer to **int32** |  | [optional] 
**RepeatIntervalMs** | **int32** |  | 

## Methods

### NewHealthSubStreamSnapshot

`func NewHealthSubStreamSnapshot(type_ string, repeatIntervalMs int32, ) *HealthSubStreamSnapshot`

NewHealthSubStreamSnapshot instantiates a new HealthSubStreamSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthSubStreamSnapshotWithDefaults

`func NewHealthSubStreamSnapshotWithDefaults() *HealthSubStreamSnapshot`

NewHealthSubStreamSnapshotWithDefaults instantiates a new HealthSubStreamSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthSubStreamSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthSubStreamSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthSubStreamSnapshot) SetType(v string)`

SetType sets Type field to given value.


### GetExpiryIntervalMs

`func (o *HealthSubStreamSnapshot) GetExpiryIntervalMs() int32`

GetExpiryIntervalMs returns the ExpiryIntervalMs field if non-nil, zero value otherwise.

### GetExpiryIntervalMsOk

`func (o *HealthSubStreamSnapshot) GetExpiryIntervalMsOk() (*int32, bool)`

GetExpiryIntervalMsOk returns a tuple with the ExpiryIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryIntervalMs

`func (o *HealthSubStreamSnapshot) SetExpiryIntervalMs(v int32)`

SetExpiryIntervalMs sets ExpiryIntervalMs field to given value.

### HasExpiryIntervalMs

`func (o *HealthSubStreamSnapshot) HasExpiryIntervalMs() bool`

HasExpiryIntervalMs returns a boolean if a field has been set.

### GetRepeatIntervalMs

`func (o *HealthSubStreamSnapshot) GetRepeatIntervalMs() int32`

GetRepeatIntervalMs returns the RepeatIntervalMs field if non-nil, zero value otherwise.

### GetRepeatIntervalMsOk

`func (o *HealthSubStreamSnapshot) GetRepeatIntervalMsOk() (*int32, bool)`

GetRepeatIntervalMsOk returns a tuple with the RepeatIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIntervalMs

`func (o *HealthSubStreamSnapshot) SetRepeatIntervalMs(v int32)`

SetRepeatIntervalMs sets RepeatIntervalMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


