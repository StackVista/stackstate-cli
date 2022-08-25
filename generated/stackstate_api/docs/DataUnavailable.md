# DataUnavailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**RequestTimeMs** | **int32** |  | 
**AvailableSinceMs** | **int32** |  | 

## Methods

### NewDataUnavailable

`func NewDataUnavailable(type_ string, message string, requestTimeMs int32, availableSinceMs int32, ) *DataUnavailable`

NewDataUnavailable instantiates a new DataUnavailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUnavailableWithDefaults

`func NewDataUnavailableWithDefaults() *DataUnavailable`

NewDataUnavailableWithDefaults instantiates a new DataUnavailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataUnavailable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataUnavailable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataUnavailable) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *DataUnavailable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DataUnavailable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DataUnavailable) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestTimeMs

`func (o *DataUnavailable) GetRequestTimeMs() int32`

GetRequestTimeMs returns the RequestTimeMs field if non-nil, zero value otherwise.

### GetRequestTimeMsOk

`func (o *DataUnavailable) GetRequestTimeMsOk() (*int32, bool)`

GetRequestTimeMsOk returns a tuple with the RequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeMs

`func (o *DataUnavailable) SetRequestTimeMs(v int32)`

SetRequestTimeMs sets RequestTimeMs field to given value.


### GetAvailableSinceMs

`func (o *DataUnavailable) GetAvailableSinceMs() int32`

GetAvailableSinceMs returns the AvailableSinceMs field if non-nil, zero value otherwise.

### GetAvailableSinceMsOk

`func (o *DataUnavailable) GetAvailableSinceMsOk() (*int32, bool)`

GetAvailableSinceMsOk returns a tuple with the AvailableSinceMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSinceMs

`func (o *DataUnavailable) SetAvailableSinceMs(v int32)`

SetAvailableSinceMs sets AvailableSinceMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


