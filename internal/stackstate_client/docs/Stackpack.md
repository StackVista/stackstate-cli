# Stackpack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]StackpackConfigurationsInner**](StackpackConfigurationsInner.md) |  | [optional] 
**LatestVersion** | Pointer to [**StackpackLatestVersion**](StackpackLatestVersion.md) |  | [optional] 
**NextVersion** | Pointer to [**StackpackLatestVersion**](StackpackLatestVersion.md) |  | [optional] 

## Methods

### NewStackpack

`func NewStackpack() *Stackpack`

NewStackpack instantiates a new Stackpack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackpackWithDefaults

`func NewStackpackWithDefaults() *Stackpack`

NewStackpackWithDefaults instantiates a new Stackpack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Stackpack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stackpack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stackpack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Stackpack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Stackpack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Stackpack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Stackpack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Stackpack) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *Stackpack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Stackpack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Stackpack) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Stackpack) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConfigurations

`func (o *Stackpack) GetConfigurations() []StackpackConfigurationsInner`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Stackpack) GetConfigurationsOk() (*[]StackpackConfigurationsInner, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Stackpack) SetConfigurations(v []StackpackConfigurationsInner)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Stackpack) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Stackpack) GetLatestVersion() StackpackLatestVersion`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Stackpack) GetLatestVersionOk() (*StackpackLatestVersion, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Stackpack) SetLatestVersion(v StackpackLatestVersion)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Stackpack) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetNextVersion

`func (o *Stackpack) GetNextVersion() StackpackLatestVersion`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *Stackpack) GetNextVersionOk() (*StackpackLatestVersion, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *Stackpack) SetNextVersion(v StackpackLatestVersion)`

SetNextVersion sets NextVersion field to given value.

### HasNextVersion

`func (o *Stackpack) HasNextVersion() bool`

HasNextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


