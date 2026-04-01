# PresentationFiltersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]PresentationFilter**](PresentationFilter.md) | Presentation filters in display order. Earlier items are primary (filter bar), later items are secondary (\&quot;More\&quot; section).  | 
**MenuSection** | **string** | Label for the section within \&quot;More\&quot; tab when secondary filters are present. | 

## Methods

### NewPresentationFiltersResponse

`func NewPresentationFiltersResponse(filters []PresentationFilter, menuSection string, ) *PresentationFiltersResponse`

NewPresentationFiltersResponse instantiates a new PresentationFiltersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationFiltersResponseWithDefaults

`func NewPresentationFiltersResponseWithDefaults() *PresentationFiltersResponse`

NewPresentationFiltersResponseWithDefaults instantiates a new PresentationFiltersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *PresentationFiltersResponse) GetFilters() []PresentationFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PresentationFiltersResponse) GetFiltersOk() (*[]PresentationFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PresentationFiltersResponse) SetFilters(v []PresentationFilter)`

SetFilters sets Filters field to given value.


### GetMenuSection

`func (o *PresentationFiltersResponse) GetMenuSection() string`

GetMenuSection returns the MenuSection field if non-nil, zero value otherwise.

### GetMenuSectionOk

`func (o *PresentationFiltersResponse) GetMenuSectionOk() (*string, bool)`

GetMenuSectionOk returns a tuple with the MenuSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuSection

`func (o *PresentationFiltersResponse) SetMenuSection(v string)`

SetMenuSection sets MenuSection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


