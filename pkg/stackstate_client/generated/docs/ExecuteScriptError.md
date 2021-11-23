# ExecuteScriptError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**Errors** | [**[]ExecuteScriptSyntaxErrorsErrors**](ExecuteScriptSyntaxErrorsErrors.md) |  | 
**Location** | Pointer to [**ScriptLocation**](ScriptLocation.md) |  | [optional] 
**Reason** | **string** |  | 
**ActualReturnType** | **string** |  | 
**ExpectedReturnType** | **string** |  | 
**TimeoutMs** | **int32** |  | 
**Error** | [**ExecuteScriptError**](ExecuteScriptError.md) |  | 

## Methods

### NewExecuteScriptError

`func NewExecuteScriptError(type_ string, message string, errors []ExecuteScriptSyntaxErrorsErrors, reason string, actualReturnType string, expectedReturnType string, timeoutMs int32, error_ ExecuteScriptError, ) *ExecuteScriptError`

NewExecuteScriptError instantiates a new ExecuteScriptError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteScriptErrorWithDefaults

`func NewExecuteScriptErrorWithDefaults() *ExecuteScriptError`

NewExecuteScriptErrorWithDefaults instantiates a new ExecuteScriptError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExecuteScriptError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecuteScriptError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecuteScriptError) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ExecuteScriptError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExecuteScriptError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExecuteScriptError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ExecuteScriptError) GetErrors() []ExecuteScriptSyntaxErrorsErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExecuteScriptError) GetErrorsOk() (*[]ExecuteScriptSyntaxErrorsErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExecuteScriptError) SetErrors(v []ExecuteScriptSyntaxErrorsErrors)`

SetErrors sets Errors field to given value.


### GetLocation

`func (o *ExecuteScriptError) GetLocation() ScriptLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExecuteScriptError) GetLocationOk() (*ScriptLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExecuteScriptError) SetLocation(v ScriptLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExecuteScriptError) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetReason

`func (o *ExecuteScriptError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ExecuteScriptError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ExecuteScriptError) SetReason(v string)`

SetReason sets Reason field to given value.


### GetActualReturnType

`func (o *ExecuteScriptError) GetActualReturnType() string`

GetActualReturnType returns the ActualReturnType field if non-nil, zero value otherwise.

### GetActualReturnTypeOk

`func (o *ExecuteScriptError) GetActualReturnTypeOk() (*string, bool)`

GetActualReturnTypeOk returns a tuple with the ActualReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualReturnType

`func (o *ExecuteScriptError) SetActualReturnType(v string)`

SetActualReturnType sets ActualReturnType field to given value.


### GetExpectedReturnType

`func (o *ExecuteScriptError) GetExpectedReturnType() string`

GetExpectedReturnType returns the ExpectedReturnType field if non-nil, zero value otherwise.

### GetExpectedReturnTypeOk

`func (o *ExecuteScriptError) GetExpectedReturnTypeOk() (*string, bool)`

GetExpectedReturnTypeOk returns a tuple with the ExpectedReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReturnType

`func (o *ExecuteScriptError) SetExpectedReturnType(v string)`

SetExpectedReturnType sets ExpectedReturnType field to given value.


### GetTimeoutMs

`func (o *ExecuteScriptError) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ExecuteScriptError) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ExecuteScriptError) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.


### GetError

`func (o *ExecuteScriptError) GetError() ExecuteScriptError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExecuteScriptError) GetErrorOk() (*ExecuteScriptError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExecuteScriptError) SetError(v ExecuteScriptError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


