# CheckError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimeStamp** | Pointer to **int64** |  | [optional] 
**Message** | **string** |  | 
**DetailedMessage** | **string** |  | 
**State** | [**ErrorStateValue**](ErrorStateValue.md) |  | 

## Methods

### NewCheckError

`func NewCheckError(type_ string, message string, detailedMessage string, state ErrorStateValue, ) *CheckError`

NewCheckError instantiates a new CheckError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckErrorWithDefaults

`func NewCheckErrorWithDefaults() *CheckError`

NewCheckErrorWithDefaults instantiates a new CheckError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckError) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CheckError) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckError) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckError) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimeStamp

`func (o *CheckError) GetLastUpdateTimeStamp() int64`

GetLastUpdateTimeStamp returns the LastUpdateTimeStamp field if non-nil, zero value otherwise.

### GetLastUpdateTimeStampOk

`func (o *CheckError) GetLastUpdateTimeStampOk() (*int64, bool)`

GetLastUpdateTimeStampOk returns a tuple with the LastUpdateTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeStamp

`func (o *CheckError) SetLastUpdateTimeStamp(v int64)`

SetLastUpdateTimeStamp sets LastUpdateTimeStamp field to given value.

### HasLastUpdateTimeStamp

`func (o *CheckError) HasLastUpdateTimeStamp() bool`

HasLastUpdateTimeStamp returns a boolean if a field has been set.

### GetMessage

`func (o *CheckError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetailedMessage

`func (o *CheckError) GetDetailedMessage() string`

GetDetailedMessage returns the DetailedMessage field if non-nil, zero value otherwise.

### GetDetailedMessageOk

`func (o *CheckError) GetDetailedMessageOk() (*string, bool)`

GetDetailedMessageOk returns a tuple with the DetailedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedMessage

`func (o *CheckError) SetDetailedMessage(v string)`

SetDetailedMessage sets DetailedMessage field to given value.


### GetState

`func (o *CheckError) GetState() ErrorStateValue`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CheckError) GetStateOk() (*ErrorStateValue, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CheckError) SetState(v ErrorStateValue)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


