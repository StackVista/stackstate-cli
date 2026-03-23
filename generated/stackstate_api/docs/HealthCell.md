# HealthCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | [**HealthStateValue**](HealthStateValue.md) |  | 

## Methods

### NewHealthCell

`func NewHealthCell(type_ string, state HealthStateValue, ) *HealthCell`

NewHealthCell instantiates a new HealthCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCellWithDefaults

`func NewHealthCellWithDefaults() *HealthCell`

NewHealthCellWithDefaults instantiates a new HealthCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthCell) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCell) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCell) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *HealthCell) GetState() HealthStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HealthCell) GetStateOk() (*HealthStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HealthCell) SetState(v HealthStateValue)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


