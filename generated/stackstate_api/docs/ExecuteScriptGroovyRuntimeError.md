# ExecuteScriptGroovyRuntimeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Message** | **string** |  | 
**Location** | Pointer to [**ScriptLocation**](ScriptLocation.md) |  | [optional] 

## Methods

### NewExecuteScriptGroovyRuntimeError

`func NewExecuteScriptGroovyRuntimeError(type_ string, message string, ) *ExecuteScriptGroovyRuntimeError`

NewExecuteScriptGroovyRuntimeError instantiates a new ExecuteScriptGroovyRuntimeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteScriptGroovyRuntimeErrorWithDefaults

`func NewExecuteScriptGroovyRuntimeErrorWithDefaults() *ExecuteScriptGroovyRuntimeError`

NewExecuteScriptGroovyRuntimeErrorWithDefaults instantiates a new ExecuteScriptGroovyRuntimeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExecuteScriptGroovyRuntimeError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecuteScriptGroovyRuntimeError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecuteScriptGroovyRuntimeError) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *ExecuteScriptGroovyRuntimeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExecuteScriptGroovyRuntimeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExecuteScriptGroovyRuntimeError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLocation

`func (o *ExecuteScriptGroovyRuntimeError) GetLocation() ScriptLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExecuteScriptGroovyRuntimeError) GetLocationOk() (*ScriptLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExecuteScriptGroovyRuntimeError) SetLocation(v ScriptLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExecuteScriptGroovyRuntimeError) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


