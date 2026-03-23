# ReadyStatusCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Ready** | **int32** |  | 
**Total** | **int32** |  | 
**Status** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 

## Methods

### NewReadyStatusCell

`func NewReadyStatusCell(type_ string, ready int32, total int32, ) *ReadyStatusCell`

NewReadyStatusCell instantiates a new ReadyStatusCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadyStatusCellWithDefaults

`func NewReadyStatusCellWithDefaults() *ReadyStatusCell`

NewReadyStatusCellWithDefaults instantiates a new ReadyStatusCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReadyStatusCell) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReadyStatusCell) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReadyStatusCell) SetType(v string)`

SetType sets Type field to given value.


### GetReady

`func (o *ReadyStatusCell) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ReadyStatusCell) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ReadyStatusCell) SetReady(v int32)`

SetReady sets Ready field to given value.


### GetTotal

`func (o *ReadyStatusCell) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReadyStatusCell) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReadyStatusCell) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetStatus

`func (o *ReadyStatusCell) GetStatus() HealthStateValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReadyStatusCell) GetStatusOk() (*HealthStateValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReadyStatusCell) SetStatus(v HealthStateValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReadyStatusCell) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


