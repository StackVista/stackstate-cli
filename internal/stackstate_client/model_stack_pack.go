/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
)

// StackPack struct for StackPack
type StackPack struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Version string `json:"version"`
	Logo *string `json:"logo,omitempty"`
	Categories []string `json:"categories"`
	IsNew bool `json:"isNew"`
	OverviewUrl *string `json:"overviewUrl,omitempty"`
	DetailedOverviewUrl *string `json:"detailedOverviewUrl,omitempty"`
	ResourcesUrl *string `json:"resourcesUrl,omitempty"`
	Faqs *[]StackPackFaqs `json:"faqs,omitempty"`
	ConfigurationUrls *[][]string `json:"configurationUrls,omitempty"`
	ReleaseStatus string `json:"releaseStatus"`
	IsCompatible bool `json:"isCompatible"`
}

// NewStackPack instantiates a new StackPack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackPack(name string, displayName string, version string, categories []string, isNew bool, releaseStatus string, isCompatible bool) *StackPack {
	this := StackPack{}
	this.Name = name
	this.DisplayName = displayName
	this.Version = version
	this.Categories = categories
	this.IsNew = isNew
	this.ReleaseStatus = releaseStatus
	this.IsCompatible = isCompatible
	return &this
}

// NewStackPackWithDefaults instantiates a new StackPack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackPackWithDefaults() *StackPack {
	this := StackPack{}
	return &this
}

// GetName returns the Name field value
func (o *StackPack) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StackPack) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *StackPack) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *StackPack) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetVersion returns the Version field value
func (o *StackPack) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *StackPack) SetVersion(v string) {
	o.Version = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *StackPack) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *StackPack) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *StackPack) SetLogo(v string) {
	o.Logo = &v
}

// GetCategories returns the Categories field value
func (o *StackPack) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetCategoriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *StackPack) SetCategories(v []string) {
	o.Categories = v
}

// GetIsNew returns the IsNew field value
func (o *StackPack) GetIsNew() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetIsNewOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsNew, true
}

// SetIsNew sets field value
func (o *StackPack) SetIsNew(v bool) {
	o.IsNew = v
}

// GetOverviewUrl returns the OverviewUrl field value if set, zero value otherwise.
func (o *StackPack) GetOverviewUrl() string {
	if o == nil || o.OverviewUrl == nil {
		var ret string
		return ret
	}
	return *o.OverviewUrl
}

// GetOverviewUrlOk returns a tuple with the OverviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetOverviewUrlOk() (*string, bool) {
	if o == nil || o.OverviewUrl == nil {
		return nil, false
	}
	return o.OverviewUrl, true
}

// HasOverviewUrl returns a boolean if a field has been set.
func (o *StackPack) HasOverviewUrl() bool {
	if o != nil && o.OverviewUrl != nil {
		return true
	}

	return false
}

// SetOverviewUrl gets a reference to the given string and assigns it to the OverviewUrl field.
func (o *StackPack) SetOverviewUrl(v string) {
	o.OverviewUrl = &v
}

// GetDetailedOverviewUrl returns the DetailedOverviewUrl field value if set, zero value otherwise.
func (o *StackPack) GetDetailedOverviewUrl() string {
	if o == nil || o.DetailedOverviewUrl == nil {
		var ret string
		return ret
	}
	return *o.DetailedOverviewUrl
}

// GetDetailedOverviewUrlOk returns a tuple with the DetailedOverviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetDetailedOverviewUrlOk() (*string, bool) {
	if o == nil || o.DetailedOverviewUrl == nil {
		return nil, false
	}
	return o.DetailedOverviewUrl, true
}

// HasDetailedOverviewUrl returns a boolean if a field has been set.
func (o *StackPack) HasDetailedOverviewUrl() bool {
	if o != nil && o.DetailedOverviewUrl != nil {
		return true
	}

	return false
}

// SetDetailedOverviewUrl gets a reference to the given string and assigns it to the DetailedOverviewUrl field.
func (o *StackPack) SetDetailedOverviewUrl(v string) {
	o.DetailedOverviewUrl = &v
}

// GetResourcesUrl returns the ResourcesUrl field value if set, zero value otherwise.
func (o *StackPack) GetResourcesUrl() string {
	if o == nil || o.ResourcesUrl == nil {
		var ret string
		return ret
	}
	return *o.ResourcesUrl
}

// GetResourcesUrlOk returns a tuple with the ResourcesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetResourcesUrlOk() (*string, bool) {
	if o == nil || o.ResourcesUrl == nil {
		return nil, false
	}
	return o.ResourcesUrl, true
}

// HasResourcesUrl returns a boolean if a field has been set.
func (o *StackPack) HasResourcesUrl() bool {
	if o != nil && o.ResourcesUrl != nil {
		return true
	}

	return false
}

// SetResourcesUrl gets a reference to the given string and assigns it to the ResourcesUrl field.
func (o *StackPack) SetResourcesUrl(v string) {
	o.ResourcesUrl = &v
}

// GetFaqs returns the Faqs field value if set, zero value otherwise.
func (o *StackPack) GetFaqs() []StackPackFaqs {
	if o == nil || o.Faqs == nil {
		var ret []StackPackFaqs
		return ret
	}
	return *o.Faqs
}

// GetFaqsOk returns a tuple with the Faqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetFaqsOk() (*[]StackPackFaqs, bool) {
	if o == nil || o.Faqs == nil {
		return nil, false
	}
	return o.Faqs, true
}

// HasFaqs returns a boolean if a field has been set.
func (o *StackPack) HasFaqs() bool {
	if o != nil && o.Faqs != nil {
		return true
	}

	return false
}

// SetFaqs gets a reference to the given []StackPackFaqs and assigns it to the Faqs field.
func (o *StackPack) SetFaqs(v []StackPackFaqs) {
	o.Faqs = &v
}

// GetConfigurationUrls returns the ConfigurationUrls field value if set, zero value otherwise.
func (o *StackPack) GetConfigurationUrls() [][]string {
	if o == nil || o.ConfigurationUrls == nil {
		var ret [][]string
		return ret
	}
	return *o.ConfigurationUrls
}

// GetConfigurationUrlsOk returns a tuple with the ConfigurationUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPack) GetConfigurationUrlsOk() (*[][]string, bool) {
	if o == nil || o.ConfigurationUrls == nil {
		return nil, false
	}
	return o.ConfigurationUrls, true
}

// HasConfigurationUrls returns a boolean if a field has been set.
func (o *StackPack) HasConfigurationUrls() bool {
	if o != nil && o.ConfigurationUrls != nil {
		return true
	}

	return false
}

// SetConfigurationUrls gets a reference to the given [][]string and assigns it to the ConfigurationUrls field.
func (o *StackPack) SetConfigurationUrls(v [][]string) {
	o.ConfigurationUrls = &v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *StackPack) GetReleaseStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetReleaseStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *StackPack) SetReleaseStatus(v string) {
	o.ReleaseStatus = v
}

// GetIsCompatible returns the IsCompatible field value
func (o *StackPack) GetIsCompatible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompatible
}

// GetIsCompatibleOk returns a tuple with the IsCompatible field value
// and a boolean to check if the value has been set.
func (o *StackPack) GetIsCompatibleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsCompatible, true
}

// SetIsCompatible sets field value
func (o *StackPack) SetIsCompatible(v bool) {
	o.IsCompatible = v
}

func (o StackPack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["isNew"] = o.IsNew
	}
	if o.OverviewUrl != nil {
		toSerialize["overviewUrl"] = o.OverviewUrl
	}
	if o.DetailedOverviewUrl != nil {
		toSerialize["detailedOverviewUrl"] = o.DetailedOverviewUrl
	}
	if o.ResourcesUrl != nil {
		toSerialize["resourcesUrl"] = o.ResourcesUrl
	}
	if o.Faqs != nil {
		toSerialize["faqs"] = o.Faqs
	}
	if o.ConfigurationUrls != nil {
		toSerialize["configurationUrls"] = o.ConfigurationUrls
	}
	if true {
		toSerialize["releaseStatus"] = o.ReleaseStatus
	}
	if true {
		toSerialize["isCompatible"] = o.IsCompatible
	}
	return json.Marshal(toSerialize)
}

type NullableStackPack struct {
	value *StackPack
	isSet bool
}

func (v NullableStackPack) Get() *StackPack {
	return v.value
}

func (v *NullableStackPack) Set(val *StackPack) {
	v.value = val
	v.isSet = true
}

func (v NullableStackPack) IsSet() bool {
	return v.isSet
}

func (v *NullableStackPack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackPack(val *StackPack) *NullableStackPack {
	return &NullableStackPack{value: val, isSet: true}
}

func (v NullableStackPack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackPack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

