# ComponentMetricSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Metrics** | [**[]BoundMetric**](BoundMetric.md) |  | 
**DefaultExpanded** | **bool** |  | 

## Methods

### NewComponentMetricSection

`func NewComponentMetricSection(sectionId string, title string, metrics []BoundMetric, defaultExpanded bool, ) *ComponentMetricSection`

NewComponentMetricSection instantiates a new ComponentMetricSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentMetricSectionWithDefaults

`func NewComponentMetricSectionWithDefaults() *ComponentMetricSection`

NewComponentMetricSectionWithDefaults instantiates a new ComponentMetricSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *ComponentMetricSection) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ComponentMetricSection) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ComponentMetricSection) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetTitle

`func (o *ComponentMetricSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ComponentMetricSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ComponentMetricSection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ComponentMetricSection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentMetricSection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentMetricSection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentMetricSection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetrics

`func (o *ComponentMetricSection) GetMetrics() []BoundMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ComponentMetricSection) GetMetricsOk() (*[]BoundMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ComponentMetricSection) SetMetrics(v []BoundMetric)`

SetMetrics sets Metrics field to given value.


### GetDefaultExpanded

`func (o *ComponentMetricSection) GetDefaultExpanded() bool`

GetDefaultExpanded returns the DefaultExpanded field if non-nil, zero value otherwise.

### GetDefaultExpandedOk

`func (o *ComponentMetricSection) GetDefaultExpandedOk() (*bool, bool)`

GetDefaultExpandedOk returns a tuple with the DefaultExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpanded

`func (o *ComponentMetricSection) SetDefaultExpanded(v bool)`

SetDefaultExpanded sets DefaultExpanded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


