# TooManyAnomaliesError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**NumberOfMatches** | Pointer to **int64** |  | [optional] 
**MaxAllowed** | Pointer to **int64** |  | [optional] 

## Methods

### NewTooManyAnomaliesError

`func NewTooManyAnomaliesError(message string, ) *TooManyAnomaliesError`

NewTooManyAnomaliesError instantiates a new TooManyAnomaliesError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyAnomaliesErrorWithDefaults

`func NewTooManyAnomaliesErrorWithDefaults() *TooManyAnomaliesError`

NewTooManyAnomaliesErrorWithDefaults instantiates a new TooManyAnomaliesError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TooManyAnomaliesError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TooManyAnomaliesError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TooManyAnomaliesError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNumberOfMatches

`func (o *TooManyAnomaliesError) GetNumberOfMatches() int64`

GetNumberOfMatches returns the NumberOfMatches field if non-nil, zero value otherwise.

### GetNumberOfMatchesOk

`func (o *TooManyAnomaliesError) GetNumberOfMatchesOk() (*int64, bool)`

GetNumberOfMatchesOk returns a tuple with the NumberOfMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMatches

`func (o *TooManyAnomaliesError) SetNumberOfMatches(v int64)`

SetNumberOfMatches sets NumberOfMatches field to given value.

### HasNumberOfMatches

`func (o *TooManyAnomaliesError) HasNumberOfMatches() bool`

HasNumberOfMatches returns a boolean if a field has been set.

### GetMaxAllowed

`func (o *TooManyAnomaliesError) GetMaxAllowed() int64`

GetMaxAllowed returns the MaxAllowed field if non-nil, zero value otherwise.

### GetMaxAllowedOk

`func (o *TooManyAnomaliesError) GetMaxAllowedOk() (*int64, bool)`

GetMaxAllowedOk returns a tuple with the MaxAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowed

`func (o *TooManyAnomaliesError) SetMaxAllowed(v int64)`

SetMaxAllowed sets MaxAllowed field to given value.

### HasMaxAllowed

`func (o *TooManyAnomaliesError) HasMaxAllowed() bool`

HasMaxAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


