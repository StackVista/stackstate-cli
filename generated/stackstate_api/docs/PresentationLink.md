# PresentationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **string** | Stable identity key for merging across presentations, analogous to fieldId and resourceId. | 
**Title** | **string** | CEL expression that returns the displayed link title. | 
**Target** | **string** | CEL expression that returns a relative or fully-qualified link target. | 
**OpenInNewTab** | Pointer to **bool** | Whether the link opens in a new tab. Defaults to false. | [optional] [default to false]
**Tooltip** | Pointer to **string** | Optional CEL expression that returns tooltip text. | [optional] 
**Filter** | Pointer to **string** | Optional CEL boolean expression deciding whether the link is shown. Missing filters default to true. | [optional] 
**Order** | **float64** | Display order. Higher value means it shows first in UI. | 

## Methods

### NewPresentationLink

`func NewPresentationLink(linkId string, title string, target string, order float64, ) *PresentationLink`

NewPresentationLink instantiates a new PresentationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationLinkWithDefaults

`func NewPresentationLinkWithDefaults() *PresentationLink`

NewPresentationLinkWithDefaults instantiates a new PresentationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *PresentationLink) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *PresentationLink) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *PresentationLink) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetTitle

`func (o *PresentationLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationLink) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTarget

`func (o *PresentationLink) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PresentationLink) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PresentationLink) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetOpenInNewTab

`func (o *PresentationLink) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *PresentationLink) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *PresentationLink) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *PresentationLink) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

### GetTooltip

`func (o *PresentationLink) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *PresentationLink) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *PresentationLink) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *PresentationLink) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetFilter

`func (o *PresentationLink) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PresentationLink) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PresentationLink) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PresentationLink) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOrder

`func (o *PresentationLink) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationLink) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationLink) SetOrder(v float64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


