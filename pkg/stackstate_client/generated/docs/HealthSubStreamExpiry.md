# HealthSubStreamExpiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExpiryIntervalMs** | **int32** |  | 
**RepeatIntervalMs** | **int32** |  | 

## Methods

### NewHealthSubStreamExpiry

`func NewHealthSubStreamExpiry(type_ string, expiryIntervalMs int32, repeatIntervalMs int32, ) *HealthSubStreamExpiry`

NewHealthSubStreamExpiry instantiates a new HealthSubStreamExpiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthSubStreamExpiryWithDefaults

`func NewHealthSubStreamExpiryWithDefaults() *HealthSubStreamExpiry`

NewHealthSubStreamExpiryWithDefaults instantiates a new HealthSubStreamExpiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthSubStreamExpiry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthSubStreamExpiry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthSubStreamExpiry) SetType(v string)`

SetType sets Type field to given value.


### GetExpiryIntervalMs

`func (o *HealthSubStreamExpiry) GetExpiryIntervalMs() int32`

GetExpiryIntervalMs returns the ExpiryIntervalMs field if non-nil, zero value otherwise.

### GetExpiryIntervalMsOk

`func (o *HealthSubStreamExpiry) GetExpiryIntervalMsOk() (*int32, bool)`

GetExpiryIntervalMsOk returns a tuple with the ExpiryIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryIntervalMs

`func (o *HealthSubStreamExpiry) SetExpiryIntervalMs(v int32)`

SetExpiryIntervalMs sets ExpiryIntervalMs field to given value.


### GetRepeatIntervalMs

`func (o *HealthSubStreamExpiry) GetRepeatIntervalMs() int32`

GetRepeatIntervalMs returns the RepeatIntervalMs field if non-nil, zero value otherwise.

### GetRepeatIntervalMsOk

`func (o *HealthSubStreamExpiry) GetRepeatIntervalMsOk() (*int32, bool)`

GetRepeatIntervalMsOk returns a tuple with the RepeatIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIntervalMs

`func (o *HealthSubStreamExpiry) SetRepeatIntervalMs(v int32)`

SetRepeatIntervalMs sets RepeatIntervalMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


