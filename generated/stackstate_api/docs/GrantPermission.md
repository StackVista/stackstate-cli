# GrantPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** |  |
**ResourceName** | **string** |  |

## Methods

### NewGrantPermission

`func NewGrantPermission(permission string, resourceName string, ) *GrantPermission`

NewGrantPermission instantiates a new GrantPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPermissionWithDefaults

`func NewGrantPermissionWithDefaults() *GrantPermission`

NewGrantPermissionWithDefaults instantiates a new GrantPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *GrantPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *GrantPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *GrantPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetResourceName

`func (o *GrantPermission) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *GrantPermission) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *GrantPermission) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
