# PresentationHighlightMetricsSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DefaultExpanded** | Pointer to **bool** |  | [optional] 
**Order** | **float64** |  | 
**Metrics** | Pointer to [**[]OrderedComponentPresentationMetric**](OrderedComponentPresentationMetric.md) |  | [optional] 

## Methods

### NewPresentationHighlightMetricsSection

`func NewPresentationHighlightMetricsSection(sectionId string, title string, order float64, ) *PresentationHighlightMetricsSection`

NewPresentationHighlightMetricsSection instantiates a new PresentationHighlightMetricsSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationHighlightMetricsSectionWithDefaults

`func NewPresentationHighlightMetricsSectionWithDefaults() *PresentationHighlightMetricsSection`

NewPresentationHighlightMetricsSectionWithDefaults instantiates a new PresentationHighlightMetricsSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *PresentationHighlightMetricsSection) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PresentationHighlightMetricsSection) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PresentationHighlightMetricsSection) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetTitle

`func (o *PresentationHighlightMetricsSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationHighlightMetricsSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationHighlightMetricsSection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *PresentationHighlightMetricsSection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PresentationHighlightMetricsSection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PresentationHighlightMetricsSection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PresentationHighlightMetricsSection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultExpanded

`func (o *PresentationHighlightMetricsSection) GetDefaultExpanded() bool`

GetDefaultExpanded returns the DefaultExpanded field if non-nil, zero value otherwise.

### GetDefaultExpandedOk

`func (o *PresentationHighlightMetricsSection) GetDefaultExpandedOk() (*bool, bool)`

GetDefaultExpandedOk returns a tuple with the DefaultExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpanded

`func (o *PresentationHighlightMetricsSection) SetDefaultExpanded(v bool)`

SetDefaultExpanded sets DefaultExpanded field to given value.

### HasDefaultExpanded

`func (o *PresentationHighlightMetricsSection) HasDefaultExpanded() bool`

HasDefaultExpanded returns a boolean if a field has been set.

### GetOrder

`func (o *PresentationHighlightMetricsSection) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationHighlightMetricsSection) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationHighlightMetricsSection) SetOrder(v float64)`

SetOrder sets Order field to given value.


### GetMetrics

`func (o *PresentationHighlightMetricsSection) GetMetrics() []OrderedComponentPresentationMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PresentationHighlightMetricsSection) GetMetricsOk() (*[]OrderedComponentPresentationMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PresentationHighlightMetricsSection) SetMetrics(v []OrderedComponentPresentationMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PresentationHighlightMetricsSection) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


