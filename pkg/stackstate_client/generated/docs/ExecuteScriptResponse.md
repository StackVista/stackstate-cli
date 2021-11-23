# ExecuteScriptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **map[string]interface{}** | Can be of any type: boolean, number, string, array, object or null. | 

## Methods

### NewExecuteScriptResponse

`func NewExecuteScriptResponse(result map[string]interface{}, ) *ExecuteScriptResponse`

NewExecuteScriptResponse instantiates a new ExecuteScriptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteScriptResponseWithDefaults

`func NewExecuteScriptResponseWithDefaults() *ExecuteScriptResponse`

NewExecuteScriptResponseWithDefaults instantiates a new ExecuteScriptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ExecuteScriptResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExecuteScriptResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExecuteScriptResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *ExecuteScriptResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ExecuteScriptResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


