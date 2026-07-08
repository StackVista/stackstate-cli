# ComponentPresentationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**LinkId** | **string** | Identity key for this link. | 
**Title** | **string** | Displayed link title, or the link identity when the title could not be resolved. | 
**Target** | **string** | Relative or fully-qualified link target. | 
**OpenInNewTab** | **bool** | Whether the link opens in a new tab. | 
**Tooltip** | Pointer to **string** | Optional tooltip text. | [optional] 
**Message** | **string** | Reason why the link could not be resolved. | 

## Methods

### NewComponentPresentationLink

`func NewComponentPresentationLink(type_ string, linkId string, title string, target string, openInNewTab bool, message string, ) *ComponentPresentationLink`

NewComponentPresentationLink instantiates a new ComponentPresentationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPresentationLinkWithDefaults

`func NewComponentPresentationLinkWithDefaults() *ComponentPresentationLink`

NewComponentPresentationLinkWithDefaults instantiates a new ComponentPresentationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentPresentationLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentPresentationLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentPresentationLink) SetType(v string)`

SetType sets Type field to given value.


### GetLinkId

`func (o *ComponentPresentationLink) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ComponentPresentationLink) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ComponentPresentationLink) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetTitle

`func (o *ComponentPresentationLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ComponentPresentationLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ComponentPresentationLink) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTarget

`func (o *ComponentPresentationLink) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ComponentPresentationLink) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ComponentPresentationLink) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetOpenInNewTab

`func (o *ComponentPresentationLink) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *ComponentPresentationLink) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *ComponentPresentationLink) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.


### GetTooltip

`func (o *ComponentPresentationLink) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ComponentPresentationLink) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ComponentPresentationLink) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ComponentPresentationLink) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetMessage

`func (o *ComponentPresentationLink) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ComponentPresentationLink) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ComponentPresentationLink) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


