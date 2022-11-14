# StackPackIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**DisplayName** | **string** |  |
**Logo** | Pointer to **string** |  | [optional]
**BrandIcon** | Pointer to **string** |  | [optional]
**Categories** | **[]string** |  |
**IsNew** | **bool** |  |
**OverviewUrl** | Pointer to **string** |  | [optional]
**ResourceUrl** | Pointer to **string** |  | [optional]
**Faqs** | [**[]FAQ**](FAQ.md) |  |

## Methods

### NewStackPackIntegration

`func NewStackPackIntegration(name string, displayName string, categories []string, isNew bool, faqs []FAQ, ) *StackPackIntegration`

NewStackPackIntegration instantiates a new StackPackIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackPackIntegrationWithDefaults

`func NewStackPackIntegrationWithDefaults() *StackPackIntegration`

NewStackPackIntegrationWithDefaults instantiates a new StackPackIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StackPackIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackPackIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackPackIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *StackPackIntegration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StackPackIntegration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StackPackIntegration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLogo

`func (o *StackPackIntegration) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *StackPackIntegration) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *StackPackIntegration) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *StackPackIntegration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetBrandIcon

`func (o *StackPackIntegration) GetBrandIcon() string`

GetBrandIcon returns the BrandIcon field if non-nil, zero value otherwise.

### GetBrandIconOk

`func (o *StackPackIntegration) GetBrandIconOk() (*string, bool)`

GetBrandIconOk returns a tuple with the BrandIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIcon

`func (o *StackPackIntegration) SetBrandIcon(v string)`

SetBrandIcon sets BrandIcon field to given value.

### HasBrandIcon

`func (o *StackPackIntegration) HasBrandIcon() bool`

HasBrandIcon returns a boolean if a field has been set.

### GetCategories

`func (o *StackPackIntegration) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *StackPackIntegration) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *StackPackIntegration) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetIsNew

`func (o *StackPackIntegration) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *StackPackIntegration) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *StackPackIntegration) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.


### GetOverviewUrl

`func (o *StackPackIntegration) GetOverviewUrl() string`

GetOverviewUrl returns the OverviewUrl field if non-nil, zero value otherwise.

### GetOverviewUrlOk

`func (o *StackPackIntegration) GetOverviewUrlOk() (*string, bool)`

GetOverviewUrlOk returns a tuple with the OverviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewUrl

`func (o *StackPackIntegration) SetOverviewUrl(v string)`

SetOverviewUrl sets OverviewUrl field to given value.

### HasOverviewUrl

`func (o *StackPackIntegration) HasOverviewUrl() bool`

HasOverviewUrl returns a boolean if a field has been set.

### GetResourceUrl

`func (o *StackPackIntegration) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *StackPackIntegration) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *StackPackIntegration) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *StackPackIntegration) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### GetFaqs

`func (o *StackPackIntegration) GetFaqs() []FAQ`

GetFaqs returns the Faqs field if non-nil, zero value otherwise.

### GetFaqsOk

`func (o *StackPackIntegration) GetFaqsOk() (*[]FAQ, bool)`

GetFaqsOk returns a tuple with the Faqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaqs

`func (o *StackPackIntegration) SetFaqs(v []FAQ)`

SetFaqs sets Faqs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
