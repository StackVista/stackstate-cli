# OverviewErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ErrorMessage** | **string** | Error message describing why the overview request failed. | 

## Methods

### NewOverviewErrorResponse

`func NewOverviewErrorResponse(type_ string, errorMessage string, ) *OverviewErrorResponse`

NewOverviewErrorResponse instantiates a new OverviewErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewErrorResponseWithDefaults

`func NewOverviewErrorResponseWithDefaults() *OverviewErrorResponse`

NewOverviewErrorResponseWithDefaults instantiates a new OverviewErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverviewErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverviewErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverviewErrorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetErrorMessage

`func (o *OverviewErrorResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OverviewErrorResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OverviewErrorResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


