# ReadyStatusProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ReadyNumber** | **string** | Cel expression that returns a number | 
**TotalNumber** | **string** | Cel expression that returns a number | 
**ReadyStatus** | Pointer to **string** | Cel expression that returns a string that represents a valid HealthState | [optional] 

## Methods

### NewReadyStatusProjection

`func NewReadyStatusProjection(type_ string, readyNumber string, totalNumber string, ) *ReadyStatusProjection`

NewReadyStatusProjection instantiates a new ReadyStatusProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadyStatusProjectionWithDefaults

`func NewReadyStatusProjectionWithDefaults() *ReadyStatusProjection`

NewReadyStatusProjectionWithDefaults instantiates a new ReadyStatusProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReadyStatusProjection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReadyStatusProjection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReadyStatusProjection) SetType(v string)`

SetType sets Type field to given value.


### GetReadyNumber

`func (o *ReadyStatusProjection) GetReadyNumber() string`

GetReadyNumber returns the ReadyNumber field if non-nil, zero value otherwise.

### GetReadyNumberOk

`func (o *ReadyStatusProjection) GetReadyNumberOk() (*string, bool)`

GetReadyNumberOk returns a tuple with the ReadyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyNumber

`func (o *ReadyStatusProjection) SetReadyNumber(v string)`

SetReadyNumber sets ReadyNumber field to given value.


### GetTotalNumber

`func (o *ReadyStatusProjection) GetTotalNumber() string`

GetTotalNumber returns the TotalNumber field if non-nil, zero value otherwise.

### GetTotalNumberOk

`func (o *ReadyStatusProjection) GetTotalNumberOk() (*string, bool)`

GetTotalNumberOk returns a tuple with the TotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumber

`func (o *ReadyStatusProjection) SetTotalNumber(v string)`

SetTotalNumber sets TotalNumber field to given value.


### GetReadyStatus

`func (o *ReadyStatusProjection) GetReadyStatus() string`

GetReadyStatus returns the ReadyStatus field if non-nil, zero value otherwise.

### GetReadyStatusOk

`func (o *ReadyStatusProjection) GetReadyStatusOk() (*string, bool)`

GetReadyStatusOk returns a tuple with the ReadyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyStatus

`func (o *ReadyStatusProjection) SetReadyStatus(v string)`

SetReadyStatus sets ReadyStatus field to given value.

### HasReadyStatus

`func (o *ReadyStatusProjection) HasReadyStatus() bool`

HasReadyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


