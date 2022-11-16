# FullStackPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Version** | **string** |  | 
**Logo** | Pointer to **string** |  | [optional] 
**Categories** | **[]string** |  | 
**IsNew** | **bool** |  | 
**IsMultiConfig** | **bool** |  | 
**OverviewUrl** | Pointer to **string** |  | [optional] 
**DetailedOverviewUrl** | Pointer to **string** |  | [optional] 
**ResourcesUrl** | Pointer to **string** |  | [optional] 
**Faqs** | [**[]FAQ**](FAQ.md) |  | 
**ConfigurationUrls** | **[][]string** |  | 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | 
**Steps** | [**[]StackPackStep**](StackPackStep.md) |  | 
**Integrations** | [**[]StackPackIntegration**](StackPackIntegration.md) |  | 
**ReleaseNotes** | Pointer to **string** |  | [optional] 
**UpgradeInstructions** | Pointer to **string** |  | [optional] 
**Configurations** | [**[]StackPackConfiguration**](StackPackConfiguration.md) |  | 
**NextVersion** | Pointer to [**FullStackPack**](FullStackPack.md) |  | [optional] 
**LatestVersion** | Pointer to [**FullStackPack**](FullStackPack.md) |  | [optional] 
**CanUninstall** | **bool** |  | 
**IsCompatible** | **bool** |  | 

## Methods

### NewFullStackPack

`func NewFullStackPack(name string, displayName string, version string, categories []string, isNew bool, isMultiConfig bool, faqs []FAQ, configurationUrls [][]string, releaseStatus ReleaseStatus, steps []StackPackStep, integrations []StackPackIntegration, configurations []StackPackConfiguration, canUninstall bool, isCompatible bool, ) *FullStackPack`

NewFullStackPack instantiates a new FullStackPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullStackPackWithDefaults

`func NewFullStackPackWithDefaults() *FullStackPack`

NewFullStackPackWithDefaults instantiates a new FullStackPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FullStackPack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullStackPack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullStackPack) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *FullStackPack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FullStackPack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FullStackPack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetVersion

`func (o *FullStackPack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FullStackPack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FullStackPack) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLogo

`func (o *FullStackPack) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *FullStackPack) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *FullStackPack) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *FullStackPack) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCategories

`func (o *FullStackPack) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FullStackPack) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FullStackPack) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetIsNew

`func (o *FullStackPack) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *FullStackPack) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *FullStackPack) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.


### GetIsMultiConfig

`func (o *FullStackPack) GetIsMultiConfig() bool`

GetIsMultiConfig returns the IsMultiConfig field if non-nil, zero value otherwise.

### GetIsMultiConfigOk

`func (o *FullStackPack) GetIsMultiConfigOk() (*bool, bool)`

GetIsMultiConfigOk returns a tuple with the IsMultiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiConfig

`func (o *FullStackPack) SetIsMultiConfig(v bool)`

SetIsMultiConfig sets IsMultiConfig field to given value.


### GetOverviewUrl

`func (o *FullStackPack) GetOverviewUrl() string`

GetOverviewUrl returns the OverviewUrl field if non-nil, zero value otherwise.

### GetOverviewUrlOk

`func (o *FullStackPack) GetOverviewUrlOk() (*string, bool)`

GetOverviewUrlOk returns a tuple with the OverviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewUrl

`func (o *FullStackPack) SetOverviewUrl(v string)`

SetOverviewUrl sets OverviewUrl field to given value.

### HasOverviewUrl

`func (o *FullStackPack) HasOverviewUrl() bool`

HasOverviewUrl returns a boolean if a field has been set.

### GetDetailedOverviewUrl

`func (o *FullStackPack) GetDetailedOverviewUrl() string`

GetDetailedOverviewUrl returns the DetailedOverviewUrl field if non-nil, zero value otherwise.

### GetDetailedOverviewUrlOk

`func (o *FullStackPack) GetDetailedOverviewUrlOk() (*string, bool)`

GetDetailedOverviewUrlOk returns a tuple with the DetailedOverviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedOverviewUrl

`func (o *FullStackPack) SetDetailedOverviewUrl(v string)`

SetDetailedOverviewUrl sets DetailedOverviewUrl field to given value.

### HasDetailedOverviewUrl

`func (o *FullStackPack) HasDetailedOverviewUrl() bool`

HasDetailedOverviewUrl returns a boolean if a field has been set.

### GetResourcesUrl

`func (o *FullStackPack) GetResourcesUrl() string`

GetResourcesUrl returns the ResourcesUrl field if non-nil, zero value otherwise.

### GetResourcesUrlOk

`func (o *FullStackPack) GetResourcesUrlOk() (*string, bool)`

GetResourcesUrlOk returns a tuple with the ResourcesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesUrl

`func (o *FullStackPack) SetResourcesUrl(v string)`

SetResourcesUrl sets ResourcesUrl field to given value.

### HasResourcesUrl

`func (o *FullStackPack) HasResourcesUrl() bool`

HasResourcesUrl returns a boolean if a field has been set.

### GetFaqs

`func (o *FullStackPack) GetFaqs() []FAQ`

GetFaqs returns the Faqs field if non-nil, zero value otherwise.

### GetFaqsOk

`func (o *FullStackPack) GetFaqsOk() (*[]FAQ, bool)`

GetFaqsOk returns a tuple with the Faqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaqs

`func (o *FullStackPack) SetFaqs(v []FAQ)`

SetFaqs sets Faqs field to given value.


### GetConfigurationUrls

`func (o *FullStackPack) GetConfigurationUrls() [][]string`

GetConfigurationUrls returns the ConfigurationUrls field if non-nil, zero value otherwise.

### GetConfigurationUrlsOk

`func (o *FullStackPack) GetConfigurationUrlsOk() (*[][]string, bool)`

GetConfigurationUrlsOk returns a tuple with the ConfigurationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUrls

`func (o *FullStackPack) SetConfigurationUrls(v [][]string)`

SetConfigurationUrls sets ConfigurationUrls field to given value.


### GetReleaseStatus

`func (o *FullStackPack) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *FullStackPack) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *FullStackPack) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetSteps

`func (o *FullStackPack) GetSteps() []StackPackStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FullStackPack) GetStepsOk() (*[]StackPackStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FullStackPack) SetSteps(v []StackPackStep)`

SetSteps sets Steps field to given value.


### GetIntegrations

`func (o *FullStackPack) GetIntegrations() []StackPackIntegration`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *FullStackPack) GetIntegrationsOk() (*[]StackPackIntegration, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *FullStackPack) SetIntegrations(v []StackPackIntegration)`

SetIntegrations sets Integrations field to given value.


### GetReleaseNotes

`func (o *FullStackPack) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *FullStackPack) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *FullStackPack) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *FullStackPack) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetUpgradeInstructions

`func (o *FullStackPack) GetUpgradeInstructions() string`

GetUpgradeInstructions returns the UpgradeInstructions field if non-nil, zero value otherwise.

### GetUpgradeInstructionsOk

`func (o *FullStackPack) GetUpgradeInstructionsOk() (*string, bool)`

GetUpgradeInstructionsOk returns a tuple with the UpgradeInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInstructions

`func (o *FullStackPack) SetUpgradeInstructions(v string)`

SetUpgradeInstructions sets UpgradeInstructions field to given value.

### HasUpgradeInstructions

`func (o *FullStackPack) HasUpgradeInstructions() bool`

HasUpgradeInstructions returns a boolean if a field has been set.

### GetConfigurations

`func (o *FullStackPack) GetConfigurations() []StackPackConfiguration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *FullStackPack) GetConfigurationsOk() (*[]StackPackConfiguration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *FullStackPack) SetConfigurations(v []StackPackConfiguration)`

SetConfigurations sets Configurations field to given value.


### GetNextVersion

`func (o *FullStackPack) GetNextVersion() FullStackPack`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *FullStackPack) GetNextVersionOk() (*FullStackPack, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *FullStackPack) SetNextVersion(v FullStackPack)`

SetNextVersion sets NextVersion field to given value.

### HasNextVersion

`func (o *FullStackPack) HasNextVersion() bool`

HasNextVersion returns a boolean if a field has been set.

### GetLatestVersion

`func (o *FullStackPack) GetLatestVersion() FullStackPack`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *FullStackPack) GetLatestVersionOk() (*FullStackPack, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *FullStackPack) SetLatestVersion(v FullStackPack)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *FullStackPack) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetCanUninstall

`func (o *FullStackPack) GetCanUninstall() bool`

GetCanUninstall returns the CanUninstall field if non-nil, zero value otherwise.

### GetCanUninstallOk

`func (o *FullStackPack) GetCanUninstallOk() (*bool, bool)`

GetCanUninstallOk returns a tuple with the CanUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUninstall

`func (o *FullStackPack) SetCanUninstall(v bool)`

SetCanUninstall sets CanUninstall field to given value.


### GetIsCompatible

`func (o *FullStackPack) GetIsCompatible() bool`

GetIsCompatible returns the IsCompatible field if non-nil, zero value otherwise.

### GetIsCompatibleOk

`func (o *FullStackPack) GetIsCompatibleOk() (*bool, bool)`

GetIsCompatibleOk returns a tuple with the IsCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompatible

`func (o *FullStackPack) SetIsCompatible(v bool)`

SetIsCompatible sets IsCompatible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


