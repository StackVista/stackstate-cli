# ResolvedComponentPresentationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**LinkId** | **string** | Identity key for this link. | 
**Title** | **string** | Displayed link title. | 
**Target** | **string** | Relative or fully-qualified link target. | 
**OpenInNewTab** | **bool** | Whether the link opens in a new tab. | 
**Tooltip** | Pointer to **string** | Optional tooltip text. | [optional] 

## Methods

### NewResolvedComponentPresentationLink

`func NewResolvedComponentPresentationLink(type_ string, linkId string, title string, target string, openInNewTab bool, ) *ResolvedComponentPresentationLink`

NewResolvedComponentPresentationLink instantiates a new ResolvedComponentPresentationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolvedComponentPresentationLinkWithDefaults

`func NewResolvedComponentPresentationLinkWithDefaults() *ResolvedComponentPresentationLink`

NewResolvedComponentPresentationLinkWithDefaults instantiates a new ResolvedComponentPresentationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResolvedComponentPresentationLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResolvedComponentPresentationLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResolvedComponentPresentationLink) SetType(v string)`

SetType sets Type field to given value.


### GetLinkId

`func (o *ResolvedComponentPresentationLink) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ResolvedComponentPresentationLink) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ResolvedComponentPresentationLink) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetTitle

`func (o *ResolvedComponentPresentationLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResolvedComponentPresentationLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResolvedComponentPresentationLink) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTarget

`func (o *ResolvedComponentPresentationLink) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ResolvedComponentPresentationLink) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ResolvedComponentPresentationLink) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetOpenInNewTab

`func (o *ResolvedComponentPresentationLink) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *ResolvedComponentPresentationLink) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *ResolvedComponentPresentationLink) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.


### GetTooltip

`func (o *ResolvedComponentPresentationLink) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ResolvedComponentPresentationLink) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ResolvedComponentPresentationLink) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ResolvedComponentPresentationLink) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


