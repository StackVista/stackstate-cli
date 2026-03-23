# OverviewFetchTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**UsedTimeoutSeconds** | **int64** | Timeout used for the overview request (seconds). | 

## Methods

### NewOverviewFetchTimeout

`func NewOverviewFetchTimeout(type_ string, usedTimeoutSeconds int64, ) *OverviewFetchTimeout`

NewOverviewFetchTimeout instantiates a new OverviewFetchTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewFetchTimeoutWithDefaults

`func NewOverviewFetchTimeoutWithDefaults() *OverviewFetchTimeout`

NewOverviewFetchTimeoutWithDefaults instantiates a new OverviewFetchTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewFetchTimeout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewFetchTimeout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewFetchTimeout) SetType(v string)`

SetType sets Type field to given value.


### GetUsedTimeoutSeconds

`func (o *OverviewFetchTimeout) GetUsedTimeoutSeconds() int64`

GetUsedTimeoutSeconds returns the UsedTimeoutSeconds field if non-nil, zero value otherwise.

### GetUsedTimeoutSecondsOk

`func (o *OverviewFetchTimeout) GetUsedTimeoutSecondsOk() (*int64, bool)`

GetUsedTimeoutSecondsOk returns a tuple with the UsedTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTimeoutSeconds

`func (o *OverviewFetchTimeout) SetUsedTimeoutSeconds(v int64)`

SetUsedTimeoutSeconds sets UsedTimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


