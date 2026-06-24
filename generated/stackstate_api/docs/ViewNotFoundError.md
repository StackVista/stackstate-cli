# ViewNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A detailed error message describing why the operation failed. | 
**ViewIdOrIdentifier** | **string** |  | 

## Methods

### NewViewNotFoundError

`func NewViewNotFoundError(message string, viewIdOrIdentifier string, ) *ViewNotFoundError`

NewViewNotFoundError instantiates a new ViewNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewNotFoundErrorWithDefaults

`func NewViewNotFoundErrorWithDefaults() *ViewNotFoundError`

NewViewNotFoundErrorWithDefaults instantiates a new ViewNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ViewNotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ViewNotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ViewNotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetViewIdOrIdentifier

`func (o *ViewNotFoundError) GetViewIdOrIdentifier() string`

GetViewIdOrIdentifier returns the ViewIdOrIdentifier field if non-nil, zero value otherwise.

### GetViewIdOrIdentifierOk

`func (o *ViewNotFoundError) GetViewIdOrIdentifierOk() (*string, bool)`

GetViewIdOrIdentifierOk returns a tuple with the ViewIdOrIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIdOrIdentifier

`func (o *ViewNotFoundError) SetViewIdOrIdentifier(v string)`

SetViewIdOrIdentifier sets ViewIdOrIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


