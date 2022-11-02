/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// StackPackIntegration struct for StackPackIntegration
type StackPackIntegration struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Logo        *string  `json:"logo,omitempty"`
	BrandIcon   *string  `json:"brandIcon,omitempty"`
	Categories  []string `json:"categories"`
	IsNew       bool     `json:"isNew"`
	OverviewUrl *string  `json:"overviewUrl,omitempty"`
	ResourceUrl *string  `json:"resourceUrl,omitempty"`
	Faqs        []FAQ    `json:"faqs"`
}

// NewStackPackIntegration instantiates a new StackPackIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackPackIntegration(name string, displayName string, categories []string, isNew bool, faqs []FAQ) *StackPackIntegration {
	this := StackPackIntegration{}
	this.Name = name
	this.DisplayName = displayName
	this.Categories = categories
	this.IsNew = isNew
	this.Faqs = faqs
	return &this
}

// NewStackPackIntegrationWithDefaults instantiates a new StackPackIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackPackIntegrationWithDefaults() *StackPackIntegration {
	this := StackPackIntegration{}
	return &this
}

// GetName returns the Name field value
func (o *StackPackIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StackPackIntegration) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *StackPackIntegration) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *StackPackIntegration) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *StackPackIntegration) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *StackPackIntegration) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *StackPackIntegration) SetLogo(v string) {
	o.Logo = &v
}

// GetBrandIcon returns the BrandIcon field value if set, zero value otherwise.
func (o *StackPackIntegration) GetBrandIcon() string {
	if o == nil || o.BrandIcon == nil {
		var ret string
		return ret
	}
	return *o.BrandIcon
}

// GetBrandIconOk returns a tuple with the BrandIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetBrandIconOk() (*string, bool) {
	if o == nil || o.BrandIcon == nil {
		return nil, false
	}
	return o.BrandIcon, true
}

// HasBrandIcon returns a boolean if a field has been set.
func (o *StackPackIntegration) HasBrandIcon() bool {
	if o != nil && o.BrandIcon != nil {
		return true
	}

	return false
}

// SetBrandIcon gets a reference to the given string and assigns it to the BrandIcon field.
func (o *StackPackIntegration) SetBrandIcon(v string) {
	o.BrandIcon = &v
}

// GetCategories returns the Categories field value
func (o *StackPackIntegration) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *StackPackIntegration) SetCategories(v []string) {
	o.Categories = v
}

// GetIsNew returns the IsNew field value
func (o *StackPackIntegration) GetIsNew() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetIsNewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsNew, true
}

// SetIsNew sets field value
func (o *StackPackIntegration) SetIsNew(v bool) {
	o.IsNew = v
}

// GetOverviewUrl returns the OverviewUrl field value if set, zero value otherwise.
func (o *StackPackIntegration) GetOverviewUrl() string {
	if o == nil || o.OverviewUrl == nil {
		var ret string
		return ret
	}
	return *o.OverviewUrl
}

// GetOverviewUrlOk returns a tuple with the OverviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetOverviewUrlOk() (*string, bool) {
	if o == nil || o.OverviewUrl == nil {
		return nil, false
	}
	return o.OverviewUrl, true
}

// HasOverviewUrl returns a boolean if a field has been set.
func (o *StackPackIntegration) HasOverviewUrl() bool {
	if o != nil && o.OverviewUrl != nil {
		return true
	}

	return false
}

// SetOverviewUrl gets a reference to the given string and assigns it to the OverviewUrl field.
func (o *StackPackIntegration) SetOverviewUrl(v string) {
	o.OverviewUrl = &v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *StackPackIntegration) GetResourceUrl() string {
	if o == nil || o.ResourceUrl == nil {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetResourceUrlOk() (*string, bool) {
	if o == nil || o.ResourceUrl == nil {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *StackPackIntegration) HasResourceUrl() bool {
	if o != nil && o.ResourceUrl != nil {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *StackPackIntegration) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetFaqs returns the Faqs field value
func (o *StackPackIntegration) GetFaqs() []FAQ {
	if o == nil {
		var ret []FAQ
		return ret
	}

	return o.Faqs
}

// GetFaqsOk returns a tuple with the Faqs field value
// and a boolean to check if the value has been set.
func (o *StackPackIntegration) GetFaqsOk() ([]FAQ, bool) {
	if o == nil {
		return nil, false
	}
	return o.Faqs, true
}

// SetFaqs sets field value
func (o *StackPackIntegration) SetFaqs(v []FAQ) {
	o.Faqs = v
}

func (o StackPackIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.BrandIcon != nil {
		toSerialize["brandIcon"] = o.BrandIcon
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
	if o.ResourceUrl != nil {
		toSerialize["resourceUrl"] = o.ResourceUrl
	}
	if true {
		toSerialize["faqs"] = o.Faqs
	}
	return json.Marshal(toSerialize)
}

type NullableStackPackIntegration struct {
	value *StackPackIntegration
	isSet bool
}

func (v NullableStackPackIntegration) Get() *StackPackIntegration {
	return v.value
}

func (v *NullableStackPackIntegration) Set(val *StackPackIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableStackPackIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableStackPackIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackPackIntegration(val *StackPackIntegration) *NullableStackPackIntegration {
	return &NullableStackPackIntegration{value: val, isSet: true}
}

func (v NullableStackPackIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackPackIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
