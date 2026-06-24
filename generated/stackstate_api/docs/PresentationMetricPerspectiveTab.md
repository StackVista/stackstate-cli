# PresentationMetricPerspectiveTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TabId** | **string** |  | 
**Title** | **string** |  | 
**Order** | **float64** |  | 
**Sections** | Pointer to [**[]PresentationMetricPerspectiveSection**](PresentationMetricPerspectiveSection.md) |  | [optional] 

## Methods

### NewPresentationMetricPerspectiveTab

`func NewPresentationMetricPerspectiveTab(tabId string, title string, order float64, ) *PresentationMetricPerspectiveTab`

NewPresentationMetricPerspectiveTab instantiates a new PresentationMetricPerspectiveTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationMetricPerspectiveTabWithDefaults

`func NewPresentationMetricPerspectiveTabWithDefaults() *PresentationMetricPerspectiveTab`

NewPresentationMetricPerspectiveTabWithDefaults instantiates a new PresentationMetricPerspectiveTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTabId

`func (o *PresentationMetricPerspectiveTab) GetTabId() string`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *PresentationMetricPerspectiveTab) GetTabIdOk() (*string, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *PresentationMetricPerspectiveTab) SetTabId(v string)`

SetTabId sets TabId field to given value.


### GetTitle

`func (o *PresentationMetricPerspectiveTab) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationMetricPerspectiveTab) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationMetricPerspectiveTab) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOrder

`func (o *PresentationMetricPerspectiveTab) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationMetricPerspectiveTab) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationMetricPerspectiveTab) SetOrder(v float64)`

SetOrder sets Order field to given value.


### GetSections

`func (o *PresentationMetricPerspectiveTab) GetSections() []PresentationMetricPerspectiveSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *PresentationMetricPerspectiveTab) GetSectionsOk() (*[]PresentationMetricPerspectiveSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *PresentationMetricPerspectiveTab) SetSections(v []PresentationMetricPerspectiveSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *PresentationMetricPerspectiveTab) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


