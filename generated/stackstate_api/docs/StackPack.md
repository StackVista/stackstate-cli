# StackPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Version** | **string** |  | 
**Logo** | Pointer to **string** |  | [optional] 
**Categories** | **[]string** |  | 
**IsNew** | **bool** |  | 
**OverviewUrl** | Pointer to **string** |  | [optional] 
**DetailedOverviewUrl** | Pointer to **string** |  | [optional] 
**ResourcesUrl** | Pointer to **string** |  | [optional] 
**Faqs** | Pointer to [**[]StackPackFaqs**](StackPackFaqs.md) |  | [optional] 
**ConfigurationUrls** | Pointer to **[][]string** |  | [optional] 
**ReleaseStatus** | **string** |  | 
**IsCompatible** | **bool** |  | 

## Methods

### NewStackPack

`func NewStackPack(name string, displayName string, version string, categories []string, isNew bool, releaseStatus string, isCompatible bool, ) *StackPack`

NewStackPack instantiates a new StackPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackPackWithDefaults

`func NewStackPackWithDefaults() *StackPack`

NewStackPackWithDefaults instantiates a new StackPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StackPack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackPack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackPack) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *StackPack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StackPack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StackPack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetVersion

`func (o *StackPack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StackPack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StackPack) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLogo

`func (o *StackPack) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *StackPack) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *StackPack) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *StackPack) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCategories

`func (o *StackPack) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *StackPack) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *StackPack) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetIsNew

`func (o *StackPack) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *StackPack) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *StackPack) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.


### GetOverviewUrl

`func (o *StackPack) GetOverviewUrl() string`

GetOverviewUrl returns the OverviewUrl field if non-nil, zero value otherwise.

### GetOverviewUrlOk

`func (o *StackPack) GetOverviewUrlOk() (*string, bool)`

GetOverviewUrlOk returns a tuple with the OverviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewUrl

`func (o *StackPack) SetOverviewUrl(v string)`

SetOverviewUrl sets OverviewUrl field to given value.

### HasOverviewUrl

`func (o *StackPack) HasOverviewUrl() bool`

HasOverviewUrl returns a boolean if a field has been set.

### GetDetailedOverviewUrl

`func (o *StackPack) GetDetailedOverviewUrl() string`

GetDetailedOverviewUrl returns the DetailedOverviewUrl field if non-nil, zero value otherwise.

### GetDetailedOverviewUrlOk

`func (o *StackPack) GetDetailedOverviewUrlOk() (*string, bool)`

GetDetailedOverviewUrlOk returns a tuple with the DetailedOverviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedOverviewUrl

`func (o *StackPack) SetDetailedOverviewUrl(v string)`

SetDetailedOverviewUrl sets DetailedOverviewUrl field to given value.

### HasDetailedOverviewUrl

`func (o *StackPack) HasDetailedOverviewUrl() bool`

HasDetailedOverviewUrl returns a boolean if a field has been set.

### GetResourcesUrl

`func (o *StackPack) GetResourcesUrl() string`

GetResourcesUrl returns the ResourcesUrl field if non-nil, zero value otherwise.

### GetResourcesUrlOk

`func (o *StackPack) GetResourcesUrlOk() (*string, bool)`

GetResourcesUrlOk returns a tuple with the ResourcesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesUrl

`func (o *StackPack) SetResourcesUrl(v string)`

SetResourcesUrl sets ResourcesUrl field to given value.

### HasResourcesUrl

`func (o *StackPack) HasResourcesUrl() bool`

HasResourcesUrl returns a boolean if a field has been set.

### GetFaqs

`func (o *StackPack) GetFaqs() []StackPackFaqs`

GetFaqs returns the Faqs field if non-nil, zero value otherwise.

### GetFaqsOk

`func (o *StackPack) GetFaqsOk() (*[]StackPackFaqs, bool)`

GetFaqsOk returns a tuple with the Faqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaqs

`func (o *StackPack) SetFaqs(v []StackPackFaqs)`

SetFaqs sets Faqs field to given value.

### HasFaqs

`func (o *StackPack) HasFaqs() bool`

HasFaqs returns a boolean if a field has been set.

### GetConfigurationUrls

`func (o *StackPack) GetConfigurationUrls() [][]string`

GetConfigurationUrls returns the ConfigurationUrls field if non-nil, zero value otherwise.

### GetConfigurationUrlsOk

`func (o *StackPack) GetConfigurationUrlsOk() (*[][]string, bool)`

GetConfigurationUrlsOk returns a tuple with the ConfigurationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUrls

`func (o *StackPack) SetConfigurationUrls(v [][]string)`

SetConfigurationUrls sets ConfigurationUrls field to given value.

### HasConfigurationUrls

`func (o *StackPack) HasConfigurationUrls() bool`

HasConfigurationUrls returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *StackPack) GetReleaseStatus() string`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *StackPack) GetReleaseStatusOk() (*string, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *StackPack) SetReleaseStatus(v string)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetIsCompatible

`func (o *StackPack) GetIsCompatible() bool`

GetIsCompatible returns the IsCompatible field if non-nil, zero value otherwise.

### GetIsCompatibleOk

`func (o *StackPack) GetIsCompatibleOk() (*bool, bool)`

GetIsCompatibleOk returns a tuple with the IsCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompatible

`func (o *StackPack) SetIsCompatible(v bool)`

SetIsCompatible sets IsCompatible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


