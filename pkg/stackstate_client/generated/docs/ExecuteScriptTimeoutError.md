# ExecuteScriptTimeoutError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**TimeoutMs** | **int32** |  | 

## Methods

### NewExecuteScriptTimeoutError

`func NewExecuteScriptTimeoutError(type_ string, message string, timeoutMs int32, ) *ExecuteScriptTimeoutError`

NewExecuteScriptTimeoutError instantiates a new ExecuteScriptTimeoutError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteScriptTimeoutErrorWithDefaults

`func NewExecuteScriptTimeoutErrorWithDefaults() *ExecuteScriptTimeoutError`

NewExecuteScriptTimeoutErrorWithDefaults instantiates a new ExecuteScriptTimeoutError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExecuteScriptTimeoutError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecuteScriptTimeoutError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecuteScriptTimeoutError) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ExecuteScriptTimeoutError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExecuteScriptTimeoutError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExecuteScriptTimeoutError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimeoutMs

`func (o *ExecuteScriptTimeoutError) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ExecuteScriptTimeoutError) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ExecuteScriptTimeoutError) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


