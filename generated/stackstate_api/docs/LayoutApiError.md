# LayoutApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewLayoutApiError

`func NewLayoutApiError(statusCode string, message string, ) *LayoutApiError`

NewLayoutApiError instantiates a new LayoutApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutApiErrorWithDefaults

`func NewLayoutApiErrorWithDefaults() *LayoutApiError`

NewLayoutApiErrorWithDefaults instantiates a new LayoutApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *LayoutApiError) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *LayoutApiError) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *LayoutApiError) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *LayoutApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LayoutApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LayoutApiError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


