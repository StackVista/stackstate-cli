# DashboardNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A detailed error message describing why the operation failed. | 
**DashboardIdOrUrn** | **string** |  | 

## Methods

### NewDashboardNotFoundError

`func NewDashboardNotFoundError(message string, dashboardIdOrUrn string, ) *DashboardNotFoundError`

NewDashboardNotFoundError instantiates a new DashboardNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardNotFoundErrorWithDefaults

`func NewDashboardNotFoundErrorWithDefaults() *DashboardNotFoundError`

NewDashboardNotFoundErrorWithDefaults instantiates a new DashboardNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DashboardNotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DashboardNotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DashboardNotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDashboardIdOrUrn

`func (o *DashboardNotFoundError) GetDashboardIdOrUrn() string`

GetDashboardIdOrUrn returns the DashboardIdOrUrn field if non-nil, zero value otherwise.

### GetDashboardIdOrUrnOk

`func (o *DashboardNotFoundError) GetDashboardIdOrUrnOk() (*string, bool)`

GetDashboardIdOrUrnOk returns a tuple with the DashboardIdOrUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardIdOrUrn

`func (o *DashboardNotFoundError) SetDashboardIdOrUrn(v string)`

SetDashboardIdOrUrn sets DashboardIdOrUrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


