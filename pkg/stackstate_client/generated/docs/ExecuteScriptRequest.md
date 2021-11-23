# ExecuteScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutMs** | Pointer to **int32** |  | [optional] 
**Script** | **string** |  | 
**ArgumentsScript** | Pointer to **string** |  | [optional] 

## Methods

### NewExecuteScriptRequest

`func NewExecuteScriptRequest(script string, ) *ExecuteScriptRequest`

NewExecuteScriptRequest instantiates a new ExecuteScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteScriptRequestWithDefaults

`func NewExecuteScriptRequestWithDefaults() *ExecuteScriptRequest`

NewExecuteScriptRequestWithDefaults instantiates a new ExecuteScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutMs

`func (o *ExecuteScriptRequest) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ExecuteScriptRequest) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ExecuteScriptRequest) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *ExecuteScriptRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### GetScript

`func (o *ExecuteScriptRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ExecuteScriptRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ExecuteScriptRequest) SetScript(v string)`

SetScript sets Script field to given value.


### GetArgumentsScript

`func (o *ExecuteScriptRequest) GetArgumentsScript() string`

GetArgumentsScript returns the ArgumentsScript field if non-nil, zero value otherwise.

### GetArgumentsScriptOk

`func (o *ExecuteScriptRequest) GetArgumentsScriptOk() (*string, bool)`

GetArgumentsScriptOk returns a tuple with the ArgumentsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentsScript

`func (o *ExecuteScriptRequest) SetArgumentsScript(v string)`

SetArgumentsScript sets ArgumentsScript field to given value.

### HasArgumentsScript

`func (o *ExecuteScriptRequest) HasArgumentsScript() bool`

HasArgumentsScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


