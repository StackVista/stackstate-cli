# OverviewDataUnavailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**UnavailableAtEpochMs** | **int64** | Time when data became unavailable (epoch ms). | 
**AvailableSinceEpochMs** | **int64** | Time when data became available again (epoch ms). | 

## Methods

### NewOverviewDataUnavailable

`func NewOverviewDataUnavailable(type_ string, unavailableAtEpochMs int64, availableSinceEpochMs int64, ) *OverviewDataUnavailable`

NewOverviewDataUnavailable instantiates a new OverviewDataUnavailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewDataUnavailableWithDefaults

`func NewOverviewDataUnavailableWithDefaults() *OverviewDataUnavailable`

NewOverviewDataUnavailableWithDefaults instantiates a new OverviewDataUnavailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewDataUnavailable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewDataUnavailable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewDataUnavailable) SetType(v string)`

SetType sets Type field to given value.


### GetUnavailableAtEpochMs

`func (o *OverviewDataUnavailable) GetUnavailableAtEpochMs() int64`

GetUnavailableAtEpochMs returns the UnavailableAtEpochMs field if non-nil, zero value otherwise.

### GetUnavailableAtEpochMsOk

`func (o *OverviewDataUnavailable) GetUnavailableAtEpochMsOk() (*int64, bool)`

GetUnavailableAtEpochMsOk returns a tuple with the UnavailableAtEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableAtEpochMs

`func (o *OverviewDataUnavailable) SetUnavailableAtEpochMs(v int64)`

SetUnavailableAtEpochMs sets UnavailableAtEpochMs field to given value.


### GetAvailableSinceEpochMs

`func (o *OverviewDataUnavailable) GetAvailableSinceEpochMs() int64`

GetAvailableSinceEpochMs returns the AvailableSinceEpochMs field if non-nil, zero value otherwise.

### GetAvailableSinceEpochMsOk

`func (o *OverviewDataUnavailable) GetAvailableSinceEpochMsOk() (*int64, bool)`

GetAvailableSinceEpochMsOk returns a tuple with the AvailableSinceEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSinceEpochMs

`func (o *OverviewDataUnavailable) SetAvailableSinceEpochMs(v int64)`

SetAvailableSinceEpochMs sets AvailableSinceEpochMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


