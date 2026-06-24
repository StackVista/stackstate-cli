# PresentationMetricPerspectiveSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **string** |  | 
**Title** | **string** |  | 
**Order** | **float64** |  | 
**Metrics** | Pointer to [**[]OrderedComponentPresentationMetric**](OrderedComponentPresentationMetric.md) |  | [optional] 

## Methods

### NewPresentationMetricPerspectiveSection

`func NewPresentationMetricPerspectiveSection(sectionId string, title string, order float64, ) *PresentationMetricPerspectiveSection`

NewPresentationMetricPerspectiveSection instantiates a new PresentationMetricPerspectiveSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationMetricPerspectiveSectionWithDefaults

`func NewPresentationMetricPerspectiveSectionWithDefaults() *PresentationMetricPerspectiveSection`

NewPresentationMetricPerspectiveSectionWithDefaults instantiates a new PresentationMetricPerspectiveSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *PresentationMetricPerspectiveSection) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PresentationMetricPerspectiveSection) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PresentationMetricPerspectiveSection) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetTitle

`func (o *PresentationMetricPerspectiveSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationMetricPerspectiveSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationMetricPerspectiveSection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOrder

`func (o *PresentationMetricPerspectiveSection) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PresentationMetricPerspectiveSection) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PresentationMetricPerspectiveSection) SetOrder(v float64)`

SetOrder sets Order field to given value.


### GetMetrics

`func (o *PresentationMetricPerspectiveSection) GetMetrics() []OrderedComponentPresentationMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PresentationMetricPerspectiveSection) GetMetricsOk() (*[]OrderedComponentPresentationMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PresentationMetricPerspectiveSection) SetMetrics(v []OrderedComponentPresentationMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PresentationMetricPerspectiveSection) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


