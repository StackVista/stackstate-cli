# ExecuteScriptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **map[string]interface{}** | Contains a &#x60;value&#x60; and &#x60;_type&#x60; property. This is described this way, because &#x60;value&#x60; can hold any JSON type and most open api generators (but not all) have problems with such types or are inconsistent with the OpenAPI spec. The &#x60;_type&#x60; property describes the original type of the result value during script execution. The &#x60;value&#x60; property holds the resulting value serialized to JSON. The &#x60;value&#x60; property can be of any JSON type, i.e. null, undefined, boolean, number, string, array or object.  | 

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


