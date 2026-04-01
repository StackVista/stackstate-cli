# ComponentPresentationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterId** | **string** | Stable identity used for deduplication and cross-stackpack compatibility.  | 
**DisplayName** | Pointer to [**FilterName**](FilterName.md) |  | [optional] 
**MenuSection** | Pointer to **string** | Label to group filters in the UI.  | [optional] 
**Filter** | Pointer to [**ComponentPresentationFilterDefinition**](ComponentPresentationFilterDefinition.md) |  | [optional] 

## Methods

### NewComponentPresentationFilter

`func NewComponentPresentationFilter(filterId string, ) *ComponentPresentationFilter`

NewComponentPresentationFilter instantiates a new ComponentPresentationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPresentationFilterWithDefaults

`func NewComponentPresentationFilterWithDefaults() *ComponentPresentationFilter`

NewComponentPresentationFilterWithDefaults instantiates a new ComponentPresentationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterId

`func (o *ComponentPresentationFilter) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *ComponentPresentationFilter) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *ComponentPresentationFilter) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.


### GetDisplayName

`func (o *ComponentPresentationFilter) GetDisplayName() FilterName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComponentPresentationFilter) GetDisplayNameOk() (*FilterName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComponentPresentationFilter) SetDisplayName(v FilterName)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ComponentPresentationFilter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMenuSection

`func (o *ComponentPresentationFilter) GetMenuSection() string`

GetMenuSection returns the MenuSection field if non-nil, zero value otherwise.

### GetMenuSectionOk

`func (o *ComponentPresentationFilter) GetMenuSectionOk() (*string, bool)`

GetMenuSectionOk returns a tuple with the MenuSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuSection

`func (o *ComponentPresentationFilter) SetMenuSection(v string)`

SetMenuSection sets MenuSection field to given value.

### HasMenuSection

`func (o *ComponentPresentationFilter) HasMenuSection() bool`

HasMenuSection returns a boolean if a field has been set.

### GetFilter

`func (o *ComponentPresentationFilter) GetFilter() ComponentPresentationFilterDefinition`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ComponentPresentationFilter) GetFilterOk() (*ComponentPresentationFilterDefinition, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ComponentPresentationFilter) SetFilter(v ComponentPresentationFilterDefinition)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ComponentPresentationFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


