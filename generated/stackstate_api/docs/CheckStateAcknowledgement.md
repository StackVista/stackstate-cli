# CheckStateAcknowledgement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimeStamp** | Pointer to **int64** |  | [optional] 
**AcknowledgedTimestamp** | **int64** |  | 
**Message** | **string** |  | 

## Methods

### NewCheckStateAcknowledgement

`func NewCheckStateAcknowledgement(type_ string, acknowledgedTimestamp int64, message string, ) *CheckStateAcknowledgement`

NewCheckStateAcknowledgement instantiates a new CheckStateAcknowledgement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckStateAcknowledgementWithDefaults

`func NewCheckStateAcknowledgementWithDefaults() *CheckStateAcknowledgement`

NewCheckStateAcknowledgementWithDefaults instantiates a new CheckStateAcknowledgement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckStateAcknowledgement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckStateAcknowledgement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckStateAcknowledgement) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CheckStateAcknowledgement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckStateAcknowledgement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckStateAcknowledgement) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckStateAcknowledgement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimeStamp

`func (o *CheckStateAcknowledgement) GetLastUpdateTimeStamp() int64`

GetLastUpdateTimeStamp returns the LastUpdateTimeStamp field if non-nil, zero value otherwise.

### GetLastUpdateTimeStampOk

`func (o *CheckStateAcknowledgement) GetLastUpdateTimeStampOk() (*int64, bool)`

GetLastUpdateTimeStampOk returns a tuple with the LastUpdateTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeStamp

`func (o *CheckStateAcknowledgement) SetLastUpdateTimeStamp(v int64)`

SetLastUpdateTimeStamp sets LastUpdateTimeStamp field to given value.

### HasLastUpdateTimeStamp

`func (o *CheckStateAcknowledgement) HasLastUpdateTimeStamp() bool`

HasLastUpdateTimeStamp returns a boolean if a field has been set.

### GetAcknowledgedTimestamp

`func (o *CheckStateAcknowledgement) GetAcknowledgedTimestamp() int64`

GetAcknowledgedTimestamp returns the AcknowledgedTimestamp field if non-nil, zero value otherwise.

### GetAcknowledgedTimestampOk

`func (o *CheckStateAcknowledgement) GetAcknowledgedTimestampOk() (*int64, bool)`

GetAcknowledgedTimestampOk returns a tuple with the AcknowledgedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTimestamp

`func (o *CheckStateAcknowledgement) SetAcknowledgedTimestamp(v int64)`

SetAcknowledgedTimestamp sets AcknowledgedTimestamp field to given value.


### GetMessage

`func (o *CheckStateAcknowledgement) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckStateAcknowledgement) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckStateAcknowledgement) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


