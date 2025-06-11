# DashboardAuthorizationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A detailed error message describing why the operation failed. | 
**DashboardIdOrUrn** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewDashboardAuthorizationError

`func NewDashboardAuthorizationError(message string, dashboardIdOrUrn string, userName string, ) *DashboardAuthorizationError`

NewDashboardAuthorizationError instantiates a new DashboardAuthorizationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardAuthorizationErrorWithDefaults

`func NewDashboardAuthorizationErrorWithDefaults() *DashboardAuthorizationError`

NewDashboardAuthorizationErrorWithDefaults instantiates a new DashboardAuthorizationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DashboardAuthorizationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DashboardAuthorizationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DashboardAuthorizationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDashboardIdOrUrn

`func (o *DashboardAuthorizationError) GetDashboardIdOrUrn() string`

GetDashboardIdOrUrn returns the DashboardIdOrUrn field if non-nil, zero value otherwise.

### GetDashboardIdOrUrnOk

`func (o *DashboardAuthorizationError) GetDashboardIdOrUrnOk() (*string, bool)`

GetDashboardIdOrUrnOk returns a tuple with the DashboardIdOrUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardIdOrUrn

`func (o *DashboardAuthorizationError) SetDashboardIdOrUrn(v string)`

SetDashboardIdOrUrn sets DashboardIdOrUrn field to given value.


### GetUserName

`func (o *DashboardAuthorizationError) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DashboardAuthorizationError) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DashboardAuthorizationError) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


