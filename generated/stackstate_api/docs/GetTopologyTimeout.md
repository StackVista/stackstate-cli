# GetTopologyTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**TimeoutSeconds** | **int32** |  | 

## Methods

### NewGetTopologyTimeout

`func NewGetTopologyTimeout(type_ string, message string, timeoutSeconds int32, ) *GetTopologyTimeout`

NewGetTopologyTimeout instantiates a new GetTopologyTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTopologyTimeoutWithDefaults

`func NewGetTopologyTimeoutWithDefaults() *GetTopologyTimeout`

NewGetTopologyTimeoutWithDefaults instantiates a new GetTopologyTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetTopologyTimeout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTopologyTimeout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTopologyTimeout) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *GetTopologyTimeout) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetTopologyTimeout) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetTopologyTimeout) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimeoutSeconds

`func (o *GetTopologyTimeout) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *GetTopologyTimeout) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *GetTopologyTimeout) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


