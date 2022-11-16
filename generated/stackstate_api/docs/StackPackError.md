# StackPackError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retryable** | **bool** |  | 
**Action** | Pointer to **string** |  | [optional] 
**Error** | **map[string]interface{}** |  | 

## Methods

### NewStackPackError

`func NewStackPackError(retryable bool, error_ map[string]interface{}, ) *StackPackError`

NewStackPackError instantiates a new StackPackError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackPackErrorWithDefaults

`func NewStackPackErrorWithDefaults() *StackPackError`

NewStackPackErrorWithDefaults instantiates a new StackPackError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryable

`func (o *StackPackError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *StackPackError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *StackPackError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.


### GetAction

`func (o *StackPackError) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StackPackError) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StackPackError) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *StackPackError) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetError

`func (o *StackPackError) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StackPackError) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StackPackError) SetError(v map[string]interface{})`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


