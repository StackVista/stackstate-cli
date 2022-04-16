# Sstackpack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]SstackpackConfigurations**](SstackpackConfigurations.md) |  | [optional] 
**LatestVersion** | Pointer to [**SstackpackLatestVersion**](SstackpackLatestVersion.md) |  | [optional] 
**NextVersion** | Pointer to [**SstackpackLatestVersion**](SstackpackLatestVersion.md) |  | [optional] 

## Methods

### NewSstackpack

`func NewSstackpack() *Sstackpack`

NewSstackpack instantiates a new Sstackpack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSstackpackWithDefaults

`func NewSstackpackWithDefaults() *Sstackpack`

NewSstackpackWithDefaults instantiates a new Sstackpack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Sstackpack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sstackpack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sstackpack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sstackpack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Sstackpack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Sstackpack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Sstackpack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Sstackpack) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *Sstackpack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Sstackpack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Sstackpack) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Sstackpack) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConfigurations

`func (o *Sstackpack) GetConfigurations() []SstackpackConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Sstackpack) GetConfigurationsOk() (*[]SstackpackConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Sstackpack) SetConfigurations(v []SstackpackConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Sstackpack) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Sstackpack) GetLatestVersion() SstackpackLatestVersion`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Sstackpack) GetLatestVersionOk() (*SstackpackLatestVersion, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Sstackpack) SetLatestVersion(v SstackpackLatestVersion)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Sstackpack) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetNextVersion

`func (o *Sstackpack) GetNextVersion() SstackpackLatestVersion`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *Sstackpack) GetNextVersionOk() (*SstackpackLatestVersion, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *Sstackpack) SetNextVersion(v SstackpackLatestVersion)`

SetNextVersion sets NextVersion field to given value.

### HasNextVersion

`func (o *Sstackpack) HasNextVersion() bool`

HasNextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


